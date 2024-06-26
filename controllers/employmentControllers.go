package controllers

// import (
// 	"Learnium/adapters"
// 	"Learnium/database"
// 	"Learnium/logger"
// 	"Learnium/models"
// 	"Learnium/serializers"
// 	"Learnium/utils"
// 	"context"
// 	"encoding/json"
// 	"fmt"
// 	"github.com/gofiber/fiber/v2"
// 	"github.com/google/uuid"
// 	"github.com/jinzhu/copier"
// 	"go.uber.org/zap"
// 	"gorm.io/gorm"
// 	"time"
// )

// // EmploymentConfigurationCreateController /* Creating employment the configuration have to be set for the job first before creating the job itself*/
// func EmploymentConfigurationCreateController(c *fiber.Ctx) error {
// 	/* this is used to create employment configuration*/
// 	schoolCode := c.Query("school_code")

// 	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// 	defer cancel()

// 	// get the logged-in user
// 	user := c.Locals("user").(models.User)

// 	// Open the database connection
// 	db := database.DBConnection()
// 	defer database.CloseDB(db)

// 	validator := adapters.NewValidate()

// 	// get the school
// 	var school models.School
// 	school, err := school.RetrieveBySchoolCode(ctx, db, schoolCode)
// 	if err != nil {
// 		return c.Status(404).JSON(Response{
// 			Message: "school with this id does not exists",
// 			Success: false,
// 			Detail:  err,
// 		})
// 	}

// 	isOk := school.IsSchoolAdminOrOwner(ctx, db, schoolCode, user.ID)
// 	if !isOk {
// 		return c.Status(400).JSON(Response{Message: "you dont have permission", Success: false, Detail: err})
// 	}

// 	// Parse the request body
// 	var requestBody models.EmploymentConfiguration
// 	if err := c.BodyParser(&requestBody); err != nil {

// 		return c.Status(400).JSON(Response{Message: "Invalid request body", Success: false, Detail: err})

// 	}

// 	// call the validate function in the request serializer for login
// 	if err := validator.ValidateData(requestBody); err != nil {
// 		return c.Status(400).JSON(Response{Message: err, Success: false, Detail: err})

// 	}

// 	// also validate the custom field data
// 	if requestBody.EmploymentCustomFields != nil {
// 		for _, customField := range *requestBody.EmploymentCustomFields {
// 			if err := validator.ValidateData(customField); err != nil {
// 				return c.Status(400).JSON(Response{Message: err, Success: false, Detail: err})

// 			}
// 		}
// 	}

// 	// also validate the additional questions
// 	if requestBody.AdditionalQuestions != nil {
// 		for _, additionalQuestion := range *requestBody.AdditionalQuestions {
// 			if err := validator.ValidateData(additionalQuestion); err != nil {
// 				return c.Status(400).JSON(Response{Message: err, Success: false, Detail: err})

// 			}
// 		}
// 	}

// 	employmentConfiguration := requestBody
// 	employmentConfiguration.SchoolID = &school.ID

// 	err = db.WithContext(ctx).Model(&employmentConfiguration).Create(&employmentConfiguration).Error
// 	if err != nil {
// 		return c.Status(500).JSON(Response{Message: "Error creating employment configuration", Success: false, Detail: err.Error()})

// 	}
// 	return c.Status(201).JSON(employmentConfiguration)
// }

// func EmploymentConfigurationListController(c *fiber.Ctx) error {
// 	/* this is used to create employment configuration*/
// 	var employmentConfigurations []models.EmploymentConfiguration
// 	var total int64
// 	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// 	defer cancel()

// 	// Open the database connection
// 	db := database.DBConnection()
// 	defer database.CloseDB(db)

// 	schoolCode := c.Query("school_code")

// 	page := c.QueryInt("page", 1)    // default to page 1 if not provided
// 	limit := c.QueryInt("limit", 10) // default to 10 items per page if not provided

// 	// get the logged-in user
// 	user := c.Locals("user").(models.User)
// 	// get the school
// 	var school models.School
// 	school, err := school.RetrieveBySchoolCode(ctx, db, schoolCode)
// 	if err != nil {
// 		return c.Status(404).JSON(Response{
// 			Message: "school with this id does not exists",
// 			Success: false,
// 			Detail:  err,
// 		})
// 	}

