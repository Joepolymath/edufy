package controllers

// import (
// 	"Learnium/adapters"
// 	"Learnium/database"
// 	"Learnium/logger"
// 	"Learnium/models"
// 	"Learnium/serializers"
// 	"context"
// 	"github.com/gofiber/fiber/v2"
// 	"github.com/google/uuid"
// 	"go.uber.org/zap"
// 	"gorm.io/gorm"
// 	"strings"
// 	"time"
// )

// func StudentListController(c *fiber.Ctx) error {
// 	schoolCode := c.Query("school_code")
// 	page := c.QueryInt("page", 1)    // Page number, default to 1
// 	limit := c.QueryInt("limit", 10) // Number of records per page, default to 10

// 	var students []serializers.StudentListSerializer
// 	var totalStudents int64
// 	var school models.School

// 	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// 	defer cancel()

// 	// Open the database connection
// 	db := database.DBConnection()
// 	defer database.CloseDB(db)

// 	search := c.Query("search")
// 	status := c.Query("status")
// 	if status != "" {
// 		status = strings.ToUpper(status)
// 	}

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

// 	// Calculate the offset based on the page and limit
// 	offset := (page - 1) * limit

// 	// Query for a paginated student and count a total student
// 	query := db.WithContext(ctx).
// 		Model(&models.Student{}).
// 		Where("school_id = ?", school.ID).
// 		Preload("Class").
// 		Preload("User").
// 		Preload("EnrollmentAdmission").
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
// 	err = query.Limit(limit).Find(&students).Error

// 	if err != nil {
// 		if err != gorm.ErrRecordNotFound {
// 			return c.Status(500).JSON(Response{Message: "Error on our end filtering student %s", Success: false, Detail: err})
// 		}
// 	}

// 	// Count a total student
// 	db.Model(&models.Student{}).Where("school_id = ?", school.ID).Count(&totalStudents)

// 	// Check if there are more records (has_next)
// 	hasMore := (offset + limit) < int(totalStudents)

// 	response := fiber.Map{
// 		"total":    totalStudents,
// 		"page":     page,
// 		"offset":   offset,
// 		"has_next": hasMore,
// 		"data":     students,
// 	}

// 	return c.Status(200).JSON(response)
// }

// func StudentNoteCreateController(c *fiber.Ctx) error {
// 	/* This is used to create not by the teacher of the students*/

// 	var student models.Student
// 	var staff models.Staff
// 	var school models.School
// 	var note models.Note
// 	var requestBody serializers.NoteCreateSerializer

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
// 		return c.Status(400).JSON(Response{Message: "School with this code does not exist", Success: false, Detail: err.Error()})

// 	}

// 	staff, err = staff.RetrieveByUserAndSchool(ctx, db, school.ID, user.ID)
// 	if err != nil {
// 		return c.Status(400).JSON(Response{Message: "staff does not exist with this user", Success: false, Detail: err})

// 	}

// 	err = c.BodyParser(&requestBody)
// 	if err != nil {
// 		logger.Error(ctx, "Error parsing request body", zap.Error(err))
// 		return c.Status(400).JSON(Response{Message: "invalid request body", Success: false, Detail: err.Error()})
// 	}

// 	vErr := validateAdapter.ValidateData(&requestBody)
// 	if vErr != nil {
// 		return c.Status(400).JSON(Response{Message: vErr, Success: false, Detail: vErr})

// 	}

// 	student, err = student.RetrieveStudentByIDAndSchool(ctx, db, *requestBody.StudentID, school.ID)
// 	if err != nil {
// 		return c.Status(400).JSON(Response{Message: "student with this  id does not  exists", Success: false, Detail: err})

// 	}

// 	note = models.Note{
// 		SchoolID:  &school.ID,
// 		StudentID: &student.ID,
// 		StaffID:   &staff.ID,
// 		NoteType:  requestBody.NoteType,
// 		Share:     requestBody.Share,
// 		Title:     requestBody.Title,
// 		Note:      requestBody.Note,
// 	}

