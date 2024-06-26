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
// 	"strings"
// 	"time"
// )

// // EnrollmentConfigurationCreateController /* Creating enrollment the configuration have to be set for the job first before creating the job itself*/
// func EnrollmentConfigurationCreateController(c *fiber.Ctx) error {
// 	/* this is used to create enrollment configuration*/
// 	schoolCode := c.Query("school_code")

// 	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// 	defer cancel()

// 	// get the logged-in user
// 	user := c.Locals("user").(models.User)

// 	// Open the database connection
// 	db := database.DBConnection()
// 	defer database.CloseDB(db)

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
// 	var requestBody models.EnrollmentConfiguration
// 	if err := c.BodyParser(&requestBody); err != nil {
// 		return c.Status(400).JSON(Response{Message: "invalid request body", Success: false, Detail: err.Error()})
// 	}

// 	// call the validate function in the request serializer for login
// 	validator := adapters.NewValidate()
// 	err2 := validator.ValidateData(&requestBody)

// 	if err2 != nil {
// 		return c.Status(400).JSON(Response{Message: err, Success: false, Detail: err})
// 	}

// 	// also validate the custom field data
// 	if requestBody.EnrollmentCustomFields != nil {
// 		for _, customField := range *requestBody.EnrollmentCustomFields {

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

// 	enrollmentConfiguration := requestBody
// 	enrollmentConfiguration.SchoolID = &school.ID

// 	err = db.WithContext(ctx).Model(&enrollmentConfiguration).Create(&enrollmentConfiguration).Error
// 	if err != nil {
// 		return c.Status(500).JSON(Response{Message: "Error creating enrollment configuration", Success: false, Detail: err.Error()})

// 	}
// 	return c.Status(201).JSON(enrollmentConfiguration)
// }

// func EnrollmentConfigurationListController(c *fiber.Ctx) error {
// 	/* this is used to create enrollment configuration*/
// 	schoolCode := c.Query("school_code")

// 	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// 	defer cancel()

// 	// Open the database connection
// 	db := database.DBConnection()
// 	defer database.CloseDB(db)

// 	var school models.School
// 	school, err := school.RetrieveBySchoolCode(ctx, db, schoolCode)
// 	if err != nil {
// 		return c.Status(404).JSON(Response{
// 			Message: "school with this id does not exists",
// 			Success: false,
// 			Detail:  err,
// 		})
// 	}

// 	// make the enrollment configuration list
// 	var enrollmentConfigurations []models.EnrollmentConfiguration
// 	err = db.WithContext(ctx).Model(&models.EnrollmentConfiguration{}).Where("school_id = ?", school.ID).
// 		Order("timestamp desc").
// 		Find(&enrollmentConfigurations).Error
// 	if err != nil {
// 		if err == gorm.ErrRecordNotFound {
// 			return c.Status(404).JSON(Response{Message: "Error fetching enrollment configurations", Success: false, Detail: err.Error()})
// 		} else {
// 			return c.Status(500).JSON(Response{Message: "Error deleting enrollment configuration", Success: false, Detail: err.Error()})

// 		}
// 	}
// 	return c.Status(200).JSON(enrollmentConfigurations)
// }

// func EnrollmentConfigurationDetailController(c *fiber.Ctx) error {
// 	schoolCode := c.Query("school_code")
// 	enrollmentConfigID := c.Params("id")

// 	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// 	defer cancel()

// 	// Open the database connection
// 	db := database.DBConnection()
// 	defer database.CloseDB(db)

// 	var school models.School
// 	school, err := school.RetrieveBySchoolCode(ctx, db, schoolCode)
// 	if err != nil {
// 		return c.Status(404).JSON(Response{
// 			Message: "school with this id does not exists",
// 			Success: false,
// 			Detail:  err,
// 		})
// 	}

// 	var enrollmentConfiguration models.EnrollmentConfiguration
// 	err = db.WithContext(ctx).Model(&enrollmentConfiguration).Where("id = ?", enrollmentConfigID).First(&enrollmentConfiguration).Error
// 	if err != nil {
// 		if err == gorm.ErrRecordNotFound {
// 			return c.Status(404).JSON(Response{Message: "Error fetching enrollment configurations", Success: false, Detail: err.Error()})
// 		} else {
// 			return c.Status(500).JSON(Response{Message: "Error deleting enrollment configuration", Success: false, Detail: err.Error()})

// 		}
// 	}
// 	return c.Status(200).JSON(enrollmentConfiguration)
// }