// 	isOk := school.IsSchoolAdminOrOwner(ctx, db, schoolCode, user.ID)
// 	if !isOk {
// 		return c.Status(400).JSON(Response{Message: "you dont have permission", Success: false, Detail: err})
// 	}

// 	// make the employment configuration list
// 	db.WithContext(ctx).Model(&models.EmploymentConfiguration{}).Where("school_id = ?", school.ID).
// 		Count(&total)

// 	err = db.Scopes(adapters.NewPaginateAdapter(limit, page).PaginatedResult).WithContext(ctx).Model(&models.EmploymentConfiguration{}).Where("school_id = ?", school.ID).
// 		Order("timestamp desc").
// 		Find(&employmentConfigurations).Error
// 	if err != nil {
// 		if err != gorm.ErrRecordNotFound {
// 			return c.Status(500).JSON(Response{Message: "Error fetching employment configurations", Success: false, Detail: err.Error()})

// 		}
// 	}
// 	return c.Status(200).JSON(fiber.Map{"total": total, "data": employmentConfigurations})
// }

// func EmploymentConfigurationDetailController(c *fiber.Ctx) error {
// 	schoolCode := c.Query("school_code")
// 	employmentConfigID := c.Params("id")

// 	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// 	defer cancel()

// 	// Open the database connection
// 	db := database.DBConnection()
// 	defer database.CloseDB(db)

// 	// get the logged-in user
// 	user := c.Locals("user").(models.User)
// 	// get the school
// 	var school models.School
// 	school, err := school.RetrieveBySchoolCode(ctx, db, schoolCode)
// 	if err != nil {
// 		return c.Status(404).JSON(Response{
// 			Message: "school with this id does not exists",
// 			Success: false,
// 			Detail:  err,
// 		})
// 	}

// 	isOk := school.IsSchoolAdminOrOwner(ctx, db, schoolCode, user.ID)
// 	if !isOk {
// 		return c.Status(400).JSON(Response{Message: "you dont have permission", Success: false, Detail: err})
// 	}

// 	var employmentConfiguration models.EmploymentConfiguration
// 	err = db.WithContext(ctx).Model(&employmentConfiguration).Where("id = ?", employmentConfigID).First(&employmentConfiguration).Error
// 	if err != nil {
// 		if err == gorm.ErrRecordNotFound {
// 			return c.Status(404).JSON(Response{Message: "Error fetching employment configurations", Success: false, Detail: err.Error()})
// 		} else {
// 			return c.Status(500).JSON(Response{Message: "Error fetching employment configurations", Success: false, Detail: err.Error()})
// 		}
// 	}
// 	return c.Status(200).JSON(employmentConfiguration)
// }

// func EmploymentConfigurationUpdateController(c *fiber.Ctx) error {
// 	/* this is used to create employment configuration*/
// 	schoolCode := c.Query("school_code")
// 	employmentConfigID := c.Params("id")

// 	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// 	defer cancel()

// 	// Open the database connection
// 	db := database.DBConnection()
// 	defer database.CloseDB(db)

// 	// get the logged-in user
// 	user := c.Locals("user").(models.User)
// 	validator := adapters.NewValidate()

// 	// get the school
// 	var school models.School
// 	school, err := school.RetrieveBySchoolCode(ctx, db, schoolCode)
// 	if err != nil {
// 		return c.Status(404).JSON(Response{
// 			Message: "school with this id does not exists",
// 			Success: false,
// 			Detail:  err,
// 		})
// 	}

// 	isOk := school.IsSchoolAdminOrOwner(ctx, db, schoolCode, user.ID)
// 	if !isOk {
// 		return c.Status(400).JSON(Response{Message: "you dont have permission", Success: false, Detail: err})
// 	}

// 	// Parse the request body
// 	var requestBody models.EmploymentConfiguration
// 	if err := c.BodyParser(&requestBody); err != nil {

// 		return c.Status(400).JSON(Response{Message: "Invalid request body", Success: false, Detail: err.Error()})

// 	}
// 	// call the validate function in the request serializer for login
// 	if err := validator.ValidateData(requestBody); err != nil {
// 		return c.Status(400).JSON(Response{Message: err, Success: false, Detail: err})

// 	}
// 	employmentConfiguration := requestBody

// 	err = db.WithContext(ctx).Model(&employmentConfiguration).Where("id = ?", employmentConfigID).Updates(&employmentConfiguration).Error
// 	if err != nil {
// 		return c.Status(500).JSON(Response{Message: "Error creating employment configuration", Success: false, Detail: err.Error()})