// 	err = db.WithContext(ctx).Model(&note).Create(&note).Preload("Staff.User").First(&note).Error
// 	if err != nil {
// 		logger.Error(ctx, "Error creating student ", zap.Error(err))
// 		return c.Status(400).JSON(Response{Message: "error creating  note", Success: false, Detail: err})

// 	}

// 	return c.Status(201).JSON(note)
// }

// func StudentNoteListController(c *fiber.Ctx) error {
// 	/* this is used to list all the note tied to a student */

// 	var student models.Student
// 	var staff models.Staff
// 	var school models.School
// 	var notes []models.Note

// 	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// 	defer cancel()

// 	// Open the database connection
// 	db := database.DBConnection()
// 	defer database.CloseDB(db)

// 	page := c.QueryInt("page", 1)    // Page number, default to 1
// 	limit := c.QueryInt("limit", 10) // Number of records per page, default to 10

// 	user := c.Locals("user").(models.User)
// 	studentID, err := uuid.Parse(c.Params("id"))
// 	if err != nil {
// 		return c.Status(400).JSON(Response{Message: "Invalid studentID passed ", Success: false, Detail: err})

// 	}

// 	search := c.Query("search")
// 	schoolCode := c.Query("school_code")

// 	school, err = school.RetrieveBySchoolCode(ctx, db, schoolCode)
// 	if err != nil {
// 		return c.Status(400).JSON(Response{Message: "School with this code does not exist", Success: false, Detail: err.Error()})

// 	}

// 	student, err = student.RetrieveStudentByUserIDAndSchool(db, ctx, user.ID, school.ID)
// 	if err != nil {
// 		if err == gorm.ErrRecordNotFound {
// 			isOk := school.IsSchoolAdminOrOwner(ctx, db, schoolCode, user.ID)
// 			if !isOk {
// 				staff, err = staff.RetrieveByUserAndSchool(ctx, db, school.ID, user.ID)
// 				if err != nil {
// 					return c.Status(400).JSON(Response{Message: "staff does not exist with this user", Success: false, Detail: err})

// 				}
// 			}
// 		} else {
// 			return c.Status(400).JSON(Response{Message: "You dont have access to view this note", Success: false, Detail: err})

// 		}
// 	}

// 	if student.ID != uuid.Nil {
// 		if student.ID != studentID {
// 			return c.Status(400).JSON(Response{Message: "You dont have access to view this note", Success: false, Detail: err})
// 		}
// 	}

// 	query := db.Scopes(adapters.NewPaginateAdapter(limit, page).PaginatedResult).WithContext(ctx).Model(&models.Note{}).
// 		Where("student_id =?", studentID).
// 		Where("school_id =?", school.ID).
// 		Preload("Staff.User").
// 		Preload("Student.User")
// 	if search != "" {
// 		query = query.Where("CONCAT(users.first_name, ' ', users.last_name) ILIKE ?", "%"+search+"%")
// 	}

// 	err = query.Find(&notes).Error
// 	if err != nil {
// 		logger.Error(ctx, "error occured filtering notes", zap.Error(err))
// 		return c.Status(400).JSON(Response{Message: "error occured filtering notes", Success: false, Detail: err})
// 	}

// 	var total int64
// 	db.Model(&models.Note{}).
// 		Where("student_id =?", studentID).
// 		Where("school_id =?", school.ID).Count(&total)

// 	return c.Status(200).JSON(fiber.Map{"data": notes, "total": total})
// }

// func StudentNoteUpdateController(c *fiber.Ctx) error {
// 	/* this is used to update the note*/
// 	var staff models.Staff
// 	var school models.School
// 	var note models.Note
// 	var requestBody serializers.NoteUpdateSerializer

// 	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// 	defer cancel()