// func EnrollmentConfigurationUpdateController(c *fiber.Ctx) error {
// 	/* this is used to create enrollment configuration*/
// 	schoolCode := c.Query("school_code")
// 	enrollmentConfigID := c.Params("id")

// 	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// 	defer cancel()

// 	// Open the database connection
// 	db := database.DBConnection()
// 	defer database.CloseDB(db)

// 	// get the logged-in user
// 	user := c.Locals("user").(models.User)
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
// 	var requestBody models.EnrollmentConfiguration
// 	if err := c.BodyParser(&requestBody); err != nil {
// 		return c.Status(400).JSON(Response{Message: "invalid request body", Success: false, Detail: err.Error()})
// 	}
// 	// call the validate function in the request serializer for login
// 	validator := adapters.NewValidate()
// 	err2 := validator.ValidateData(&requestBody)

// 	if err2 != nil {
// 		return c.Status(400).JSON(Response{Message: err, Success: false, Detail: err})
// 	}

// 	enrollmentConfiguration := requestBody

// 	err = db.WithContext(ctx).Model(&enrollmentConfiguration).Where("id = ?", enrollmentConfigID).Updates(&enrollmentConfiguration).First(&enrollmentConfiguration).Error
// 	if err != nil {
// 		return c.Status(500).JSON(Response{Message: "Error creating enrollment configuration", Success: false, Detail: err.Error()})

// 	}
// 	return c.Status(200).JSON(enrollmentConfiguration)
// }

// func EnrollmentConfigurationDeleteController(c *fiber.Ctx) error {
// 	schoolCode := c.Query("school_code")
// 	enrollmentConfigID := c.Params("id")

// 	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// 	defer cancel()

// 	// Open the database connection
// 	db := database.DBConnection()
// 	defer database.CloseDB(db)

// 	// get the logged-in user
// 	user := c.Locals("user").(models.User)
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

// 	var enrollmentConfiguration models.EnrollmentConfiguration
// 	err = db.WithContext(ctx).Model(&enrollmentConfiguration).Where("id = ?", enrollmentConfigID).First(&enrollmentConfiguration).Delete(&enrollmentConfiguration).Error
// 	if err != nil {
// 		if err == gorm.ErrRecordNotFound {
// 			return c.Status(404).JSON(Response{Message: "enrollment configuration with this id does not exists", Success: false, Detail: err.Error()})

// 		}
// 		return c.Status(500).JSON(Response{Message: "Error deleting enrollment configuration", Success: false, Detail: err.Error()})

// 	}
// 	return c.Status(204).JSON(Response{Message: "enrollment configuration deleted successfully", Success: true, Detail: nil})

// }

// func EnrollmentApplicationController(c *fiber.Ctx) error {
// 	/*
// 		this is used to apply to job with the enrollment id*/

// 	// get the enrollment id
// 	var enrollmentConfig models.EnrollmentConfiguration
// 	var enrollmentAdmission models.EnrollmentAdmission
// 	var requestBody serializers.EnrollmentApplicationRequestSerializer

// 	schoolCode := c.Query("school_code")
// 	validator := adapters.NewValidate()
// 	fileUpload := adapters.NewFileUpload()

// 	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// 	defer cancel()

// 	// Open the database connection
// 	db := database.DBConnection()
// 	defer database.CloseDB(db)

// 	if err := c.BodyParser(&requestBody); err != nil {
// 		return c.Status(400).JSON(Response{Message: "invalid request body", Success: false, Detail: err.Error()})
// 	}
// 	err := json.Unmarshal([]byte(c.FormValue("enrollment_custom_field_answers")), &requestBody.EnrollmentCustomFieldAnswers)
// 	if err != nil {
// 		logger.Error(ctx, "Error converting custom field answers to json", zap.Error(err))
// 		return c.Status(400).JSON(Response{Message: "error converting custom field answers to json", Success: false, Detail: err.Error()})

// 	}
// 	err = json.Unmarshal([]byte(c.FormValue("additional_questions_answer")), &requestBody.AdditionalQuestionsAnswer)
// 	if err != nil {
// 		logger.Error(ctx, "Error converting additional question answers to json", zap.Error(err))
// 		return c.Status(400).JSON(Response{Message: "error converting custom field answers to json", Success: false, Detail: err.Error()})

// 	}

