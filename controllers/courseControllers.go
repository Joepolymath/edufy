package controllers

// import (
// 	"Learnium/adapters"
// 	"Learnium/database"
// 	"Learnium/logger"
// 	"Learnium/models"
// 	"Learnium/serializers"
// 	"Learnium/utils"
// 	"context"
// 	"fmt"
// 	"github.com/gofiber/fiber/v2"
// 	"github.com/google/uuid"
// 	"github.com/jinzhu/copier"
// 	"go.uber.org/zap"
// 	"gorm.io/gorm"
// 	"gorm.io/gorm/clause"
// 	"time"
// )

// func AdminCreateCourseController(c *fiber.Ctx) error {
// 	// Note: since we are using multipart/format-data uuid cant be binded to I used form value in there

// 	var course models.Course
// 	var requestBody serializers.CourseCreateRequestSerializer
// 	var staff models.Staff
// 	var class models.Class

// 	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// 	defer cancel()

// 	// Open the database connection
// 	db := database.DBConnection()
// 	defer database.CloseDB(db)

// 	schoolCode := c.Query("school_code")
// 	// get the logged-in user
// 	user := c.Locals("user").(models.User)
// 	validator := adapters.NewValidate()
// 	fileUpload := adapters.NewFileUpload()

// 	// get the school
// 	var school models.School
// 	school, err := school.RetrieveBySchoolCode(ctx, db, schoolCode)
// 	if err != nil {
// 		return c.Status(404).JSON(Response{Message: "school with this id does not exists", Success: false, Detail: err.Error()})
// 	}

// 	isOk := school.IsSchoolAdminOrOwner(ctx, db, schoolCode, user.ID)
// 	if !isOk {
// 		return c.Status(400).JSON(Response{Message: "you dont have permission", Success: false, Detail: err.Error()})

// 	}

// 	// get the request data
// 	if err = c.BodyParser(&requestBody); err != nil {
// 		return c.Status(400).JSON(Response{Message: "invalid json passed", Success: false, Detail: err.Error()})

// 	}
// 	requestBody.Image, err = fileUpload.UploadFile("image", c)
// 	requestBody.Video, err = fileUpload.UploadFile("video", c)

// 	// check the class id
// 	if requestBody.StaffID != nil {
// 		err := db.WithContext(ctx).Model(&models.Staff{}).Where("id = ?", requestBody.StaffID).Where("school_id = ?", school.ID).First(&staff).Error
// 		if err != nil {
// 			return c.Status(404).JSON(Response{Message: "staff with this id does not exists", Success: false, Detail: err.Error()})
// 		}
// 	}

// 	// check the class id
// 	if requestBody.ClassID != nil {
// 		err := db.WithContext(ctx).Model(&models.Class{}).Where("id = ?", requestBody.ClassID).Where("school_id = ?", school.ID).First(&class).Error
// 		if err != nil {
// 			return c.Status(400).JSON(Response{Message: "class with this id does not exists", Success: false, Detail: err.Error()})

// 		}
// 	}

// 	err2 := validator.ValidateData(&requestBody)
// 	if err2 != nil {
// 		return c.Status(400).JSON(Response{Message: err2, Success: false, Detail: err2})

// 	}

// 	err = copier.Copy(&course, &requestBody)
// 	if err != nil {
// 		return c.Status(500).JSON(
// 			Response{Message: "error copying data", Success: false, Detail: err})
// 	}

// 	// adding nil to prevent issues when creating
// 	course.Staff = nil
// 	course.Class = nil
// 	course.SchoolID = &school.ID
// 	err = db.WithContext(ctx).Model(&course).Create(&course).Preload("Staff").Preload("Class").First(&course).Error
// 	if err != nil {
// 		return c.Status(500).JSON(Response{Message: "Error on our end we are working on it.", Success: false, Detail: err.Error()})

// 	}

// 	return c.Status(201).JSON(course)
// }

// func StaffCurriculumCreateController(c *fiber.Ctx) error {

// 	var course models.Course
// 	var curriculum models.Curriculum
// 	var requestBody serializers.CurriculumCreateSerializer
// 	var staff models.Staff
// 	var school models.School
// 	var curriculumCount int64

// 	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// 	defer cancel()

// 	// Open the database connection
// 	db := database.DBConnection()
// 	defer database.CloseDB(db)

// 	validator := adapters.NewValidate()
// 	fileUpload := adapters.NewFileUpload()
// 	pointer := adapters.NewPointer()

// 	schoolCode := c.Query("school_code")
// 	// get the logged-in user
// 	user := c.Locals("user").(models.User)

// 	school, err := school.RetrieveBySchoolCode(ctx, db, schoolCode)
// 	if err != nil {
// 		return c.Status(404).JSON(Response{Message: "school with this id does not exists", Success: false, Detail: err.Error()})
// 	}

// 	// get staff by id
// 	staff, err = staff.CheckUserStaff(ctx, db, user.ID, school.ID)
// 	if err != nil {
// 		return c.Status(400).JSON(
// 			Response{Message: "user is currently not a staff.", Success: false, Detail: err.Error()})
// 	}

// 	if err = c.BodyParser(&requestBody); err != nil {
// 		return c.Status(400).JSON(Response{Message: "invalid json passed", Success: false, Detail: err.Error()})
// 	}

// 	// upload the image
// 	requestBody.File, err = fileUpload.UploadFile("file", c)

// 	// validate the data
// 	err2 := validator.ValidateData(&requestBody)
// 	if err2 != nil {
// 		return c.Status(400).JSON(Response{Message: err2, Success: false, Detail: err2})

// 	}

// 	// validate the course id
// 	if requestBody.CourseID != nil {
// 		err := db.WithContext(ctx).Model(&course).Where("school_id = ?", school.ID).Where("id = ?", requestBody.CourseID).First(&course).Error
// 		if err != nil {
// 			return c.Status(400).JSON(Response{Message: "course with this id does not exist", Success: false, Detail: err.Error()})

// 		}
// 	}

// 	if requestBody.Order == nil {
// 		err = db.WithContext(ctx).Model(&models.Curriculum{}).Where("course_id =?", requestBody.CourseID).Count(&curriculumCount).Error
// 		if err != nil {
// 			return c.Status(400).JSON(Response{Message: "error getting order for curriculum", Success: false, Detail: err.Error()})
// 		}
// 		requestBody.Order = pointer.IntPointer(int(curriculumCount) + 1)
// 	}

