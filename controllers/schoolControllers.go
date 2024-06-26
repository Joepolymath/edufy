package controllers

// import (
// 	"Learnium/adapters"
// 	"Learnium/database"
// 	"Learnium/logger"
// 	"Learnium/models"
// 	"Learnium/serializers"
// 	"Learnium/utils"

// 	// "Learnium/utils"
// 	"context"
// 	"fmt"
// 	"time"

// 	"github.com/gofiber/fiber/v2"
// 	"github.com/google/uuid"
// 	"github.com/jinzhu/copier"
// 	"go.uber.org/zap"
// 	"gorm.io/gorm"
// )

// func SchoolTypeCreateController(c *fiber.Ctx) error {
// 	/* This is used to list the school in which a user is the owner*/
// 	var err error
// 	var requestBody serializers.SchoolTypeCreateRequestSerializer
// 	var schoolType models.SchoolType

// 	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// 	defer cancel()

// 	// Open the database connection
// 	db := database.DBConnection()
// 	defer database.CloseDB(db)

// 	validateAdapter := adapters.NewValidate()
// 	fileUpload := adapters.NewFileUpload()

// 	// logged-in user
// 	user := c.Locals("user").(models.User)

// 	_, err = user.IsSuperUserOrStaff(db, ctx, user.ID)
// 	if err != nil {
// 		return c.Status(400).JSON(Response{Message: err.Error(), Success: false, Detail: err.Error()})
// 	}

// 	if err := c.BodyParser(&requestBody); err != nil {
// 		return c.Status(400).JSON(Response{Message: "invalid request body", Success: false, Detail: err.Error()})
// 	}

// 	requestBody.Image, err = fileUpload.UploadFile("image", c)

// 	errors := validateAdapter.ValidateData(&requestBody)
// 	if errors != nil {
// 		return c.Status(400).JSON(Response{
// 			Message: errors, Success: false, Detail: errors,
// 		})
// 	}

// 	err = copier.Copy(&schoolType, requestBody)
// 	if err != nil {
// 		return c.Status(400).JSON(Response{Message: "error cooying data", Success: false, Detail: err.Error()})
// 	}

// 	err = db.WithContext(ctx).Model(&models.SchoolType{}).Create(&schoolType).Error
// 	if err != nil {
// 		return c.Status(500).JSON(Response{Message: "Error creating school type from our end",
// 			Success: false, Detail: err.Error()})
// 	}

// 	return c.Status(201).JSON(schoolType)
// }

// func SchoolTypeListController(c *fiber.Ctx) error {
// 	var schoolTypes []models.SchoolType

// 	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// 	defer cancel()

// 	// Open the database connection
// 	db := database.DBConnection()
// 	defer database.CloseDB(db)

// 	err := db.WithContext(ctx).Model(&models.SchoolType{}).Find(&schoolTypes).Error
// 	if err != nil {
// 		return c.Status(400).JSON(Response{Message: "error from our end getting school types", Success: false, Detail: err.Error()})
// 	}
// 	// it does not need to be paginated
// 	return c.Status(200).JSON(schoolTypes)
// }

// func SchoolTypeUpdateController(c *fiber.Ctx) error {
// 	var schoolType models.SchoolType
// 	var requestBody serializers.SchoolTypeUpdateRequestSerializer

// 	schoolTypeID, err := uuid.Parse(c.Params("id"))
// 	if err != nil {
// 		return c.Status(400).JSON(Response{Message: "Invalid uuid passed", Success: false, Detail: err.Error()})
// 	}

// 	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// 	defer cancel()

// 	// Open the database connection
// 	db := database.DBConnection()
// 	defer database.CloseDB(db)

// 	validateAdapter := adapters.NewValidate()
// 	fileUpload := adapters.NewFileUpload()

// 	// logged-in user
// 	user := c.Locals("user").(models.User)

// 	_, err = user.IsSuperUserOrStaff(db, ctx, user.ID)
// 	if err != nil {
// 		return c.Status(400).JSON(Response{Message: err.Error(), Success: false, Detail: err.Error()})
// 	}

// 	if err := c.BodyParser(&requestBody); err != nil {
// 		return c.Status(400).JSON(Response{Message: "invalid request body", Success: false, Detail: err.Error()})
// 	}

// 	requestBody.Image, err = fileUpload.UploadFile("image", c)

// 	errors := validateAdapter.ValidateData(&requestBody)
// 	if errors != nil {
// 		return c.Status(400).JSON(Response{
// 			Message: errors, Success: false, Detail: errors,
// 		})
// 	}
// 	err = db.WithContext(ctx).Model(&models.SchoolType{}).Where("id", schoolTypeID).Updates(&requestBody).First(&schoolType).Error
// 	if err != nil {
// 		if err == gorm.ErrRecordNotFound {
// 			return c.Status(404).JSON(Response{
// 				Message: "error school type with this id  does not exist",
// 				Success: false,
// 				Detail:  err.Error(),
// 			})
// 		}
// 		return c.Status(400).JSON(Response{Message: "error on our end getting this school type", Success: false, Detail: err.Error()})
// 	}

// 	return c.Status(200).JSON(schoolType)
// }

// func SchoolTypeDeleteController(c *fiber.Ctx) error {
// 	var schoolType models.SchoolType

// 	schoolTypeID, err := uuid.Parse(c.Params("id"))
// 	if err != nil {
// 		return c.Status(400).JSON(Response{Message: "Invalid uuid passed", Success: false, Detail: err.Error()})
// 	}

// 	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// 	defer cancel()

// 	// Open the database connection
// 	db := database.DBConnection()
// 	defer database.CloseDB(db)

// 	// logged-in user
// 	user := c.Locals("user").(models.User)

// 	_, err = user.IsSuperUserOrStaff(db, ctx, user.ID)
// 	if err != nil {
// 		return c.Status(400).JSON(Response{Message: err.Error(), Success: false, Detail: err.Error()})
// 	}

// 	schoolType.ID = schoolTypeID
// 	err = db.WithContext(ctx).Model(&models.SchoolType{}).Where("id", schoolTypeID).Delete(&schoolType).Error
// 	if err != nil {
// 		if err == gorm.ErrRecordNotFound {
// 			return c.Status(404).JSON(Response{
// 				Message: "error school type with this id  does not exist",
// 				Success: false,
// 				Detail:  err.Error(),
// 			})
// 		}
// 		return c.Status(400).JSON(Response{Message: "error on our end getting this school type", Success: false, Detail: err.Error()})
// 	}

// 	return c.Status(204).JSON(Response{
// 		Message: "Successfully deleted school type",
// 		Success: false,
// 		Detail:  err.Error(),
// 	})
// }

// func SchoolCreateController(c *fiber.Ctx) error {
// 	/* This is used for creating school and making the user creating the school as the admin or owner of the schoool
// 	 */
// 	var (
// 		err error
// 	)

