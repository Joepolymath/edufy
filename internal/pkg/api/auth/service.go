package auth

import (
	"Learnium/internal/database"
	"Learnium/internal/pkg/models"
	"Learnium/internal/utils"
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	// "github.com/patrickmn/go-cache"
)

// describe service to be injected by learnium auth controller
type IAuthService interface {
	RegisterUser(ctx context.Context, email, password string) (models.User, error)
	AuthenticateUser(ctx context.Context, email, password string) (models.User, error)
	SetRegistrationOTP(ctx context.Context, email string) string
	ValidateRegistrationOTP(ctx context.Context, email string, providedOtp string) (*models.User, error)
	GenerateJwtTokens(ctx context.Context, ID string) (signedToken string,
		refreshToken string, err error)
	ValidateJwtToken(signedToken string) (claims *TokenPayload, msg string)
	GetUser(ctx context.Context, id string) (models.User, error)
	GetUserByEmail(ctx context.Context, email string) (models.User, error)
	UpdateUser(ctx context.Context, user models.User) (models.User, error)
}

// authentication exposing methods relevant to operations within the learnium auth controller
type AuthService struct {
	repository   IUserRepository
	memStore     database.IRedisDriver
	jwtSecretKey string
}

// create new auth service instance
func NewAuthService(repository IUserRepository, memStore database.IRedisDriver) IAuthService {
	jwtSecretKey := os.Getenv("SECRET_KEY")
	authSrv := &AuthService{
		repository,
		memStore,
		jwtSecretKey,
	}
	service := IAuthService(authSrv)
	return service
}

// register user by creating user record in users table
func (srv *AuthService) RegisterUser(ctx context.Context, email, password string) (models.User, error) {
	hashedPwd := utils.HashSecret(password)

	user := &models.User{
		Email:    email,
		Password: &hashedPwd,
	}
	return srv.repository.CreateUser(ctx, *user)

}

// validate user credentials and return user data
func (srv *AuthService) AuthenticateUser(ctx context.Context, email, password string) (models.User, error) {
	existingUser, err := srv.repository.GetUserByEmail(ctx, email)
	if err != nil {
		return models.User{}, err
	}
	verifiedPassword, err := utils.VerifySecret(*existingUser.Password, password)
	if err != nil {
		return models.User{}, err
	} else if !verifiedPassword {
		return models.User{}, ErrPassword
	} else if !existingUser.EmailVerified {
		return models.User{}, ErrUnverified
	}
	return existingUser, nil
}

// get user with id
func (srv *AuthService) GetUser(ctx context.Context, id string) (models.User, error) {
	return srv.repository.GetUserByID(ctx, id)

}

// get user with email
func (srv *AuthService) GetUserByEmail(ctx context.Context, email string) (models.User, error) {
	return srv.repository.GetUserByEmail(ctx, email)

}

// create account registration otp associated with user account by email
func (srv *AuthService) SetRegistrationOTP(ctx context.Context, email string) string {

	otp := utils.GenerateOTP(4)
	// cache the otp
	cacheKey := fmt.Sprintf("registration/otp/%s", email)
	srv.memStore.SetValue(ctx, cacheKey, otp, 10*time.Minute)
	return otp
}

// validate account registration otp associated with user account by email
// returns user model
func (srv *AuthService) ValidateRegistrationOTP(ctx context.Context, email string, providedOtp string) (*models.User, error) {
	// fetch associated otp from cache
	cacheKey := fmt.Sprintf("registration/otp/%s", email)
	registeredOtp, err := srv.memStore.GetValueString(ctx, cacheKey)
	if err != nil {
		return nil, err
	}
	if *registeredOtp != providedOtp {
		return nil, ErrOTP
	}
	user, err := srv.repository.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, ErrRecordNotFound
	}
	srv.memStore.DeleteValue(ctx, cacheKey)
	user.EmailVerified = true
	update, err := srv.UpdateUser(ctx, user)
	if err != nil {
		log.Println("Error while updating user email verification status::: ", err.Error())
		return &user, nil
	}

	return &update, nil
}

// generate access and refresh tokens for user to consume platform APIs
func (srv *AuthService) GenerateJwtTokens(ctx context.Context, ID string) (signedToken string,
	refreshToken string, err error) {
	// Update the access claims
	claims := &TokenPayload{
		ID: ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * time.Duration(24*30)).Unix(),
		},
	}
	// Update the refresh claims
	refreshClaims := &TokenPayload{
		ID: ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * time.Duration(24*30)).Unix(),
		},
	}

	/* Update the signed token with the claims provided for it and also the secret key*/
	SignedToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(srv.jwtSecretKey))
	if err != nil {
		return "", "", err
	}
	/* Update the refresh token with the refresh claims provided*/
	refreshToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(srv.jwtSecretKey))
	if err != nil {
		return "", "", err
	}

	return SignedToken, refreshToken, nil
}

// validate access token provided by user
func (srv *AuthService) ValidateJwtToken(signedToken string) (claims *TokenPayload, msg string) {
	/*
		This is used to validate the token

	*/
	token, err := jwt.ParseWithClaims(
		signedToken,
		&TokenPayload{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(srv.jwtSecretKey), nil
		},
	)
	if err != nil {
		msg = err.Error()
		return
	}
	claims, ok := token.Claims.(*TokenPayload)
	if !ok {
		msg = fmt.Sprint("the token is invalid")
		msg = err.Error()
		return
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		msg = fmt.Sprintf("token is expired")
		msg = err.Error()
		return
	}
	return claims, msg
}

func (srv *AuthService) UpdateUser(ctx context.Context, user models.User) (models.User, error) {
	return srv.repository.UpdateUser(ctx, user)
}