// 	curriculum = models.Curriculum{
// 		CourseID:    requestBody.CourseID,
// 		StaffID:     &staff.ID,
// 		SchoolID:    &school.ID,
// 		Name:        requestBody.Name,
// 		Description: requestBody.Description,
// 		File:        requestBody.File,
// 		Order:       requestBody.Order,
// 	}

// 	// create the lesson
// 	err = db.WithContext(ctx).Model(&models.Curriculum{}).Create(&curriculum).Preload("Course").First(&curriculum).Error
// 	if err != nil {
// 		return c.Status(500).JSON(
// 			Response{Message: "error on our end creating curriculum we are working on it", Success: false, Detail: err})
// 	}

// 	return c.Status(201).JSON(curriculum)
// }

// func StaffCurriculumListController(c *fiber.Ctx) error {
// 	var course models.Course
// 	var curriculums []models.Curriculum
// 	var school models.School

// 	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// 	defer cancel()

// 	// Open the database connection
// 	db := database.DBConnection()
// 	defer database.CloseDB(db)

// 	user := c.Locals("user").(models.User)
// 	schoolCode := c.Query("school_code")
// 	school, err := school.RetrieveBySchoolCode(ctx, db, schoolCode)
// 	if err != nil {
// 		return c.Status(404).JSON(Response{Message: "school with this id does not exists", Success: false, Detail: err.Error()})
// 	}
// 	courseID, err := uuid.Parse(c.Params("id"))
// 	if err != nil {
// 		return c.Status(400).JSON(Response{Message: "Invalid course id passed", Success: false, Detail: err.Error()})
// 	}

// 	err = db.WithContext(ctx).Model(&course).Where("school_id =?", school.ID).Where("id = ?", courseID).First(&course).Error
// 	if err != nil {
// 		return c.Status(400).JSON(Response{Message: "could not find course", Detail: err.Error(), Success: false})
// 	}

// 	isOK := school.IsSchoolAdminOrOwnerOrStaff(ctx, db, schoolCode, user.ID)
// 	if !isOK {
// 		return c.Status(400).JSON(Response{Message: "You dont have permission to perform this action", Success: false, Detail: err.Error()})
// 	}

// 	// check if user is a staff or admin
// 	err = db.WithContext(ctx).Model(&models.Curriculum{}).
// 		Where("school_id =?", school.ID).Where("course_id =?", course.ID).Order(clause.OrderByColumn{Column: clause.Column{Name: "order"}, Desc: false}).
// 		Find(&curriculums).Error
// 	if err != nil {
// 		return c.Status(400).JSON(Response{Message: "error getting curriculum", Detail: err.Error(), Success: false})
// 	}

// 	return c.Status(200).JSON(curriculums)
// }

// func StaffCurriculumDetailController(c *fiber.Ctx) error {
// 	var school models.School
// 	var curriculum models.Curriculum
// 	var lessons []models.Lesson
// 	var serializer serializers.CurriculumDetailSerializer

// 	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// 	defer cancel()

// 	// Open the database connection
// 	db := database.DBConnection()
// 	defer database.CloseDB(db)

// 	user := c.Locals("user").(models.User)
// 	schoolCode := c.Query("school_code")
// 	school, err := school.RetrieveBySchoolCode(ctx, db, schoolCode)
// 	if err != nil {
// 		return c.Status(404).JSON(Response{Message: "school with this id does not exists", Success: false, Detail: err.Error()})
// 	}
// 	curriculumID, err := uuid.Parse(c.Params("id"))
// 	if err != nil {
// 		return c.Status(400).JSON(Response{Message: "invalid lesson id", Success: false, Detail: err.Error()})
// 	}
// 	isOK := school.IsSchoolAdminOrOwnerOrStaff(ctx, db, schoolCode, user.ID)
// 	if !isOK {
// 		return c.Status(400).JSON(Response{Message: "You dont have permission to perform this action", Success: false, Detail: err.Error()})
// 	}

// 	err = db.WithContext(ctx).Model(&models.Curriculum{}).Where("id = ?", curriculumID).Where("school_id =?", school.ID).Preload("Staff.User").First(&curriculum).Error
// 	if err != nil {
// 		if err == gorm.ErrRecordNotFound {
// 			return c.Status(404).JSON(Response{Message: "curriculum with this id does not exist", Success: false, Detail: err.Error()})
// 		}
// 		return c.Status(400).JSON(Response{Message: "error occured getting curriculum with this id", Success: false, Detail: err.Error()})
// 	}

// 	err = db.WithContext(ctx).Model(&models.Lesson{}).Where("curriculum_id = ?", curriculum.ID).
// 		Order(clause.OrderByColumn{Column: clause.Column{Name: "order"}, Desc: false}).Find(&lessons).Error
// 	if err != nil {
// 		return c.Status(400).JSON(Response{Message: "error occured getting lessons", Success: false, Detail: err.Error()})
// 	}

// 	serializer = serializers.CurriculumDetailSerializer{
// 		ID:          &curriculum.ID,
// 		SchoolID:    curriculum.SchoolID,
// 		Staff:       curriculum.Staff,
// 		StaffID:     curriculum.StaffID,
// 		Name:        curriculum.Name,
// 		Description: curriculum.Description,
// 		Order:       curriculum.Order,
// 		Lessons:     lessons,
// 		Timestamp:   curriculum.Timestamp,
// 	}

// 	return c.Status(200).JSON(serializer)
// }

// func StaffLessonCreateController(c *fiber.Ctx) error {

// 	var curriculum models.Curriculum
// 	var requestBody serializers.LessonCreateSerializer
// 	var staff models.Staff
// 	var lessonCount int64

// 	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// 	defer cancel()

// 	// Open the database connection
// 	db := database.DBConnection()
// 	defer database.CloseDB(db)

// 	validator := adapters.NewValidate()
// 	fileUpload := adapters.NewFileUpload()
// 	pointer := adapters.NewPointer()

// 	schoolCode := c.Query("school_code")
// 	// get the logged-in user
// 	user := c.Locals("user").(models.User)

// 	// get the school
// 	var school models.School
// 	school, err := school.RetrieveBySchoolCode(ctx, db, schoolCode)
// 	if err != nil {
// 		return c.Status(404).JSON(Response{Message: "school with this id does not exists", Success: false, Detail: err.Error()})
// 	}

// 	// get staff by id
// 	staff, err = staff.CheckUserStaff(ctx, db, user.ID, school.ID)
// 	if err != nil {
// 		return c.Status(400).JSON(
// 			Response{Message: "user is currently not a staff.", Success: false, Detail: err})
// 	}