// 	// Open the database connection
// 	db := database.DBConnection()
// 	defer database.CloseDB(db)

// 	validateAdapter := adapters.NewValidate()
// 	user := c.Locals("user").(models.User)
// 	schoolCode := c.Query("school_code")
// 	noteID, err := uuid.Parse(c.Params("id"))
// 	if err != nil {
// 		return c.Status(400).JSON(Response{Message: "Invalid note id passed", Success: false, Detail: err})
// 	}

// 	school, err = school.RetrieveBySchoolCode(ctx, db, schoolCode)
// 	if err != nil {
// 		return c.Status(404).JSON(Response{
// 			Message: "school with this id does not exists",
// 			Success: false,
// 			Detail:  err,
// 		})
// 	}

// 	// school admin or staff
// 	isOk := school.IsSchoolAdminOrOwner(ctx, db, schoolCode, user.ID)
// 	if !isOk {
// 		staff, err = staff.RetrieveByUserAndSchool(ctx, db, school.ID, user.ID)
// 		if err != nil {
// 			return c.Status(400).JSON(Response{Message: "user does not exist as a staff in this school", Success: false, Detail: err})
// 		}
// 	}
// 	err = c.BodyParser(&requestBody)
// 	if err != nil {
// 		logger.Error(ctx, "Error parsing request body", zap.Error(err))
// 		return c.Status(400).JSON(Response{Message: "invalid request body", Success: false, Detail: err.Error()})
// 	}

// 	vErr := validateAdapter.ValidateData(&requestBody)
// 	if vErr != nil {
// 		return c.Status(400).JSON(Response{Message: vErr, Success: false, Detail: vErr})
// 	}

// 	err = db.WithContext(ctx).Model(&note).Where("id = ?", noteID).Where("school_id =?", school.ID).Updates(&requestBody).
// 		Preload("Staff.User").
// 		Preload("Student.User").First(&note).Error
// 	if err != nil {
// 		if err == gorm.ErrRecordNotFound {
// 			return c.Status(400).JSON(Response{Message: "note with this id does not exist", Success: false, Detail: err})

// 		}
// 		return c.Status(400).JSON(
// 			Response{Message: "An error occured updating  note", Success: false, Detail: err})
// 	}
// 	return c.Status(200).JSON(note)
// }

// func StudentNoteDeleteController(c *fiber.Ctx) error {
// 	/* this is used to update the note*/
// 	var staff models.Staff
// 	var school models.School
// 	var note models.Note

// 	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// 	defer cancel()

// 	// Open the database connection
// 	db := database.DBConnection()
// 	defer database.CloseDB(db)

// 	user := c.Locals("user").(models.User)
// 	schoolCode := c.Query("school_code")
// 	noteID, err := uuid.Parse(c.Params("id"))
// 	if err != nil {
// 		return c.Status(400).JSON(
// 			Response{Message: "Invalid note id passed", Success: false, Detail: err})
// 	}

// 	school, err = school.RetrieveBySchoolCode(ctx, db, schoolCode)
// 	if err != nil {
// 		return c.Status(404).JSON(Response{
// 			Message: "school with this id does not exists",
// 			Success: false,
// 			Detail:  err,
// 		})
// 	}

// 	// school admin or staff
// 	isOk := school.IsSchoolAdminOrOwner(ctx, db, schoolCode, user.ID)
// 	if !isOk {
// 		staff, err = staff.RetrieveByUserAndSchool(ctx, db, school.ID, user.ID)
// 		if err != nil {
// 			return c.Status(400).JSON(
// 				Response{Message: "user does not exist as a staff in this school", Success: false, Detail: err})
// 		}
// 	}