// 	requestBody.BirthCertificate, err = fileUpload.UploadFile("birth_certificate", c)
// 	requestBody.ImmunizationRecord, err = fileUpload.UploadFile("immunization_record", c)
// 	requestBody.ProofOfAddress, err = fileUpload.UploadFile("proof_of_address", c)
// 	requestBody.ReportCard, err = fileUpload.UploadFile("report_card", c)
// 	requestBody.OtherSupportingDocument, err = fileUpload.UploadFile("other_supporting_document", c)

// 	requestBody.SchoolID = enrollmentConfig.SchoolID
// 	err2 := validator.ValidateData(&requestBody)
// 	if err2 != nil {
// 		return c.Status(400).JSON(Response{Message: err, Success: false, Detail: err})
// 	}

// 	var school models.School
// 	school, err = school.RetrieveBySchoolCode(ctx, db, schoolCode)
// 	if err != nil {
// 		return c.Status(404).JSON(Response{
// 			Message: "school with this id does not exists",
// 			Success: false,
// 			Detail:  err,
// 		})
// 	}

// 	// get the school enrolment config
// 	err = db.WithContext(ctx).Model(&enrollmentConfig).Where("school_id =? ", school.ID).First(&enrollmentConfig).Error
// 	if err != nil {
// 		if err == gorm.ErrRecordNotFound {
// 			return c.Status(400).JSON(
// 				Response{Message: "Enrolment config with this school code does not exist", Success: false, Detail: err})
// 		}
// 		return c.Status(500).JSON(
// 			Response{Message: "An error occured trying to get the first enrollment configuration", Success: false, Detail: err})
// 	}

// 	err = copier.Copy(&enrollmentAdmission, &requestBody)
// 	if err != nil {
// 		return c.Status(500).JSON(Response{Message: "error copying data", Success: false, Detail: err.Error()})

// 	}

// 	enrollmentAdmission.EnrollmentConfigurationID = &enrollmentConfig.ID
// 	enrollmentAdmission.SchoolID = &school.ID
// 	err = db.WithContext(ctx).Model(&enrollmentAdmission).Create(&enrollmentAdmission).Error
// 	if err != nil {
// 		return c.Status(500).JSON(Response{Message: "Error creating enrollment application", Success: false, Detail: err.Error()})

// 	}

// 	return c.Status(200).JSON(enrollmentAdmission)
// }

// func EnrollmentApplicantListController(c *fiber.Ctx) error {
// 	/*This is used to view all applicants that have applied to a job */
// 	var total int64
// 	var enrollmentAdmissions []models.EnrollmentAdmission

// 	schoolCode := c.Query("school_code")
// 	search := c.Query("search")
// 	status := c.Query("status")

// 	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// 	defer cancel()
// 	// Open the database connection
// 	db := database.DBConnection()
// 	defer database.CloseDB(db)

// 	// get the logged-in user
// 	user := c.Locals("user").(models.User)
// 	page := c.QueryInt("page", 1)    // default to page 1 if not provided
// 	limit := c.QueryInt("limit", 10) // default to 10 items per page if not provided

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

// 	db.WithContext(ctx).Model(&models.EnrollmentAdmission{}).
// 		Where("school_id = ?", school.ID).Count(&total)

// 	query := db.Scopes(adapters.NewPaginateAdapter(limit, page).PaginatedResult).WithContext(ctx).Model(&models.EnrollmentAdmission{}).
// 		Where("school_id = ?", school.ID)

// 	if search != "" {
// 		query = query.Where("full_name ILIKE ? or status ILIKE ? or email ILIKE ?", "%"+search+"%", "%"+search+"%", "%"+search+"%")
// 	}

// 	if status != "" {
// 		// Add status condition when it's not empty
// 		query = query.Where("status = ?", status)
// 	}

// 	err = query.Find(&enrollmentAdmissions).Error

// 	if err != nil {
// 		if err == gorm.ErrRecordNotFound {
// 			return c.Status(200).JSON(enrollmentAdmissions)
// 		}
// 		return c.Status(500).JSON(Response{Message: "An error occured filtering for the  applicants", Success: false, Detail: err.Error()})

// 	}

// 	return c.Status(200).JSON(fiber.Map{"total": total, "data": enrollmentAdmissions})
// }

// func EnrollmentApplicantDetailController(c *fiber.Ctx) error {
// 	/*This is used to view all applicants that have applied to a job */
// 	schoolCode := c.Query("school_code")
// 	_enrolmentApplicantID := c.Params("id")
// 	enrolmentApplicantID, err := uuid.Parse(_enrolmentApplicantID)
// 	if err != nil {
// 		return c.Status(400).JSON(Response{Message: "Not a valid UUID passed in  params", Success: false, Detail: err.Error()})