// 	var requestBody serializers.SchoolCreateRequestSerializer
// 	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// 	defer cancel()

// 	// Open the database connection
// 	db := database.DBConnection()
// 	defer database.CloseDB(db)

// 	// logged-in user
// 	user := c.Locals("user").(models.User)
// 	validateAdapters := adapters.NewValidate()
// 	fileUpload := adapters.NewFileUpload()

// 	if err := c.BodyParser(&requestBody); err != nil {
// 		return c.Status(400).JSON(Response{Message: "invalid request body", Success: false, Detail: err.Error()})
// 	}

// 	logoUrl, err := fileUpload.UploadFile("logo", c)
// 	if err != nil {
// 		fmt.Println(err)
// 		return c.Status(500).JSON(Response{Message: "server error", Success: false, Detail: err.Error()})
// 	}
// 	documentUrl, err := fileUpload.UploadFile("document", c)
// 	if err != nil {
// 		fmt.Println(err)
// 		return c.Status(500).JSON(Response{Message: "server error", Success: false, Detail: err.Error()})
// 	}

// 	requestBody.Logo = logoUrl
// 	requestBody.Document = documentUrl

// 	// call the validate function in the request serializer for login
// 	if err := validateAdapters.ValidateData(&requestBody); err != nil {
// 		return c.Status(500).JSON(Response{Message: err, Success: false, Detail: err})
// 	}

// 	var school models.School
// 	// // initialize the data for creating the school
// 	err = copier.Copy(&school, &requestBody)
// 	if err != nil {
// 		return c.Status(500).JSON(Response{Message: "Error copying data ", Success: false, Detail: err.Error()})
// 	}

// 	// // generate code
// 	schoolCode := utils.GenerateSchoolCode(db, ctx)

// 	school.SchoolCode = &schoolCode
// 	school.OwnerID = &user.ID

// 	// log.Println("the user id ", user.ID)

// 	err2 := db.WithContext(ctx).Model(school).Preload("Owner").Create(&school).Error
// 	if err2 != nil {
// 		return c.Status(500).JSON(Response{Message: "Error Creating the school", Success: false, Detail: err2})

// 	}

// 	return c.Status(200).JSON(school)
// }

// func SchoolOwnerListCreatedController(c *fiber.Ctx) error {
// 	/* This is used to list the school in which a user is the owner*/
// 	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// 	defer cancel()

// 	// Open the database connection
// 	db := database.DBConnection()
// 	defer database.CloseDB(db)

// 	// logged-in user
// 	user := c.Locals("user").(models.User)

// 	var schools []models.School
// 	var school models.School
// 	err := db.WithContext(ctx).Model(&school).Where("owner_id = ?", user.ID).Find(&schools).Error

// 	if err != nil {
// 		return c.Status(500).JSON(Response{Message: "An error occurred when filtering  for the schools", Success: false, Detail: err})
// 	}

// 	return c.Status(200).JSON(schools)
// }

// func SchoolUpdateController(c *fiber.Ctx) error {
// 	/* this is used to update the school*/
// 	// Parse the request body
// 	var requestBody serializers.SchoolUpdateRequestSerializer

// 	schoolId := c.Params("id")
// 	if schoolId == "" {
// 		return c.Status(400).JSON(
// 			Response{Message: "school id missing", Success: false, Detail: nil})
// 	}

// 	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// 	defer cancel()

// 	// Open the database connection
// 	db := database.DBConnection()
// 	defer database.CloseDB(db)

// 	// logged-in user
// 	user := c.Locals("user").(models.User)
// 	validateAdapters := adapters.NewValidate()

// 	if err := c.BodyParser(&requestBody); err != nil {
// 		return c.Status(400).JSON(Response{Message: "invalid request body", Success: false, Detail: err.Error()})
// 	}
// 	// call the validate function in the request serializer for login
// 	if err := validateAdapters.ValidateData(&requestBody); err != nil {
// 		return c.Status(400).JSON(Response{Message: err, Success: false, Detail: err})
// 	}
// 	var school models.School

// 	err := db.WithContext(ctx).Model(&school).
// 		Where("owner_id = ?", user.ID).
// 		Where("id = ?", schoolId).
// 		First(&school).Error
// 	if err != nil {
// 		return c.Status(404).JSON(Response{Message: "School not found or it does not belong to you ", Success: false, Detail: err.Error()})

// 	}

// 	err = db.WithContext(ctx).Model(&school).Where(&school).Updates(&requestBody).Error

// 	if err != nil {
// 		return c.Status(500).JSON(Response{Message: "error updating the school", Success: false, Detail: err.Error()})

// 	}
// 	return c.Status(200).JSON(school)

// }

// func ClassCreateController(c *fiber.Ctx) error {
// 	/* this is used to create a class for a school in which a school can have multiple classes*/
// 	schoolCode := c.Query("school_code")
// 	var school models.School

// 	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// 	defer cancel()

// 	// Open the database connection
// 	db := database.DBConnection()
// 	defer database.CloseDB(db)

// 	// logged-in user
// 	user := c.Locals("user").(models.User)

// 	validateAdapter := adapters.NewValidate()
// 	// first check if the user has access to create a class
// 	isOk := school.IsSchoolAdminOrOwner(ctx, db, schoolCode, user.ID)
// 	if !isOk {
// 		return c.Status(400).JSON(Response{Message: "You dont have access to create class with this school code", Success: false, Detail: nil})
// 	}
// 	// you have to be the owner of a school before you would be able to create a class for it

// 	school, err := school.RetrieveBySchoolCode(ctx, db, schoolCode)
// 	// if the school does not exist
// 	if err == gorm.ErrRecordNotFound {
// 		return c.Status(404).JSON(Response{Message: "School not found", Success: false, Detail: nil})
// 	}

// 	// Parse the request body
// 	var requestBody serializers.ClassCreateUpdateSerializer
// 	if err := c.BodyParser(&requestBody); err != nil {
// 		return c.Status(400).JSON(Response{Message: "invalid request body", Success: false, Detail: err.Error()})
// 	}
// 	// call the validate function in the request serializer for login
// 	if err2 := validateAdapter.ValidateData(&requestBody); err2 != nil {
// 		return c.Status(400).JSON(Response{Message: err, Success: false, Detail: err.Error()})
// 	}

// 	class := models.Class{
// 		Name:     requestBody.Name,
// 		SchoolID: &school.ID,
// 	}
// 	err = db.WithContext(ctx).Model(&class).Preload("School").Create(&class).First(&class).Error
// 	if err != nil {
// 		return c.Status(500).JSON(Response{Message: "Error creating school", Success: false, Detail: nil})

// 	}

// 	return c.Status(201).JSON(class)
// }

// func ClassUpdateController(c *fiber.Ctx) error {
// 	/* this is used to update a class */
// 	schoolCode := c.Query("school_code")
// 	classID := c.Params("id")
// 	var school models.School
// 	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// 	defer cancel()

