package controllers

// import (
// 	"Learnium/adapters"
// 	"Learnium/database"
// 	"Learnium/logger"
// 	"Learnium/models"
// 	"Learnium/serializers"
// 	"context"
// 	"encoding/json"
// 	"errors"
// 	"fmt"
// 	"github.com/gofiber/fiber/v2"
// 	"github.com/google/uuid"
// 	"github.com/jinzhu/copier"
// 	"go.uber.org/zap"
// 	"gorm.io/gorm"
// 	"reflect"
// 	"strconv"
// 	"strings"
// 	"time"
// )

// func CourseTaskCreateController(c *fiber.Ctx) error {
// 	var requestBody serializers.TaskCreateRequestSerializer
// 	var course models.Course
// 	var staff models.Staff
// 	var school models.School
// 	var enrolledCourses []models.EnrolledCourse
// 	var studentTask models.StudentTask

// 	validator := adapters.NewValidate()
// 	fileUpload := adapters.NewFileUpload()

// 	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// 	defer cancel()

// 	// Open the database connection
// 	db := database.DBConnection()
// 	defer database.CloseDB(db)

// 	// get the logged-in user
// 	user := c.Locals("user").(models.User)
// 	schoolCode := c.Query("school_code")

// 	err := c.BodyParser(&requestBody)
// 	if err != nil {
// 		return c.Status(400).JSON(
// 			Response{Message: "Invalid json data", Success: false, Detail: err.Error()})
// 	}
// 	requestBody.File, err = fileUpload.UploadFile("file", c)

// 	err2 := validator.ValidateData(&requestBody)
// 	if err2 != nil {
// 		return c.Status(400).JSON(Response{Message: err2, Success: false, Detail: err2})
// 	}

// 	// check the course if its exists
// 	course, err = course.Retrieve(ctx, db, *requestBody.CourseID)
// 	if err != nil {
// 		return c.Status(400).JSON(Response{Message: "course with this id does not exist", Success: false, Detail: err.Error()})
// 	}

// 	school, err = school.RetrieveBySchoolCode(ctx, db, schoolCode)
// 	if err != nil {
// 		return c.Status(400).JSON(
// 			Response{Message: "Invalid school code ", Success: false, Detail: err.Error()})
// 	}

// 	// get the user staff
// 	staff, err = staff.RetrieveByUserAndSchool(ctx, db, school.ID, user.ID)
// 	if err != nil {

// 		return c.Status(400).JSON(
// 			Response{Message: "User is not a staff in this school", Success: false, Detail: err.Error()})
// 	}

// 	courseTask := models.CourseTask{
// 		CourseID:    requestBody.CourseID,
// 		SchoolID:    &school.ID,
// 		Title:       requestBody.Title,
// 		Description: requestBody.Description,
// 		TotalPoint:  requestBody.TotalPoint,
// 		File:        requestBody.File,
// 		DueDate:     requestBody.DueDate,
// 		TaskType:    requestBody.TaskType,
// 	}

// 	// create the course task
// 	err = db.WithContext(ctx).Model(&courseTask).Create(&courseTask).Preload("Course").First(&courseTask).Error
// 	if err != nil {
// 		logger.Error(ctx, "Error creating task", zap.Error(err))
// 		return c.Status(400).JSON(Response{Message: "An error occurred creating task", Success: false, Detail: err.Error()})

// 	}

// 	// Get all the list of enrolled courses
// 	err = db.WithContext(ctx).Model(&models.EnrolledCourse{}).
// 		Where("course_id = ?", courseTask.CourseID).
// 		Where("status =?", "ENROLLED").
// 		Where("end_date > ?", time.Now().UTC()).Find(&enrolledCourses).Error
// 	if err != nil {
// 		logger.Error(ctx, "error getting enrolled student in a course ", zap.Error(err))
// 		return c.Status(400).JSON(Response{Message: "An error occurred getting all the student enrolled in this course", Success: false, Detail: err.Error()})

// 	}

// 	// loop through the enrolled course and create the task for the student
// 	for _, enrolledCourse := range enrolledCourses {
// 		err := studentTask.CreateStudentTask(ctx, db, school.ID, *courseTask.CourseID, *enrolledCourse.StudentID)
// 		if err != nil {
// 			logger.Error(ctx, "Error creating students task under course task create", zap.Error(err))
// 			// delete the  course task
// 			db.WithContext(ctx).Model(&models.CourseTask{}).Delete(&courseTask)
// 			return c.Status(400).JSON(Response{Message: "error creating student task", Success: false, Detail: err.Error()})

// 		}
// 	}
// 	return c.Status(201).JSON(courseTask)
// }

// func CourseTaskDetailController(c *fiber.Ctx) error {
// 	var courseTask models.CourseTask
// 	var school models.School
// 	var staff models.Staff
// 	var questions []models.Question

// 	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// 	defer cancel()

// 	// Open the database connection
// 	db := database.DBConnection()
// 	defer database.CloseDB(db)

// 	// get the logged-in user
// 	user := c.Locals("user").(models.User)
// 	schoolCode := c.Query("school_code")
// 	courseID, err := uuid.Parse(c.Params("course_id"))
// 	if err != nil {
// 		return c.Status(400).JSON(
// 			Response{Message: "Invalid id passed", Success: false, Detail: err.Error()})
// 	}
// 	courseTaskID, err := uuid.Parse(c.Params("id"))
// 	if err != nil {
// 		return c.Status(400).JSON(
// 			Response{Message: "Invalid id passed", Success: false, Detail: err.Error()})
// 	}

// 	// get the school
// 	school, err = school.RetrieveBySchoolCode(ctx, db, schoolCode)
// 	if err != nil {
// 		return c.Status(400).JSON(
// 			Response{Message: "an error occured  getting school with code", Success: false, Detail: err.Error()})
// 	}
// 	// get the user staff
// 	staff, err = staff.RetrieveByUserAndSchool(ctx, db, school.ID, user.ID)
// 	if err != nil {

// 		return c.Status(400).JSON(
// 			Response{Message: "User is not a staff in this school", Success: false, Detail: err.Error()})
// 	}

// 	// get the course tasks
// 	err = db.WithContext(ctx).Model(&models.CourseTask{}).
// 		Where("id = ?", courseTaskID).
// 		Where("course_id = ?", courseID).
// 		First(&courseTask).Error
// 	if err != nil {
// 		return c.Status(400).JSON(
// 			Response{Message: "an error occured getting course tasks", Success: false, Detail: err.Error()})
// 	}

// 	// get the questions
// 	err = db.WithContext(ctx).Model(&models.Question{}).
// 		Where("course_task_id = ?", courseTaskID).
// 		Preload("Options").Find(&questions).Error
// 	if err != nil {
// 		logger.Error(ctx, "An error occured  getting task questions", zap.Error(err))
// 		return c.Status(500).JSON(
// 			Response{Message: "An error occured", Success: false, Detail: err.Error()})
// 	}