// 	if err = c.BodyParser(&requestBody); err != nil {
// 		return c.Status(400).JSON(Response{Message: "invalid json passed", Success: false, Detail: err.Error()})
// 	}

// 	// upload the image
// 	requestBody.File, err = fileUpload.UploadFile("file", c)

// 	// validate the data
// 	err2 := validator.ValidateData(&requestBody)
// 	if err2 != nil {
// 		return c.Status(400).JSON(Response{Message: err2, Success: false, Detail: err2})

// 	}
// 	// validate the course id
// 	if requestBody.CurriculumID != nil {
// 		err := db.WithContext(ctx).Model(&curriculum).Where("school_id = ?", school.ID).Where("id = ?", requestBody.CurriculumID).First(&curriculum).Error
// 		if err != nil {
// 			return c.Status(400).JSON(Response{Message: "curriculum with this id does not exist", Success: false, Detail: err.Error()})

// 		}
// 	}
// 	if requestBody.Order == nil {
// 		err = db.WithContext(ctx).Model(&models.Lesson{}).Where("curriculum_id =?", requestBody.CurriculumID).Count(&lessonCount).Error
// 		if err != nil {
// 			return c.Status(400).JSON(Response{Message: "error getting order for lesson", Success: false, Detail: err.Error()})
// 		}
// 		requestBody.Order = pointer.IntPointer(int(lessonCount) + 1)
// 	}
// 	lesson := models.Lesson{
// 		CurriculumID: requestBody.CurriculumID,
// 		StaffID:      &staff.ID,
// 		Name:         requestBody.Name,
// 		Description:  requestBody.Description,
// 		File:         requestBody.File,
// 		Order:        requestBody.Order,
// 		SchoolID:     &school.ID,
// 	}

// 	// create the lesson
// 	err = db.WithContext(ctx).Model(&models.Lesson{}).Create(&lesson).Error
// 	if err != nil {
// 		return c.Status(500).JSON(
// 			Response{Message: "error on our end creating lesson we are working on it", Success: false, Detail: err})
// 	}

// 	// update a total lesson
// 	err = lesson.UpdateCurriculumLessonCount(db)
// 	if err != nil {
// 		return c.Status(400).JSON(Response{Message: "error updating curriculum lesson count", Success: false, Detail: err.Error()})
// 	}

// 	return c.Status(201).JSON(lesson)
// }

// func StaffLessonUpdateController(c *fiber.Ctx) error {

// 	var requestBody serializers.LessonUpdateSerializer
// 	var school models.School
// 	var lesson models.Lesson

// 	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// 	defer cancel()

// 	// Open the database connection
// 	db := database.DBConnection()
// 	defer database.CloseDB(db)

// 	validator := adapters.NewValidate()
// 	fileUpload := adapters.NewFileUpload()

// 	user := c.Locals("user").(models.User)
// 	schoolCode := c.Query("school_code")
// 	school, err := school.RetrieveBySchoolCode(ctx, db, schoolCode)
// 	if err != nil {
// 		return c.Status(404).JSON(Response{Message: "school with this id does not exists", Success: false, Detail: err.Error()})
// 	}
// 	lessonID, err := uuid.Parse(c.Params("id"))
// 	if err != nil {
// 		return c.Status(400).JSON(Response{Message: "invalid lesson id", Success: false, Detail: err.Error()})
// 	}
// 	isOK := school.IsSchoolAdminOrOwnerOrStaff(ctx, db, schoolCode, user.ID)
// 	if !isOK {
// 		return c.Status(400).JSON(Response{Message: "You dont have permission to perform this action", Success: false, Detail: err.Error()})
// 	}

// 	err = c.BodyParser(&requestBody)
// 	if err != nil {
// 		return c.Status(400).JSON(Response{Message: "invalid json passed", Success: false, Detail: err.Error()})
// 	}
// 	requestBody.File, err = fileUpload.UploadFile("file", c)

// 	vErr := validator.ValidateData(&requestBody)
// 	if vErr != nil {
// 		return c.Status(400).JSON(Response{Message: vErr, Success: false, Detail: vErr})
// 	}

// 	// get the lesson update the lesson and get it again
// 	err = db.WithContext(ctx).Model(&models.Lesson{}).Where("id = ?", lessonID).
// 		Where("school_id =?", school.ID).Updates(&requestBody).First(&lesson).Error
// 	if err != nil {
// 		if err == gorm.ErrRecordNotFound {
// 			return c.Status(400).JSON(Response{Message: "lesson not found", Success: false, Detail: err.Error()})
// 		}
// 		return c.Status(400).JSON(Response{Message: "error updating lesson", Success: false, Detail: err.Error()})
// 	}
// 	// update a total lesson
// 	err = lesson.UpdateCurriculumCompletedCount(db)
// 	if err != nil {
// 		return c.Status(400).JSON(Response{Message: "error updating curriculum lesson completed count", Success: false, Detail: err.Error()})
// 	}

// 	return c.Status(200).JSON(lesson)
// }

// func StaffLessonListController(c *fiber.Ctx) error {
// 	var curriculum models.Curriculum
// 	var lessons []models.Lesson
// 	var school models.School

// 	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// 	defer cancel()

// 	// Open the database connection
// 	db := database.DBConnection()
// 	defer database.CloseDB(db)

// 	user := c.Locals("user").(models.User)
// 	schoolCode := c.Query("school_code")
// 	school, err := school.RetrieveBySchoolCode(ctx, db, schoolCode)
// 	if err != nil {
// 		return c.Status(404).JSON(Response{Message: "school with this id does not exists", Success: false, Detail: err.Error()})
// 	}
// 	curriculumID, err := uuid.Parse(c.Params("id"))
// 	if err != nil {
// 		return c.Status(400).JSON(Response{Message: "Invalid course id passed", Success: false, Detail: err.Error()})
// 	}

// 	err = db.WithContext(ctx).Model(&models.Curriculum{}).Where("school_id =?", school.ID).Where("id = ?", curriculumID).First(&curriculum).Error
// 	if err != nil {
// 		return c.Status(400).JSON(Response{Message: "could not find curriculum", Detail: err.Error(), Success: false})
// 	}

// 	isOK := school.IsSchoolAdminOrOwnerOrStaff(ctx, db, schoolCode, user.ID)
// 	if !isOK {
// 		return c.Status(400).JSON(Response{Message: "You dont have permission to perform this action", Success: false, Detail: err.Error()})
// 	}