// 	// Open the database connection
// 	db := database.DBConnection()
// 	defer database.CloseDB(db)

// 	// logged-in user
// 	user := c.Locals("user").(models.User)
// 	validateAdapter := adapters.NewValidate()

// 	// first check if the user has access to create a class
// 	isOk := school.IsSchoolAdminOrOwner(ctx, db, schoolCode, user.ID)

// 	if !isOk {
// 		return c.Status(400).JSON(Response{Message: "You dont have access to create class with this school code", Success: false, Detail: nil})

// 	}
// 	// you have to be the owner of a school before you would be able to create a class for it

// 	school, err := school.RetrieveBySchoolCode(ctx, db, schoolCode)
// 	// if the school does not exist
// 	if err != nil {
// 		return c.Status(404).JSON(Response{Message: "School not found", Success: false, Detail: err.Error()})
// 	}

// 	var class models.Class
// 	// filter to get the class
// 	err = db.WithContext(ctx).Model(&class).
// 		Where("id = ?", classID).
// 		Where("school_id = ?", school.ID).
// 		First(&class).Error

// 	if err != nil {
// 		return c.Status(404).JSON(Response{Message: "staff with this id does not exists", Success: false, Detail: err.Error()})

// 	}

// 	// Parse the request body
// 	var requestBody serializers.ClassCreateUpdateSerializer
// 	if err := c.BodyParser(&requestBody); err != nil {
// 		return c.Status(400).JSON(Response{Message: "invalid request body", Success: false, Detail: err.Error()})
// 	}
// 	// call the validate function in the request serializer for login
// 	if err2 := validateAdapter.ValidateData(&requestBody); err2 != nil {
// 		return c.Status(400).JSON(Response{Message: err, Success: false, Detail: err.Error()})
// 	}

// 	// update the class
// 	err = db.WithContext(ctx).Model(&class).
// 		Where("id = ?", classID).
// 		Where("school_id = ?", school.ID).
// 		Updates(&requestBody).First(&class).Error

// 	if err != nil {
// 		return c.Status(500).JSON(
// 			Response{Message: "Error updating class", Success: false, Detail: err})
// 	}
// 	return c.Status(200).JSON(class)

// }

// func ClassListController(c *fiber.Ctx) error {
// 	/* this is used to list all the class available on the school  */
// 	var classes []models.Class
// 	var total int64
// 	var school models.School

// 	schoolCode := c.Query("school_code")
// 	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// 	defer cancel()

// 	// Open the database connection
// 	db := database.DBConnection()
// 	defer database.CloseDB(db)

// 	page := c.QueryInt("page", 1)    // default to page 1 if not provided
// 	limit := c.QueryInt("limit", 10) // default to 10 items per page if not provided

// 	school, err := school.RetrieveBySchoolCode(ctx, db, schoolCode)

// 	if err != nil {
// 		return c.Status(404).JSON(Response{
// 			Message: "school with this id does not exists",
// 			Success: false,
// 			Detail:  err,
// 		})
// 	}

// 	db.WithContext(ctx).Model(&models.Class{}).Where("school_id = ?", school.ID).Count(&total)
// 	err = db.Scopes(adapters.NewPaginateAdapter(limit, page).PaginatedResult).WithContext(ctx).Model(&models.Class{}).Where("school_id = ?", school.ID).Find(&classes).Error
// 	if err != nil {
// 		return c.Status(400).JSON(
// 			Response{Message: "school with this id does not exists", Success: false, Detail: err})
// 	}
// 	return c.Status(200).JSON(fiber.Map{"total": total, "data": classes})
// }

// func ClassDeleteController(c *fiber.Ctx) error {
// 	/* this is used to delete a class, and you must be an admin to be able to do this */
// 	schoolCode := c.Query("school_code")
// 	classID := c.Params("id")

// 	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// 	defer cancel()

// 	// get the logged-in user
// 	user := c.Locals("user").(models.User)

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

// 	isOk := school.IsSchoolAdminOrOwner(ctx, db, schoolCode, user.ID)
// 	if !isOk {
// 		return c.Status(400).JSON(Response{Message: "you dont have permission", Success: false, Detail: err.Error()})
// 	}

// 	var class models.Class
// 	err = db.WithContext(ctx).Model(&class).Where("id = ?", classID).Where("school_id = ?", school.ID).First(&class).Delete(&class).Error
// 	if err != nil {
// 		return c.Status(400).JSON(Response{Message: err.Error(), Success: false, Detail: err.Error()})

// 	}

// 	return c.Status(204).JSON(
// 		Response{Message: "Successfully deleted the class", Success: true, Detail: nil})

// }

// func RoleCreateController(c *fiber.Ctx) error {
// 	/* this is used to create a role for users when applying  to job and also when adding the  teachers */
// 	schoolCode := c.Query("school_code")

// 	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// 	defer cancel()

// 	// get the logged-in user
// 	user := c.Locals("user").(models.User)

// 	// Open the database connection
// 	db := database.DBConnection()
// 	defer database.CloseDB(db)

// 	validator := adapters.NewValidate()

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
// 		return c.Status(400).JSON(Response{Message: "you dont have permission", Success: false, Detail: err.Error()})
// 	}

// 	// Parse the request body
// 	var requestBody models.Role
// 	if err := c.BodyParser(&requestBody); err != nil {
// 		return c.Status(400).JSON(Response{Message: "invalid request body", Success: false, Detail: err.Error()})
// 	}
// 	// call the validate function in the request serializer for login
// 	if vErr := validator.ValidateData(requestBody); vErr != nil {
// 		return c.Status(400).JSON(Response{Message: vErr, Success: false, Detail: vErr})
// 	}

// 	// create the role
// 	requestBody.SchoolID = &school.ID
// 	err = db.WithContext(ctx).Model(&requestBody).Create(&requestBody).Error
// 	if err != nil {
// 		return c.Status(500).JSON(Response{Message: "Error creating role", Success: false, Detail: err})

// 	}

// 	return c.Status(201).JSON(requestBody)
// }

// func RoleListController(c *fiber.Ctx) error {
// 	var total int64
// 	var roles []models.Role
// 	var school models.School

// 	schoolCode := c.Query("school_code")

// 	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// 	defer cancel()

// 	// Open the database connection
// 	db := database.DBConnection()
// 	defer database.CloseDB(db)
// 	page := c.QueryInt("page", 1)    // default to page 1 if not provided
// 	limit := c.QueryInt("limit", 10) // default to 10 items per page if not provided

// 	school, err := school.RetrieveBySchoolCode(ctx, db, schoolCode)
// 	if err != nil {
// 		return c.Status(404).JSON(Response{
// 			Message: "school with this id does not exists",
// 			Success: false,
// 			Detail:  err,
// 		})
// 	}

// 	db.WithContext(ctx).Model(&roles).Where("school_id = ?", school.ID).Count(&total)