// 	serializer := serializers.CourseTaskDetailSerializer{
// 		ID:           &courseTask.ID,
// 		CourseID:     courseTask.CourseID,
// 		Title:        courseTask.Title,
// 		Description:  courseTask.Description,
// 		TotalPoint:   courseTask.TotalPoint,
// 		File:         courseTask.File,
// 		ClassAverage: courseTask.ClassAverage,
// 		TotalStudent: courseTask.TotalStudent,
// 		DueDate:      courseTask.DueDate,
// 		Questions:    questions,
// 	}

// 	return c.Status(200).JSON(serializer)
// }

// func CourseTaskUpdateController(c *fiber.Ctx) error {
// 	validator := adapters.NewValidate()
// 	fileUpload := adapters.NewFileUpload()

// 	var requestBody serializers.TaskUpdateRequestSerializer

// 	var courseTask models.CourseTask
// 	var school models.School
// 	var staff models.Staff
// 	var questions []models.Question

// 	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// 	defer cancel()

// 	// Open the database connection
// 	db := database.DBConnection()
// 	defer database.CloseDB(db)

// 	// get the logged-in user
// 	user := c.Locals("user").(models.User)
// 	schoolCode := c.Query("school_code")
// 	courseTaskID, err := uuid.Parse(c.Params("id"))
// 	if err != nil {
// 		return c.Status(400).JSON(
// 			Response{Message: "Invalid id passed", Success: false, Detail: err.Error()})
// 	}

// 	// get the school
// 	school, err = school.RetrieveBySchoolCode(ctx, db, schoolCode)
// 	if err != nil {
// 		return c.Status(400).JSON(
// 			Response{Message: "an error occured  getting school with code", Success: false, Detail: err.Error()})
// 	}

// 	// get the user staff
// 	staff, err = staff.RetrieveByUserAndSchool(ctx, db, school.ID, user.ID)
// 	if err != nil {

// 		return c.Status(400).JSON(
// 			Response{Message: "User is not a staff in this school", Success: false, Detail: err.Error()})
// 	}

// 	err = c.BodyParser(&requestBody)
// 	if err != nil {
// 		return c.Status(400).JSON(
// 			Response{Message: "Invalid json data", Success: false, Detail: err.Error()})
// 	}
// 	requestBody.File, err = fileUpload.UploadFile("file", c)

// 	err2 := validator.ValidateData(&requestBody)
// 	if err2 != nil {
// 		return c.Status(400).JSON(
// 			Response{Message: err2, Success: false, Detail: err2})
// 	}

// 	// get the course tasks
// 	err = db.WithContext(ctx).Model(&models.CourseTask{}).
// 		Where("id = ?", courseTaskID).
// 		Updates(&requestBody).First(&courseTask).Error
// 	if err != nil {

// 		return c.Status(400).JSON(
// 			Response{Message: "an error occured getting course tasks", Success: false, Detail: err.Error()})
// 	}

// 	// get the questions
// 	err = db.WithContext(ctx).Model(&models.Question{}).
// 		Where("course_task_id = ?", courseTaskID).
// 		Preload("Options").Find(&questions).Error
// 	if err != nil {
// 		logger.Error(ctx, "An error occured  getting task questions", zap.Error(err))
// 		return c.Status(500).JSON(
// 			Response{Message: "An error occured getting task question ", Success: false, Detail: err.Error()})
// 	}

// 	serializer := serializers.CourseTaskDetailSerializer{
// 		ID:           &courseTask.ID,
// 		CourseID:     courseTask.CourseID,
// 		Title:        courseTask.Title,
// 		Description:  courseTask.Description,
// 		TotalPoint:   courseTask.TotalPoint,
// 		File:         courseTask.File,
// 		ClassAverage: courseTask.ClassAverage,
// 		DueDate:      courseTask.DueDate,
// 		TaskType:     courseTask.TaskType,
// 		Questions:    questions,
// 	}

// 	return c.Status(200).JSON(serializer)
// }

// func CourseTaskDeleteController(c *fiber.Ctx) error {
// 	var courseTask models.CourseTask
// 	var school models.School
// 	var student models.Student

// 	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// 	defer cancel()

// 	// Open the database connection
// 	db := database.DBConnection()
// 	defer database.CloseDB(db)

// 	// get the logged-in user
// 	user := c.Locals("user").(models.User)
// 	schoolCode := c.Query("school_code")
// 	courseTaskID, err := uuid.Parse(c.Params("id"))
// 	if err != nil {
// 		return c.Status(400).JSON(
// 			Response{Message: "Invalid id passed", Success: false, Detail: err.Error()})
// 	}

// 	// get the school
// 	school, err = school.RetrieveBySchoolCode(ctx, db, schoolCode)
// 	if err != nil {
// 		return c.Status(400).JSON(
// 			Response{Message: "an error occured  getting school with code", Success: false, Detail: err.Error()})
// 	}

// 	// get the student
// 	student, err = student.RetrieveStudentByUserIDAndSchool(db, ctx, user.ID, school.ID)
// 	if err != nil {
// 		return c.Status(400).JSON(
// 			Response{Message: "student with this id does not exists", Success: false, Detail: err.Error()})
// 	}

// 	// get the course tasks
// 	err = db.WithContext(ctx).Model(&models.CourseTask{}).
// 		Where("id = ?", courseTaskID).
// 		First(&courseTask).
// 		Delete(&courseTask).Error
// 	if err != nil {

// 		return c.Status(400).JSON(
// 			Response{Message: "an error occured getting course tasks", Success: false, Detail: err.Error()})
// 	}

// 	return c.Status(204).JSON(fiber.Map{"message": "Successfully delete course task"})
// }

// func CourseTaskQuestionCreateController(c *fiber.Ctx) error {
// 	validator := adapters.NewValidate()
// 	fileUpload := adapters.NewFileUpload()
// 	var requestBody serializers.TaskQuestionCreateRequestSerializer

// 	var staff models.Staff
// 	var school models.School
// 	var courseTask models.CourseTask
// 	var questionOptions []*models.QuestionOption

// 	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// 	defer cancel()

// 	// Open the database connection
// 	db := database.DBConnection()
// 	defer database.CloseDB(db)

// 	// get the logged-in user
// 	user := c.Locals("user").(models.User)
// 	schoolCode := c.Query("school_code")

// 	err := c.BodyParser(&requestBody)
// 	if err != nil {
// 		return c.Status(400).JSON(
// 			Response{Message: "Invalid json data", Success: false, Detail: err.Error()})
// 	}

// 	requestBody.File, err = fileUpload.UploadFile("file", c)