// 	}

// 	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// 	defer cancel()
// 	// Open the database connection
// 	db := database.DBConnection()
// 	defer database.CloseDB(db)

// 	// get the logged-in user
// 	user := c.Locals("user").(models.User)
// 	var enrollmentAdmission models.EnrollmentAdmission

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

// 	err = db.WithContext(ctx).Model(&models.EnrollmentAdmission{}).
// 		Where("school_id = ?", school.ID).
// 		Where("id =?", enrolmentApplicantID).First(&enrollmentAdmission).Error

// 	if err != nil {
// 		if err == gorm.ErrRecordNotFound {

// 			return c.Status(404).JSON(
// 				Response{Message: "enrolment applicant does not exist", Success: false, Detail: err})
// 		}
// 		return c.Status(500).JSON(
// 			Response{Message: "An error occured filtering for the  applicants", Success: false, Detail: err})
// 	}

// 	return c.Status(200).JSON(enrollmentAdmission)
// }

// func EnrollmentApplicantAcceptController(c *fiber.Ctx) error {
// 	/* this is used to accept an enrollment applicant and can only be done by one of the admins in the school*/
// 	var class models.Class

// 	var requestBody serializers.AcceptEnrollmentSerializer
// 	var conversationUser models.ConversationUser

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

// 	// get the enrollment application
// 	var enrollmentAdmission models.EnrollmentAdmission
// 	enrollmentAdmission, err = enrollmentAdmission.Retrieve(ctx, db, *requestBody.ApplicantID)
// 	if err != nil {
// 		return c.Status(400).JSON(Response{Message: "enrolment application with this id does not exists", Success: false, Detail: err.Error()})

// 	}

// 	if requestBody.ClassID != nil {
// 		class, err = class.Retrieve(ctx, db, *requestBody.ClassID)
// 		if err != nil {
// 			return c.Status(404).JSON(Response{Message: "staff with this id does not exists", Success: false, Detail: err.Error()})

// 		}
// 	}

// 	names := strings.Split(*enrollmentAdmission.FullName, " ")
// 	firstName := names[0]
// 	lastName := ""
// 	if len(names) > 1 {
// 		lastName = names[1]
// 	}

// 	// generate password
// 	password, err := userUtils.GenerateRandomPassword(8)
// 	if err != nil {
// 		return c.Status(400).JSON(Response{Message: "an error occured generating password", Success: false, Detail: err.Error()})

// 	}

// 	// create a user
// 	isFalse := false
// 	currentTime := time.Now()

// 	newUser := models.User{
// 		FirstName:   &firstName,
// 		LastName:    &lastName,
// 		IsStaff:     &isFalse,
// 		IsSuperUser: &isFalse,
// 		IsVerified:  &isFalse,
// 		PhoneNumber: enrollmentAdmission.Phone,
// 		Email:       enrollmentAdmission.Email,
// 		CreatedAt:   &currentTime,
// 		Password:    &password,
// 	}
// 	newUser.ID = uuid.New() // set the id of the user because we use a base model

// 	// check if a user exists with the mail
// 	newUser, userExist, err := newUser.CheckIfEmailExists(db, ctx, *newUser.Email)
// 	if !userExist {
// 		newUser, err = userUtils.CreateUser(db, ctx, newUser, true)
// 		if err != nil {
// 			return c.Status(400).JSON(Response{Message: "you dont have permission", Success: false, Detail: err.Error()})
// 		}
// 	}

// 	// check if the user is a student
// 	var studentExistCount int64
// 	err = db.WithContext(ctx).Model(&models.Student{}).Where("school_id =?", school.ID).Where("user_id =?", newUser.ID).Count(&studentExistCount).Error
// 	if err != nil {
// 		if err != gorm.ErrRecordNotFound {
// 			return c.Status(400).JSON(Response{Message: "error checking if student exist", Success: false, Detail: err.Error()})
// 		}
// 	}

// 	if studentExistCount > 0 {
// 		return c.Status(400).JSON(Response{Message: "User have already been accepted", Success: false, Detail: err.Error()})
// 	}