// 	err = db.Scopes(adapters.NewPaginateAdapter(limit, page).PaginatedResult).WithContext(ctx).Model(&roles).Where("school_id = ?", school.ID).Find(&roles).Error
// 	if err != nil {
// 		if err == gorm.ErrRecordNotFound {
// 			return c.Status(200).JSON(roles)
// 		}
// 		return c.Status(500).JSON(
// 			Response{Message: "An error occured", Success: false, Detail: err})
// 	}

// 	return c.Status(200).JSON(fiber.Map{"total": total, "roles": roles})
// }

// func RoleDetailController(c *fiber.Ctx) error {
// 	schoolCode := c.Query("school_code")
// 	roleID := c.Params("id")

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

// 	var role models.Role
// 	err = db.WithContext(ctx).Model(&role).Where("school_id = ?", school.ID).Where("id = ?", roleID).First(&role).Error
// 	if err != nil {
// 		if err == gorm.ErrRecordNotFound {
// 			return c.Status(404).JSON(Response{Message: "role with this id does not exists", Success: false, Detail: err.Error()})
// 		}
// 		return c.Status(500).JSON(Response{Message: "an error occured ", Success: false, Detail: err.Error()})
// 	}

// 	return c.Status(200).JSON(role)
// }

// func RoleUpdateController(c *fiber.Ctx) error {
// 	schoolCode := c.Query("school_code")
// 	roleID := c.Params("id")

// 	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// 	defer cancel()

// 	// get the logged-in user
// 	user := c.Locals("user").(models.User)
// 	validator := adapters.NewValidate()

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

// 	isOk := school.IsSchoolAdminOrOwner(ctx, db, schoolCode, user.ID)
// 	if !isOk {
// 		return c.Status(400).JSON(Response{Message: "you dont have permission", Success: false, Detail: err.Error()})
// 	}

// 	// Parse the request body
// 	var requestBody models.Role
// 	if err := c.BodyParser(&requestBody); err != nil {
// 		return c.Status(400).JSON(Response{Message: "invalid request body", Success: false, Detail: err.Error()})
// 	}
// 	// call the validate function in the request serializer for login
// 	if vErr := validator.ValidateData(requestBody); vErr != nil {
// 		return c.Status(400).JSON(Response{Message: vErr, Success: false, Detail: vErr})
// 	}

// 	var role models.Role
// 	role = requestBody
// 	role.SchoolID = &school.ID

// 	err = db.WithContext(ctx).Model(&role).Where("school_id = ?", school.ID).Where("id = ?", roleID).Updates(&role).First(&role).Error
// 	if err != nil {
// 		if err == gorm.ErrRecordNotFound {
// 			return c.Status(404).JSON(Response{Message: "role with this id does not exists", Success: false, Detail: err.Error()})

// 		}
// 		return c.Status(400).JSON(Response{Message: "An error occured", Success: false, Detail: err.Error()})

// 	}

// 	return c.Status(200).JSON(role)
// }

// func RoleDeleteController(c *fiber.Ctx) error {
// 	schoolCode := c.Query("school_code")
// 	roleID := c.Params("id")

// 	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// 	defer cancel()

// 	// get the logged-in user
// 	user := c.Locals("user").(models.User)

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

// 	isOk := school.IsSchoolAdminOrOwner(ctx, db, schoolCode, user.ID)
// 	if !isOk {
// 		return c.Status(400).JSON(Response{Message: "you dont have permission", Success: false, Detail: err.Error()})
// 	}

// 	var role models.Role
// 	role.SchoolID = &school.ID

// 	err = db.WithContext(ctx).Model(&role).Where("school_id = ?", school.ID).Where("id = ?", roleID).Delete(&role).Error
// 	if err != nil {
// 		if err == gorm.ErrRecordNotFound {
// 			return c.Status(404).JSON(Response{Message: "role with this id does not exists", Success: false, Detail: err.Error()})
// 		}
// 		return c.Status(500).JSON(
// 			Response{Message: "An error occured", Success: false, Detail: err})
// 	}

// 	return c.Status(204).JSON(Response{Message: "Successfully deleted the role", Success: true, Detail: nil})
// }

// func CategoryCreateController(c *fiber.Ctx) error {
// 	/* this is used to create a role for users when applying  to job and also when adding the  teachers */
// 	schoolCode := c.Query("school_code")

// 	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// 	defer cancel()

// 	// get the logged-in user
// 	user := c.Locals("user").(models.User)
// 	validator := adapters.NewValidate()

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

// 	isOk := school.IsSchoolAdminOrOwner(ctx, db, schoolCode, user.ID)
// 	if !isOk {
// 		return c.Status(400).JSON(Response{Message: "you dont have permission", Success: false, Detail: err.Error()})
// 	}

// 	// Parse the request body
// 	var requestBody models.Category
// 	if err := c.BodyParser(&requestBody); err != nil {
// 		return c.Status(400).JSON(Response{Message: "invalid request body", Success: false, Detail: err.Error()})
// 	}
// 	// call the validate function in the request serializer for login
// 	if vErr := validator.ValidateData(requestBody); vErr != nil {
// 		return c.Status(400).JSON(Response{Message: vErr, Success: false, Detail: vErr})
// 	}

// 	// create the role
// 	requestBody.SchoolID = &school.ID
// 	err = db.WithContext(ctx).Model(&requestBody).Create(&requestBody).Error
// 	if err != nil {
// 		return c.Status(500).JSON(Response{Message: "Error creating role", Success: false, Detail: err.Error()})
// 	}

// 	return c.Status(201).JSON(requestBody)
// }

// func CategoryListController(c *fiber.Ctx) error {
// 	schoolCode := c.Query("school_code")
// 	var categories []models.Category
// 	var total int64
// 	var school models.School

// 	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// 	defer cancel()

// 	// Open the database connection
// 	db := database.DBConnection()
// 	defer database.CloseDB(db)

// 	page := c.QueryInt("page", 1)    // default to page 1 if not provided
// 	limit := c.QueryInt("limit", 10) // default to 10 items per page if not provided

// 	school, err := school.RetrieveBySchoolCode(ctx, db, schoolCode)
// 	if err != nil {
// 		return c.Status(404).JSON(Response{
// 			Message: "school with this id does not exists",
// 			Success: false,
// 			Detail:  err,
// 		})
// 	}

// 	db.WithContext(ctx).Model(&categories).Where("school_id = ?", school.ID).Count(&total)
// 	err = db.Scopes(adapters.NewPaginateAdapter(limit, page).PaginatedResult).WithContext(ctx).Model(&categories).Where("school_id = ?", school.ID).Find(&categories).Error
// 	if err != nil {
// 		if err == gorm.ErrRecordNotFound {
// 			return c.Status(200).JSON(categories)
// 		}
// 		return c.Status(500).JSON(Response{Message: "An error occured", Success: false, Detail: err})