// 	}
// 	return c.Status(200).JSON(employmentConfiguration)
// }

// func EmploymentConfigurationDeleteController(c *fiber.Ctx) error {
// 	schoolCode := c.Query("school_code")
// 	employmentConfigID := c.Params("id")

// 	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// 	defer cancel()

// 	// Open the database connection
// 	db := database.DBConnection()
// 	defer database.CloseDB(db)

// 	// get the logged-in user
// 	user := c.Locals("user").(models.User)
// 	// get the school
// 	var school models.School
// 	school, err := school.RetrieveBySchoolCode(ctx, db, schoolCode)
// 	if err != nil {
// 		return c.Status(404).JSON(Response{
// 			Message: "school with this id does not exists",
// 			Success: false,
// 			Detail:  err,
// 		})
// 	}

// 	isOk := school.IsSchoolAdminOrOwner(ctx, db, schoolCode, user.ID)
// 	if !isOk {
// 		return c.Status(400).JSON(Response{Message: "you dont have permission", Success: false, Detail: err})
// 	}

// 	var employmentConfiguration models.EmploymentConfiguration
// 	err = db.WithContext(ctx).Model(&employmentConfiguration).Where("id = ?", employmentConfigID).Delete(&employmentConfiguration).Error
// 	if err != nil {
// 		if err == gorm.ErrRecordNotFound {
// 			return c.Status(404).JSON(Response{Message: "employment configuration with this id does not exists", Success: false, Detail: err.Error()})

// 		}
// 		return c.Status(500).JSON(Response{Message: "Error deleting employment configuration", Success: false, Detail: err.Error()})

// 	}

// 	return c.Status(204).JSON(Response{Message: "employment configuration deleted successfully", Success: true, Detail: nil})
// }

// func EmploymentCreateController(c *fiber.Ctx) error {
// 	/* this is used to create employment configuration*/
// 	var requestBody models.Employment
// 	var employmentConfiguration models.EmploymentConfiguration
// 	var role models.Role
// 	var category models.Category

// 	schoolCode := c.Query("school_code")

// 	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// 	defer cancel()
// 	// Open the database connection
// 	db := database.DBConnection()
// 	defer database.CloseDB(db)

// 	// get the logged-in user
// 	user := c.Locals("user").(models.User)
// 	validator := adapters.NewValidate()

// 	// get the school
// 	var school models.School
// 	school, err := school.RetrieveBySchoolCode(ctx, db, schoolCode)
// 	if err != nil {
// 		return c.Status(404).JSON(Response{
// 			Message: "school with this id does not exists",
// 			Success: false,
// 			Detail:  err,
// 		})
// 	}

// 	isOk := school.IsSchoolAdminOrOwner(ctx, db, schoolCode, user.ID)
// 	if !isOk {
// 		return c.Status(400).JSON(Response{Message: "you dont have permission", Success: false, Detail: err})
// 	}

// 	// Parse the request body
// 	if err := c.BodyParser(&requestBody); err != nil {

// 		return c.Status(400).JSON(Response{Message: "Invalid request body", Success: false, Detail: err.Error()})
// 	}
// 	// call the validate function in the request serializer for login
// 	if err := validator.ValidateData(requestBody); err != nil {
// 		return c.Status(400).JSON(Response{Message: err, Success: false, Detail: err})

// 	}

// 	// check if the employment configuration id exists on the database
// 	err = db.WithContext(ctx).Model(&employmentConfiguration).Where("id = ?", requestBody.EmploymentConfigurationID).First(&employmentConfiguration).Error
// 	if err != nil {
// 		if err == gorm.ErrRecordNotFound {
// 			return c.Status(500).JSON(Response{Message: "employment configuration with this id does not exists", Success: false, Detail: err.Error()})

// 		} else {
// 			return c.Status(500).JSON(Response{Message: "an error occured filtering for the employment configuration", Success: false, Detail: err.Error()})

// 		}
// 	}

// 	err = db.WithContext(ctx).Model(&role).Where("id = ?", requestBody.RoleID).First(&role).Error
// 	if err != nil {
// 		if err == gorm.ErrRecordNotFound {
// 			return c.Status(404).JSON(Response{Message: "role with this id does not exists", Success: false, Detail: err.Error()})
// 		}
// 		logger.Error(ctx, "Error  filtering for the role on creating job", zap.Error(err))
// 		return c.Status(500).JSON(Response{Message: "an error occured filtering for the role ", Success: false, Detail: err.Error()})