// 	// validate the request body
// 	err2 := validator.ValidateData(&requestBody)
// 	if err2 != nil {
// 		return c.Status(404).JSON(
// 			Response{Message: err2, Success: false, Detail: err2})
// 	}

// 	// check if the task type is choice then validate
// 	if requestBody.QuestionType != nil {
// 		switch *requestBody.QuestionType {
// 		case "OPTION":
// 			options := c.FormValue("options")
// 			err = json.Unmarshal([]byte(options), &requestBody.Options)
// 			if err != nil {
// 				return c.Status(400).JSON(Response{Message: "an error occurred binding option data", Success: false, Detail: err.Error()})
// 			}
// 			// validate the option data
// 			err2 = validator.ValidateData(&requestBody.Options)
// 			if err2 != nil {
// 				logger.Error(ctx, " error binding json", zap.Error(err))
// 				return c.Status(400).JSON(Response{Message: "an error occurred binding option data", Success: false, Detail: err.Error()})

// 			}

// 			for _, option := range *requestBody.Options {
// 				err2 = validator.ValidateData(&option)
// 				if err2 != nil {
// 					return c.Status(400).JSON(Response{Message: err2, Success: false, Detail: err2})

// 				}
// 			}

// 			// add the options
// 			if requestBody.Options != nil {
// 				for _, option := range *requestBody.Options {
// 					questionOptions = append(questionOptions, &models.QuestionOption{
// 						Option:    option.Option,
// 						IsCorrect: option.IsCorrect,
// 						SchoolID:  &school.ID,
// 					})
// 				}
// 			}
// 		case "FILL_IN_THE_BLANK":
// 			fillInBlankAnswer := c.FormValue("fill_in_blank_answer")
// 			err = json.Unmarshal([]byte(fillInBlankAnswer), &requestBody.FillInBlankAnswer)
// 			if err != nil {
// 				return c.Status(400).JSON(Response{Message: "an error occurred binding fillInBlankAnswer data", Success: false, Detail: err.Error()})
// 			}
// 		}

// 	}

// 	school, err = school.RetrieveBySchoolCode(ctx, db, schoolCode)
// 	if err != nil {
// 		return c.Status(400).JSON(
// 			Response{Message: "Invalid school code ", Success: false, Detail: err.Error()})
// 	}

// 	// get the user staff
// 	staff, err = staff.RetrieveByUserAndSchool(ctx, db, school.ID, user.ID)
// 	if err != nil {

// 		return c.Status(400).JSON(
// 			Response{Message: "User is not a staff in this school", Success: false, Detail: err.Error()})
// 	}

// 	// check if the assigment exists
// 	err = db.WithContext(ctx).Model(&courseTask).Where("id = ?", requestBody.CourseTaskID).First(&courseTask).Error
// 	if err != nil {
// 		return c.Status(400).JSON(
// 			Response{Message: "Course task with this id does not exists", Success: false, Detail: err.Error()})
// 	}

// 	// create the question
// 	question := models.Question{
// 		SchoolID:          &school.ID,
// 		CourseTaskID:      requestBody.CourseTaskID,
// 		Question:          requestBody.Question,
// 		File:              requestBody.File,
// 		Point:             requestBody.Point,
// 		BooleanAnswer:     requestBody.BooleanAnswer,
// 		FillInBlankAnswer: requestBody.FillInBlankAnswer,
// 		Options:           questionOptions,
// 		QuestionType:      requestBody.QuestionType,
// 	}

// 	// create the question
// 	err = db.WithContext(ctx).Model(&question).Create(&question).First(&question).Error
// 	if err != nil {
// 		logger.Error(ctx, " Error creating question on CourseTaskQuestionCreateController", zap.Error(err))

// 		return c.Status(500).JSON(
// 			Response{Message: "an error from our end creating question we are working on it ", Success: false, Detail: err.Error()})
// 	}
// 	return c.Status(201).JSON(question)
// }

// func CourseTaskQuestionUpdateController(c *fiber.Ctx) error {
// 	validator := adapters.NewValidate()
// 	fileUpload := adapters.NewFileUpload()

// 	var requestBody serializers.TaskQuestionUpdateRequestSerializer

// 	var staff models.Staff
// 	var school models.School
// 	var question models.Question

// 	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// 	defer cancel()

// 	// Open the database connection
// 	db := database.DBConnection()
// 	defer database.CloseDB(db)

// 	// get the logged-in user
// 	user := c.Locals("user").(models.User)
// 	schoolCode := c.Query("school_code")
// 	questionID, err := uuid.Parse(c.Params("id"))
// 	if err != nil {
// 		return c.Status(400).JSON(
// 			Response{Message: "Invalid id passed", Success: false, Detail: err.Error()})
// 	}

// 	err = c.BodyParser(&requestBody)
// 	if err != nil {
// 		return c.Status(400).JSON(
// 			Response{Message: "Invalid json data", Success: false, Detail: err.Error()})
// 	}

// 	requestBody.File, err = fileUpload.UploadFile("file", c)

// 	// validate the request body
// 	err2 := validator.ValidateData(&requestBody)
// 	if err2 != nil {
// 		return c.Status(400).JSON(
// 			Response{Message: err2, Success: false, Detail: err2})
// 	}

// 	school, err = school.RetrieveBySchoolCode(ctx, db, schoolCode)
// 	if err != nil {
// 		return c.Status(400).JSON(
// 			Response{Message: "Invalid school code ", Success: false, Detail: err.Error()})
// 	}

// 	// get the user staff
// 	staff, err = staff.RetrieveByUserAndSchool(ctx, db, school.ID, user.ID)
// 	if err != nil {

// 		return c.Status(400).JSON(
// 			Response{Message: "User is not a staff in this school", Success: false, Detail: err.Error()})
// 	}

// 	// update the question
// 	err = db.WithContext(ctx).Model(&question).Where("id =?", questionID).Updates(&requestBody).First(&question).Error
// 	if err != nil {

// 		return c.Status(500).JSON(
// 			Response{Message: "an error from our end creating question we are working on it ", Success: false, Detail: err.Error()})
// 	}
// 	return c.Status(200).JSON(question)
// }

// func CourseTaskQuestionDeleteController(c *fiber.Ctx) error {

// 	var staff models.Staff
// 	var school models.School
// 	var question models.Question

// 	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// 	defer cancel()

// 	// Open the database connection
// 	db := database.DBConnection()
// 	defer database.CloseDB(db)

// 	// get the logged-in user
// 	user := c.Locals("user").(models.User)
// 	schoolCode := c.Query("school_code")
// 	questionID, err := uuid.Parse(c.Params("id"))
// 	if err != nil {
// 		return c.Status(400).JSON(
// 			Response{Message: "Invalid id passed", Success: false, Detail: err.Error()})
// 	}