// 	}

// 	return c.Status(200).JSON(fiber.Map{"data": categories, "total": total})
// }

// func CategoryDetailController(c *fiber.Ctx) error {
// 	schoolCode := c.Query("school_code")
// 	categoryID := c.Params("id")

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

// 	var category models.Category
// 	err = db.WithContext(ctx).Model(&category).Where("school_id = ?", school.ID).Where("id = ?", categoryID).First(&category).Error
// 	if err != nil {
// 		if err == gorm.ErrRecordNotFound {

// 			return c.Status(404).JSON(Response{Message: "category with this id does not exists", Success: false, Detail: err})
// 		}
// 		return c.Status(500).JSON(
// 			Response{Message: "An error occured", Success: false, Detail: err})
// 	}

// 	return c.Status(200).JSON(category)
// }

// func CategoryUpdateController(c *fiber.Ctx) error {
// 	schoolCode := c.Query("school_code")
// 	categoryID := c.Params("id")

// 	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// 	defer cancel()

// 	// get the logged-in user
// 	user := c.Locals("user").(models.User)

// 	// Open the database connection
// 	db := database.DBConnection()
// 	defer database.CloseDB(db)
// 	validator := adapters.NewValidate()

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
// 		return c.Status(400).JSON(Response{Message: "you dont have permission", Success: false, Detail: err.Error()})
// 	}

// 	// Parse the request body
// 	var requestBody models.Category
// 	if err := c.BodyParser(&requestBody); err != nil {
// 		return c.Status(400).JSON(Response{Message: "invalid request body", Success: false, Detail: err.Error()})
// 	}
// 	// call the validate function in the request serializer for login
// 	if vErr := validator.ValidateData(requestBody); vErr != nil {
// 		return c.Status(400).JSON(Response{Message: vErr, Success: false, Detail: vErr})
// 	}

// 	var category models.Category
// 	category = requestBody
// 	category.SchoolID = &school.ID

// 	err = db.WithContext(ctx).Model(&category).Where("school_id = ?", school.ID).Where("id = ?", categoryID).Updates(&category).First(&category).Error
// 	if err != nil {
// 		if err == gorm.ErrRecordNotFound {

// 			return c.Status(404).JSON(Response{Message: "category with this id does not exists", Success: false, Detail: err.Error()})
// 		}
// 		return c.Status(500).JSON(
// 			Response{Message: "An error occured", Success: false, Detail: err})
// 	}

// 	return c.Status(200).JSON(category)
// }

// func CategoryDeleteController(c *fiber.Ctx) error {
// 	schoolCode := c.Query("school_code")
// 	categoryID := c.Params("id")

// 	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// 	defer cancel()

// 	// get the logged-in user
// 	user := c.Locals("user").(models.User)

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

// 	isOk := school.IsSchoolAdminOrOwner(ctx, db, schoolCode, user.ID)
// 	if !isOk {
// 		return c.Status(400).JSON(Response{Message: "you dont have permission", Success: false, Detail: err.Error()})
// 	}

// 	var category models.Category
// 	category.SchoolID = &school.ID

// 	err = db.WithContext(ctx).Model(&category).Where("school_id = ?", school.ID).Where("id = ?", categoryID).First(&category).Delete(&category).Error
// 	if err != nil {
// 		if err == gorm.ErrRecordNotFound {
// 			return c.Status(404).JSON(Response{Message: "category with this id does not exists", Success: false, Detail: err.Error()})
// 		}
// 		return c.Status(500).JSON(Response{Message: "An error occured", Success: false, Detail: err.Error()})
// 	}

// 	return c.Status(204).JSON(Response{Message: "Successfully deleted the category", Success: true, Detail: nil})
// }

// func SchoolAdminListController(c *fiber.Ctx) error {
// 	/* Used to list all the admins available in the school */
// 	var school models.School
// 	var admins []models.Admin
// 	var total int64

// 	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// 	defer cancel()

// 	// Open the database connection
// 	db := database.DBConnection()
// 	defer database.CloseDB(db)

// 	page := c.QueryInt("page", 1)    // default to page 1 if not provided
// 	limit := c.QueryInt("limit", 10) // default to 10 items per page if not provided

// 	schoolCode := c.Query("school_code")

// 	user := c.Locals("user").(models.User)
// 	school, err := school.RetrieveBySchoolCode(ctx, db, schoolCode)
// 	if err != nil {
// 		return c.Status(400).JSON(
// 			Response{Message: "Invalid school code ", Success: false, Detail: err})
// 	}
// 	// check if the user is the owner of the school
// 	if *school.OwnerID != user.ID {
// 		return c.Status(400).JSON(Response{Message: "you dont have permission", Success: false, Detail: err.Error()})
// 	}

// 	db.Model(&models.Admin{}).Where("school_id = ?", school.ID).Count(&total)
// 	err = db.Scopes(adapters.NewPaginateAdapter(limit, page).PaginatedResult).
// 		Model(&models.Admin{}).
// 		Where("school_id = ?", school.ID).
// 		Preload("User").
// 		Find(&admins).Error
// 	if err != nil {
// 		logger.Error(ctx, "an error occurred retrieving the school  admin ", zap.Error(err))
// 		return c.Status(400).JSON(Response{Message: "An error occured retrieving school admins", Success: false, Detail: err.Error()})
// 	}

// 	return c.Status(200).JSON(fiber.Map{"total": total, "data": admins})
// }

// func SchoolAdminAddController(c *fiber.Ctx) error {
// 	/* This is used to add user to the school admin */
// 	var school models.School
// 	var admin models.Admin
// 	var foundUser models.User
// 	var requestBody serializers.SchoolAdminCreateRequestSerializer

// 	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// 	defer cancel()

// 	// Open the database connection
// 	db := database.DBConnection()
// 	defer database.CloseDB(db)

// 	validate := adapters.NewValidate()
// 	schoolCode := c.Query("school_code")

// 	err := c.BodyParser(&requestBody)
// 	if err != nil {
// 		return c.Status(400).JSON(Response{Message: "Invalid data passed", Success: false, Detail: err})

// 	}
// 	vErr := validate.ValidateData(&requestBody)
// 	if vErr != nil {
// 		return c.Status(400).JSON(Response{Message: vErr, Success: false, Detail: vErr})
// 	}

// 	//	check if it is a valid school
// 	user := c.Locals("user").(models.User)
// 	school, err = school.RetrieveBySchoolCode(ctx, db, schoolCode)
// 	if err != nil {
// 		return c.Status(400).JSON(
// 			Response{Message: "Invalid school code ", Success: false, Detail: err})
// 	}
// 	// check if the user is the owner of the school
// 	if *school.OwnerID != user.ID {
// 		return c.Status(400).JSON(Response{Message: "you dont have permission", Success: false, Detail: err.Error()})
// 	}