// 	err = db.WithContext(ctx).Model(&models.Lesson{}).Where("curriculum_id = ?", curriculum.ID).
// 		Order(clause.OrderByColumn{Column: clause.Column{Name: "order"}, Desc: false}).
// 		Find(&lessons).Error
// 	if err != nil {
// 		return c.Status(400).JSON(Response{Message: "error getting lessons", Success: false, Detail: err.Error()})
// 	}

// 	return c.Status(200).JSON(lessons)
// }

// func StaffCourseDetailController(c *fiber.Ctx) error {
// 	var curriculumSerializer []serializers.CurriculumDetailSerializer

// 	schoolCode := c.Query("school_code")
// 	_courseID := c.Params("id")
// 	courseID, err := uuid.Parse(_courseID)
// 	if err != nil {
// 		return c.Status(400).JSON(
// 			Response{Message: "invalid course id", Success: false, Detail: err})
// 	}

// 	var courseDetailSerializer serializers.CourseDetailSerializer
// 	var course models.Course

// 	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// 	defer cancel()

// 	// Open the database connection
// 	db := database.DBConnection()
// 	defer database.CloseDB(db)

// 	// get the logged-in user
// 	_ = c.Locals("user").(models.User)

// 	// get the school
// 	var school models.School
// 	school, err = school.RetrieveBySchoolCode(ctx, db, schoolCode)
// 	if err != nil {
// 		return c.Status(404).JSON(Response{Message: "school with this id does not exists", Success: false, Detail: err.Error()})
// 	}

// 	err = db.WithContext(ctx).Model(&models.Course{}).
// 		Where("id = ?", courseID).
// 		Where("school_id = ?", school.ID).
// 		Preload("Class").
// 		Preload("Staff.User").First(&course).Error
// 	if err != nil {
// 		if err == gorm.ErrRecordNotFound {
// 			return c.Status(400).JSON(Response{Message: "course with this id does not exist", Success: false, Detail: err.Error()})
// 		}
// 		return c.Status(500).JSON(Response{Message: "Error on our end we are working on it.", Success: false, Detail: err.Error()})
// 	}

// 	// get the curriculums
// 	err = db.WithContext(ctx).Model(&models.Curriculum{}).Where("course_id = ?", courseID).
// 		Order(clause.OrderByColumn{Column: clause.Column{Name: "order"}, Desc: false}).Find(&curriculumSerializer).Error
// 	if err != nil {
// 		return c.Status(400).JSON(
// 			Response{Message: "Error getting course lessons", Success: false, Detail: err})
// 	}

// 	for i, curriculum := range curriculumSerializer {
// 		var lessons []models.Lesson
// 		err = db.WithContext(ctx).Model(&models.Lesson{}).Where("curriculum_id = ?", curriculum.ID).
// 			Order(clause.OrderByColumn{Column: clause.Column{Name: "order"}, Desc: false}).Find(&lessons).Error
// 		if err != nil {
// 			return c.Status(400).JSON(
// 				Response{Message: "Error getting course lessons", Success: false, Detail: err})
// 		}
// 		curriculumSerializer[i].Lessons = lessons

// 	}

// 	// map course to courseDetailSerializer
// 	courseDetailSerializer = serializers.CourseDetailSerializer{
// 		ID:                   &course.ID,
// 		School:               course.School,
// 		SchoolID:             course.SchoolID,
// 		Class:                course.Class,
// 		ClassID:              course.ClassID,
// 		Staff:                course.Staff,
// 		StaffID:              course.StaffID,
// 		Name:                 course.Name,
// 		Description:          course.Description,
// 		Image:                course.Image,
// 		Video:                course.Video,
// 		Curriculums:          curriculumSerializer,
// 		Timestamp:            course.Timestamp,
// 		Performance:          course.Performance,
// 		Attendance:           course.Attendance,
// 		EnrolledStudentCount: course.EnrolledStudentCount,
// 	}

// 	return c.Status(200).JSON(courseDetailSerializer)

// }

// func AdminCourseListAllController(c *fiber.Ctx) error {
// 	schoolCode := c.Query("school_code")

// 	page := c.QueryInt("page", 1)    // default to page 1 if not provided
// 	limit := c.QueryInt("limit", 10) // default to 10 items per page if not provided

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
// 		return c.Status(404).JSON(Response{Message: "school with this id does not exists", Success: false, Detail: err})
// 	}

// 	isOk := school.IsSchoolAdminOrOwner(ctx, db, schoolCode, user.ID)
// 	if !isOk {
// 		return c.Status(400).JSON(Response{Message: "you dont have permission", Success: false, Detail: nil})
// 	}

// 	var total int64
// 	db.Model(&models.Course{}).Where("school_id = ?", school.ID).Count(&total)

// 	var courses []models.Course
// 	err = db.Scopes(adapters.NewPaginateAdapter(limit, page).PaginatedResult).
// 		WithContext(ctx).
// 		Model(&models.Course{}).
// 		Preload("Class").
// 		Preload("Staff").
// 		Preload("Staff.User").
// 		Where("school_id = ?", school.ID).
// 		Find(&courses).Error
// 	if err != nil {
// 		return c.Status(400).JSON(
// 			Response{Message: "Error getting course for this school ", Success: false, Detail: err})
// 	}

// 	return c.Status(200).JSON(fiber.Map{
// 		"total": total,
// 		"data":  courses,
// 	})
// }

// func StaffEnrollCourseCreateController(c *fiber.Ctx) error {
// 	var requestBody serializers.EnrollCourseRequestSerializer
// 	var studentTask models.StudentTask
// 	var checkEnrolledCourse models.EnrolledCourse
// 	var course models.Course

// 	schoolCode := c.Query("school_code")

// 	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// 	defer cancel()

// 	// Open the database connection
// 	db := database.DBConnection()
// 	defer database.CloseDB(db)

// 	// get the logged-in user
// 	user := c.Locals("user").(models.User)
// 	validator := adapters.NewValidate()

// 	if err := c.BodyParser(&requestBody); err != nil {
// 		return c.Status(400).JSON(Response{Message: "invalid json passed", Success: false, Detail: err.Error()})
// 	}
// 	err2 := validator.ValidateData(&requestBody)
// 	if err2 != nil {
// 		return c.Status(400).JSON(Response{Message: err2, Success: false, Detail: err2})

// 	}
// 	if requestBody.EndDate != nil {
// 		currentUTC := time.Now().UTC()
// 		if requestBody.EndDate.Before(currentUTC) {
// 			return c.Status(400).JSON(Response{Message: "end date cannot be less than current date", Success: false, Detail: err2})
// 		}
// 	}

