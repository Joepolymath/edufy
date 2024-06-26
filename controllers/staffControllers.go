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
// 	"github.com/google/uuid"
// 	"go.uber.org/zap"
// 	"gorm.io/gorm"
// 	"strings"
// 	"time"
// )

// func StaffListController(c *fiber.Ctx) error {
// 	var staffs []models.Staff
// 	var totalStaffs int64
// 	var school models.School

// 	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// 	defer cancel()

// 	// Open the database connection
// 	db := database.DBConnection()
// 	defer database.CloseDB(db)

// 	schoolCode := c.Query("school_code")
// 	search := c.Query("search")
// 	status := c.Query("status")
// 	if status != "" {
// 		status = strings.ToUpper(status)
// 	}

// 	page := c.QueryInt("page", 1)    // Page number, default to 1
// 	limit := c.QueryInt("limit", 10) // Number of records per page, default to 10

// 	school, err := school.RetrieveBySchoolCode(ctx, db, schoolCode)
// 	if err != nil {
// 		return c.Status(400).JSON(Response{Message: "School with this code does not exist", Success: false, Detail: err.Error()})

// 	}

// 	// Get the logged-in user
// 	user := c.Locals("user").(models.User)

// 	isOk := school.IsSchoolAdminOrOwner(ctx, db, schoolCode, user.ID)
// 	if !isOk {
// 		return c.Status(400).JSON(
// 			Response{Message: "you dont have permission", Success: false, Detail: err})
// 	}

// 	// Calculate the offset based on the page and limit
// 	offset := (page - 1) * limit

// 	// Query for paginated staffs and count total staffs
// 	query := db.WithContext(ctx).
// 		Model(&models.Staff{}).
// 		Where("school_id = ?", school.ID).
// 		Preload("Classes").
// 		Preload("Role").
// 		Preload("User").
// 		Offset(offset).
// 		Limit(limit)

// 	// Check if the status is provided
// 	if status != "" {
// 		// Add status condition when it's not empty
// 		query = query.Where("status = ?", status)
// 	}

// 	if search != "" {
// 		// Add search condition for first name or last name
// 		query = query.Joins("JOIN users ON staffs.user_id = users.id").
// 			Where("users.first_name ILIKE ? OR users.last_name ILIKE ?", "%"+search+"%", "%"+search+"%")
// 	}
// 	// Run the query
// 	err = query.Find(&staffs).Error

// 	if err != nil {
// 		if err != gorm.ErrRecordNotFound {
// 			return c.Status(500).JSON(Response{Message: "Error on our end filtering staffs %s", Success: false, Detail: err.Error()})
// 		}
// 		return c.Status(400).JSON(Response{Message: "Error on our end filtering staffs %s", Success: false, Detail: err.Error()})
// 	}

// 	// Count total staffs
// 	err = db.Model(&models.Staff{}).Where("school_id = ?", school.ID).Count(&totalStaffs).Error
// 	if err != nil {
// 		return c.Status(500).JSON(Response{
// 			Message: "error getting staff counts", Detail: err, Success: false,
// 		})
// 	}

// 	return c.Status(200).JSON(fiber.Map{"total": totalStaffs, "data": staffs})
// }

// func StaffUpdateController(c *fiber.Ctx) error {
// 	schoolCode := c.Query("school_code")
// 	staffID := c.Params("id")

// 	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// 	defer cancel()

// 	// Open the database connection
// 	db := database.DBConnection()
// 	defer database.CloseDB(db)
// 	validator := adapters.NewValidate()

// 	var school models.School
// 	school, err := school.RetrieveBySchoolCode(ctx, db, schoolCode)
// 	if err != nil {
// 		return c.Status(400).JSON(Response{Message: "School with this code does not exist", Success: false, Detail: err.Error()})

// 	}

// 	// Get the logged-in user
// 	user := c.Locals("user").(models.User)

// 	isOk := school.IsSchoolAdminOrOwner(ctx, db, schoolCode, user.ID)
// 	if !isOk {
// 		return c.Status(400).JSON(
// 			Response{Message: "you dont have permission", Success: false, Detail: err})
// 	}

// 	var requestBody serializers.StaffUpdateSerializer
// 	err = c.BodyParser(&requestBody)
// 	if err != nil {
// 		return c.Status(400).JSON(
// 			Response{Message: "Error parsing request body", Success: false, Detail: err})
// 	}

// 	// validate the data
// 	err2 := validator.ValidateData(&requestBody)
// 	if err2 != nil {
// 		return c.Status(400).JSON(Response{Message: err2, Success: false, Detail: err2})