// 	err = db.WithContext(ctx).Model(&note).Where("id = ?", noteID).Where("school_id =?", school.ID).Delete(&note).First(&note).Error
// 	if err != nil {
// 		if err == gorm.ErrRecordNotFound {
// 			return c.Status(400).JSON(Response{Message: "note with this id does not exist", Success: false, Detail: err})
// 		}
// 		return c.Status(400).JSON(
// 			Response{Message: "An error occured updating  note", Success: false, Detail: err})
// 	}
// 	return c.Status(204).JSON(fiber.Map{"message": "Successfully delete note"})
// }

// func StudentNoteDetailController(c *fiber.Ctx) error {
// 	/* this is used to update the note*/
// 	var staff models.Staff
// 	var school models.School
// 	var note models.Note

// 	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// 	defer cancel()

// 	// Open the database connection
// 	db := database.DBConnection()
// 	defer database.CloseDB(db)

// 	user := c.Locals("user").(models.User)
// 	schoolCode := c.Query("school_code")
// 	noteID, err := uuid.Parse(c.Params("id"))
// 	if err != nil {
// 		return c.Status(400).JSON(
// 			Response{Message: "Invalid note id passed", Success: false, Detail: err})
// 	}

// 	school, err = school.RetrieveBySchoolCode(ctx, db, schoolCode)
// 	if err != nil {
// 		return c.Status(404).JSON(Response{
// 			Message: "school with this id does not exists",
// 			Success: false,
// 			Detail:  err,
// 		})
// 	}

// 	err = db.WithContext(ctx).Model(&note).Where("id = ?", noteID).Where("school_id =?", school.ID).
// 		Preload("Staff.User").
// 		Preload("Student.User").
// 		First(&note).Error
// 	if err != nil {
// 		if err == gorm.ErrRecordNotFound {
// 			return c.Status(400).JSON(Response{Message: "note with this id does not exist", Success: false, Detail: err})
// 		}
// 		return c.Status(400).JSON(Response{Message: "nAn error occured updating  note", Success: false, Detail: err})
// 	}

// 	// school admin or staff
// 	isOk := school.IsSchoolAdminOrOwner(ctx, db, schoolCode, user.ID)
// 	if !isOk {
// 		staff, err = staff.RetrieveByUserAndSchool(ctx, db, school.ID, user.ID)
// 		if err != nil {
// 			if note.Student.User.ID != user.ID {
// 				return c.Status(400).JSON(Response{Message: "you dont have permission", Success: false, Detail: err})
// 			}
// 		}
// 	}

// 	return c.Status(200).JSON(note)
// }

// func ClinicVisitationListController(c *fiber.Ctx) error {
// 	/* This is used to list all times a user has  visited a clinic*/
// 	var clinicVisitations []models.ClinicVisitation
// 	var school models.School
// 	var staff models.Staff
// 	var student models.Student

// 	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// 	defer cancel()

// 	// Open the database connection
// 	db := database.DBConnection()
// 	defer database.CloseDB(db)

// 	search := c.Query("search")
// 	schoolCode := c.Query("school_code")
// 	page := c.QueryInt("page", 1)    // Page number, default to 1
// 	limit := c.QueryInt("limit", 10) // Number of records per page, default to 10
// 	studentID, err := uuid.Parse(c.Params("id"))
// 	if err != nil {
// 		return c.Status(400).JSON(
// 			Response{Message: "Invalid id passed", Success: false, Detail: err})
// 	}

// 	school, err = school.RetrieveBySchoolCode(ctx, db, schoolCode)
// 	if err != nil {
// 		return c.Status(404).JSON(Response{
// 			Message: "school with this id does not exists",
// 			Success: false,
// 			Detail:  err,
// 		})
// 	}

// 	user := c.Locals("user").(models.User)

// 	// if you are not an admin,staff and also not the user you are trying to get the clinic visitation, for then you are not authorized
// 	student, err = student.RetrieveStudentByIDAndSchool(ctx, db, studentID, school.ID)
// 	if err != nil {
// 		isOk := school.IsSchoolAdminOrOwner(ctx, db, schoolCode, user.ID)
// 		if isOk == false {
// 			staff, err = staff.RetrieveByUserAndSchool(ctx, db, school.ID, user.ID)
// 			if err != nil {
// 				return c.Status(400).JSON(Response{Message: "you dont have permission", Success: false, Detail: err})
// 			}
// 		}
// 	}