// 	//	check if the user exists
// 	err = db.WithContext(ctx).Model(&foundUser).Where("id = ?", requestBody.UserID).First(&foundUser).Error
// 	if err != nil {
// 		if err == gorm.ErrRecordNotFound {

// 			return c.Status(400).JSON(Response{Message: "User does not exist with this id ", Success: false, Detail: err})
// 		}
// 		logger.Error(ctx, "An error occurred filtering user on admin add to check user exists", zap.Error(err))
// 		return c.Status(500).JSON(Response{Message: "an error occurred we are working on it", Success: false, Detail: err.Error()})
// 	}

// 	// create the admin
// 	status := "ACTIVE"
// 	admin.Status = &status
// 	admin.UserID = &foundUser.ID
// 	admin.SchoolID = &school.ID
// 	err = db.WithContext(ctx).Model(&admin).Create(&admin).Preload("School").Preload("User").First(&admin).Error
// 	if err != nil {
// 		logger.Error(ctx, "An error occured creating the admin ", zap.Error(err))
// 		return c.Status(500).JSON(Response{Message: "An error occurred we are working on it", Success: false, Detail: err})
// 	}

// 	return c.Status(200).JSON(admin)
// }

// func SchoolAdminDetailController(c *fiber.Ctx) error {
// 	/* this is used to retrieve the detail of the admin */
// 	var school models.School
// 	var admin models.Admin

// 	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// 	defer cancel()

// 	// Open the database connection
// 	db := database.DBConnection()
// 	defer database.CloseDB(db)

// 	schoolCode := c.Query("school_code")

// 	adminID, err := uuid.Parse(c.Params("id"))
// 	if err != nil {
// 		return c.Status(400).JSON(
// 			Response{Message: "Invalid id passed", Success: false, Detail: err})
// 	}

// 	//	check if it is a valid school
// 	user := c.Locals("user").(models.User)
// 	school, err = school.RetrieveBySchoolCode(ctx, db, schoolCode)
// 	if err != nil {
// 		return c.Status(400).JSON(
// 			Response{Message: "Invalid school code ", Success: false, Detail: err})
// 	}
// 	// check if the user is the owner of the school
// 	if *school.OwnerID != user.ID {
// 		return c.Status(400).JSON(Response{Message: "you dont have permission", Success: false, Detail: err.Error()})
// 	}

// 	err = db.WithContext(ctx).Model(&admin).Where("id =?", adminID).Preload("School").Preload("User").First(&admin).Error
// 	if err != nil {
// 		if err == gorm.ErrRecordNotFound {
// 			return c.Status(400).JSON(Response{Message: "admin with this  id does not exist", Success: false, Detail: err.Error()})
// 		}
// 		return c.Status(400).JSON(Response{Message: "an error occurred on our end we are working on it", Success: false, Detail: err.Error()})
// 	}

// 	return c.Status(200).JSON(admin)
// }

// func SchoolAdminUpdateController(c *fiber.Ctx) error {
// 	/* this is used to update school admin */
// 	var school models.School
// 	var admin models.Admin
// 	var requestBody serializers.SchoolAdminUpdateRequestSerializer

// 	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// 	defer cancel()

// 	// Open the database connection
// 	db := database.DBConnection()
// 	defer database.CloseDB(db)

// 	validate := adapters.NewValidate()
// 	schoolCode := c.Query("school_code")
// 	adminID, err := uuid.Parse(c.Params("id"))
// 	if err != nil {
// 		return c.Status(400).JSON(
// 			Response{Message: "Invalid id passed", Success: false, Detail: err})
// 	}

// 	err = c.BodyParser(&requestBody)
// 	if err != nil {
// 		return c.Status(400).JSON(Response{Message: "Invalid data passed", Success: false, Detail: err})
// 	}
// 	vErr := validate.ValidateData(&requestBody)
// 	if vErr != nil {
// 		return c.Status(400).JSON(Response{Message: vErr, Success: false, Detail: vErr})
// 	}

// 	//	check if it is a valid school
// 	user := c.Locals("user").(models.User)
// 	school, err = school.RetrieveBySchoolCode(ctx, db, schoolCode)
// 	if err != nil {
// 		return c.Status(400).JSON(
// 			Response{Message: "Invalid school code ", Success: false, Detail: err})
// 	}
// 	// check if the user is the owner of the school
// 	if *school.OwnerID != user.ID {
// 		return c.Status(400).JSON(Response{Message: "you dont have permission", Success: false, Detail: err.Error()})
// 	}

// 	err = db.WithContext(ctx).Model(&admin).Where("id = ?", adminID).Update("status", requestBody.Status).
// 		Preload("School").Preload("User").First(&admin).Error

// 	if err != nil {
// 		if err == gorm.ErrRecordNotFound {
// 			return c.Status(400).JSON(Response{Message: "Admin with this id does not exists", Success: false, Detail: err.Error()})
// 		}
// 		logger.Error(ctx, "An error occurred updating the admin status ", zap.Error(err))
// 		return c.Status(400).JSON(Response{Message: "Error updating the admin status ", Success: false, Detail: err.Error()})
// 	}
// 	return c.Status(200).JSON(admin)
// }

// func EventCreateController(c *fiber.Ctx) error {
// 	/* this is used to create event */
// 	var requestBody serializers.EventCreateRequestSerializer
// 	var school models.School
// 	var staff models.Staff
// 	var event models.Event

// 	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// 	defer cancel()

// 	// Open the database connection
// 	db := database.DBConnection()
// 	defer database.CloseDB(db)

// 	// get the logged-in user
// 	user := c.Locals("user").(models.User)
// 	validator := adapters.NewValidate()

// 	schoolCode := c.Query("school_code")
// 	school, err := school.RetrieveBySchoolCode(ctx, db, schoolCode)
// 	if err != nil {
// 		return c.Status(404).JSON(Response{
// 			Message: "school with this id does not exists",
// 			Success: false,
// 			Detail:  err.Error(),
// 		})
// 	}

// 	if err := c.BodyParser(&requestBody); err != nil {
// 		return c.Status(400).JSON(Response{Message: "invalid request body", Success: false, Detail: err.Error()})
// 	}
// 	// call the validate function in the request serializer for login
// 	if vErr := validator.ValidateData(requestBody); vErr != nil {
// 		return c.Status(400).JSON(Response{Message: vErr, Success: false, Detail: vErr})
// 	}

// 	// validate the uuids for the staff
// 	for _, participantID := range requestBody.EventsParticipants {
// 		staff, err = staff.RetrieveByIDAndSchool(ctx, db, *participantID, school.ID)
// 		if err != nil {
// 			return c.Status(400).JSON(Response{Message: "invalid staff participant provided", Success: false, Detail: err.Error()})
// 		}

// 	}
// 	isOk := school.IsSchoolAdminOrOwner(ctx, db, schoolCode, user.ID)
// 	if !isOk {
// 		return c.Status(400).JSON(Response{Message: "You dont have access to create class with this school code",
// 			Success: false, Detail: nil})
// 	}