// 	}

// 	err = db.WithContext(ctx).Model(&category).Where("id = ?", requestBody.CategoryID).First(&category).Error
// 	if err != nil {
// 		if err == gorm.ErrRecordNotFound {
// 			return c.Status(400).JSON(Response{Message: "category with this id does not exists", Success: false, Detail: err.Error()})

// 		}
// 		logger.Error(ctx, "Error  filtering for the category on creating job", zap.Error(err))
// 		return c.Status(500).JSON(Response{Message: "an error occured filtering for the category ", Success: false, Detail: err.Error()})
// 	}

// 	requestBody.SchoolID = &school.ID
// 	// save the job info the database
// 	err = db.WithContext(ctx).Model(&requestBody).Create(&requestBody).Error
// 	if err != nil {
// 		return c.Status(500).JSON(Response{Message: "an error occured creating the employment", Success: false, Detail: err.Error()})

// 	}
// 	// add the school
// 	requestBody.School = &school
// 	return c.Status(201).JSON(requestBody)
// }

// func EmploymentUpdateController(c *fiber.Ctx) error {
// 	/* the is used to update employment */
// 	schoolCode := c.Query("school_code")
// 	employmentID := c.Params("id")

// 	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// 	defer cancel()
// 	// Open the database connection
// 	db := database.DBConnection()
// 	defer database.CloseDB(db)

// 	// get the logged-in user
// 	user := c.Locals("user").(models.User)
// 	validator := adapters.NewValidate()

// 	var employment models.Employment

// 	// get the school
// 	var school models.School
// 	school, err := school.RetrieveBySchoolCode(ctx, db, schoolCode)
// 	if err != nil {
// 		return c.Status(404).JSON(Response{
// 			Message: "school with this id does not exists",
// 			Success: false,
// 			Detail:  err,
// 		})
// 	}

// 	isOk := school.IsSchoolAdminOrOwner(ctx, db, schoolCode, user.ID)
// 	if !isOk {
// 		return c.Status(400).JSON(Response{Message: "you dont have permission", Success: false, Detail: nil})

// 	}

// 	// making a partial update
// 	err = db.WithContext(ctx).Model(&employment).
// 		Where("id = ?", employmentID).
// 		Where("school_id = ?", school.ID).First(&employment).Error

// 	if err != nil {
// 		if err == gorm.ErrRecordNotFound {
// 			return c.Status(400).JSON(Response{Message: "employment with this id does not exists", Success: false, Detail: err.Error()})

// 		}
// 		return c.Status(500).JSON(Response{Message: "an error occured filtering for the employment", Success: false, Detail: err.Error()})

// 	}

// 	// Parse the request body
// 	if err := c.BodyParser(&employment); err != nil {
// 		return c.Status(400).JSON(Response{Message: "invalid request body", Success: false, Detail: err.Error()})
// 	}

// 	// call the validate function in the request serializer for login
// 	if err := validator.ValidateData(employment); err != nil {
// 		return c.Status(400).JSON(Response{Message: err, Success: false, Detail: err})

// 	}

// 	employment.School = &school

// 	// get the employment
// 	employmentConfig := models.EmploymentConfiguration{}
// 	employmentConfig, err = employmentConfig.Retrieve(ctx, db, *employment.EmploymentConfigurationID)
// 	if err != nil {
// 		return c.Status(400).JSON(Response{Message: "an error occured filtering for the employmentConfig", Success: false, Detail: err.Error()})

// 	}
// 	employment.EmploymentConfiguration = &employmentConfig

// 	return c.Status(200).JSON(employment)
// }

// func EmploymentDetailController(c *fiber.Ctx) error {
// 	/* this is used to create employment configuration*/
// 	employmentID := c.Params("id")

// 	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// 	defer cancel()
// 	// Open the database connection
// 	db := database.DBConnection()
// 	defer database.CloseDB(db)

// 	var employment models.Employment
// 	err := db.WithContext(ctx).Model(&employment).
// 		Where("id = ?", employmentID).First(&employment).Error
// 	if err != nil {
// 		if err == gorm.ErrRecordNotFound {
// 			return c.Status(400).JSON(Response{Message: "employment with this id does not exists", Success: false, Detail: err.Error()})

