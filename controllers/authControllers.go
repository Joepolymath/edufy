package controllers

// import (
// 	"Learnium/adapters"
// 	"Learnium/database"
// 	"Learnium/logger"
// 	"Learnium/models"
// 	"Learnium/serializers"
// 	"Learnium/utils"
// 	"context"
// 	"github.com/gofiber/fiber/v2"
// 	"github.com/jinzhu/copier"
// 	"go.uber.org/zap"
// 	"gorm.io/gorm"
// 	"time"
// )

// func SignupController(c *fiber.Ctx) error {
// 	var user models.User
// 	var foundUser models.User
// 	var requestBody serializers.SignupRequestSerializer
// 	var authResponse serializers.AuthenticationResponseSerializer

// 	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// 	defer cancel()

// 	db := database.DBConnection()
// 	defer database.CloseDB(db)

// 	validator := adapters.NewValidate()
// 	userUtils := utils.NewUserUtils()

// 	// Parse the request body
// 	if err := c.BodyParser(&requestBody); err != nil {
// 		return c.Status(400).JSON(Response{Message: "invalid request body", Success: false, Detail: err.Error()})
// 	}
// 	// Validate the signup request struct
// 	if err2 := validator.ValidateData(&requestBody); err2 != nil {
// 		return c.Status(400).JSON(Response{Message: err2, Success: false, Detail: err2})

// 	}

// 	// check if the user Email  exists
// 	err := db.WithContext(ctx).Where("email ILIKE ?", requestBody.Email).First(&foundUser).Error
// 	if err != nil {
// 		if err != gorm.ErrRecordNotFound {
// 			return c.Status(400).JSON(Response{Message: "An error occured check user email", Success: false, Detail: err.Error()})
// 		}
// 	}

// 	// check if the Email already exists
// 	if foundUser.Email != nil {
// 		return c.Status(400).JSON(Response{Message: "Email already exists", Success: false, Detail: nil})

// 	}

// 	err = copier.Copy(&user, &requestBody)
// 	if err != nil {
// 		logger.Error(ctx, "Error marshaling data", zap.Error(err))
// 		return c.Status(500).JSON(Response{Message: "error copying data", Success: false, Detail: err.Error()})

// 	}

// 	user, err = userUtils.CreateUser(db, ctx, user, false)
// 	if err != nil {
// 		return c.Status(400).JSON(Response{Message: err.Error(), Success: false, Detail: err.Error()})
// 	}

// 	if err := authResponse.InitializeData(ctx, &user); err != nil {
// 		return c.Status(500).JSON(Response{Message: err.Error(), Success: false, Detail: err.Error()})

// 	}
// 	return c.Status(201).JSON(authResponse)
// }

// /*LoginController // This is used to log in a user*/
// func LoginController(c *fiber.Ctx) error {
// 	var foundUser models.User
// 	var user models.User
// 	var authResponse serializers.AuthenticationResponseSerializer

// 	// the timeout which is used for accessing this route it must not be more than 100 seconds
// 	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// 	defer cancel()

// 	// Open the database connection
// 	db := database.DBConnection()
// 	defer database.CloseDB(db)

// 	validateAdapters := adapters.NewValidate()

// 	// Parse the request body
// 	var requestBody serializers.LoginRequestSerializer
// 	if err := c.BodyParser(&requestBody); err != nil {
// 		return c.Status(400).JSON(Response{Message: "invalid request body", Success: false, Detail: err.Error()})
// 	}
// 	// call the validate function in the request serializer for login
// 	if vErr := validateAdapters.ValidateData(&requestBody); vErr != nil {
// 		return c.Status(400).JSON(Response{Message: vErr, Success: false, Detail: vErr})
// 	}

// 	// use the Email to find the user
// 	err := db.WithContext(ctx).Model(&foundUser).Where("email ILIKE ?", requestBody.Email).First(&foundUser).Error
// 	if err != nil {
// 		return c.Status(400).JSON(Response{Message: "Login credentials not correct", Success: false, Detail: err})

// 	}

// 	// Let's check the user password if It's the same as the password passed
// 	check, err := utils.VerifyPassword(*foundUser.Password, requestBody.Password)
// 	if err != nil && check == false {
// 		return c.Status(400).JSON(Response{Message: "Password might be invalid or Email", Success: false, Detail: err.Error()})
// 	}

// 	err = db.WithContext(ctx).Model(&user).
// 		Where(&models.User{Email: foundUser.Email}).First(&foundUser).Error
// 	if err != nil {
// 		return c.Status(400).JSON(Response{Message: err.Error(), Success: false, Detail: err.Error()})

// 	}

// 	// update last login
// 	err = db.WithContext(ctx).Model(&user).Where("email ILIKE ?", requestBody.Email).Update("last_login", time.Now()).Error
// 	if err != nil {
// 		return c.Status(500).JSON(Response{Message: "Error updating last login", Success: false, Detail: err.Error()})

// 	}

// 	// return the error if the data was not initialized
// 	if err := authResponse.InitializeData(ctx, &foundUser); err != nil {
// 		return c.Status(500).JSON(Response{Message: err.Error(), Success: false, Detail: err.Error()})

