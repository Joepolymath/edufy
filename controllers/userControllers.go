package controllers

// import (
// 	"Learnium/adapters"
// 	"Learnium/config"
// 	"Learnium/database"
// 	"Learnium/logger"
// 	"Learnium/models"
// 	"Learnium/serializers"
// 	"context"
// 	"github.com/gofiber/fiber/v2"
// 	"github.com/google/uuid"
// 	"go.uber.org/zap"
// 	"gorm.io/gorm"
// 	"time"
// )

// func UserDetailController(c *fiber.Ctx) error {
// 	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// 	defer cancel()

// 	// Open the database connection
// 	db := database.DBConnection()
// 	defer database.CloseDB(db)

// 	// logged-in user
// 	user := c.Locals("user").(models.User)

// 	var profile models.Profile

// 	logger.Info(ctx, "user id", zap.String("userId", user.ID.String()))

// 	// Check if a profile already exists for the user
// 	err := db.WithContext(ctx).Where("user_id = ?", user.ID).Preload("User").First(&profile).Error

// 	if err != gorm.ErrRecordNotFound {
// 		// A profile already exists for the user
// 		return c.Status(200).JSON(profile)
// 	}

// 	// No profile found, create a new one
// 	profile.UserID = &user.ID
// 	err = db.WithContext(ctx).Preload("User").Create(&profile).Error

// 	if err != nil {
// 		return c.Status(500).JSON(
// 			Response{Message: "Failed to create a profile for the user", Success: false, Detail: err})
// 	}

// 	// Marshal userProfileDetail to JSON and return it as a response
// 	return c.Status(200).JSON(profile)
// }
// func UserAndProfileUpdateInfoController(c *fiber.Ctx) error {
// 	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// 	defer cancel()

// 	// Open the database connection
// 	db := database.DBConnection()
// 	defer database.CloseDB(db)

// 	// logged-in user
// 	user := c.Locals("user").(models.User)
// 	validator := adapters.NewValidate()
// 	fileUpload := adapters.NewFileUpload()

// 	var requestBody serializers.UserAndProfileUpdateRequestSerializer

// 	var profile models.Profile

// 	// Check if a profile already exists for the user
// 	err := db.WithContext(ctx).Where("user_id = ?", user.ID).First(&profile).Error

// 	if err == gorm.ErrRecordNotFound {
// 		// No profile found, create a new one
// 		profile.UserID = &user.ID
// 		if err := db.WithContext(ctx).Create(&profile).Error; err != nil {
// 			return c.Status(500).JSON(
// 				Response{Message: "Failed to create a profile for the user", Success: false, Detail: err})
// 		}
// 	} else if err != nil {
// 		// Handle other errors when querying for the profile

// 		return c.Status(500).JSON(
// 			Response{Message: "Error retrieving user profile", Success: false, Detail: err})
// 	}

// 	if err := c.BodyParser(&requestBody); err != nil {
// 		return c.Status(400).JSON(Response{Message: "invalid request body", Success: false, Detail: err.Error()})

// 	}
// 	// call the validate function in the request serializer for update
// 	if err := validator.ValidateData(&requestBody); err != nil {
// 		return c.Status(400).JSON(
// 			Response{Message: err, Success: false, Detail: err})
// 	}

// 	requestBody.ProfileImage, err = fileUpload.UploadFile("profile_image", c)

// 	// Update the user info
// 	err = db.WithContext(ctx).Model(&user).Where("id = ?", user.ID).Updates(&requestBody).Error

// 	if err != nil {

// 		return c.Status(400).JSON(Response{Message: "There was an error updating the user info", Success: false, Detail: err.Error()})
// 	}

// 	// Update the profile info
// 	err = db.WithContext(ctx).Model(&profile).Where("user_id = ?", user.ID).Preload("User").Updates(&requestBody).First(&profile).Error

// 	if err != nil {

// 		return c.Status(500).JSON(Response{Message: "There was an error updating the user info", Success: false, Detail: err.Error()})

// 	}

// 	return c.Status(200).JSON(profile)
// }
// func HealthInfoUpdateController(c *fiber.Ctx) error {
// 	/*
// 		This is used to update the health info of the user
// 	*/
// 	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// 	defer cancel()

// 	// Open the database connection
// 	db := database.DBConnection()
// 	defer database.CloseDB(db)

// 	// logged-in user
// 	user := c.Locals("user").(models.User)
// 	validator := adapters.NewValidate()

// 	var health models.Health

// 	// Check if a health record already exists for the user
// 	err := db.WithContext(ctx).Where("user_id = ?", user.ID).First(&health).Error