// 	// get the school
// 	var school models.School
// 	school, err := school.RetrieveBySchoolCode(ctx, db, schoolCode)
// 	if err != nil {
// 		return c.Status(404).JSON(Response{Message: "school with this id does not exists", Success: false, Detail: err.Error()})
// 	}

// 	isOk := school.IsSchoolAdminOrOwner(ctx, db, schoolCode, user.ID)
// 	if !isOk {
// 		return c.Status(400).JSON(Response{Message: "you dont have permission", Success: false, Detail: err.Error()})
// 	}

// 	// check if the student exists
// 	var student models.Student
// 	err = db.WithContext(ctx).Model(&models.Student{}).
// 		Where("id = ?", requestBody.StudentID).
// 		Where("school_id = ?", school.ID).
// 		First(&student).Error
// 	if err != nil {
// 		return c.Status(400).JSON(
// 			Response{Message: "Student with this id does not exists", Success: false, Detail: err})
// 	}

// 	// check if the course exists
// 	err = db.WithContext(ctx).Model(&models.Course{}).
// 		Where("id = ?", requestBody.CourseID).
// 		Where("school_id = ?", school.ID).
// 		First(&course).Error
// 	if err != nil {
// 		return c.Status(400).JSON(Response{Message: "course with this id does not exists", Success: false, Detail: err.Error()})

// 	}

// 	// check if the student has already registered that course before
// 	err = db.WithContext(ctx).Model(&models.EnrolledCourse{}).
// 		Where("course_id = ?", requestBody.CourseID).
// 		Where("student_id = ?", requestBody.StudentID).
// 		Where("end_date > ?", time.Now().UTC()).
// 		Where("status = ?", "ENROLLED").First(&checkEnrolledCourse).Error
// 	if err == nil {
// 		return c.Status(400).JSON(Response{Message: "course had already being registered", Success: false, Detail: err.Error()})
// 	}

// 	// create the enrolled course
// 	enrolledCourse := models.EnrolledCourse{
// 		StudentID: requestBody.StudentID,
// 		CourseID:  requestBody.CourseID,
// 		StartDate: requestBody.StartDate,
// 		EndDate:   requestBody.EndDate,
// 		SchoolID:  &school.ID,
// 	}

// 	err = db.WithContext(ctx).Model(&models.EnrolledCourse{}).Create(&enrolledCourse).
// 		Preload("Student").
// 		Preload("Student.User").
// 		Preload("Course").First(&enrolledCourse).Error
// 	if err != nil {
// 		return c.Status(500).JSON(Response{Message: "an error occured ", Success: false, Detail: err.Error()})
// 	}

// 	// create the student task
// 	err = studentTask.CreateStudentTask(ctx, db, school.ID, *enrolledCourse.CourseID, *enrolledCourse.StudentID)
// 	if err != nil {
// 		db.WithContext(ctx).Model(&enrolledCourse).Delete(&enrolledCourse) // delete the enrolled course
// 		logger.Error(ctx, "this is used to create the task for the student", zap.Error(err))
// 		return c.Status(400).JSON(
// 			Response{Message: "an error occurred creating the task for the student", Success: false, Detail: err})
// 	}

// 	return c.Status(201).JSON(enrolledCourse)
// }

// func StudentEnrollCourseListController(c *fiber.Ctx) error {

// 	schoolCode := c.Query("school_code")
// 	status := c.Query("status", "ENROLLED")

// 	page := c.QueryInt("page", 1)    // default to page 1 if not provided
// 	limit := c.QueryInt("limit", 10) // default to 10 items per page if not provided

// 	err := utils.ValidateQueryStatus(status)
// 	if err != nil {
// 		return c.Status(400).JSON(
// 			Response{Message: "Invalid status provided", Success: false, Detail: err})
// 	}

// 	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// 	defer cancel()

// 	// Open the database connection
// 	db := database.DBConnection()
// 	defer database.CloseDB(db)

// 	// get the logged-in user
// 	user := c.Locals("user").(models.User)

// 	var school models.School
// 	school, err = school.RetrieveBySchoolCode(ctx, db, schoolCode)
// 	if err != nil {
// 		return c.Status(404).JSON(Response{Message: "school with this id does not exists", Success: false, Detail: err.Error()})
// 	}

// 	var student models.Student
// 	student, err = student.RetrieveStudentByUserIDAndSchool(db, ctx, user.ID, school.ID)
// 	if err != nil {
// 		return c.Status(400).JSON(
// 			Response{Message: "student with this id does not exists", Success: false, Detail: err})
// 	}

// 	var total int64
// 	db.Model(&models.EnrolledCourse{}).
// 		Where("student_id =?", student.ID).
// 		Where("status = ?", status).Count(&total)

// 	// get the enrolled course of the student
// 	var enrolledCourses []models.EnrolledCourse
// 	err = db.Scopes(adapters.NewPaginateAdapter(limit, page).PaginatedResult).
// 		WithContext(ctx).Model(&models.EnrolledCourse{}).
// 		Where("student_id =?", student.ID).
// 		Where("status = ?", status).
// 		Preload("Course").
// 		Where("end_date > ?", time.Now().UTC()).
// 		Preload("Course.Staff").
// 		Preload("Course.Staff.User").
// 		Find(&enrolledCourses).Error

// 	return c.Status(200).JSON(fiber.Map{
// 		"total": total,
// 		"data":  enrolledCourses,
// 	})
// }

// func StaffOrAdminCourseEnrolledStudentListController(c *fiber.Ctx) error {
// 	/* */
// 	var school models.School
// 	var course models.Course
// 	var enrolledCourses []models.EnrolledCourse

// 	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// 	defer cancel()

// 	// Open the database connection
// 	db := database.DBConnection()
// 	defer database.CloseDB(db)

// 	user := c.Locals("user").(models.User)

// 	page := c.QueryInt("page", 1)    // default to page 1 if not provided
// 	limit := c.QueryInt("limit", 10) // default to 10 items per page if not provided
// 	schoolCode := c.Query("school_code")
// 	search := c.Query("search")
// 	status := c.Query("status")

// 	school, err := school.RetrieveBySchoolCode(ctx, db, schoolCode)
// 	if err != nil {
// 		return c.Status(404).JSON(Response{Message: "school with this id does not exists", Success: false, Detail: err.Error()})
// 	}

// 	courseID, err := uuid.Parse(c.Params("id"))
// 	if err != nil {
// 		return c.Status(400).JSON(Response{Message: "invalid staff id passed", Success: false, Detail: err.Error()})
// 	}

