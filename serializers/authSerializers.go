package serializers

// import (
// 	"Learnium/adapters"
// 	"Learnium/logger"
// 	userServiceModel "Learnium/models"
// 	"Learnium/utils"
// 	"context"
// 	"github.com/google/uuid"
// 	"github.com/jinzhu/copier"
// 	"go.uber.org/zap"
// 	"time"
// )

// // SignupRequestSerializer /* This is used to get the request json from signup on a user
// // This serializer accepts a post-request */
// type SignupRequestSerializer struct {
// 	Email    string `json:"email"  validate:"required,max=250,min=2"`
// 	Password string `json:"password"  validate:"required,max=250,min=5"`
// }

// // LoginRequestSerializer /* This is used to get request from the user using both email or phone number one must be pased*/
// type LoginRequestSerializer struct {
// 	Email    string `json:"email" validate:"required"`
// 	Password string `json:"password"  validate:"required,max=250,min=5"`
// }

// // AuthenticationResponseSerializer /*  This is the response we send to the user once he successfully sign up */
// type AuthenticationResponseSerializer struct {
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

// func (authResponse *AuthenticationResponseSerializer) InitializeData(ctx context.Context, user *userServiceModel.User) error {
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

// // OtpRequestSerializer /* This is used to request otp from the frontend */
// type OtpRequestSerializer struct {
// 	Email string `json:"email" validate:""`
// }

// // OtpValidateRequestSerializer  /* This is used to request otp from the frontend */
// type OtpValidateRequestSerializer struct {
// 	Email string `json:"email" validate:"max=20,min=4"`
// 	Otp   string `json:"otp" validate:"max=4,min=4"`
// }

// // ForgotPasswordValidateRequestSerializer  /* This is used to request otp from the frontend */
// type ForgotPasswordValidateRequestSerializer struct {
// 	Email    string `json:"email" validate:"max=20,min=4"`
// 	Password string `json:"password" validate:"max=50,min=4"`
// 	Otp      string `json:"otp" validate:"max=4,min=4"`
// }