// 	event = models.Event{
// 		SchoolID:    &school.ID,
// 		Time:        requestBody.Time,
// 		Title:       requestBody.Title,
// 		Description: requestBody.Description,
// 		EventType:   requestBody.EventType,
// 	}

// 	err = db.WithContext(ctx).Model(&models.Event{}).Create(&event).First(&event).Error
// 	if err != nil {
// 		return c.Status(500).JSON(Response{Message: "error occured creating event on our end", Success: false, Detail: err.Error()})
// 	}
// 	// validate the uuids for the staff
// 	for _, participantID := range requestBody.EventsParticipants {
// 		// Add the staff participant to the EventsParticipants slice
// 		err = db.WithContext(ctx).Model(&models.EventParticipant{}).Create(&models.EventParticipant{
// 			EventID: &event.ID,
// 			StaffID: participantID,
// 		}).Error
// 		if err != nil {
// 			return c.Status(500).JSON(Response{Message: "error creating event participant", Success: false, Detail: err.Error()})
// 		}
// 	}

// 	err = db.WithContext(ctx).Model(&models.Event{}).Where("id =?", event.ID).
// 		Preload("EventsParticipants").
// 		Preload("EventsParticipants.User").First(&event).Error
// 	if err != nil {
// 		return c.Status(400).JSON(Response{Message: "error getting newly create event with participants", Success: false, Detail: err.Error()})
// 	}

// 	return c.Status(201).JSON(event)
// }

// func EventListController(c *fiber.Ctx) error {
// 	var school models.School
// 	var events []models.EventParticipant
// 	var total int64

// 	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// 	defer cancel()

// 	// Open the database connection
// 	db := database.DBConnection()
// 	defer database.CloseDB(db)
// 	// get the logged-in user
// 	_ = c.Locals("user").(models.User)

// 	page := c.QueryInt("page", 1)    // default to page 1 if not provided
// 	limit := c.QueryInt("limit", 10) // default to 10 items per page if not provided

// 	schoolCode := c.Query("school_code")
// 	school, err := school.RetrieveBySchoolCode(ctx, db, schoolCode)
// 	if err != nil {
// 		return c.Status(404).JSON(Response{
// 			Message: "school with this code does not exists",
// 			Success: false,
// 			Detail:  err.Error(),
// 		})
// 	}

// 	err = db.WithContext(ctx).Model(&models.Event{}).
// 		Where("school_id =?", school.ID).
// 		Count(&total).Error
// 	if err != nil {
// 		return c.Status(500).JSON(Response{Message: "error getting total", Detail: err.Error(), Success: false})
// 	}
// 	err = db.Scopes(adapters.NewPaginateAdapter(limit, page).PaginatedResult).WithContext(ctx).Model(&models.Event{}).
// 		Where("school_id =?", school.ID).Find(&events).Error
// 	if err != nil {
// 		return c.Status(500).JSON(Response{Message: "error getting events on our end", Detail: err.Error(), Success: false})
// 	}

// 	return c.Status(200).JSON(fiber.Map{"total": total, "data": events})

// }

// func UpdateAnnouncementConfigurationController(c *fiber.Ctx) error {
// 	var requestBody serializers.AnnouncementConfigurationRequestSerializer
// 	var school models.School
// 	var announcementConfiguration models.AnnouncementConfiguration
// 	var foundAnnouncementConfiguration models.AnnouncementConfiguration

// 	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// 	defer cancel()

// 	// Open the database connection
// 	db := database.DBConnection()
// 	defer database.CloseDB(db)
// 	// get the logged-in user
// 	user := c.Locals("user").(models.User)
// 	validator := adapters.NewValidate()

// 	schoolCode := c.Query("school_code")
// 	school, err := school.RetrieveBySchoolCode(ctx, db, schoolCode)
// 	if err != nil {
// 		return c.Status(404).JSON(Response{
// 			Message: "school with this id does not exists",
// 			Success: false,
// 			Detail:  err.Error(),
// 		})
// 	}

// 	if err := c.BodyParser(&requestBody); err != nil {
// 		return c.Status(400).JSON(Response{Message: "invalid request body", Success: false, Detail: err.Error()})
// 	}
// 	// call the validate function in the request serializer for login
// 	if vErr := validator.ValidateData(requestBody); vErr != nil {
// 		return c.Status(400).JSON(Response{Message: vErr, Success: false, Detail: vErr})
// 	}

// 	isOk := school.IsSchoolAdminOrOwner(ctx, db, schoolCode, user.ID)
// 	if !isOk {
// 		return c.Status(400).JSON(Response{Message: "You dont have access to create class with this school code",
// 			Success: false, Detail: nil})
// 	}

// 	announcementConfiguration = models.AnnouncementConfiguration{
// 		SchoolID:          &school.ID,
// 		ToRecipientPortal: requestBody.ToRecipientPortal,
// 		ToEmail:           requestBody.ToEmail,
// 		ReceiveCopy:       requestBody.ReceiveCopy,
// 		SendFromOrder:     requestBody.SendFromOrder,
// 		AnnouncementTag:   requestBody.AnnouncementTag,
// 		NotificationType:  requestBody.NotificationType,
// 		Layout:            requestBody.Layout,
// 	}

// 	err = db.WithContext(ctx).Model(&models.AnnouncementConfiguration{}).Where("school_id =?", school.ID).First(&foundAnnouncementConfiguration).Error
// 	if err != nil {
// 		if err == gorm.ErrRecordNotFound {
// 			err = db.WithContext(ctx).Model(&models.AnnouncementConfiguration{}).Create(&announcementConfiguration).Error
// 			if err != nil {
// 				return c.Status(500).JSON(Response{Message: "error creating announcement configuration ", Success: false, Detail: err.Error()})
// 			}
// 		}
// 	} else {
// 		err = db.WithContext(ctx).Model(models.AnnouncementConfiguration{}).Updates(&announcementConfiguration).Error
// 	}

// 	return c.Status(201).JSON(announcementConfiguration)

// }

// func UpdateNotificationConfigurationController(c *fiber.Ctx) error {
// 	var requestBody serializers.NotificationConfigurationRequestSerializer
// 	var school models.School
// 	var notificationConfiguration models.NotificationConfiguration
// 	var foundNotificationConfiguration models.NotificationConfiguration

// 	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// 	defer cancel()

// 	// Open the database connection
// 	db := database.DBConnection()
// 	defer database.CloseDB(db)
// 	// get the logged-in user
// 	user := c.Locals("user").(models.User)
// 	validator := adapters.NewValidate()

// 	schoolCode := c.Query("school_code")
// 	school, err := school.RetrieveBySchoolCode(ctx, db, schoolCode)
// 	if err != nil {
// 		return c.Status(404).JSON(Response{
// 			Message: "school with this id does not exists",
// 			Success: false,
// 			Detail:  err.Error(),
// 		})
// 	}