// 	// check if the course is valid
// 	err = db.WithContext(ctx).Model(&course).Where("id = ?", courseID).Preload("Staff").Preload("Staff.User").Find(&course).Error
// 	if err != nil {
// 		if err == gorm.ErrRecordNotFound {
// 			return c.Status(400).JSON(Response{Message: "course with this id does not exist", Success: false, Detail: err.Error()})
// 		}
// 		logger.Error(ctx, "error retrieving course", zap.Error(err))
// 		return c.Status(400).JSON(Response{Message: "course with this id does not exist", Success: false, Detail: err.Error()})
// 	}

// 	isOk := school.IsSchoolAdminOrOwner(ctx, db, schoolCode, user.ID)
// 	if !isOk {
// 		if *course.Staff.UserID != user.ID {
// 			return c.Status(400).JSON(Response{Message: "you dont have permission", Success: false, Detail: err.Error()})
// 		}
// 	}
// 	query := db.Scopes(adapters.NewPaginateAdapter(limit, page).PaginatedResult).WithContext(ctx).Model(&models.EnrolledCourse{}).
// 		Where("course_id = ?", courseID).Preload("Student.User").Preload("Course")

// 	if status != "" {
// 		query = query.Where("status =?", status)
// 	}

// 	if search != "" {
// 		query = query.Where("student_id IN (SELECT id FROM students WHERE students.user_id IN "+"(SELECT id FROM users WHERE users.first_name ILIKE ? OR users.last_name ILIKE ?))", "%"+search+"%", "%"+search+"%")
// 	}

// 	err = query.Find(&enrolledCourses).Error
// 	if err != nil {
// 		logger.Error(ctx, "error getting course for student registered to it", zap.Error(err))
// 		return c.Status(500).JSON(Response{Message: "Error getting  enrolled courses.", Success: false, Detail: err.Error()})
// 	}

// 	var total int64
// 	db.Model(&models.EnrolledCourse{}).
// 		Where("course_id = ?", courseID).Where("status =?", "ENROLLED").Count(&total)

// 	return c.Status(200).JSON(fiber.Map{
// 		"total": total,
// 		"data":  enrolledCourses,
// 	})
// }

// func StaffCoursesListController(c *fiber.Ctx) error {
// 	var school models.School
// 	var staff models.Staff
// 	var courses []models.Course

// 	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// 	defer cancel()

// 	// Open the database connection
// 	db := database.DBConnection()
// 	defer database.CloseDB(db)

// 	user := c.Locals("user").(models.User)
// 	page := c.QueryInt("page", 1)

// 	limit := c.QueryInt("limit", 10)
// 	staffID, err := uuid.Parse(c.Params("id"))
// 	if err != nil {
// 		return c.Status(400).JSON(Response{Message: "invalid staff id passed", Success: false, Detail: err.Error()})
// 	}

// 	staff, err = staff.Retrieve(ctx, db, staffID)
// 	if err != nil {
// 		return c.Status(404).JSON(Response{Message: "staff with this id does not exists", Success: false, Detail: err.Error()})
// 	}
// 	schoolCode := c.Query("school_code")
// 	school, err = school.RetrieveBySchoolCode(ctx, db, schoolCode)
// 	if err != nil {
// 		return c.Status(404).JSON(Response{Message: "school with this id does not exists", Success: false, Detail: err.Error()})
// 	}
// 	search := c.Query("search")

// 	isOk := school.IsSchoolAdminOrOwner(ctx, db, schoolCode, user.ID)
// 	if !isOk {

// 		if *staff.UserID != user.ID {
// 			return c.Status(400).JSON(Response{Message: "you dont have permission", Success: false, Detail: err.Error()})
// 		}
// 	}

// 	query := db.Scopes(adapters.NewPaginateAdapter(limit, page).PaginatedResult).WithContext(ctx).Model(&models.Course{}).
// 		Where("staff_id = ?", staffID).Where("school_id = ?", school.ID).Preload("Class")

// 	if search != "" {
// 		query = query.Where("name LIKE ? OR  description ILIKE ?", "%"+search+"%", "%"+search+"%")
// 	}

// 	err = query.Find(&courses).Error
// 	if err != nil {
// 		logger.Error(ctx, "error finding courses ", zap.Error(err))
// 		return c.Status(400).JSON(Response{Message: "error getting courses", Success: false, Detail: err.Error()})
// 	}
// 	var total int64
// 	db.Model(&models.Course{}).
// 		Where("staff_id = ?", staffID).Where("school_id = ?", school.ID).Count(&total)

// 	return c.Status(200).JSON(fiber.Map{
// 		"total": total,
// 		"data":  courses,
// 	})

// }

// func CourseScheduleCreateController(c *fiber.Ctx) error {
// 	var requestBody serializers.CourseScheduleRequestSerializer
// 	var course models.Course
// 	var school models.School
// 	var courseSchedule models.CourseSchedule
// 	var staff models.Staff

// 	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// 	defer cancel()

// 	// Open the database connection
// 	db := database.DBConnection()
// 	defer database.CloseDB(db)

// 	validateAdapters := adapters.NewValidate()

// 	user := c.Locals("user").(models.User)
// 	schoolCode := c.Query("school_code")
// 	school, err := school.RetrieveBySchoolCode(ctx, db, schoolCode)
// 	if err != nil {
// 		return c.Status(400).JSON(Response{Message: "school with this code does not exists", Success: false, Detail: err.Error()})
// 	}
// 	err = c.BodyParser(&requestBody)
// 	if err != nil {
// 		return c.Status(400).JSON(Response{Message: "invalid json passed", Success: false, Detail: err.Error()})
// 	}

// 	vErr := validateAdapters.ValidateData(&requestBody)
// 	if vErr != nil {
// 		return c.Status(400).JSON(Response{Message: "course with this id does not exists", Success: false, Detail: vErr})
// 	}

// 	course, err = course.RetrieveSchoolCourse(ctx, db, *requestBody.CourseID, school.ID)
// 	if err != nil {
// 		return c.Status(400).JSON(Response{Message: "course with this id does not exists", Success: false, Detail: err.Error()})
// 	}
// 	staff, err = staff.RetrieveByIDAndSchool(ctx, db, *requestBody.StaffID, school.ID)
// 	if err != nil {
// 		return c.Status(400).JSON(Response{Message: "staff with this id does not exists", Success: false, Detail: err.Error()})
// 	}

// 	isOk := school.IsSchoolAdminOrOwner(ctx, db, schoolCode, user.ID)
// 	if !isOk {
// 		return c.Status(400).JSON(Response{Message: "you dont have permission to create schedule", Success: false, Detail: err.Error()})
// 	}