// 	school, err = school.RetrieveBySchoolCode(ctx, db, schoolCode)
// 	if err != nil {
// 		return c.Status(400).JSON(
// 			Response{Message: "Invalid school code ", Success: false, Detail: err.Error()})
// 	}

// 	// get the user staff
// 	staff, err = staff.RetrieveByUserAndSchool(ctx, db, school.ID, user.ID)
// 	if err != nil {

// 		return c.Status(400).JSON(
// 			Response{Message: "User is not a staff in this school", Success: false, Detail: err.Error()})
// 	}

// 	// delete the question
// 	err = db.WithContext(ctx).Model(&question).Where("id =?", questionID).First(&question).Delete(&question).Error
// 	if err != nil {
// 		logger.Error(ctx, "an error occured deleting question ", zap.Error(err))
// 		return c.Status(500).JSON(
// 			Response{Message: "an error from our end deleting question we are working on it ", Success: false, Detail: err.Error()})
// 	}
// 	return c.Status(204).JSON(fiber.Map{"message": "Successfully delete question"})
// }

// func CourseTaskQuestionDetailController(c *fiber.Ctx) error {

// 	var staff models.Staff
// 	var school models.School
// 	var question models.Question

// 	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// 	defer cancel()

// 	// Open the database connection
// 	db := database.DBConnection()
// 	defer database.CloseDB(db)

// 	// get the logged-in user
// 	user := c.Locals("user").(models.User)
// 	schoolCode := c.Query("school_code")
// 	questionID, err := uuid.Parse(c.Params("id"))
// 	if err != nil {
// 		return c.Status(400).JSON(
// 			Response{Message: "Invalid id passed", Success: false, Detail: err.Error()})
// 	}

// 	school, err = school.RetrieveBySchoolCode(ctx, db, schoolCode)
// 	if err != nil {
// 		return c.Status(400).JSON(
// 			Response{Message: "Invalid school code ", Success: false, Detail: err.Error()})
// 	}

// 	// get the user staff
// 	staff, err = staff.RetrieveByUserAndSchool(ctx, db, school.ID, user.ID)
// 	if err != nil {

// 		return c.Status(400).JSON(
// 			Response{Message: "User is not a staff in this school", Success: false, Detail: err.Error()})
// 	}

// 	// delete the question
// 	err = db.WithContext(ctx).Model(&question).Where("id =?", questionID).Preload("Options").First(&question).Error
// 	if err != nil {

// 		return c.Status(500).JSON(
// 			Response{Message: "an error from our end creating question we are working on it ", Success: false, Detail: err.Error()})
// 	}
// 	return c.Status(200).JSON(question)
// }

// func CourseTaskListController(c *fiber.Ctx) error {

// 	var courseTasks []models.CourseTask
// 	var school models.School
// 	var staff models.Staff

// 	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// 	defer cancel()

// 	// Open the database connection
// 	db := database.DBConnection()
// 	defer database.CloseDB(db)

// 	// get the logged-in user
// 	user := c.Locals("user").(models.User)
// 	schoolCode := c.Query("school_code")
// 	courseID, err := uuid.Parse(c.Params("id"))
// 	if err != nil {
// 		return c.Status(400).JSON(
// 			Response{Message: "Invalid id passed", Success: false, Detail: err.Error()})
// 	}

// 	page := c.QueryInt("page", 1)    // default to page 1 if not provided
// 	limit := c.QueryInt("limit", 10) // default to 10 items per page if not provided

// 	// get the school
// 	school, err = school.RetrieveBySchoolCode(ctx, db, schoolCode)
// 	if err != nil {
// 		return c.Status(400).JSON(
// 			Response{Message: "an error occured  getting school with code", Success: false, Detail: err.Error()})
// 	}
// 	// get the user staff
// 	staff, err = staff.RetrieveByUserAndSchool(ctx, db, school.ID, user.ID)
// 	if err != nil {

// 		return c.Status(400).JSON(
// 			Response{Message: "User is not a staff in this school", Success: false, Detail: err.Error()})
// 	}

// 	var total int64
// 	db.Model(&models.CourseTask{}).
// 		Where("course_id =?", courseID).Count(&total)

// 	// get the course tasks
// 	err = db.Scopes(adapters.NewPaginateAdapter(limit, page).PaginatedResult).
// 		WithContext(ctx).Model(&models.CourseTask{}).Where("course_id = ?", courseID).Find(&courseTasks).Error
// 	if err != nil {

// 		return c.Status(400).JSON(
// 			Response{Message: "an error occured getting course tasks", Success: false, Detail: err.Error()})
// 	}
// 	return c.Status(200).JSON(fiber.Map{
// 		"total": total,
// 		"data":  courseTasks,
// 	})
// }

// func CourseTaskOptionUpdateController(c *fiber.Ctx) error {
// 	var questionOption models.QuestionOption
// 	var requestBody serializers.QuestionOptionRequestSerializer
// 	var staff models.Staff
// 	var school models.School

// 	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// 	defer cancel()

// 	// Open the database connection
// 	db := database.DBConnection()
// 	defer database.CloseDB(db)

// 	// get the logged-in user
// 	validateAdapters := adapters.NewValidate()
// 	user := c.Locals("user").(models.User)
// 	schoolCode := c.Query("school_code")
// 	optionID, err := uuid.Parse(c.Params("id"))
// 	if err != nil {
// 		return c.Status(400).JSON(
// 			Response{Message: "Invalid id passed", Success: false, Detail: err.Error()})
// 	}

// 	if err := c.BodyParser(&requestBody); err != nil {
// 		return c.Status(400).JSON(
// 			Response{Message: "invalid data passed ", Success: false, Detail: err.Error()})
// 	}

// 	// validate the data
// 	err2 := validateAdapters.ValidateData(&requestBody)
// 	if err2 != nil {
// 		return c.Status(400).JSON(Response{Message: err, Success: false, Detail: err.Error()})
// 	}

// 	// get the school
// 	school, err = school.RetrieveBySchoolCode(ctx, db, schoolCode)
// 	if err != nil {
// 		return c.Status(400).JSON(
// 			Response{Message: "an error occured  getting school with code", Success: false, Detail: err.Error()})
// 	}
// 	// get the user staff
// 	staff, err = staff.RetrieveByUserAndSchool(ctx, db, school.ID, user.ID)
// 	if err != nil {

// 		return c.Status(400).JSON(
// 			Response{Message: "User is not a staff in this school", Success: false, Detail: err.Error()})
// 	}

// 	err = db.WithContext(ctx).Model(&questionOption).Where("id = ?", optionID).Updates(&requestBody).First(&questionOption).Error
// 	if err != nil {
// 		if err == gorm.ErrRecordNotFound {
// 			return c.Status(400).JSON(
// 				Response{Message: "Option with this id does not exists", Success: false, Detail: err.Error()})
// 		}
// 		return c.Status(500).JSON(
// 			Response{Message: "an error occured updating question option", Success: false, Detail: err.Error()})
// 	}

