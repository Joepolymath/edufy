package auth

import (
	"github.com/dgrijalva/jwt-go"
)

// SignupRequestSerializer /* This is used to get the request json from signup
type SignupDto struct {
	Email    string `json:"email" validate:"required,email,max=50"`
	Password string `json:"password" validate:"required,min=8,max=20,strongPassword"`
}

// LoginRequestSerializer /* This is used to get request from the user using both email or phone number one must be pased*/
type LoginDto struct {
	Email    string `json:"email" validate:"required,email,max=250"`
	Password string `json:"password" validate:"required"`
}

type UpdatePersonalInfoDto struct {
	FirstName   string `json:"firstName" validate:"required,min=2,max=50"`
	LastName    string `json:"lastName" validate:"required,min=2,max=50"`
	Address     string `json:"address" validate:"required,min=5,max=100"`
	PhoneNumber string `json:"phoneNumber" validate:"required,min=10,max=15,numeric"`
}

// // AuthenticationResponseSerializer /*  This is the response we send to the user once he successfully sign up */
// type AuthenticationResponseDto struct {
// 	ID           uuid.UUID `json:"id" `
// 	FirstName    string    `json:"first_name" validate:"required,max=250,min=2"`
// 	LastName     string    `json:"last_name" validate:"required,max=250,min=2"`
// 	Email        string    `json:"email"  validate:"required,max=250,min=2"`
// 	IsStaff      bool      `json:"is_staff" `
// 	IsSuperUser  bool      `json:"is_super_user" `
// 	IsVerified   bool      `json:"is_verified" `
// 	AccessToken  string    `json:"access_token,omitempty"`
// 	RefreshToken string    `json:"refresh_token,omitempty"`
// 	Timestamp    time.Time `json:"timestamp" `
// }

// func (authResponse *AuthenticationResponseDto) InitializeData(ctx context.Context, user *userServiceModel.User) error {
// 	pointerAdapters := adapters.NewPointer()

// 	err := copier.Copy(&authResponse, &user)
// 	if err != nil {
// 		logger.Error(ctx, "Error copying  user info in auth response", zap.Error(err))
// 		return err
// 	}

// 	// Since the signup process does not contain first name and last name i have to use empty
// 	if user.FirstName == nil {
// 		user.FirstName = pointerAdapters.StringPointer("")
// 	}
// 	if user.LastName == nil {
// 		user.LastName = pointerAdapters.StringPointer("")
// 	}

// 	token, refreshToken, err := utils.GenerateAllToken(ctx, *user.FirstName, *user.LastName, *user.Password, *&user.ID)
// 	if err != nil {
// 		return err
// 	}
// 	authResponse.AccessToken = token
// 	authResponse.RefreshToken = refreshToken

// 	return nil
// }

// OtpRequestSerializer /* This is used to request otp from the frontend */
type OtpRequestSerializer struct {
	Email string `json:"email" validate:""`
}

// OtpValidateRequestSerializer  /* This is used to request otp from the frontend */
type ValidateOtpDto struct {
	Email string `json:"email" validate:"required,email,max=50"`
	Otp   string `json:"otp" validate:"max=4,min=4"`
}

// type ValidateOtpResponseDto struct {
// 	Email string `json:"email" validate:"max=20,min=4"`
// 	Otp   int    `json:"otp" validate:"max=4,min=4"`
// }

// ForgotPasswordValidateRequestSerializer  /* This is used to request otp from the frontend */
type ForgotPasswordValidateRequestSerializer struct {
	Email    string `json:"email" validate:"max=20,min=4"`
	Password string `json:"password" validate:"max=50,min=4"`
	Otp      string `json:"otp" validate:"max=4,min=4"`
}

// create access token payloads bearing user identity and token validity
type TokenPayload struct {
	ID string
	jwt.StandardClaims
}