// 	courseSchedule = models.CourseSchedule{
// 		SchoolID:     &school.ID,
// 		CourseID:     &course.ID,
// 		StaffID:      requestBody.StaffID,
// 		ScheduleDate: requestBody.ScheduleDate,
// 		FromTime:     requestBody.FromTime,
// 		ToTime:       requestBody.FromTime,
// 	}

// 	err = db.WithContext(ctx).Model(&courseSchedule).Create(&courseSchedule).Preload("Course").First(&courseSchedule).Error
// 	if err != nil {
// 		return c.Status(500).JSON(Response{Message: "error creating course schedule", Success: false, Detail: err.Error()})
// 	}
// 	return c.Status(200).JSON(courseSchedule)
// }

// func StaffScheduleListController(c *fiber.Ctx) error {
// 	/* this is used to list all the schedules available to a staff */
// 	var staff models.Staff
// 	var courseSchedule []models.CourseSchedule
// 	var school models.School

// 	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// 	defer cancel()

// 	// Open the database connection
// 	db := database.DBConnection()
// 	defer database.CloseDB(db)

// 	user := c.Locals("user").(models.User)
// 	schoolCode := c.Query("school_code")
// 	staffID, err := uuid.Parse(c.Params("id"))
// 	if err != nil {
// 		return c.Status(400).JSON(Response{Message: "invalid staff id passed", Success: false, Detail: err})
// 	}
// 	school, err = school.RetrieveBySchoolCode(ctx, db, schoolCode)
// 	if err != nil {
// 		return c.Status(400).JSON(Response{Message: "error retrieving school", Detail: err, Success: false})
// 	}
// 	isOk := school.IsSchoolAdminOrOwner(ctx, db, schoolCode, user.ID)
// 	if !isOk {
// 		return c.Status(400).JSON(Response{Message: "you dont have permission", Success: false, Detail: err.Error()})

// 	}

// 	err = db.WithContext(ctx).Model(&staff).Where("id =? ", staffID).Where("school_id =?", school.ID).First(&staff).Error
// 	if err != nil {
// 		return c.Status(400).JSON(Response{Message: "error retrieving staff", Detail: err, Success: false})
// 	}

// 	err = db.WithContext(ctx).Model(&models.CourseSchedule{}).Where("staff_id = ?", staffID).
// 		Where("school_id = ?", school.ID).Preload("Course").Find(&courseSchedule).Error
// 	if err != nil {
// 		return c.Status(400).JSON(Response{Message: "error retrieving course schedule", Detail: err, Success: false})
// 	}

// 	return c.Status(200).JSON(courseSchedule)

// }

// func StaffListEnrolledStudentsInCourseController(c *fiber.Ctx) error {
// 	/* this is used to list all students that have enrolled in a course*/
// 	var students []models.Student
// 	var school models.School
// 	var course models.Course
// 	var enrolledCourses []models.EnrolledCourse

// 	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// 	defer cancel()

// 	// Open the database connection
// 	db := database.DBConnection()
// 	defer database.CloseDB(db)

// 	user := c.Locals("user").(models.User)
// 	schoolCode := c.Query("school_code")
// 	courseID, err := uuid.Parse(c.Params("id"))
// 	if err != nil {
// 		return c.Status(400).JSON(Response{Message: "invalid course id passed", Success: false, Detail: err})
// 	}
// 	school, err = school.RetrieveBySchoolCode(ctx, db, schoolCode)
// 	if err != nil {
// 		return c.Status(400).JSON(Response{Message: "error retrieving school", Detail: err, Success: false})
// 	}
// 	isOk := school.IsSchoolAdminOrOwner(ctx, db, schoolCode, user.ID)
// 	if !isOk {
// 		return c.Status(400).JSON(Response{Message: "you dont have permission", Success: false, Detail: err.Error()})
// 	}

// 	course, err = course.RetrieveSchoolCourse(ctx, db, courseID, school.ID)
// 	if err != nil {
// 		return c.Status(400).JSON(Response{Message: "error retrieving course", Detail: err, Success: false})
// 	}

// 	err = db.WithContext(ctx).
// 		Model(&models.EnrolledCourse{}).
// 		Where("school_id = ?", school.ID).
// 		Where("course_id = ?", course.ID).
// 		Where("status ILIKE ?", "ENROLLED").
// 		Where("end_date > ?", time.Now().UTC()). // Use UTC() to ensure consistent timezone
// 		Preload("Student.User").
// 		Find(&enrolledCourses).Error
// 	if err != nil {
// 		return c.Status(500).JSON(Response{Message: "error retrieving enrolled courses", Detail: err.Error(), Success: false})
// 	}

// 	for _, enrolledCourse := range enrolledCourses {
// 		students = append(students, *enrolledCourse.Student)
// 	}

// 	return c.Status(200).JSON(students)
// }

// func StaffRegisterStudentAttendanceController(c *fiber.Ctx) error {
// 	// Implementation for staff to register student attendance in a course
// 	var requestBody serializers.RegisterStudentPresenceRequestSerializer
// 	var school models.School
// 	var attendance models.Attendance
// 	var studentAttendance models.StudentAttendance
// 	var course models.Course

// 	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// 	defer cancel()

// 	// Open the database connection
// 	db := database.DBConnection()
// 	defer database.CloseDB(db)

// 	validateAdapter := adapters.NewValidate()

// 	user := c.Locals("user").(models.User)
// 	schoolCode := c.Query("school_code")

// 	school, err := school.RetrieveBySchoolCode(ctx, db, schoolCode)
// 	if err != nil {
// 		return c.Status(400).JSON(Response{Message: "error retrieving school", Detail: err, Success: false})
// 	}
// 	isOk := school.IsSchoolAdminOrOwner(ctx, db, schoolCode, user.ID)
// 	if !isOk {
// 		return c.Status(400).JSON(Response{Message: "you dont have permission", Success: false, Detail: err.Error()})
// 	}

// 	err = c.BodyParser(&requestBody)
// 	if err != nil {
// 		return c.Status(400).JSON(Response{Message: "invalid json passed", Success: false, Detail: err.Error()})
// 	}

// 	vErr := validateAdapter.ValidateData(&requestBody)
// 	if vErr != nil {
// 		return c.Status(200).JSON(Response{Message: vErr, Success: false, Detail: vErr})
// 	}