// 		}
// 		return c.Status(400).JSON(Response{Message: "an error occured filtering for the employment", Success: false, Detail: err.Error()})

// 	}

// 	// get the school
// 	school := models.School{}
// 	school, err = school.Retrieve(ctx, db, *employment.SchoolID)
// 	if err != nil {
// 		return c.Status(400).JSON(Response{Message: "an error occured filtering for the school", Success: false, Detail: err.Error()})

// 	}
// 	employment.School = &school

// 	// get the employment
// 	employmentConfig := models.EmploymentConfiguration{}
// 	employmentConfig, err = employmentConfig.Retrieve(ctx, db, *employment.EmploymentConfigurationID)
// 	if err != nil {
// 		return c.Status(400).JSON(Response{Message: "an error occured filtering for the employmentConfig", Success: false, Detail: err.Error()})

// 	}
// 	employment.EmploymentConfiguration = &employmentConfig

// 	return c.Status(200).JSON(employment)
// }

// func EmploymentListController(c *fiber.Ctx) error {
// 	var school models.School
// 	var employments []models.Employment
// 	var total int64

// 	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// 	defer cancel()
// 	// Open the database connection
// 	db := database.DBConnection()
// 	defer database.CloseDB(db)

// 	page := c.QueryInt("page", 1)    // default to page 1 if not provided
// 	limit := c.QueryInt("limit", 10) // default to 10 items per page if not provided
// 	schoolCode := c.Query("school_code")

// 	// get the school
// 	school, err := school.RetrieveBySchoolCode(ctx, db, schoolCode)
// 	if err != nil {
// 		return c.Status(400).JSON(Response{Message: "an error occured  getting school with code.", Success: false, Detail: err.Error()})

// 	}

// 	db.Model(&models.Employment{}).
// 		Where("school_id", school.ID).Count(&total)

// 	err = db.Scopes(adapters.NewPaginateAdapter(limit, page).PaginatedResult).WithContext(ctx).Model(&models.Employment{}).Where("school_id", school.ID).Find(&employments).Error
// 	if err != nil {
// 		if err == gorm.ErrRecordNotFound {
// 			return c.Status(400).JSON(Response{Message: "employment with this id does not exists.", Success: false, Detail: err.Error()})

// 		}
// 		return c.Status(400).JSON(Response{Message: "an error occured filtering for the employment.", Success: false, Detail: err.Error()})

// 	}

// 	return c.Status(200).JSON(fiber.Map{
// 		"total": total,
// 		"data":  employments,
// 	})
// }

// func EmploymentDeleteController(c *fiber.Ctx) error {
// 	/*
// 		this is used to delete the employment which was created
// 	*/
// 	schoolCode := c.Query("school_code")
// 	employmentConfigID := c.Params("id")

// 	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// 	defer cancel()

// 	// Open the database connection
// 	db := database.DBConnection()
// 	defer database.CloseDB(db)

// 	// get the logged-in user
// 	user := c.Locals("user").(models.User)
// 	// get the school
// 	var school models.School
// 	school, err := school.RetrieveBySchoolCode(ctx, db, schoolCode)
// 	if err != nil {
// 		return c.Status(404).JSON(Response{
// 			Message: "school with this id does not exists",
// 			Success: false,
// 			Detail:  err,
// 		})
// 	}

// 	isOk := school.IsSchoolAdminOrOwner(ctx, db, schoolCode, user.ID)
// 	if !isOk {
// 		return c.Status(400).JSON(Response{Message: "you dont have permission", Success: false, Detail: err})
// 	}

// 	var employment models.Employment
// 	err = db.WithContext(ctx).Model(&employment).Where("id = ?", employmentConfigID).Where("school_id = ?", school.ID).First(&employment).Delete(&employment).Error
// 	if err != nil {
// 		if err == gorm.ErrRecordNotFound {
// 			return c.Status(404).JSON(Response{Message: "employment with this id does not exists", Success: false, Detail: err.Error()})

// 		}
// 		return c.Status(400).JSON(Response{Message: "Error deleting employment.", Success: false, Detail: err.Error()})

// 	}
// 	return c.Status(400).JSON(Response{Message: "employment deleted successfully.", Success: false, Detail: err.Error()})

// }

// func EmploymentApplicationController(c *fiber.Ctx) error {
// 	/*
// 		this is used to apply to job with the employment id*/