// 	}

// 	return c.Status(200).JSON(authResponse)
// }

// func RequestOTPController(c *fiber.Ctx) error {

// 	var requestBody serializers.OtpRequestSerializer
// 	var foundUser models.User

// 	// the timeout which is used for accessing this route it must not be more than 100 seconds
// 	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// 	defer cancel()

// 	// Open the database connection
// 	db := database.DBConnection()
// 	defer database.CloseDB(db)

// 	validateAdapters := adapters.NewValidate()

// 	if err := c.BodyParser(&requestBody); err != nil {
// 		return c.Status(400).JSON(Response{Message: "invalid request body", Success: false, Detail: err.Error()})
// 	}

// 	// call the validate function in the request serializer for requesting otp
// 	if err := validateAdapters.ValidateData(&requestBody); err != nil {

// 		return c.Status(400).JSON(Response{Message: err, Success: false, Detail: err})

// 	}

// 	// find the user we are requesting the otp for and also check if an error exists raise it also
// 	err := db.WithContext(ctx).Model(foundUser).Where("email ILIKE ?", requestBody.Email).First(&foundUser).Error
// 	if err != nil {
// 		return c.Status(400).JSON(Response{Message: "User does not exists", Success: false, Detail: err.Error()})

// 	}

// 	err = utils.SendOTP(*foundUser.Email)
// 	if err != nil {
// 		return c.Status(400).JSON(Response{Message: "Error sending otp", Success: false, Detail: err.Error()})

// 	}

// 	return c.Status(200).JSON(Response{Message: "Successfully sent otp to your Email", Success: true, Detail: nil})

// }

// // VerifyUserController /* This is used to verify the user Email if not verified  */
// func VerifyUserController(c *fiber.Ctx) error {
// 	var foundUser models.User
// 	var requestBody serializers.OtpValidateRequestSerializer

// 	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// 	defer cancel()

// 	// Open the database connection
// 	db := database.DBConnection()
// 	defer database.CloseDB(db)

// 	// Parse the request body
// 	if err := c.BodyParser(&requestBody); err != nil {
// 		return c.Status(400).JSON(Response{Message: "invalid request body", Success: false, Detail: err.Error()})
// 	}

// 	// find the user we are requesting the otp for
// 	err := db.WithContext(ctx).Model(foundUser).Where("email ILIKE ?", requestBody.Email).First(&foundUser).Error
// 	if err != nil {
// 		return c.Status(400).JSON(Response{Message: "Email does not exist or an error occurred ", Success: false, Detail: err.Error()})

// 	}

// 	err = utils.ValidateOTP(*foundUser.Email, requestBody.Otp)
// 	if err != nil {
// 		return c.Status(400).JSON(Response{Message: err.Error(), Success: false, Detail: err.Error()})

// 	}

// 	if foundUser.IsVerified != nil {
// 		*foundUser.IsVerified = true
// 	}

// 	err = db.WithContext(ctx).Model(foundUser).Where("email ILIKE ?", requestBody.Email).
// 		Updates(models.User{IsVerified: foundUser.IsVerified}).Error
// 	if err != nil {
// 		return c.Status(400).JSON(Response{Message: "error updating the user", Success: false, Detail: err.Error()})
// 	}

// 	return c.Status(200).JSON(Response{Message: "Successfully verified the user", Success: true, Detail: nil})

// }

// func ForgotPasswordController(c *fiber.Ctx) error {
// 	var foundUser models.User
// 	var requestBody serializers.ForgotPasswordValidateRequestSerializer

// 	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// 	defer cancel()

// 	// Open the database connection
// 	db := database.DBConnection()
// 	defer database.CloseDB(db)

// 	// Parse the request body
// 	if err := c.BodyParser(&requestBody); err != nil {
// 		return c.Status(400).JSON(Response{Message: "invalid request body", Success: false, Detail: err.Error()})
// 	}

// 	// find the user we are requesting the otp for
// 	err := db.WithContext(ctx).Model(foundUser).Where("email ILIKE ?", requestBody.Email).First(&foundUser).Error
// 	if err != nil {
// 		return c.Status(400).JSON(Response{Message: "Email does not exist or an error occurred ", Success: false, Detail: err.Error()})

// 	}

// 	//validate the otp
// 	err = utils.ValidateOTP(*foundUser.Email, requestBody.Otp)
// 	if err != nil {
// 		return c.Status(400).JSON(Response{Message: err.Error(), Success: false, Detail: err.Error()})

// 	}

// 	// hash the password
// 	password := utils.HashPassword(requestBody.Password)
// 	foundUser.Password = &password
// 	err = db.WithContext(ctx).Model(foundUser).Where("email ILIKE ?", requestBody.Email).
// 		Updates(models.User{Password: foundUser.Password}).Error

// 	// if there is an error then
// 	if err != nil {
// 		return c.Status(400).JSON(Response{Message: "Error updating user password", Success: false, Detail: err.Error()})

// 	}

// 	return c.Status(200).JSON(Response{Message: "Successfully Update user password", Success: true, Detail: nil})

// }