// 	}

// 	// validate the data's in the database
// 	err = requestBody.ValidateData(db, ctx)
// 	if err != nil {
// 		return c.Status(404).JSON(Response{Message: err.Error(), Success: false, Detail: err.Error()})
// 	}

// 	// update the staff
// 	var staff models.Staff
// 	err = db.WithContext(ctx).Model(&staff).
// 		Where("id = ? AND school_id = ?", staffID, school.ID).
// 		Preload("User").Preload("Category").Preload("Role").Preload("Classes").
// 		Updates(&requestBody).First(&staff).Error
// 	if err != nil {
// 		if err == gorm.ErrRecordNotFound {
// 			return c.Status(404).JSON(Response{Message: "staff with this id does not exists", Success: false, Detail: err.Error()})
// 		}
// 		return c.Status(404).JSON(Response{Message: "staff with this id does not exists", Success: false, Detail: err.Error()})
// 	}

// 	return c.Status(200).JSON(staff)
// }

// func StaffDetailController(c *fiber.Ctx) error {
// 	var staff models.Staff

// 	schoolCode := c.Query("school_code")
// 	staffID := c.Params("id")

// 	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// 	defer cancel()

// 	// Open the database connection
// 	db := database.DBConnection()
// 	defer database.CloseDB(db)

// 	var school models.School
// 	school, err := school.RetrieveBySchoolCode(ctx, db, schoolCode)
// 	if err != nil {
// 		return c.Status(400).JSON(Response{Message: "School with this code does not exist", Success: false, Detail: err.Error()})

// 	}

// 	// Get the logged-in user
// 	user := c.Locals("user").(models.User)

// 	isOk := school.IsSchoolAdminOrOwner(ctx, db, schoolCode, user.ID)
// 	if !isOk {
// 		return c.Status(400).JSON(
// 			Response{
// 				Message: "you dont have permission",
// 				Success: false,
// 				Detail:  err,
// 			})
// 	}

// 	err = db.WithContext(ctx).Model(&models.Staff{}).Where("id = ? AND school_id = ?", staffID, school.ID).
// 		Preload("User").Preload("Category").Preload("Role").Preload("Classes").First(&staff).Error
// 	if err != nil {
// 		return c.Status(400).JSON(Response{Message: "Staff with this ID does not exist", Success: false, Detail: err.Error()})

// 	}

// 	return c.Status(200).JSON(staff)
// }

// func StaffRatingCreateController(c *fiber.Ctx) error {
// 	var requestBody serializers.StaffRatingCreateRequestSerializer
// 	var school models.School
// 	var staff models.Staff
// 	var staffRating models.StaffRating

// 	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// 	defer cancel()

// 	// Open the database connection
// 	db := database.DBConnection()
// 	defer database.CloseDB(db)

// 	validateAdapters := adapters.NewValidate()

// 	user := c.Locals("user").(models.User)
// 	schoolCode := c.Query("school_code")

// 	err := c.BodyParser(&requestBody)
// 	if err != nil {
// 		return c.Status(400).JSON(Response{Message: "invalid request body", Success: false, Detail: err.Error()})
// 	}

// 	vErr := validateAdapters.ValidateData(&requestBody)
// 	if vErr != nil {
// 		return c.Status(400).JSON(Response{Message: vErr, Success: false, Detail: vErr})

// 	}

// 	school, err = school.RetrieveBySchoolCode(ctx, db, schoolCode)
// 	if err != nil {
// 		return c.Status(400).JSON(Response{Message: "School with this code does not exist", Success: false, Detail: err.Error()})

// 	}
// 	staff, err = staff.Retrieve(ctx, db, *requestBody.StaffID)
// 	if err != nil {
// 		return c.Status(400).JSON(Response{Message: "School with this code does not exist", Success: false, Detail: err.Error()})

// 	}

// 	staffRating = models.StaffRating{
// 		SchoolID: &school.ID,
// 		StaffID:  &staff.ID,
// 		UserID:   &user.ID,
// 		Message:  requestBody.Message,
// 		Rating:   requestBody.Rating,
// 	}

// 	err = db.WithContext(ctx).Model(&staffRating).
// 		Where("school_id =?", school.ID).
// 		Where("staff_id =?", staff.ID).
// 		Where("user_id =?", user.ID).Preload("User").First(&staffRating).Error
// 	if staffRating.ID != uuid.Nil {
// 		return c.Status(400).JSON(Response{Message: "you are not allowed to rate staff twice", Success: false, Detail: err})
// 	}

