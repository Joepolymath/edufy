package serializers

// import (
// 	"Learnium/models"
// 	"Learnium/utils"
// 	"context"
// 	"github.com/google/uuid"
// 	"gorm.io/gorm"
// )

// type StaffSerializer struct {
// 	ID                      *uuid.UUID       `json:"id" `
// 	StaffCode               *string          `json:"staffCode,omitempty" `
// 	SchoolID                *uuid.UUID       `json:"school_id"`
// 	User                    *User            `json:"user" ` //	the one-to-one relationship
// 	UserID                  *uuid.UUID       `json:"user_id" `
// 	EmploymentApplicationID *uuid.UUID       `json:"employment_application_id" ` // the application id
// 	Classes                 *[]models.Class  `json:"classes" `                   // the class which the teacher is in control of
// 	Role                    *models.Role     `json:"role" `
// 	RoleID                  *uuid.UUID       `json:"role_id" `
// 	Category                *models.Category `json:"category" `
// 	CategoryID              *uuid.UUID       `json:"category_id" `
// 	Qualification           *string          `json:"qualification" `
// 	NetSalary               *float64         `json:"net_salary" ` // this would be updated anytime the salary is updated
// 	Status                  *string          `json:"status"`
// }

// type StaffUpdateSerializer struct {
// 	ClassID       *uuid.UUID   `json:"class_id" ` // the class which the teacher is in control of
// 	Role          *models.Role `json:"role" `
// 	RoleID        *uuid.UUID   `json:"role_id" `
// 	CategoryID    *uuid.UUID   `json:"category_id" `
// 	Qualification *string      `json:"qualification" `
// 	NetSalary     *float64     `json:"net_salary" ` // this would be updated anytime the salary is updated
// 	StaffType     *string      `json:"staff_type" validate:"oneof=FULL_TIME PART_TIME"`
// 	Status        *string      `json:"status" validate:"oneof=ACTIVE INACTIVE"`
// }

// func (staff *StaffUpdateSerializer) ValidateData(db *gorm.DB, ctx context.Context) error {

// 	// check the role id
// 	if staff.RoleID != nil {
// 		var role models.Role
// 		err := db.WithContext(ctx).Where("id = ?", staff.RoleID).First(&role).Error
// 		if err != nil {
// 			return err
// 		}
// 	}

// 	// check the classID
// 	if staff.ClassID != nil {
// 		var class models.Class
// 		err := db.WithContext(ctx).Where("id = ?", staff.ClassID).First(&class).Error
// 		if err != nil {
// 			return err
// 		}
// 	}

// 	// check the category id
// 	if staff.CategoryID != nil {
// 		var category models.Category
// 		err := db.WithContext(ctx).Where("id = ?", staff.CategoryID).First(&category).Error
// 		if err != nil {
// 			return err
// 		}
// 	}
// 	return nil
// }

// type StaffRatingCreateRequestSerializer struct {
// 	StaffID *uuid.UUID `json:"staff_id" validate:"required"`
// 	Message *string    `json:"message"  validate:"required"`
// 	Rating  *int       `json:"rating" validate:"required,max=5"`
// }

// type StaffsAnalyticsResponseSerializer struct {
// 	Rating             []models.StaffRating      `json:"rating"`
// 	Attendance         []utils.MonthlyAttendance `json:"attendance"`
// 	ActiveStaffCount   int64                     `json:"active_staff_count"`
// 	InactiveStaffCount int64                     `json:"inactive_staff_count"`
// 	TotalStaffCount    int64                     `json:"total_staff_count"`
// 	MaleCount          int64                     `json:"male_count"`
// 	FemaleCount        int64                     `json:"female_count"`
// }

// type CoursePerformanceSerializer struct {
// 	Name        *string `json:"name" `
// 	Performance *int    `json:"performance"`
// 	Attendance  *int    `json:"attendance"`
// }

// type StaffAnalyticsResponseSerializer struct {
// 	CoursePerformance       []CoursePerformanceSerializer   `json:"course_performance"`
// 	StaffFeedbackPercentage []utils.StaffFeedbackPercentage `json:"staff_feedback_percentage"`
// }