// 	var studentCount int64
// 	err = db.WithContext(ctx).Model(&models.Student{}).Where("school_id = ?", school.ID).Count(&studentCount).Error
// 	if err != nil {
// 		return c.Status(400).JSON(Response{Message: "error getting student count", Success: false, Detail: err.Error()})
// 	}
// 	studentCode := fmt.Sprintf("NN%04d", studentCount+1)
// 	status := "PAID"
// 	student := models.Student{
// 		SchoolID:              &school.ID,
// 		UserID:                &newUser.ID,
// 		ClassID:               requestBody.ClassID,
// 		Department:            nil,
// 		Section:               nil,
// 		Status:                &status,
// 		StudentCode:           &studentCode,
// 		EnrollmentAdmissionID: requestBody.ApplicantID,
// 		Gender:                enrollmentAdmission.Gender,
// 	}
// 	// create the student
// 	err = db.WithContext(ctx).Model(&student).Preload("User").Preload("Class").Create(&student).Error
// 	if err != nil {
// 		return c.Status(400).JSON(Response{Message: "an error occured creating staff", Success: false, Detail: err.Error()})
// 	}

// 	userType := "STUDENT"
// 	IsOnline := false
// 	conversationUser = models.ConversationUser{
// 		SchoolID:  &school.ID,
// 		UserType:  &userType,
// 		UserID:    student.UserID,
// 		StudentID: &student.ID,
// 		IsOnline:  &IsOnline,
// 	}
// 	err = conversationUser.CreateConversationUser(ctx, db)
// 	if err != nil {
// 		return c.Status(400).JSON(Response{Message: "error creating conversation user", Success: false, Detail: err.Error()})

// 	}

// 	// check the status
// 	if requestBody.Status != nil {
// 		err := db.WithContext(ctx).Model(&enrollmentAdmission).Where("id = ?", enrollmentAdmission.ID).Update("status", *requestBody.Status).Error
// 		if err != nil {
// 			return c.Status(400).JSON(Response{Message: "error updating status", Success: false, Detail: err.Error()})

// 		}
// 		if *requestBody.Status == "DECLINE" {
// 			_, err = emailAdapters.SendAdmissionDecline(ctx, firstName, lastName, *enrollmentAdmission.Email, *school.Name)
// 			if err != nil {
// 				return c.Status(400).JSON(Response{Message: "error sending mail but user successfully declined", Success: false, Detail: err.Error()})

// 			}
// 			return c.Status(200).JSON(Response{Message: "Success fully decline student", Success: true, Detail: err.Error()})

// 		}
// 	}

// 	return c.Status(200).JSON(Response{Message: "applicant accepted successfully", Success: true, Detail: err.Error()})

// }

// func EnrollmentAnalyticsController(c *fiber.Ctx) error {
// 	var serializer serializers.EnrollmentAnalyticsResponseSerializer
// 	var enrollmentAdmission models.EnrollmentAdmission

// 	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// 	defer cancel()
// 	// Open the database connection
// 	db := database.DBConnection()
// 	defer database.CloseDB(db)

// 	schoolCode := c.Query("school_code")
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

// 	user := c.Locals("user").(models.User)

// 	isOk := school.IsSchoolAdminOrOwner(ctx, db, schoolCode, user.ID)
// 	if !isOk {
// 		return c.Status(400).JSON(Response{Message: "you dont have permission", Success: false, Detail: err.Error()})
// 	}

// 	err = db.WithContext(ctx).Model(&enrollmentAdmission).Where("school_id = ?", school.ID).Count(&serializer.AllApplicants).Error
// 	if err != nil {
// 		return c.Status(400).JSON(Response{Message: "error getting all count", Success: false, Detail: err.Error()})
// 	}
// 	err = db.WithContext(ctx).Model(&enrollmentAdmission).Where("school_id = ?", school.ID).Where("status =?", "APPROVE").Count(&serializer.ApprovedApplicants).Error
// 	if err != nil {
// 		return c.Status(400).JSON(Response{Message: "error getting approve count", Success: false, Detail: err.Error()})
// 	}
// 	err = db.WithContext(ctx).Model(&enrollmentAdmission).Where("school_id = ?", school.ID).Where("status =?", "DECLINE").Count(&serializer.DeclinedApplicants).Error
// 	if err != nil {
// 		return c.Status(400).JSON(Response{Message: "error getting declined count", Success: false, Detail: err.Error()})

// 	}
// 	err = db.WithContext(ctx).Model(&enrollmentAdmission).Where("school_id = ?", school.ID).Where("status =?", "DECLINE").Count(&serializer.PendingApplicants).Error
// 	if err != nil {
// 		return c.Status(400).JSON(Response{Message: "error getting pending count", Success: false, Detail: err.Error()})
// 	}
// 	return c.Status(200).JSON(fiber.Map{"data": serializer})

// }