// 	query := db.Scopes(adapters.NewPaginateAdapter(limit, page).PaginatedResult).WithContext(ctx).Model(&models.ClinicVisitation{}).
// 		Where("student_id = ?", studentID).
// 		Where("school_id = ?", school.ID).
// 		Preload("Student").
// 		Preload("School")

// 	var total int64
// 	db.WithContext(ctx).Model(&models.ClinicVisitation{}).
// 		Where("student_id = ?", studentID).
// 		Where("school_id = ?", school.ID).Count(&total)

// 	if search != "" {
// 		query = query.Where("name ILIKE ? or description ILIKE ? or doctors_name ILIKE ?", "%"+search+"%", "%"+search+"%", "%"+search+"%")
// 	}

// 	err = query.Find(&clinicVisitations).Error
// 	if err != nil {
// 		logger.Error(ctx, "error getting clinic visitation", zap.Error(err))
// 		return c.Status(400).JSON(fiber.Map{"error": "There was an error getting clinic visitation"})
// 	}
// 	return c.Status(200).JSON(fiber.Map{"data": clinicVisitations, "total": total})
// }

// func ClinicVisitationCreateController(c *fiber.Ctx) error {
// 	/* This is used to create a clinic visitation */
// 	var school models.School
// 	var staff models.Staff
// 	var student models.Student
// 	var clinicVisitation models.ClinicVisitation
// 	var requestBody serializers.ClinicVisitationCreateRequestSerializer

// 	// can be added by the staff
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
// 		return c.Status(404).JSON(Response{
// 			Message: "school with this id does not exists",
// 			Success: false,
// 			Detail:  err,
// 		})
// 	}

// 	err = c.BodyParser(&requestBody)
// 	if err != nil {
// 		logger.Error(ctx, "Error parsing request body", zap.Error(err))
// 		return c.Status(400).JSON(Response{Message: "invalid request body", Success: false, Detail: err.Error()})
// 	}

// 	vErr := validateAdapter.ValidateData(&requestBody)
// 	if vErr != nil {
// 		return c.Status(400).JSON(
// 			Response{Message: vErr, Success: false, Detail: vErr})
// 	}
// 	student, err = student.RetrieveStudentByIDAndSchool(ctx, db, *requestBody.StudentID, school.ID)
// 	if err != nil {
// 		if err == gorm.ErrRecordNotFound {
// 			isOk := school.IsSchoolAdminOrOwner(ctx, db, schoolCode, user.ID)
// 			if !isOk {
// 				staff, err = staff.RetrieveByUserAndSchool(ctx, db, school.ID, user.ID)
// 				if err != nil {
// 					return c.Status(400).JSON(fiber.Map{"error": "staff does not exist with this user"})
// 				}
// 			}
// 		} else {
// 			return c.Status(400).JSON(Response{Message: "You dont have access to view this note", Success: false, Detail: err})

// 		}
// 	}

// 	if student.ID != uuid.Nil {
// 		if student.ID != *requestBody.StudentID {
// 			return c.Status(400).JSON(Response{Message: "You dont have access to view this note", Success: false, Detail: err})

// 		}
// 	}
// 	clinicVisitation = models.ClinicVisitation{
// 		StudentID:      requestBody.StudentID,
// 		SchoolID:       &school.ID,
// 		Name:           requestBody.Name,
// 		Description:    requestBody.Description,
// 		DoctorsName:    requestBody.DoctorsName,
// 		VisitationTime: requestBody.VisitationTime,
// 		Status:         requestBody.Status,
// 	}