// 	return c.Status(200).JSON(questionOption)
// }

// func CourseTaskOptionDeleteController(c *fiber.Ctx) error {
// 	var questionOption models.QuestionOption
// 	var staff models.Staff
// 	var school models.School

// 	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// 	defer cancel()

// 	// Open the database connection
// 	db := database.DBConnection()
// 	defer database.CloseDB(db)

// 	// get the logged-in user
// 	user := c.Locals("user").(models.User)
// 	schoolCode := c.Query("school_code")
// 	optionID, err := uuid.Parse(c.Params("id"))
// 	if err != nil {
// 		return c.Status(400).JSON(
// 			Response{Message: "Invalid id passed", Success: false, Detail: err.Error()})
// 	}

// 	// get the school
// 	school, err = school.RetrieveBySchoolCode(ctx, db, schoolCode)
// 	if err != nil {
// 		return c.Status(400).JSON(
// 			Response{Message: "an error occured  getting school with code", Success: false, Detail: err.Error()})
// 	}
// 	// get the user staff
// 	staff, err = staff.RetrieveByUserAndSchool(ctx, db, school.ID, user.ID)
// 	if err != nil {

// 		return c.Status(400).JSON(
// 			Response{Message: "User is not a staff in this school", Success: false, Detail: err.Error()})
// 	}

// 	err = db.WithContext(ctx).Model(&questionOption).Where("id = ?", optionID).First(&questionOption).Delete(&questionOption).Error
// 	if err != nil {
// 		if err == gorm.ErrRecordNotFound {
// 			return c.Status(400).JSON(
// 				Response{Message: "Option with this id does not exists", Success: false, Detail: err.Error()})
// 		}
// 		return c.Status(400).JSON(
// 			Response{Message: "an error occured deleting question option", Success: false, Detail: err.Error()})
// 	}

// 	return c.Status(204).JSON(
// 		Response{Message: "Successfully delete this option", Success: false, Detail: err.Error()})
// }

// func CourseTaskOptionDetailController(c *fiber.Ctx) error {
// 	var questionOption models.QuestionOption
// 	var staff models.Staff
// 	var school models.School

// 	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// 	defer cancel()

// 	// Open the database connection
// 	db := database.DBConnection()
// 	defer database.CloseDB(db)

// 	// get the logged-in user
// 	user := c.Locals("user").(models.User)
// 	schoolCode := c.Query("school_code")
// 	optionID, err := uuid.Parse(c.Params("id"))
// 	if err != nil {
// 		return c.Status(400).JSON(
// 			Response{Message: "Invalid id passed", Success: false, Detail: err.Error()})
// 	}

// 	// get the school
// 	school, err = school.RetrieveBySchoolCode(ctx, db, schoolCode)
// 	if err != nil {
// 		return c.Status(400).JSON(
// 			Response{Message: "an error occured  getting school with code", Success: false, Detail: err.Error()})
// 	}
// 	// get the user staff
// 	staff, err = staff.RetrieveByUserAndSchool(ctx, db, school.ID, user.ID)
// 	if err != nil {

// 		return c.Status(400).JSON(
// 			Response{Message: "User is not a staff in this school", Success: false, Detail: err.Error()})
// 	}

// 	err = db.WithContext(ctx).Model(&questionOption).Where("id = ?", optionID).First(&questionOption).Error
// 	if err != nil {
// 		if err == gorm.ErrRecordNotFound {
// 			return c.Status(400).JSON(
// 				Response{Message: "Option with this id does not exists ", Success: false, Detail: err.Error()})
// 		}
// 		return c.Status(500).JSON(
// 			Response{Message: "an error occured ", Success: false, Detail: err.Error()})
// 	}

// 	return c.Status(200).JSON(questionOption)
// }

// func CourseTaskOptionCreateController(c *fiber.Ctx) error {
// 	var newOption models.QuestionOption
// 	var staff models.Staff
// 	var school models.School
// 	var requestBody serializers.QuestionOptionCreateRequestSerializer
// 	var question models.Question

// 	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// 	defer cancel()

// 	// Open the database connection
// 	db := database.DBConnection()
// 	defer database.CloseDB(db)

// 	// Get the logged-in user
// 	user := c.Locals("user").(models.User)
// 	schoolCode := c.Query("school_code")

// 	// Get the school
// 	school, err := school.RetrieveBySchoolCode(ctx, db, schoolCode)
// 	if err != nil {
// 		return c.Status(400).JSON(
// 			Response{Message: "An error occurred while getting the school with code", Success: false, Detail: err.Error()})
// 	}

// 	// Get the user staff
// 	staff, err = staff.RetrieveByUserAndSchool(ctx, db, school.ID, user.ID)
// 	if err != nil {
// 		return c.Status(400).JSON(
// 			Response{Message: "User is not a staff in this school", Success: false, Detail: err.Error()})
// 	}

// 	validateAdapters := adapters.NewValidate()

// 	if err := c.BodyParser(&requestBody); err != nil {
// 		return c.Status(400).JSON(Response{Message: "invalid request body", Success: false, Detail: err.Error()})
// 	}

// 	vErr := validateAdapters.ValidateData(&requestBody)
// 	if vErr != nil {
// 		return c.Status(400).JSON(Response{Message: vErr, Success: false, Detail: vErr})
// 	}

// 	// Find the question by ID
// 	if err := db.WithContext(ctx).Model(&question).Where("id =? ", requestBody.QuestionID).Preload("Options").First(&question).Error; err != nil {
// 		return c.Status(400).JSON(
// 			Response{Message: "Question not found", Success: false, Detail: err.Error()})
// 	}

// 	// Create the new option
// 	newOption.SchoolID = &school.ID
// 	newOption.Option = requestBody.Option
// 	newOption.IsCorrect = requestBody.IsCorrect

// 	// Begin a database transaction to ensure data consistency
// 	tx := db.Begin()
// 	defer func() {
// 		if r := recover(); r != nil {
// 			tx.Rollback()
// 		}
// 	}()

// 	// Create the new option
// 	err = tx.WithContext(ctx).Model(&newOption).Create(&newOption).First(&newOption).Error
// 	if err != nil {
// 		tx.Rollback()
// 		return c.Status(500).JSON(
// 			Response{Message: "Failed to create the new option", Success: false, Detail: err.Error()})
// 	}

// 	// Establish the many-to-many relationship between the question and the new option
// 	err = tx.WithContext(ctx).Model(&question).Association("Options").Append(&newOption)
// 	if err != nil {
// 		tx.Rollback()
// 		return c.Status(500).JSON(
// 			Response{Message: "Failed to associate the option with the question", Success: false, Detail: err.Error()})
// 	}

