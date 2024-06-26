package auth

import (
	"Learnium/internal/database"
	"Learnium/internal/pkg/adapters"
	"Learnium/internal/pkg/common"
	"Learnium/internal/pkg/models"
	"Learnium/internal/utils"
	"context"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

// auth controller processes authentication related requests
type AuthController struct {
	service      IAuthService
	mailProvider adapters.IEmailAdapter
	serializer   common.IValidator
	logger       common.ILogger
}

// create new instance of Auth controller
func NewAuthController(srv IAuthService, mail adapters.IEmailAdapter,
	validator common.IValidator, logger common.ILogger) AuthController {
	// register custom password validation function
	validator.RegisterCustomValidator("strongPassword", ValidateStrongPassword)
	// instantiate auth controller with dependencies
	controller := &AuthController{
		service:      srv,
		mailProvider: mail,
		serializer:   validator,
		logger:       logger,
	}
	return *controller
}

// register new user account with email and password
// auto send otp to user email to verify email ownership
func (cntrl *AuthController) handleRegistration(c *fiber.Ctx) error {

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	var requestBody SignupDto
	// var authResponse auth.AuthenticationResponseDto

	// validate request body
	if err := cntrl.serializer.ValidateRequestBody(c, &requestBody); err != nil {
		return c.Status(422).JSON(common.APIResponse{Message: "Invalid Request Body", Success: false, Detail: err.Error()})
	}

	// save user data to db
	savedUser, err := cntrl.service.RegisterUser(ctx, requestBody.Email, requestBody.Password)
	if err != nil && err == ErrDuplicate {
		return c.Status(409).JSON(common.APIResponse{Message: "User Registration Failed", Success: false, Detail: err.Error()})
	} else if err != nil {
		return c.Status(500).JSON(common.APIResponse{Message: "User Registration Failed", Success: false, Detail: "Internal Server Error"})
	}

	newAccountOtp := cntrl.service.SetRegistrationOTP(ctx, savedUser.Email)
	mailBody := fmt.Sprintf("Your OTP code is: %s", newAccountOtp)
	// // send otp mail to user
	err = cntrl.mailProvider.SendEmail(savedUser.Email, "OTP Code Requested", mailBody)
	if err != nil {
		cntrl.logger.Error(ctx, "error during registration otp mail sending",
			zap.String("additional_info", err.Error()))
	}

	return c.Status(201).JSON(common.APIResponse{
		Message: "User Account Registered and OTP sent to Email Address",
		Success: true})
}

// regenerate registration otp and send to email address
func (cntrl *AuthController) resendRegistrationOTP(c *fiber.Ctx) error {

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	newAccountMail := c.Params("email")
	if newAccountMail == "" || !utils.IsValidEmail(newAccountMail) {
		return c.Status(422).JSON(common.APIResponse{Message: "Failed To Resend Registration OTP", Success: false,
			Detail: "Invalid Email in Path Param"})
	}
	user, err := cntrl.service.GetUserByEmail(ctx, newAccountMail)
	if err != nil {
		return c.Status(424).JSON(common.APIResponse{Message: "Failed To Resend Registration OTP", Success: false,
			Detail: "Email Not Registered"})
	}
	fmt.Println("user in resend otp: ", user)
	if user.EmailVerified {
		return c.Status(409).JSON(common.APIResponse{Message: "Failed To Resend Registration OTP", Success: false,
			Detail: "User Already Registered and Verified"})
	}

	newAccountOtp := cntrl.service.SetRegistrationOTP(ctx, newAccountMail)
	mailBody := fmt.Sprintf("Your OTP code is: %s", newAccountOtp)
	// send otp mail to user
	err = cntrl.mailProvider.SendEmail(newAccountMail, "OTP Code Requested", mailBody)
	if err != nil {
		cntrl.logger.Error(ctx, "error during registration otp mail sending",
			zap.String("additional_info", err.Error()))
		return c.Status(500).JSON(common.APIResponse{Message: "Failed To Resend Registration OTP", Success: false,
			Detail: "Internal server Error"})

	}

	return c.Status(200).JSON(common.APIResponse{
		Message: "OTP sent to Registered Email Address",
		Success: true})
}

// verify otp provided by user for acccount registration
func (cntrl *AuthController) handleRegistrationOTPValidation(c *fiber.Ctx) error {

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	var requestBody ValidateOtpDto
	// var authResponse auth.AuthenticationResponseDto

	// validate request body
	if err := cntrl.serializer.ValidateRequestBody(c, &requestBody); err != nil {
		return c.Status(422).JSON(common.APIResponse{Message: "Invalid Request Body", Success: false, Detail: err.Error()})
	}

	// validate otp provided by user with associated email
	validatedUser, err := cntrl.service.ValidateRegistrationOTP(ctx, requestBody.Email, requestBody.Otp)
	if err == ErrOTP {
		return c.Status(401).JSON(common.APIResponse{Message: "OTP Validation Failed", Success: false, Detail: err.Error()})
	} else if err == ErrRecordNotFound {
		return c.Status(424).JSON(common.APIResponse{Message: "OTP Validation Failed", Success: false, Detail: err.Error()})
	} else if err == database.ErrInternal {
		cntrl.logger.Error(ctx, "Error while communicating with cache", zap.Error(err))
		return c.Status(500).JSON(common.APIResponse{Message: "OTP Validation Failed", Success: false, Detail: "Internal Server Error"})
	} else if err == database.ErrKeyNotFound {
		return c.Status(424).JSON(common.APIResponse{Message: "OTP Validation Failed", Success: false,
			Detail: "Either User Did Not Request For OTP or User Verified Account Already"})
	}

	accessToken, refreshToken, err := cntrl.service.GenerateJwtTokens(ctx, validatedUser.ID)
	if err != nil {
		cntrl.logger.Error(ctx, "Error during OTP Validation",
			zap.String("additional_info", err.Error()))
		return c.Status(500).JSON(common.APIResponse{Message: "OTP Validation Failed", Success: false, Detail: "Internal Server Error"})
	}

	return c.Status(201).JSON(common.APIResponse{
		Message: "OTP Validation Successful",
		Success: true,
		Data: map[string]interface{}{"user": validatedUser.ExportDetails(),
			"access_token": accessToken, "refresh_token": refreshToken},
	})
}

// handle the update of super admin personal info during onboarding
func (cntrl *AuthController) UpdateSuperAdminPersonalInfo(c *fiber.Ctx) error {

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	var requestBody UpdatePersonalInfoDto

	// validate request body
	if err := cntrl.serializer.ValidateRequestBody(c, &requestBody); err != nil {
		return c.Status(422).JSON(common.APIResponse{Message: "Invalid Request Body", Success: false, Detail: err.Error()})
	}

	// get user from context
	user, ok := c.Locals("user").(models.User)

	if !ok {
		// User not found in the context
		return fiber.ErrUnauthorized
	}

	// update user
	user.Address = &requestBody.Address
	user.FirstName = &requestBody.FirstName
	user.LastName = &requestBody.LastName
	user.PhoneNumber = &requestBody.PhoneNumber
	user.IsSuperAdmin = true

	updatedUser, err := cntrl.service.UpdateUser(ctx, user)

	if err != nil {
		cntrl.logger.Error(ctx, "Error during super admin personal info update",
			zap.String("additional_info", err.Error()))
		return c.Status(401).JSON(common.APIResponse{Message: "OTP Validation Failed", Success: false, Detail: err.Error()})
	}

	return c.Status(200).JSON(common.APIResponse{
		Message: "Super Admin Personal Info Update Successful",
		Success: true,
		Data:    updatedUser.ExportDetails(),
	})
}

func (cntrl *AuthController) handleLearniumOSLogin(c *fiber.Ctx) error {

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	var requestBody LoginDto

	// validate request body
	if err := cntrl.serializer.ValidateRequestBody(c, &requestBody); err != nil {
		return c.Status(422).JSON(common.APIResponse{Message: "Invalid Request Body", Success: false, Detail: err.Error()})
	}

	// authenticate user
	authenticatedUser, err := cntrl.service.AuthenticateUser(ctx, requestBody.Email, requestBody.Password)

	if err != nil && (err == ErrPassword || err == ErrRecordNotFound) {
		return c.Status(401).JSON(common.APIResponse{Message: "Login Failed",
			Success: false, Detail: ErrCredentials.Error()})
	} else if err != nil && err == ErrUnverified {
		return c.Status(403).JSON(common.APIResponse{Message: "Login Failed",
			Success: false, Detail: err.Error()})
	} else if err != nil {
		cntrl.logger.Error(ctx, "Error during Login", zap.String("additional_info", err.Error()))
		return c.Status(401).JSON(common.APIResponse{Message: "Login Failed", Success: false, Detail: "Internal Server Error"})
	}

	if authenticatedUser.IsSuperAdmin || authenticatedUser.IsSchoolMember {
		accessToken, refreshToken, err := cntrl.service.GenerateJwtTokens(ctx, authenticatedUser.ID)
		if err != nil {
			cntrl.logger.Error(ctx, "Error during Login", zap.String("additional_info", err.Error()))
			return c.Status(500).JSON(common.APIResponse{Message: "Login Failed", Success: false, Detail: "Internal Server Error"})
		}

		return c.Status(200).JSON(common.APIResponse{
			Message: "Login Successful",
			Success: true,
			Data: map[string]interface{}{"user": authenticatedUser.ExportDetails(),
				"access_token": accessToken, "refresh_token": refreshToken},
		})
	}

	return c.Status(403).JSON(common.APIResponse{Message: "Login Failed",
		Success: false, Detail: "User Cannot Login To Learnium OS"})
}