// 	err = db.WithContext(ctx).Model(&staffRating).Create(&staffRating).Preload("Staff").First(&staffRating).Error
// 	if err != nil {
// 		logger.Error(ctx, "error creating staff rating ", zap.Error(err))
// 		return c.Status(400).JSON(Response{Message: "an error occured creating staff rating", Success: false, Detail: err})

// 	}

// 	return c.Status(201).JSON(staffRating)
// }

// func StaffRatingListController(c *fiber.Ctx) error {
// 	var staffRating []models.StaffRating
// 	var school models.School
// 	var staff models.Staff

// 	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// 	defer cancel()

// 	// Open the database connection
// 	db := database.DBConnection()
// 	defer database.CloseDB(db)

// 	_ = c.Locals("user").(models.User)
// 	schoolCode := c.Query("school_code")
// 	staffID, err := uuid.Parse(c.Params("id"))
// 	if err != nil {
// 		return c.Status(400).JSON(Response{Message: "invalid staff id", Success: false, Detail: err.Error()})
// 	}

// 	school, err = school.RetrieveBySchoolCode(ctx, db, schoolCode)
// 	if err != nil {
// 		return c.Status(400).JSON(Response{Message: "School with this code does not exist", Success: false, Detail: err.Error()})

// 	}

// 	staff, err = staff.Retrieve(ctx, db, staffID)
// 	if err != nil {
// 		return c.Status(400).JSON(
// 			Response{Message: "staff with this id doesnt not exist", Success: false, Detail: err})
// 	}

// 	err = db.WithContext(ctx).Model(&models.StaffRating{}).Where("staff_id =?", staff.ID).Limit(10).Order("rating desc").Find(&staffRating).Error
// 	if err != nil {
// 		logger.Error(ctx, "error getting staff rating ", zap.Error(err))
// 		return c.Status(400).JSON(
// 			Response{Message: "Error getting staff rating ", Success: false, Detail: err})
// 	}

// 	return c.Status(200).JSON(staffRating)
// }

// func StaffsAnalyticController(c *fiber.Ctx) error {
// 	var staffRating []models.StaffRating
// 	var school models.School
// 	var serializer serializers.StaffsAnalyticsResponseSerializer
// 	var activeStaffCount int64
// 	var inActiveStaffCount int64
// 	var maleCount int64
// 	var femaleCount int64

// 	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// 	defer cancel()

// 	// Open the database connection
// 	db := database.DBConnection()
// 	defer database.CloseDB(db)

// 	staffUtils := utils.NewStaffUtils()

// 	user := c.Locals("user").(models.User)
// 	schoolCode := c.Query("school_code")

// 	school, err := school.RetrieveBySchoolCode(ctx, db, schoolCode)
// 	if err != nil {
// 		return c.Status(400).JSON(Response{Message: "School with this code does not exist", Success: false, Detail: err.Error()})

// 	}

// 	err = db.WithContext(ctx).Model(&staffRating).
// 		Where("school_id = ?", school.ID).
// 		Where("user_id = ?", user.ID).
// 		Order("rating desc").
// 		Limit(10).Find(&staffRating).Error
// 	if err != nil {
// 		logger.Error(ctx, "error retrieving staff rating ", zap.Error(err))
// 		return c.Status(400).JSON(Response{Message: "an error occured retrieving staff rating", Success: false, Detail: err.Error()})

// 	}

// 	err = db.WithContext(ctx).Model(&models.Staff{}).Where("school_id = ?", school.ID).
// 		Where("status ILIKE ? ", "ACTIVE").Count(&activeStaffCount).Error
// 	if err != nil {
// 		logger.Error(ctx, "error retrieving active staff count ", zap.Error(err))
// 		return c.Status(400).JSON(Response{Message: "an error occured retrieving active staff count", Success: false, Detail: err.Error()})

// 	}

// 	err = db.WithContext(ctx).Model(&models.Staff{}).Where("school_id = ?", school.ID).
// 		Where("status ILIKE ? ", "INACTIVE").Count(&inActiveStaffCount).Error
// 	if err != nil {
// 		logger.Error(ctx, "error retrieving active inactive count ", zap.Error(err))
// 		return c.Status(400).JSON(Response{Message: "an error occured retrieving inactive staff count", Success: false, Detail: err.Error()})

// 	}

// 	err = db.WithContext(ctx).
// 		Raw(`SELECT COUNT(*) FROM staffs
//          JOIN users ON staffs.user_id = users.id
//          JOIN profiles ON users.id = profiles.user_id
//          WHERE staffs.school_id = ? AND profiles.gender ILIKE  'Male'`, school.ID).
// 		Scan(&maleCount).
// 		Error