// 	if err == gorm.ErrRecordNotFound {
// 		// No health record found, create a new one
// 		health.UserID = &user.ID
// 		if err := db.WithContext(ctx).Create(&health).Error; err != nil {
// 			return c.Status(400).JSON(Response{Message: "Failed to create a health record for the user", Success: false, Detail: err.Error()})
// 		}
// 	} else if err != nil {
// 		// Handle other errors when querying for the health record
// 		return c.Status(400).JSON(Response{Message: "Error retrieving user health record", Success: false, Detail: err.Error()})
// 	}

// 	// Parse the request body
// 	var requestBody serializers.HealthUpdateRequestSerializer
// 	if err := c.BodyParser(&requestBody); err != nil {
// 		return c.Status(400).JSON(Response{Message: "invalid request body", Success: false, Detail: err.Error()})
// 	}

// 	// Call the validate function in the request serializer for the update
// 	vErr := validator.ValidateData(&requestBody)
// 	if vErr != nil {
// 		return c.Status(400).JSON(Response{Message: vErr, Success: false, Detail: vErr})
// 	}

// 	// Update the health record
// 	err2 := db.WithContext(ctx).Model(&health).Where("user_id = ?", user.ID).Updates(&requestBody).Preload("User").First(&health).Error

// 	if err2 != nil {
// 		return c.Status(400).JSON(Response{Message: "Error Updating Health info", Success: false, Detail: err.Error()})

// 	}

// 	// Return the updated health record
// 	return c.Status(200).JSON(health)
// }
// func HealthInfoDetailController(c *fiber.Ctx) error {
// 	/* this returns the detail of the health of this logged-in user*/

// 	var health models.Health

// 	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// 	defer cancel()

// 	// Open the database connection
// 	db := database.DBConnection()
// 	defer database.CloseDB(db)

// 	// logged-in user
// 	user := c.Locals("user").(models.User)

// 	health.UserID = &user.ID

// 	// Check if a health record already exists for the user
// 	err := db.WithContext(ctx).Where("user_id = ?", user.ID).Preload("User").First(&health).Error

// 	if err == gorm.ErrRecordNotFound {
// 		// No health record found
// 		err := db.WithContext(ctx).Model(&health).Where("user_id = ?", user.ID).Preload("User").Create(&health).First(&health).Error
// 		if err != nil {
// 			return c.Status(404).JSON(
// 				Response{Message: "Health info not found for the user", Success: false, Detail: err})
// 		}

// 	} else if err != nil {
// 		// Handle other errors when querying for the health record
// 		return c.Status(500).JSON(
// 			Response{Message: "Error retrieving user health info", Success: false, Detail: err})
// 	}

// 	return c.Status(200).JSON(health)
// }

// func UserInfoUpdateController(c *fiber.Ctx) error {
// 	var health models.Health
// 	var profile models.Profile
// 	var user models.User
// 	var school models.School
// 	var staff models.Staff
// 	var student models.Student
// 	var requestBody serializers.UserInfoUpdateRequestSerializer

// 	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// 	defer cancel()

// 	// Open the database connection
// 	db := database.DBConnection()
// 	defer database.CloseDB(db)

// 	validator := adapters.NewValidate()

// 	loggedInUser := c.Locals("user").(models.User)

// 	schoolCode := c.Query("school_code")

// 	school, err := school.RetrieveBySchoolCode(ctx, db, schoolCode)
// 	if err != nil {
// 		return c.Status(404).JSON(Response{
// 			Message: "school with this id does not exists",
// 			Success: false,
// 			Detail:  err,
// 		})
// 	}

// 	userID, err := uuid.Parse(c.Params("id"))
// 	if err != nil {
// 		return c.Status(400).JSON(Response{Message: "Invalid user id passed", Success: false, Detail: err.Error()})

// 	}

// 	// get the user
// 	err = db.WithContext(ctx).Model(&user).Where("id = ?", userID).First(&user).Error
// 	if err != nil {
// 		if err == gorm.ErrRecordNotFound {
// 			return c.Status(400).JSON(Response{Message: "user with this id does not exists", Success: false, Detail: err.Error()})

// 		}
// 		return c.Status(400).JSON(Response{Message: "an error occured getting user by id", Success: false, Detail: err.Error()})

// 	}

// 	if loggedInUser.ID != userID {
// 		// check if the user is an admin in the school
// 		if school.IsSchoolAdminOrOwner(ctx, db, schoolCode, loggedInUser.ID) == false {
// 			return c.Status(400).JSON(Response{Message: "You are not authorized to perform this action", Success: false, Detail: nil})
// 		}
// 	}

// 	// Parse the request body
// 	if err := c.BodyParser(&requestBody); err != nil {
// 		return c.Status(400).JSON(Response{Message: "invalid request body", Success: false, Detail: err.Error()})
// 	}

// 	vErr := validator.ValidateData(&requestBody)
// 	if vErr != nil {
// 		return c.Status(400).JSON(Response{Message: vErr, Success: false, Detail: vErr})

