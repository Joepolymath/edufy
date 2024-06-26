package setup

import (
	"Learnium/internal/database"
	"Learnium/internal/pkg/models"
	"context"
	"time"
)

var (
	HRDesc = `Manages all aspects of human resources, including recruitment, onboarding, staff development, and employee relations.
 Ensures compliance with employment laws, maintains personnel records, and supports overall staff well-being`

	AODesc = ` Manages the student admission process, including application reviews, interviews, and enrollment. 
Coordinates with prospective students and their families.`

	CounselDesc = `Provides counseling services to students. Addresses academic, personal, and emotional concerns, 
and collaborates with teachers and parents to support student well-being.`

	HWCDesc = `Promotes the health and well-being of students and staff. Coordinates health programs, organizes wellness activities, and manages health-related records.`

	ACCDesc = `Manages financial transactions, budgeting, and payroll for the school. Responsible for keeping accurate financial records.`

	ECDesc = `Manages the planning and execution of examinations. Coordinates with teachers, ensures proper exam administration, and oversees the grading process.`

	CDevDesc = `Designs and develops the school curriculum, ensuring alignment with educational standards and goals. Collaborates with other educators to enhance teaching materials.`

	CTDesc = ` Manages a specific class or group of students. Acts as the primary point of contact for parents and is responsible for monitoring the overall well-being and academic progress of the students.`

	STDesc = `Teaches a specific subject or multiple subjects. Responsible for creating lesson plans, assessing student performance, and providing feedback.`

	HODDesc = `Leads a specific academic department, overseeing curriculum, lesson planning, and the performance of teachers within the department.`

	SchoolAdminDesc = ` Manages the overall administration of the school, including student enrollment, staff management, and academic settings.
 Has the authority to configure system settings related to the school's policies.`

	// 	SysAdminDesc = `Manages the technical aspects of the school management application, including server maintenance, database management, and software updates.
	// Ensures the smooth operation of the system.`

	SuperAdminDesc = `Has access to all features and controls within the school management application. Can configure system settings, manage users, and oversee all aspects of the system.`
)

// func AutoSeedRoles(db *gorm.DB) error {
// 	// Check if any records exist in the table
// 	var count int64
// 	db.Model(&models.Role{}).Count(&count)

// 	// If no records exist, seed the table with initial data
// 	if count == 0 {
// 		initialCardTypes := []models.Role{
// 			{Name: "Super Administrator", RoleType: models.ADMIN, Description: &SuperAdminDesc},
// 			{Name: "School Administrator", RoleType: models.ADMIN, Description: &SchoolAdminDesc},

// 			{Name: "Head Of Department", RoleType: models.TEACHING_STAFF, Description: &HODDesc},
// 			{Name: "Class Teacher", RoleType: models.TEACHING_STAFF, Description: &CTDesc},
// 			{Name: "Subject Teacher", RoleType: models.TEACHING_STAFF, Description: &STDesc},
// 			{Name: "Curriculum Developer", RoleType: models.TEACHING_STAFF, Description: &CDevDesc},
// 			{Name: "Examinations Coordinator", RoleType: models.TEACHING_STAFF, Description: &ECDesc},

// 			{Name: "Accountant", RoleType: models.NON_TEACHING_STAFF, Description: &ACCDesc},
// 			{Name: "Counselor", RoleType: models.NON_TEACHING_STAFF, Description: &CounselDesc},
// 			{Name: "Admissions Officer", RoleType: models.NON_TEACHING_STAFF, Description: &AODesc},
// 			{Name: "Human Resources", RoleType: models.NON_TEACHING_STAFF, Description: &HRDesc},
// 			{Name: "Health and Wellnesss Coordinator", RoleType: models.NON_TEACHING_STAFF,
// 				Description: &HWCDesc},
// 		}

// 		if err := db.Create(&initialCardTypes).Error; err != nil {
// 			log.Println("Error auto seeding default roles::: ", err)
// 			return err
// 		}
// 	}

// 	return nil
// }

// MigrateDatabase performs database migrations for all models
func MigrateDatabase() {
	// Migrate the schema
	db := database.DBConnection()
	// close the database after the connection
	defer database.CloseDB(db)

	_, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	// if err := db.Migrator().DropTable(
	// 	&models.School{},
	// 	&models.SchoolRole{},
	// 	// &models.NewsletterSubscription{},
	// 	// &models.Role{},
	// // 	&model.SellGiftCard{},
	// // 	&model.BuyGiftCard{},
	// // 	&model.BuyUsdt{},
	// ); err != nil {
	// 	panic(err)
	// }

	// Migrate the models
	err := db.AutoMigrate(
		&models.User{},
		&models.NewsletterSubscription{},
		// &models.Profile{},
		// &models.Health{},
		// &models.SchoolType{},
		&models.School{},
		&models.SchoolRole{},
		// &models.Admin{},
		// &models.SchoolInvite{},
		// &models.AnnouncementConfiguration{},
		// &models.NotificationConfiguration{},
		// &models.Class{},
		// &models.Category{},
		// &models.Role{},
		// &models.Event{},
		// &models.Student{},
		// &models.EmploymentConfiguration{},
		// &models.Employment{},
		// &models.EmploymentApplication{},
		// &models.Staff{},
		// &models.StaffClass{},
		// &models.StaffRating{},
		// &models.Event{}, // needs the staff id and school id
		// &models.ConversationUser{},
		// &models.Conversation{},
		// &models.Message{},
		// &models.EnrollmentConfiguration{},
		// &models.EnrollmentAdmission{},
		// &models.Course{},
		// &models.Curriculum{},
		// &models.Lesson{},
		// &models.EnrolledCourse{},
		// &models.CourseSchedule{},
		// &models.Attendance{},
		// &models.StudentAttendance{},
		// &models.CourseTask{},
		// &models.Question{},
		// &models.QuestionOption{},
		// &models.QuestionQuestionOption{},
		// &models.StudentTask{},
		// &models.StudentAnswer{},
		// &models.Note{},
		// &models.ClinicVisitation{},
	)
	if err != nil {
		panic(err)
		// logger.Error(ctx, "Error migrating models", zap.Error(err))
	}

	// _ = AutoSeedRoles(db)

}