// 	// Commit the transaction
// 	tx.Commit()

// 	// Return a success response
// 	return c.Status(201).JSON(newOption)
// }

// func StudentAllCourseAssigmentListController(c *fiber.Ctx) error {
// 	/*this is used to list the student course task*/
// 	var studentTasks []models.StudentTask
// 	var school models.School
// 	var student models.Student

// 	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// 	defer cancel()

// 	// Open the database connection
// 	db := database.DBConnection()
// 	defer database.CloseDB(db)

// 	// get the logged-in user
// 	user := c.Locals("user").(models.User)
// 	schoolCode := c.Query("school_code")

// 	school, err := school.RetrieveBySchoolCode(ctx, db, schoolCode)
// 	if err != nil {
// 		return c.Status(400).JSON(
// 			Response{Message: "Invalid school code ", Success: false, Detail: err.Error()})
// 	}

// 	student, err = student.RetrieveStudentByUserIDAndSchool(db, ctx, user.ID, school.ID)
// 	if err != nil {
// 		return c.Status(400).JSON(
// 			Response{Message: "user does is not associated to this school", Success: false, Detail: err.Error()})
// 	}

// 	page := c.QueryInt("page", 1)    // default to page 1 if not provided
// 	limit := c.QueryInt("limit", 10) // default to 10 items per page if not provided

// 	var total int64
// 	db.Model(&models.StudentTask{}).
// 		Where("student_id =?", student.ID).Count(&total)

// 	// get the student tasks
// 	err = db.Scopes(adapters.NewPaginateAdapter(limit, page).PaginatedResult).
// 		WithContext(ctx).Model(&models.StudentTask{}).Where("student_id = ?", student.ID).Preload("CourseTask").Preload("Student").Find(&studentTasks).Error
// 	if err != nil {

// 		return c.Status(400).JSON(
// 			Response{Message: "an error occured getting course tasks", Success: false, Detail: err.Error()})
// 	}

// 	return c.Status(200).JSON(fiber.Map{
// 		"total": total,
// 		"data":  studentTasks,
// 	})
// }

// func StudentAnswerTaskDetailController(c *fiber.Ctx) error {
// 	/* this is used to get the detail of the student */
// 	var studentTask models.StudentTask
// 	var studentAnswers []models.StudentAnswer
// 	var school models.School
// 	var student models.Student

// 	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// 	defer cancel()

// 	// Open the database connection
// 	db := database.DBConnection()
// 	defer database.CloseDB(db)

// 	// get the logged-in user
// 	user := c.Locals("user").(models.User)
// 	schoolCode := c.Query("school_code")
// 	studentTaskID, err := uuid.Parse(c.Params("id"))
// 	if err != nil {
// 		return c.Status(400).JSON(
// 			Response{Message: "Invalid uuid passed", Success: false, Detail: err.Error()})
// 	}

// 	school, err = school.RetrieveBySchoolCode(ctx, db, schoolCode)
// 	if err != nil {
// 		return c.Status(400).JSON(
// 			Response{Message: "Invalid school code", Success: false, Detail: err.Error()})
// 	}

// 	student, err = student.RetrieveStudentByUserIDAndSchool(db, ctx, user.ID, school.ID)
// 	if err != nil {
// 		return c.Status(400).JSON(
// 			Response{Message: "user does is not associated to this school", Success: false, Detail: err.Error()})
// 	}

// 	err = db.WithContext(ctx).Model(&studentTask).
// 		Where("id = ?", studentTaskID).
// 		Where("school_id = ?", school.ID).
// 		Preload("Student").
// 		Preload("CourseTask").
// 		First(&studentTask).Error
// 	if err != nil {
// 		logger.Error(ctx, "error getting student  task detail ", zap.Error(err))
// 		return c.Status(400).JSON(
// 			Response{Message: "student with this task does not exists", Success: false, Detail: err.Error()})
// 	}
// 	// get the student answer
// 	err = db.WithContext(ctx).Model(&studentAnswers).Where("student_task_id = ?", studentTask.ID).Preload("QuestionOption").Find(&studentAnswers).Error
// 	if err != nil {
// 		logger.Error(ctx, "error getting student answer on task detail ", zap.Error(err))
// 		return c.Status(400).JSON(
// 			Response{Message: "an error occured getting student answer with this task", Success: false, Detail: err.Error()})
// 	}

// 	studentTaskDetailSerializer := serializers.StudentTaskDetailSerializer{
// 		ID:             &studentTask.ID,
// 		Student:        studentTask.Student,
// 		StudentID:      studentTask.StudentID,
// 		CourseTask:     studentTask.CourseTask,
// 		CourseTaskID:   studentTask.CourseTaskID,
// 		StartTime:      studentTask.StartTime,
// 		DueDate:        studentTask.DueDate,
// 		SubmissionTime: studentTask.SubmissionTime,
// 		TotalPoint:     studentTask.TotalPoint,
// 		Status:         studentTask.Status,

// 		StudentAnswers: studentAnswers,
// 	}

// 	return c.Status(200).JSON(studentTaskDetailSerializer)
// }

// func StudentCourseAssigmentDetailController(c *fiber.Ctx) error {
// 	/* this is used to get the detail of the student */
// 	var studentTask models.StudentTask
// 	var serializer serializers.StudentCourseTaskDetailSerializer
// 	var school models.School
// 	var student models.Student
// 	var questions []models.Question

// 	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// 	defer cancel()

// 	// Open the database connection
// 	db := database.DBConnection()
// 	defer database.CloseDB(db)

// 	// get the logged-in user
// 	user := c.Locals("user").(models.User)
// 	schoolCode := c.Query("school_code")
// 	studentTaskID, err := uuid.Parse(c.Params("id"))
// 	if err != nil {
// 		return c.Status(400).JSON(
// 			Response{Message: "Invalid uuid passed", Success: false, Detail: err.Error()})
// 	}

// 	school, err = school.RetrieveBySchoolCode(ctx, db, schoolCode)
// 	if err != nil {
// 		return c.Status(400).JSON(
// 			Response{Message: "user does is not associated to this school", Success: false, Detail: err.Error()})
// 	}

// 	student, err = student.RetrieveStudentByUserIDAndSchool(db, ctx, user.ID, school.ID)
// 	if err != nil {
// 		return c.Status(400).JSON(
// 			Response{Message: "user does is not associated to this school", Success: false, Detail: err.Error()})
// 	}

// 	err = db.WithContext(ctx).Model(&studentTask).
// 		Where("id = ?", studentTaskID).
// 		Where("school_id = ?", school.ID).
// 		Preload("Student").
// 		Preload("CourseTask").
// 		First(&studentTask).Error
// 	if err != nil {
// 		logger.Error(ctx, "error getting student  task detail ", zap.Error(err))
// 		return c.Status(400).JSON(
// 			Response{Message: "student with this task does not exists", Success: false, Detail: err.Error()})
// 	}