// 	course, err = course.RetrieveSchoolCourse(ctx, db, *requestBody.CourseID, school.ID)
// 	if err != nil {
// 		if err == gorm.ErrRecordNotFound {
// 			return c.Status(404).JSON(Response{Message: "course with this id does not exist", Success: false, Detail: err.Error()})
// 		}
// 		return c.Status(400).JSON(Response{Message: "error retrieving course", Detail: err.Error(), Success: false})
// 	}

// 	attendance = models.Attendance{
// 		SchoolID: &school.ID,
// 		CourseID: &course.ID,
// 	}
// 	err = db.WithContext(ctx).Model(&models.Attendance{}).Create(&attendance).Preload("Course").First(&attendance).Error
// 	if err != nil {
// 		return c.Status(400).JSON(Response{Message: "error creating attendance ", Detail: err.Error(), Success: false})
// 	}

// 	// mark all the student in the serializer
// 	for _, studentPresence := range requestBody.StudentPresences {

// 		// validate the student
// 		var student models.Student
// 		student, err = student.RetrieveStudentByIDAndSchool(ctx, db, *studentPresence.StudentID, school.ID)
// 		if err != nil {
// 			return c.Status(400).JSON(Response{Message: fmt.Sprintf("error retrieving student with id %s ", *studentPresence.StudentID), Detail: err.Error(), Success: false})
// 		}

// 		// check if the student is registered in the course also
// 		var enrolledCourse models.EnrolledCourse
// 		err := db.WithContext(ctx).Model(&models.EnrolledCourse{}).Where("student_id = ?", student.ID).
// 			Where("course_id = ?", course.ID).Where("end_date > ?", time.Now().UTC()).Where("status = ?", "ENROLLED").First(&enrolledCourse).Error
// 		if err != nil {
// 			if err == gorm.ErrRecordNotFound {
// 				return c.Status(400).JSON(Response{Message: "student is not registered on this course", Detail: err.Error(), Success: false})
// 			}
// 			return c.Status(400).JSON(Response{Message: "student is not registered on this course", Detail: err.Error(), Success: false})
// 		}

// 		studentAttendance = models.StudentAttendance{
// 			SchoolID:     &school.ID,
// 			StudentID:    studentPresence.StudentID,
// 			AttendanceID: &attendance.ID,
// 			Status:       studentPresence.Status,
// 		}
// 		err = db.WithContext(ctx).Model(&models.StudentAttendance{}).Create(&studentAttendance).Error
// 		if err != nil {
// 			return c.Status(400).JSON(Response{Message: fmt.Sprintf("error creating student attendance for %s", *studentPresence.StudentID),
// 				Detail: err.Error(), Success: false})
// 		}
// 	}

// 	// calculate the attendance
// 	err = attendance.CalculateAttendancePercent(ctx, db, attendance.ID)
// 	if err != nil {
// 		logger.Error(ctx, "error calculating attendance percentage", zap.Error(err))
// 		return c.Status(400).JSON(Response{Message: "error calculating attendance percentage", Detail: err.Error(), Success: false})
// 	}

// 	err = course.CalculateCourseAttendance(ctx, db, course.ID)
// 	if err != nil {
// 		logger.Error(ctx, "error calculating course attendance percentage", zap.Error(err))
// 		return c.Status(400).JSON(Response{Message: "Error calculating course attendance percentage", Success: false, Detail: err.Error()})
// 	}
// 	return c.Status(200).JSON(attendance)
// }

// func StaffCourseAnalyticsController(c *fiber.Ctx) error {
// 	//	 this is used to calculate the subject analytics
// 	var school models.School
// 	var course models.Course
// 	var serializer serializers.CourseAnalyticsSerializer

// 	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// 	defer cancel()

// 	// Open the database connection
// 	db := database.DBConnection()
// 	defer database.CloseDB(db)

// 	courseUtils := utils.NewCourseUtils()

// 	user := c.Locals("user").(models.User)
// 	schoolCode := c.Query("school_code")
// 	courseID, err := uuid.Parse(c.Params("id"))
// 	if err != nil {
// 		return c.Status(400).JSON(Response{Message: "invalid course id passed", Success: false, Detail: err})
// 	}
// 	school, err = school.RetrieveBySchoolCode(ctx, db, schoolCode)
// 	if err != nil {
// 		return c.Status(400).JSON(Response{Message: "error retrieving school", Detail: err, Success: false})
// 	}
// 	isOk := school.IsSchoolAdminOrOwnerOrStaff(ctx, db, schoolCode, user.ID)
// 	if !isOk {
// 		return c.Status(400).JSON(Response{Message: "you dont have permission", Success: false, Detail: err.Error()})
// 	}

// 	course, err = course.RetrieveSchoolCourse(ctx, db, courseID, school.ID)
// 	if err != nil {
// 		logger.Error(ctx, "error retrieving course", zap.Error(err))
// 		return c.Status(400).JSON(Response{Message: "error retrieving course", Detail: err.Error(), Success: false})
// 	}

// 	higherPerformingStudents, err := courseUtils.GetHighestPerformingStudentAnalytics(ctx, db, course.ID)
// 	if err != nil {
// 		logger.Error(ctx, "error getting highest performing students", zap.Error(err))
// 		return c.Status(400).JSON(Response{Message: "error getting highest performing students", Detail: err.Error(), Success: false})
// 	}
// 	courseCurriculumInfos, curriculumProgress, err := courseUtils.GetCourseCurriculumAnalytics(ctx, db, course.ID)
// 	if err != nil {
// 		logger.Error(ctx, "error getting course curriculum infos", zap.Error(err))
// 		return c.Status(400).JSON(Response{Message: "error getting course curriculum infos", Detail: err.Error(), Success: false})
// 	}

// 	studentPerformanceAnalytics, err := courseUtils.GetStudentPerformanceAnalytics(ctx, db, course.ID)
// 	if err != nil {
// 		logger.Error(ctx, "Error calculating Student Performance Analytics", zap.Error(err))
// 		return c.Status(500).JSON(Response{Message: "Error calculating student Performance analytics ", Success: false, Detail: err.Error()})
// 	}

// 	serializer = serializers.CourseAnalyticsSerializer{
// 		RegisteredStudents:          course.EnrolledStudentCount,
// 		Attendance:                  course.Attendance,
// 		CourseCurriculumInfo:        &courseCurriculumInfos,
// 		TeachingHours:               nil,
// 		CurriculumProgress:          &curriculumProgress,
// 		HighestPerformingStudents:   &higherPerformingStudents,
// 		StudentPerformanceAnalytics: &studentPerformanceAnalytics,
// 	}
// 	return c.Status(200).JSON(serializer)
// }