// 	var employment models.Employment
// 	var employmentApplication models.EmploymentApplication
// 	var foundEmploymentApplication models.EmploymentApplication
// 	var requestBody serializers.EmploymentApplicationRequestSerializer

// 	validator := adapters.NewValidate()
// 	fileUpload := adapters.NewFileUpload()

// 	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// 	defer cancel()

// 	// Open the database connection
// 	db := database.DBConnection()
// 	defer database.CloseDB(db)

// 	// get the employment id
// 	employmentID, err := uuid.Parse(c.Params("id"))
// 	if err != nil {
// 		return c.Status(400).JSON(Response{Message: "invalid employment id", Success: false, Detail: err.Error()})

// 	}

// 	employment, err = employment.Retrieve(ctx, db, employmentID)
// 	if err != nil {
// 		return c.Status(400).JSON(Response{Message: "employment with this id does not exists", Success: false, Detail: err.Error()})

// 	}

// 	if err = c.BodyParser(&requestBody); err != nil {
// 		return c.Status(400).JSON(Response{Message: "invalid request body", Success: false, Detail: err.Error()})
// 	}

// 	err = json.Unmarshal([]byte(c.FormValue("employment_custom_field_answers")), &requestBody.EmploymentCustomFieldAnswers)
// 	if err != nil {
// 		return c.Status(400).JSON(Response{Message: "Error mapping custom field answers", Success: false, Detail: err.Error()})

// 	}
// 	err = json.Unmarshal([]byte(c.FormValue("additional_questions_answer")), &requestBody.AdditionalQuestionsAnswer)
// 	if err != nil {
// 		return c.Status(400).JSON(Response{Message: "Error mapping additional questions answers", Success: false, Detail: err.Error()})

// 	}

// 	requestBody.Resume, err = fileUpload.UploadFile("resume", c)
// 	requestBody.CoverLetter, err = fileUpload.UploadFile("cover_letter", c)
// 	requestBody.Transcript, err = fileUpload.UploadFile("transcript", c)
// 	requestBody.TeachingCertification, err = fileUpload.UploadFile("teaching_certification", c)

// 	// call the validate function in the request serializer
// 	if err := validator.ValidateData(&requestBody); err != nil {
// 		return c.Status(400).JSON(Response{Message: err, Success: false, Detail: err})

// 	}

// 	err = copier.Copy(&employmentApplication, &requestBody)
// 	if err != nil {
// 		return c.Status(500).JSON(Response{Message: "Error from our end copying data", Success: false, Detail: err.Error()})

// 	}
// 	// todo : do validation base on the configuration provided

// 	employmentApplication.EmploymentID = &employmentID
// 	employmentApplication.SchoolID = employment.SchoolID

// 	var applicantCount int64
// 	err = db.WithContext(ctx).Model(&employmentApplication).
// 		Where("school_id = ?", employment.SchoolID).Count(&applicantCount).Error
// 	if err != nil {
// 		return c.Status(400).JSON(Response{Message: "error getting applicant count ", Success: false, Detail: err.Error()})

// 	}

// 	applicantID := fmt.Sprintf("#%05d", applicantCount+1)
// 	employmentApplication.ApplicantID = &applicantID

// 	// check if the email has applied to this job before
// 	err = db.WithContext(ctx).Model(&employmentApplication).
// 		Where("school_id = ?", employment.SchoolID).
// 		Where("email = ?", employmentApplication.Email).
// 		Where("employment_id = ?", employmentID).First(&foundEmploymentApplication).Error
// 	if err != nil {
// 		if err != gorm.ErrRecordNotFound {
// 			logger.Error(ctx, "Error  filtering for the employment application ", zap.Error(err))
// 			return c.Status(400).JSON(Response{Message: "An error occured filtering for the employment application ", Success: false, Detail: err.Error()})

// 		}
// 	}
// 	if foundEmploymentApplication.Email != nil {
// 		return c.Status(400).JSON(Response{Message: "already applied to this job ", Success: false, Detail: nil})

// 	}

// 	err = db.WithContext(ctx).Model(&employmentApplication).Create(&employmentApplication).Error
// 	if err != nil {
// 		return c.Status(500).JSON(Response{Message: "Error creating employment application", Success: false, Detail: err.Error()})

// 	}

// 	return c.Status(200).JSON(employmentApplication)
// }

// func EmploymentApplicantListController(c *fiber.Ctx) error {
// 	/*This is used to view all applicants that have applied to a job */
// 	var total int64
// 	var employmentApplicants []models.EmploymentApplication