// 	// get the questions
// 	err = db.WithContext(ctx).Model(&models.Question{}).
// 		Where("course_task_id = ?", studentTask.CourseTaskID).
// 		Preload("Options").Find(&questions).Error
// 	if err != nil {
// 		logger.Error(ctx, "An error occured  getting task questions", zap.Error(err))
// 		return c.Status(500).JSON(
// 			Response{Message: "An error occured  getting  task questions", Success: false, Detail: err.Error()})
// 	}

// 	// now let's copy the data's with copier
// 	// Copy data from studentTask to serializer
// 	err = copier.Copy(&serializer, &studentTask)
// 	if err != nil {
// 		logger.Error(ctx, "Error copying data's", zap.Error(err))
// 		return c.Status(500).JSON(Response{Message: "Error copying data's", Success: false, Detail: err.Error()})
// 	}

// 	// copy from the questions to serializers.Questions
// 	err = copier.Copy(&serializer.CourseTask.Question, &questions)
// 	if err != nil {
// 		logger.Error(ctx, "Error copying data's", zap.Error(err))
// 	}

// 	return c.Status(200).JSON(serializer)
// }
// func StudentTaskAnswerQuestionController(c *fiber.Ctx) error {
// 	var requestBody serializers.StudentAnswerRequestSerializer
// 	var school models.School
// 	var student models.Student
// 	var studentTask models.StudentTask
// 	var studentAnswer models.StudentAnswer
// 	var foundAnswer models.StudentAnswer
// 	var question models.Question
// 	var questionOption models.QuestionOption
// 	var update bool

// 	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// 	defer cancel()

// 	// Open the database connection
// 	db := database.DBConnection()
// 	defer database.CloseDB(db)

// 	pointerAdapters := adapters.NewPointer()

// 	// get the logged-in user
// 	user := c.Locals("user").(models.User)
// 	schoolCode := c.Query("school_code")
// 	studentTaskID, err := uuid.Parse(c.Params("id"))
// 	if err != nil {
// 		return c.Status(400).JSON(
// 			Response{Message: "Invalid uuid passed", Success: false, Detail: err.Error()})
// 	}

// 	school, err = school.RetrieveBySchoolCode(ctx, db, schoolCode)
// 	if err != nil {
// 		return c.Status(400).JSON(
// 			Response{Message: "Invalid school code", Success: false, Detail: err.Error()})
// 	}

// 	student, err = student.RetrieveStudentByUserIDAndSchool(db, ctx, user.ID, school.ID)
// 	if err != nil {
// 		return c.Status(400).JSON(
// 			Response{Message: "user does is not associated to this school", Success: false, Detail: err.Error()})
// 	}

// 	err = c.BodyParser(&requestBody)
// 	if err != nil {
// 		logger.Error(ctx, "invalid json passed on  StudentTaskAnswerQuestionController", zap.Error(err))
// 		return c.Status(400).JSON(
// 			Response{Message: "Invalid json passed", Success: false, Detail: err.Error()})
// 	}

// 	err = db.WithContext(ctx).Model(&studentTask).Where("id =?", studentTaskID).First(&studentTask).Error
// 	if err != nil {
// 		logger.Error(ctx, "error getting student task  on StudentTaskAnswerQuestionController", zap.Error(err))
// 		return c.Status(400).JSON(
// 			Response{Message: "Student task does not exist", Success: false, Detail: err.Error()})
// 	}

// 	if studentTask.Status != nil {
// 		if *studentTask.Status == "SUBMITTED" {
// 			return c.Status(400).JSON(Response{Message: "task has been submitted", Success: false, Detail: errors.New("task has been submitted")})
// 		}
// 	}

// 	// find the question
// 	err = db.WithContext(ctx).Model(&question).Where("id =?", requestBody.QuestionID).First(&question).Error
// 	if err != nil {
// 		logger.Error(ctx, "error getting question on  StudentTaskAnswerQuestionController ", zap.Error(err))
// 		return c.Status(400).JSON(
// 			Response{Message: "question does not exist or you dont have a pending task", Success: false, Detail: err.Error()})
// 	}

// 	// check if the student has answered this before
// 	err = db.WithContext(ctx).Model(models.StudentAnswer{}).
// 		Where("student_task_id =?", studentTaskID).
// 		Where("question_id =?", question.ID).First(&foundAnswer).Error
// 	if foundAnswer.QuestionID != nil {
// 		update = true
// 		studentAnswer = foundAnswer
// 	} else {
// 		update = false
// 		studentAnswer = models.StudentAnswer{
// 			SchoolID:      &school.ID,
// 			StudentTaskID: &studentTask.ID,
// 			QuestionID:    &question.ID,
// 		}
// 	}

// 	if question.QuestionType != nil {
// 		if *question.QuestionType != *requestBody.AnswerType {
// 			return c.Status(400).JSON(Response{Message: fmt.Sprintf("not a valid answer type for %s", *question.QuestionType),
// 				Success: false, Detail: errors.New("not a valid answer type")})
// 		}

// 		switch *question.QuestionType {
// 		case "TEXT":
// 			studentAnswer.Answer = requestBody.TextAnswer
// 			// todo: use chat gpt for the answer before giving point
// 			studentAnswer.Point = question.Point
// 		case "OPTION":
// 			// filter for the option
// 			questionOption, err := questionOption.FindQuestionOptionByQuestionAndOptionID(ctx, db, *requestBody.QuestionID, *requestBody.OptionID)
// 			if err != nil {
// 				return c.Status(400).JSON(Response{Message: "Not a valid option id passed", Success: false, Detail: err.Error()})
// 			}
// 			studentAnswer.Answer = questionOption.Option
// 			if questionOption.IsCorrect != nil {
// 				if *questionOption.IsCorrect {
// 					studentAnswer.Point = question.Point
// 				} else {
// 					studentAnswer.Point = pointerAdapters.UIntPointer(0)
// 				}
// 			}
// 		case "BOOLEAN":
// 			if question.BooleanAnswer != nil && requestBody.BooleanAnswer != nil {
// 				if question.BooleanAnswer == requestBody.BooleanAnswer {
// 					studentAnswer.Answer = pointerAdapters.StringPointer(strconv.FormatBool(*requestBody.BooleanAnswer))
// 					studentAnswer.Point = question.Point
// 				} else {
// 					studentAnswer.Answer = pointerAdapters.StringPointer(strconv.FormatBool(*requestBody.BooleanAnswer))
// 					studentAnswer.Point = pointerAdapters.UIntPointer(0)
// 				}
// 			}
// 		case "FILL_IN_THE_BLANK":
// 			if question.FillInBlankAnswer != nil && requestBody.FillInBlankAnswer != nil {
// 				answer := strings.Join(*requestBody.FillInBlankAnswer, ", ")
// 				studentAnswer.Answer = &answer
// 				questionAnswer := *question.FillInBlankAnswer
// 				requestBodyAnswer := *requestBody.FillInBlankAnswer