// 	err = db.WithContext(ctx).Model(&clinicVisitation).Create(&clinicVisitation).First(&clinicVisitation).Error
// 	if err != nil {
// 		return c.Status(400).JSON(fiber.Map{"error": "error creating clinic visitation"})
// 	}

// 	return c.Status(201).JSON(clinicVisitation)
// }

// func ClinicVisitationUpdateController(c *fiber.Ctx) error {
// 	/* This is used to update clinic visitation */
// 	var school models.School
// 	var staff models.Staff
// 	var clinicVisitation models.ClinicVisitation
// 	var requestBody serializers.ClinicVisitationUpdateRequestSerializer

// 	// can be added by the staff
// 	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// 	defer cancel()

// 	// Open the database connection
// 	db := database.DBConnection()
// 	defer database.CloseDB(db)
// 	validateAdapter := adapters.NewValidate()

// 	user := c.Locals("user").(models.User)
// 	clinicVisitationID, err := uuid.Parse(c.Params("id"))
// 	if err != nil {
// 		return c.Status(400).JSON(fiber.Map{"error": "Invalid clinic visitation id passed"})
// 	}

// 	schoolCode := c.Query("school_code")

// 	school, err = school.RetrieveBySchoolCode(ctx, db, schoolCode)
// 	if err != nil {
// 		return c.Status(404).JSON(Response{
// 			Message: "school with this id does not exists",
// 			Success: false,
// 			Detail:  err,
// 		})
// 	}

// 	err = c.BodyParser(&requestBody)
// 	if err != nil {
// 		logger.Error(ctx, "Error parsing request body", zap.Error(err))
// 		return c.Status(400).JSON(Response{Message: "invalid request body", Success: false, Detail: err.Error()})
// 	}

// 	vErr := validateAdapter.ValidateData(&requestBody)
// 	if vErr != nil {
// 		return c.Status(400).JSON(
// 			Response{Message: vErr, Success: false, Detail: vErr})
// 	}

// 	err = db.WithContext(ctx).Model(&clinicVisitation).Where("id = ?", clinicVisitationID).Where("school_id = ?", school.ID).Preload("Student").First(&clinicVisitation).Error
// 	if err != nil {
// 		if err == gorm.ErrRecordNotFound {
// 			return c.Status(400).JSON(
// 				Response{Message: "clinic visitation does not exist", Success: false, Detail: err})
// 		}
// 		logger.Error(ctx, "error getting clinic visitation", zap.Error(err))
// 		return c.Status(400).JSON(
// 			Response{Message: "error getting clinic visitation", Success: false, Detail: err})
// 	}

// 	if clinicVisitation.Student.UserID != nil {
// 		if *clinicVisitation.Student.UserID != user.ID {
// 			isOk := school.IsSchoolAdminOrOwner(ctx, db, schoolCode, user.ID)
// 			if !isOk {
// 				staff, err = staff.RetrieveByUserAndSchool(ctx, db, school.ID, user.ID)
// 				if err != nil {
// 					return c.Status(400).JSON(Response{Message: "you dont have permission", Success: false, Detail: err})
// 				}
// 			}
// 		}
// 	}

// 	err = db.WithContext(ctx).Model(&clinicVisitation).Where("id = ?", clinicVisitationID).Where("school_id = ?", school.ID).Updates(&requestBody).First(&clinicVisitation).Error
// 	if err != nil {
// 		if err == gorm.ErrRecordNotFound {
// 			return c.Status(404).JSON(
// 				Response{Message: "clinic visitation does not exist", Success: false, Detail: err})
// 		}
// 		logger.Error(ctx, "error getting clinic visitation", zap.Error(err))
// 		return c.Status(400).JSON(
// 			Response{Message: "error getting clinic visitation", Success: false, Detail: err})
// 	}

// 	return c.Status(200).JSON(clinicVisitation)
// }