// 	}

// 	// since the email has to be unique we have to check for it
// 	if requestBody.Email != nil {
// 		err = db.WithContext(ctx).Model(&user).Where("email ILIKE ?", *requestBody.Email).First(&user).Error
// 		if err != nil {
// 			if err != gorm.ErrRecordNotFound {
// 				logger.Error(ctx, "error getting user by email", zap.Error(err))
// 				return c.Status(400).JSON(Response{Message: "an error occured getting user by email", Success: false, Detail: err.Error()})

// 			}
// 		}
// 	}

// 	err = db.WithContext(ctx).Model(&profile).Where("user_id = ?", userID).Updates(&requestBody).First(&profile).Error
// 	if err != nil {
// 		return c.Status(400).JSON(Response{Message: "There was an error updating the profile info", Success: false, Detail: err.Error()})

// 	}

// 	err = db.WithContext(ctx).Model(&user).Where("id = ?", userID).Updates(&requestBody).First(&user).Error
// 	if err != nil {
// 		return c.Status(400).JSON(Response{Message: "There was an error updating the user info", Success: false, Detail: err.Error()})

// 	}

// 	err = db.WithContext(ctx).Model(&health).Where("user_id = ?", userID).Updates(&requestBody).First(&health).Error
// 	if err != nil {
// 		return c.Status(400).JSON(Response{Message: "There was an error updating the health info", Success: false, Detail: err.Error()})

// 	}

// 	if requestBody.StaffCode != nil {
// 		staff, err := staff.RetrieveByUserAndSchool(ctx, db, school.ID, userID)
// 		if err != nil {
// 			if err == gorm.ErrRecordNotFound {
// 				return c.Status(400).JSON(Response{Message: "user does not exist as a staff", Success: false, Detail: err.Error()})

// 			}
// 			return c.Status(400).JSON(Response{Message: "There was an error updating the staff info", Success: false, Detail: err.Error()})

// 		}
// 		err = db.WithContext(ctx).Model(&staff).Where("id = ?", staff.ID).Updates(&requestBody).First(&staff).Error
// 		if err != nil {
// 			return c.Status(400).JSON(Response{Message: "There was an error updating the staff info", Success: false, Detail: err.Error()})

// 		}
// 	}

// 	if requestBody.StudentCode != nil {
// 		student, err := student.RetrieveStudentByUserIDAndSchool(db, ctx, userID, school.ID)
// 		if err != nil {
// 			if err == gorm.ErrRecordNotFound {
// 				return c.Status(400).JSON(Response{Message: "User does not exist as a student", Success: false, Detail: err.Error()})

// 			}
// 			return c.Status(400).JSON(Response{Message: "There was an error updating the student info", Success: false, Detail: err.Error()})

// 		}
// 		err = db.WithContext(ctx).Model(&student).Where("id = ?", student.ID).Updates(&requestBody).First(&student).Error
// 		if err != nil {
// 			return c.Status(400).JSON(Response{Message: "There was an error updating the student info", Success: false, Detail: err.Error()})

// 		}
// 	}

// 	return c.Status(200).JSON(
// 		Response{Message: "Successfully update info", Success: true, Detail: nil})
// }

// func LearniumUpdateUserPermission(c *fiber.Ctx) error {
// 	var requestBody serializers.UserPermissionUpdateRequestSerializer

// 	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// 	defer cancel()

// 	// Open the database connection
// 	db := database.DBConnection()
// 	defer database.CloseDB(db)

// 	validateAdapters := adapters.NewValidate()

// 	loggedInUser := c.Locals("user").(models.User)

// 	cfg, err := config.Load()
// 	if err != nil {
// 		return c.Status(400).JSON(Response{Message: "Error Loading configuration", Success: false, Detail: err.Error()})
// 	}

// 	err = c.BodyParser(&requestBody)
// 	if err != nil {
// 		return c.Status(400).JSON(Response{Message: "Invalid request body", Success: false, Detail: err.Error()})
// 	}

// 	if *loggedInUser.Email != cfg.LearniumSuperUserMail {
// 		return c.Status(400).JSON(Response{Message: "You dont have the required permission", Success: false})
// 	}

// 	errors := validateAdapters.ValidateData(&requestBody)
// 	if errors != nil {
// 		return c.Status(400).JSON(Response{Message: errors, Success: false, Detail: errors})
// 	}

// 	err = db.WithContext(ctx).Model(&models.User{}).Where("email ILIKE ?", requestBody.Email).Updates(&requestBody).Error
// 	if err != nil {
// 		return c.Status(500).JSON(Response{Message: "error from our end updating the user permission", Success: false, Detail: err.Error()})
// 	}

// 	return c.Status(200).JSON(Response{Message: "Successfully update the user permission ", Success: true})
// }