// 	if err != nil {
// 		logger.Error(ctx, "error retrieving male count ", zap.Error(err))
// 		return c.Status(400).JSON(Response{Message: "an error occured retrieving male count", Success: false, Detail: err.Error()})

// 	}
// 	err = db.WithContext(ctx).
// 		Raw(`SELECT COUNT(*) FROM staffs
//          JOIN users ON staffs.user_id = users.id
//          JOIN profiles ON users.id = profiles.user_id
//          WHERE staffs.school_id = ? AND profiles.gender ILIKE  'Female'`, school.ID).
// 		Scan(&femaleCount).
// 		Error
// 	if err != nil {
// 		logger.Error(ctx, "error retrieving female count ", zap.Error(err))
// 		return c.Status(400).JSON(Response{Message: "an error occured retrieving female count", Success: false, Detail: err.Error()})
// 	}

// 	attendanceAnalytics, err := staffUtils.GetMonthlyAttendance(ctx, db, school.ID)
// 	if err != nil {
// 		logger.Error(ctx, "error retrieving attendance analytics ", zap.Error(err))
// 		return c.Status(400).JSON(Response{Message: "an error occured retrieving attendance analytics", Success: false, Detail: err.Error()})
// 	}

// 	serializer = serializers.StaffsAnalyticsResponseSerializer{
// 		Rating:             staffRating,
// 		Attendance:         attendanceAnalytics,
// 		ActiveStaffCount:   activeStaffCount,
// 		InactiveStaffCount: inActiveStaffCount,
// 		MaleCount:          maleCount,
// 		FemaleCount:        femaleCount,
// 	}

// 	return c.Status(200).JSON(serializer)
// }

// func StaffAnalyticsController(c *fiber.Ctx) error {
// 	// this is used to calculate a single staff analytics
// 	var staff models.Staff
// 	var school models.School
// 	var coursePerformance []serializers.CoursePerformanceSerializer
// 	var serializer serializers.StaffAnalyticsResponseSerializer

// 	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// 	defer cancel()

// 	// Open the database connection
// 	db := database.DBConnection()
// 	defer database.CloseDB(db)

// 	staffUtils := utils.NewStaffUtils()

// 	user := c.Locals("user").(models.User)
// 	schoolCode := c.Query("school_code")
// 	school, err := school.RetrieveBySchoolCode(ctx, db, schoolCode)
// 	if err != nil {
// 		return c.Status(400).JSON(Response{Message: "School with this code does not exist", Success: false, Detail: err.Error()})
// 	}
// 	staffId, err := uuid.Parse(c.Params("id"))
// 	if err != nil {
// 		return c.Status(400).JSON(Response{Message: "invalid staff id", Success: false, Detail: err.Error()})
// 	}
// 	staff, err = staff.RetrieveByIDAndSchool(ctx, db, staffId, school.ID)
// 	if err != nil {
// 		return c.Status(400).JSON(Response{Message: "invalid staff id", Success: false, Detail: err.Error()})
// 	}

// 	// only the staff and the owner can see his analytics
// 	isOk := school.IsSchoolAdminOrOwner(ctx, db, schoolCode, user.ID)
// 	if !isOk {
// 		if *staff.UserID != user.ID {
// 			return c.Status(400).JSON(
// 				Response{Message: "you dont have permission", Success: false, Detail: err})
// 		}
// 	}

// 	staffFeedbackPercentage, err := staffUtils.GetStaffFeedbackPercentage(ctx, db, staff.ID)
// 	if err != nil {
// 		logger.Error(ctx, "error retrieving staff feedback percentage ", zap.Error(err))
// 		return c.Status(500).JSON(Response{Message: "an error occured retrieving staff feedback percentage", Success: false, Detail: err.Error()})
// 	}

// 	err = db.WithContext(ctx).Model(&models.Course{}).Where("staff_id = ?", staff.ID).Find(&coursePerformance).Error
// 	if err != nil {
// 		logger.Error(ctx, "error retrieving course performance ", zap.Error(err))
// 		return c.Status(500).JSON(Response{Message: "an error occured retrieving course performance", Success: false, Detail: err.Error()})
// 	}

// 	serializer = serializers.StaffAnalyticsResponseSerializer{
// 		CoursePerformance:       coursePerformance,
// 		StaffFeedbackPercentage: staffFeedbackPercentage,
// 	}

// 	return c.Status(200).JSON(serializer)
// }