// 				// Convert both answers to lowercase before comparing
// 				for i, val := range questionAnswer {
// 					questionAnswer[i] = strings.ToLower(val)
// 				}
// 				for i, val := range requestBodyAnswer {
// 					requestBodyAnswer[i] = strings.ToLower(val)
// 				}

// 				// Check if the slices are equal
// 				if reflect.DeepEqual(questionAnswer, requestBodyAnswer) {
// 					studentAnswer.Point = question.Point
// 				} else {
// 					studentAnswer.Point = pointerAdapters.UIntPointer(0)
// 				}
// 			}
// 		}

// 	}

// 	if update == true {
// 		err = db.WithContext(ctx).Model(models.StudentAnswer{}).Where("id =?", studentAnswer.ID).Updates(&studentAnswer).
// 			Preload("Question").First(&studentAnswer).Error
// 		if err != nil {
// 			logger.Error(ctx, " error creating student answer on  StudentTaskAnswerQuestionController", zap.Error(err))
// 			return c.Status(400).JSON(
// 				Response{Message: "error updating student answer", Success: false, Detail: err.Error()})
// 		}
// 	} else {
// 		err = db.WithContext(ctx).Model(models.StudentAnswer{}).Create(&studentAnswer).Preload("Question").First(&studentAnswer).Error
// 		if err != nil {
// 			logger.Error(ctx, " error creating student answer on  StudentTaskAnswerQuestionController", zap.Error(err))
// 			return c.Status(400).JSON(
// 				Response{Message: "error creating student answer", Success: false, Detail: err.Error()})
// 		}
// 	}

// 	return c.Status(200).JSON(studentAnswer)
// }

// func StudentSubmitTaskController(c *fiber.Ctx) error {
// 	var studentTask models.StudentTask
// 	var studentAnswers []models.StudentAnswer
// 	var school models.School
// 	var student models.Student
// 	var courseTask models.CourseTask
// 	var requestBody serializers.StudentTaskSubmitRequestSerializer
// 	var course models.Course

// 	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// 	defer cancel()

// 	// Open the database connection
// 	db := database.DBConnection()
// 	defer database.CloseDB(db)

// 	// get the logged-in user
// 	user := c.Locals("user").(models.User)
// 	schoolCode := c.Query("school_code")
// 	validateAdapters := adapters.NewValidate()

// 	err := c.BodyParser(&requestBody)
// 	if err != nil {
// 		return c.Status(400).JSON(
// 			Response{Message: "Invalid json passed", Success: false, Detail: err.Error()})
// 	}

// 	vErr := validateAdapters.ValidateData(&requestBody)
// 	if vErr != nil {
// 		return c.Status(400).JSON(Response{Message: vErr, Success: false, Detail: vErr})
// 	}

// 	school, err = school.RetrieveBySchoolCode(ctx, db, schoolCode)
// 	if err != nil {
// 		return c.Status(400).JSON(
// 			Response{Message: "Invalid school code", Success: false, Detail: err.Error()})
// 	}

// 	student, err = student.RetrieveStudentByUserIDAndSchool(db, ctx, user.ID, school.ID)
// 	if err != nil {
// 		return c.Status(400).JSON(
// 			Response{Message: "user does is not associated to this school", Success: false, Detail: err.Error()})
// 	}

// 	// check the task if it has not been submitted, and also it is owned by the student
// 	err = db.WithContext(ctx).Model(&studentTask).
// 		Where("student_id =?", student.ID).
// 		Where("id =? ", requestBody.StudentTaskID).
// 		Where("status = ?", "PENDING").First(&studentTask).Error
// 	if err != nil {
// 		if err == gorm.ErrRecordNotFound {
// 			return c.Status(400).JSON(
// 				Response{Message: "No pending task found with this student assigment id", Success: false, Detail: err.Error()})
// 		}
// 		logger.Error(ctx, "Unable to filter for student  task ", zap.Error(err))
// 		return c.Status(500).JSON(
// 			Response{Message: "An error occurred on our end ", Success: false, Detail: err.Error()})
// 	}

// 	// submit the task and calculate the score
// 	err = db.WithContext(ctx).Model(&models.StudentAnswer{}).Where("student_task_id = ?", studentTask.ID).Find(&studentAnswers).Error
// 	if err != nil {
// 		logger.Error(ctx, "Error getting student answers", zap.Error(err))
// 	}

// 	// loop through all the answers and append the point
// 	var totalPoint uint
// 	for _, studentAnswer := range studentAnswers {
// 		if studentAnswer.Point != nil {
// 			totalPoint += uint(*studentAnswer.Point)
// 		}
// 	}
// 	studentTask.TotalPoint = &totalPoint
// 	submissionTime := time.Now()

// 	// update the student task
// 	err = db.WithContext(ctx).Model(&models.StudentTask{}).Where("id = ?", studentTask.ID).
// 		Update("total_point", totalPoint).
// 		//Update("status", "SUBMITTED").
// 		Update("submission_time", submissionTime).
// 		Preload("CourseTask").
// 		Preload("Student").
// 		First(&studentTask).Error
// 	if err != nil {
// 		logger.Error(ctx, "update student task", zap.Error(err))
// 		return c.Status(400).JSON(
// 			Response{Message: "Error updating student total point", Success: false, Detail: err.Error()})
// 	}

// 	err = studentTask.CalculateTaskPoint(ctx, db, *studentTask.StudentID, *studentTask.CourseTask.CourseID, *studentTask.CourseTask.TaskType)
// 	if err != nil {
// 		logger.Error(ctx, "error calculating task point", zap.Error(err))
// 		return c.Status(500).JSON(
// 			Response{Message: "Error calculating task point", Success: false, Detail: err.Error()})
// 	}

// 	err = courseTask.CalculateClassAverage(ctx, db, *studentTask.CourseTaskID, *studentTask.CourseTask.TotalPoint)
// 	if err != nil {
// 		logger.Error(ctx, "error calculating class average", zap.Error(err))
// 		return c.Status(500).JSON(Response{Message: "Error calculating class average", Success: false, Detail: err.Error()})
// 	}

// 	err = course.CalculateCoursePerformance(ctx, db, *studentTask.CourseTask.CourseID)
// 	if err != nil {
// 		logger.Error(ctx, "error calculating course performance", zap.Error(err))
// 		return c.Status(500).JSON(Response{Message: "Error calculating course performance", Success: false, Detail: err.Error()})
// 	}
// 	return c.Status(200).JSON(studentTask)
// }