// 	if err := c.BodyParser(&requestBody); err != nil {
// 		return c.Status(400).JSON(Response{Message: "invalid request body", Success: false, Detail: err.Error()})
// 	}
// 	// call the validate function in the request serializer for login
// 	if vErr := validator.ValidateData(requestBody); vErr != nil {
// 		return c.Status(400).JSON(Response{Message: vErr, Success: false, Detail: vErr})
// 	}

// 	isOk := school.IsSchoolAdminOrOwner(ctx, db, schoolCode, user.ID)
// 	if !isOk {
// 		return c.Status(400).JSON(Response{Message: "You dont have access to create class with this school code",
// 			Success: false, Detail: nil})
// 	}

// 	notificationConfiguration = models.NotificationConfiguration{
// 		SchoolID:          &school.ID,
// 		ToRecipientPortal: requestBody.ToRecipientPortal,
// 		ToEmail:           requestBody.ToEmail,
// 		ReceiveCopy:       requestBody.ReceiveCopy,
// 		SendFromOrder:     requestBody.SendFromOrder,
// 		AnnouncementTag:   requestBody.AnnouncementTag,
// 		NotificationType:  requestBody.NotificationType,
// 		Layout:            requestBody.Layout,
// 		AttendanceStatus:  requestBody.AttendanceStatus,
// 		Description:       requestBody.Description,
// 	}

// 	err = db.WithContext(ctx).Model(&models.NotificationConfiguration{}).Where("school_id =?", school.ID).First(&foundNotificationConfiguration).Error
// 	if err != nil {
// 		if err == gorm.ErrRecordNotFound {
// 			err = db.WithContext(ctx).Model(&models.NotificationConfiguration{}).Create(&notificationConfiguration).Error
// 			if err != nil {
// 				return c.Status(500).JSON(Response{Message: "error creating notification configuration ", Success: false, Detail: err.Error()})
// 			}
// 		}
// 	} else {
// 		err = db.WithContext(ctx).Model(models.NotificationConfiguration{}).Updates(&notificationConfiguration).Error
// 	}

// 	return c.Status(201).JSON(notificationConfiguration)

// }

// func SchoolInviteCreateController(c *fiber.Ctx) error {
// 	var requestBody serializers.SchoolInviteCreateRequestSerializer
// 	var schoolInvite models.SchoolInvite
// 	var school models.School

// 	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// 	defer cancel()

// 	// Open the database connection
// 	db := database.DBConnection()
// 	defer database.CloseDB(db)

// 	schoolCode := c.Query("school_code")
// 	user := c.Locals("user").(models.User)

// 	validateAdapter := adapters.NewValidate()
// 	pointerAdapters := adapters.NewPointer()

// 	// first check if the user has access to create a class
// 	isOk := school.IsSchoolAdminOrOwner(ctx, db, schoolCode, user.ID)
// 	if !isOk {
// 		return c.Status(400).JSON(Response{Message: "You dont have access to create class with this school code", Success: false, Detail: nil})
// 	}
// 	// you have to be the owner of a school before you would be able to create a class for it

// 	school, err := school.RetrieveBySchoolCode(ctx, db, schoolCode)
// 	// if the school does not exist
// 	if err == gorm.ErrRecordNotFound {
// 		return c.Status(404).JSON(Response{Message: "School not found", Success: false, Detail: nil})
// 	}

// 	err = c.BodyParser(&requestBody)
// 	if err != nil {
// 		return c.Status(400).JSON(Response{Message: "Invalid json passed", Success: false, Detail: err.Error()})
// 	}

// 	errors := validateAdapter.ValidateData(&requestBody)
// 	if errors != nil {
// 		return c.Status(400).JSON(Response{Message: errors, Success: false, Detail: errors})
// 	}

// 	err = copier.Copy(&schoolInvite, requestBody)
// 	if err != nil {
// 		return c.Status(500).JSON(Response{Message: "error on our end copying data", Success: false, Detail: err.Error()})
// 	}

// 	schoolInvite.SchoolID = &school.ID
// 	schoolInvite.Status = pointerAdapters.StringPointer("PENDING")

// 	err = db.WithContext(ctx).Model(&schoolInvite).Where("email ILIKE ?", schoolInvite.Email).Where("school_id = ?", schoolInvite.SchoolID).Delete(&models.SchoolInvite{}).Error
// 	if err != nil && err != gorm.ErrRecordNotFound {
// 		return c.Status(500).JSON(Response{Message: "Error on our end deleting existing school invite of the user ", Success: false, Detail: err.Error()})
// 	}

// 	err = db.WithContext(ctx).Model(&models.SchoolInvite{}).Create(&schoolInvite).First(&schoolInvite).Error
// 	if err != nil {
// 		return c.Status(400).JSON(Response{Message: "Error creating school", Success: false, Detail: err.Error()})
// 	}

// 	return c.Status(200).JSON(schoolInvite)
// }

// func SchoolInviteListController(c *fiber.Ctx) error {
// 	var schoolInvites []models.SchoolInvite
// 	var school models.School

// 	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// 	defer cancel()

// 	// Open the database connection
// 	db := database.DBConnection()
// 	defer database.CloseDB(db)

// 	schoolCode := c.Query("school_code")
// 	user := c.Locals("user").(models.User)

// 	// first check if the user has access to create a class
// 	isOk := school.IsSchoolAdminOrOwner(ctx, db, schoolCode, user.ID)
// 	if !isOk {
// 		return c.Status(400).JSON(Response{Message: "You dont have access to create class with this school code", Success: false, Detail: nil})
// 	}
// 	// you have to be the owner of a school before you would be able to create a class for it

// 	school, err := school.RetrieveBySchoolCode(ctx, db, schoolCode)
// 	// if the school does not exist
// 	if err == gorm.ErrRecordNotFound {
// 		return c.Status(404).JSON(Response{Message: "School not found", Success: false, Detail: nil})
// 	}

// 	err = db.WithContext(ctx).Model(&models.SchoolInvite{}).Where("school_id = ?", school.ID).Find(&schoolInvites).Error
// 	if err != nil {
// 		return c.Status(500).JSON(Response{Message: "error getting school invites from our end", Success: false, Detail: err.Error()})
// 	}

// 	return c.Status(200).JSON(schoolInvites)
// }

// /*
// func UserSelectSchoolRoleController(c *fiber.Ctx) error {
// 	var staff models.Staff
// 	var school models.Schoo

// 	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// 	defer cancel()

// 	// Open the database connection
// 	db := database.DBConnection()
// 	defer database.CloseDB(db)

// 	// logged-in user
// 	user := c.Locals("user").(models.User)

// 	db.WithContext(ctx).Model(&models.SchoolInvite{}).Where("email ILIKE ?",user.Email).First()

// }
// */