// 	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// 	defer cancel()
// 	// Open the database connection
// 	db := database.DBConnection()
// 	defer database.CloseDB(db)

// 	schoolCode := c.Query("school_code")

// 	employmentID, err := uuid.Parse(c.Params("id"))
// 	if err != nil {
// 		return c.Status(400).JSON(Response{Message: "not a valid uuid", Success: false, Detail: err.Error()})

// 	}
// 	page := c.QueryInt("page", 1)    // default to page 1 if not provided
// 	limit := c.QueryInt("limit", 10) // default to 10 items per page if not provided

// 	// get the logged-in user
// 	user := c.Locals("user").(models.User)

// 	// get the school
// 	var school models.School
// 	school, err = school.RetrieveBySchoolCode(ctx, db, schoolCode)
// 	if err != nil {
// 		return c.Status(404).JSON(Response{
// 			Message: "school with this id does not exists",
// 			Success: false,
// 			Detail:  err,
// 		})
// 	}

// 	isOk := school.IsSchoolAdminOrOwner(ctx, db, schoolCode, user.ID)
// 	if !isOk {
// 		return c.Status(400).JSON(Response{Message: "you dont have permission", Success: false, Detail: err})
// 	}

// 	db.WithContext(ctx).Model(&models.EmploymentApplication{}).
// 		Where("employment_id = ?", employmentID).Count(&total)

// 	err = db.Scopes(adapters.NewPaginateAdapter(limit, page).PaginatedResult).WithContext(ctx).Model(&models.EmploymentApplication{}).
// 		Where("employment_id = ?", employmentID).
// 		Where("school_id = ?", school.ID).Find(&employmentApplicants).Error

// 	if err != nil {
// 		return c.Status(500).JSON(Response{Message: "An error occured filtering for the  applicants", Success: false, Detail: err.Error()})

// 	}

// 	return c.Status(200).JSON(fiber.Map{"total": total, "data": employmentApplicants})
// }

// func EmploymentApplicantAcceptController(c *fiber.Ctx) error {
// 	/* this is used to accept an employment applicant and can only be done by one of the admins in the school*/
// 	var classes []*models.Class

// 	schoolCode := c.Query("school_code")

// 	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// 	defer cancel()
// 	// Open the database connection
// 	db := database.DBConnection()
// 	defer database.CloseDB(db)

// 	// get the logged-in user
// 	user := c.Locals("user").(models.User)
// 	userUtils := utils.NewUserUtils()
// 	emailAdapters := adapters.NewEmailAdapter()
// 	staffUtils := utils.NewStaffUtils()

// 	var requestBody serializers.AcceptEmploymentSerializer

// 	if err := c.BodyParser(&requestBody); err != nil {
// 		return c.Status(400).JSON(Response{Message: "invalid request body", Success: false, Detail: err.Error()})
// 	}

// 	// get the school
// 	var school models.School
// 	school, err := school.RetrieveBySchoolCode(ctx, db, schoolCode)
// 	if err != nil {
// 		return c.Status(404).JSON(Response{
// 			Message: "school with this id does not exists",
// 			Success: false,
// 			Detail:  err,
// 		})
// 	}

// 	isOk := school.IsSchoolAdminOrOwner(ctx, db, schoolCode, user.ID)
// 	if !isOk {
// 		return c.Status(400).JSON(Response{Message: "you dont have permission", Success: false, Detail: err})
// 	}

// 	// get the employment
// 	var employment models.Employment
// 	employment, err = employment.Retrieve(ctx, db, *requestBody.EmploymentID)
// 	if err != nil {
// 		return c.Status(400).JSON(Response{Message: "employment with this id does not exists", Success: false, Detail: err.Error()})

// 	}

// 	// get the employment application
// 	var employmentApplication models.EmploymentApplication
// 	employmentApplication, err = employmentApplication.Retrieve(ctx, db, *requestBody.ApplicantID)
// 	if err != nil {
// 		return c.Status(400).JSON(Response{Message: "employment application with this id does not exists", Success: false, Detail: err.Error()})

// 	}

// 	// get the class if its provided, but not all staff can have a class
// 	var class models.Class
// 	if requestBody.ClassID != nil {
// 		class, err = class.Retrieve(ctx, db, *requestBody.ClassID)
// 		if err != nil {
// 			return c.Status(404).JSON(Response{Message: "staff with this id does not exists", Success: false, Detail: err.Error()})