// func ClinicVisitationDetailController(c *fiber.Ctx) error {
// 	/* this is used to get the detail of the  clinic visitation*/
// 	var school models.School
// 	var staff models.Staff
// 	var clinicVisitation models.ClinicVisitation

// 	// can be added by the staff
// 	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// 	defer cancel()

// 	// Open the database connection
// 	db := database.DBConnection()
// 	defer database.CloseDB(db)

// 	user := c.Locals("user").(models.User)
// 	clinicVisitationID, err := uuid.Parse(c.Params("id"))
// 	if err != nil {
// 		return c.Status(400).JSON(fiber.Map{"error": "Invalid clinic visitation id passed"})
// 	}

// 	schoolCode := c.Query("school_code")

// 	school, err = school.RetrieveBySchoolCode(ctx, db, schoolCode)
// 	if err != nil {
// 		return c.Status(404).JSON(Response{
// 			Message: "school with this id does not exists",
// 			Success: false,
// 			Detail:  err,
// 		})
// 	}

// 	err = db.WithContext(ctx).Model(&clinicVisitation).Where("id = ?", clinicVisitationID).Where("school_id = ?", school.ID).Preload("Student").First(&clinicVisitation).Error
// 	if err != nil {
// 		if err == gorm.ErrRecordNotFound {
// 			return c.Status(404).JSON(
// 				Response{Message: "clinic visitation does not exist", Success: false, Detail: err})
// 		}
// 		logger.Error(ctx, "error getting clinic visitation", zap.Error(err))
// 		return c.Status(400).JSON(
// 			Response{Message: "error getting clinic visitation", Success: false, Detail: err})
// 	}

// 	if *clinicVisitation.Student.UserID != user.ID {
// 		isOk := school.IsSchoolAdminOrOwner(ctx, db, schoolCode, user.ID)
// 		if !isOk {
// 			staff, err = staff.RetrieveByUserAndSchool(ctx, db, school.ID, user.ID)
// 			if err != nil {
// 				return c.Status(400).JSON(Response{Message: "you dont have permission", Success: false, Detail: err})
// 			}
// 		}
// 	}

// 	return c.Status(200).JSON(clinicVisitation)
// }

// func ClinicVisitationDeleteController(c *fiber.Ctx) error {
// 	/* this is used to get the detail of the clinic visitation*/
// 	var school models.School
// 	var staff models.Staff
// 	var clinicVisitation models.ClinicVisitation

// 	// can be added by the staff
// 	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// 	defer cancel()

// 	// Open the database connection
// 	db := database.DBConnection()
// 	defer database.CloseDB(db)

// 	user := c.Locals("user").(models.User)
// 	clinicVisitationID, err := uuid.Parse(c.Params("id"))
// 	if err != nil {
// 		return c.Status(400).JSON(
// 			Response{Message: "Invalid clinic visitation id passed", Success: false, Detail: err})
// 	}

// 	schoolCode := c.Query("school_code")

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
// 		staff, err = staff.RetrieveByUserAndSchool(ctx, db, school.ID, user.ID)
// 		if err != nil {
// 			return c.Status(400).JSON(Response{Message: "you dont have permission", Success: false, Detail: err})
// 		}
// 	}
// 	err = db.WithContext(ctx).Model(&clinicVisitation).Where("id = ?", clinicVisitationID).Where("school_id = ?", school.ID).First(&clinicVisitation).Delete(&clinicVisitation).Error
// 	if err != nil {
// 		if err == gorm.ErrRecordNotFound {
// 			return c.Status(404).JSON(
// 				Response{Message: "clinic visitation does not exist", Success: false, Detail: err})
// 		}
// 		logger.Error(ctx, "error getting clinic visitation", zap.Error(err))
// 		return c.Status(400).JSON(
// 			Response{Message: "error getting clinic visitation", Success: false, Detail: err})
// 	}

// 	return c.Status(204).JSON(fiber.Map{"message": "Successfully deleted clinic visitation"})
// }