// 		}
// 		classes = append(classes, &class)
// 	}

// 	if requestBody.Status != nil {
// 		err = db.WithContext(ctx).Model(&employmentApplication).Where("school_id = ?", school.ID).Update("status", requestBody.Status).Error
// 		if err != nil {
// 			return c.Status(400).JSON(Response{Message: "error updating employment status", Success: false, Detail: err})

// 		}

// 		if *requestBody.Status == "DECLINE" {
// 			_, err := emailAdapters.SendApplicationDecline(ctx, *employmentApplication.FirstName, *employmentApplication.LastName,
// 				*employmentApplication.Email, *school.Name)
// 			if err != nil {
// 				return c.Status(400).JSON(Response{Message: "error sending mail but user successfully declined", Success: false, Detail: err.Error()})

// 			}
// 			return c.Status(200).JSON(Response{Message: "Successfully decline staff.", Success: true, Detail: err.Error()})

// 		}

// 	}

// 	// generate password
// 	password, err := userUtils.GenerateRandomPassword(8)
// 	if err != nil {
// 		return c.Status(500).JSON(Response{Message: "an error occured generating password", Success: false, Detail: err.Error()})
// 	}

// 	// create a user
// 	isFalse := false
// 	currentTime := time.Now()
// 	newUser := models.User{
// 		FirstName:   employmentApplication.FirstName,
// 		LastName:    employmentApplication.LastName,
// 		IsStaff:     &isFalse,
// 		IsSuperUser: &isFalse,
// 		IsVerified:  &isFalse,
// 		PhoneNumber: employmentApplication.Phone,
// 		Email:       employmentApplication.Email,
// 		CreatedAt:   &currentTime,
// 		Password:    &password,
// 	}
// 	newUser.ID = uuid.New() // set the id of the user because we use a base model

// 	// check if a user exists with the mail
// 	newUser, userExist, err := newUser.CheckIfEmailExists(db, ctx, *newUser.Email)
// 	if !userExist {
// 		logger.Info(ctx, "User does  not exist", zap.String("staff_user", "staff user does not exist creating user"))
// 		newUser, err = userUtils.CreateUser(db, ctx, newUser, true)
// 		if err != nil {
// 			return c.Status(400).JSON(Response{Message: err.Error(), Success: false, Detail: err.Error()})
// 		}

// 	}

// 	// check if the user is already a staff on the school
// 	err = db.WithContext(ctx).Model(&models.Staff{}).Where("user_id = ?", newUser.ID).Where("school_id = ?", school.ID).First(&models.Staff{}).Error
// 	if err == nil {
// 		return c.Status(400).JSON(Response{Message: "a staff is already attached to the user.", Success: false, Detail: nil})

// 	}

// 	// create the staff and make him active
// 	staffCode := staffUtils.GenerateStaffCode()
// 	status := "ACTIVE"
// 	staff := models.Staff{
// 		SchoolID:                &school.ID,
// 		StaffCode:               &staffCode,
// 		UserID:                  &newUser.ID,
// 		EmploymentApplicationID: &employmentApplication.ID,
// 		Classes:                 classes,
// 		RoleID:                  employment.RoleID,
// 		CategoryID:              employment.CategoryID,
// 		Qualification:           employmentApplication.OtherQualification,
// 		NetSalary:               nil,
// 		Status:                  &status,
// 	}

// 	// create the staff
// 	err = db.WithContext(ctx).Model(&staff).Create(&staff).Error
// 	if err != nil {
// 		return c.Status(500).JSON(Response{Message: "an error occured creating staff", Success: false, Detail: err.Error()})
// 	}

// 	userType := "STAFF"
// 	IsOnline := false
// 	conversationUser := models.ConversationUser{
// 		SchoolID: &school.ID,
// 		UserID:   staff.UserID,
// 		UserType: &userType,
// 		StaffID:  &staff.ID,
// 		IsOnline: &IsOnline,
// 	}
// 	err = conversationUser.CreateConversationUser(ctx, db)
// 	if err != nil {
// 		return c.Status(400).JSON(Response{Message: "error creating conversation user for staff.", Success: false, Detail: err.Error()})

// 	}

// 	return c.Status(200).JSON(Response{Message: "applicant accepted successfully", Success: false, Detail: nil})

// }
