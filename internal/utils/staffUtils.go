package utils

// import (
// 	"Learnium/database"
// 	"Learnium/models"
// 	"context"
// 	"fmt"
// 	"github.com/google/uuid"
// 	"gorm.io/gorm"
// 	"time"
// )

// type StaffUtilsInterface interface {
// 	GenerateStaffCode() string
// 	CheckIfStaffCodeExists(staffCode string) bool
// 	GetMonthlyAttendance(ctx context.Context, db *gorm.DB, schoolID uuid.UUID) ([]MonthlyAttendance, error)
// 	GetStaffFeedbackPercentage(ctx context.Context, db *gorm.DB, staffID uuid.UUID) ([]StaffFeedbackPercentage, error)
// }

// type StaffUtils struct {
// }

// func NewStaffUtils() StaffUtilsInterface {
// 	return &StaffUtils{}
// }

// func (staffUtils *StaffUtils) GenerateStaffCode() string {

// 	StaffCode := GenerateCode(6, "LM")
// 	//check if code exists
// 	StaffCodeExist := staffUtils.CheckIfStaffCodeExists(StaffCode)
// 	if StaffCodeExist {
// 		// i need to generate a new one
// 		return GenerateCode(6, "ST")
// 	}

// 	// Return the generated string
// 	return StaffCode
// }

// func (staffUtils *StaffUtils) CheckIfStaffCodeExists(staffCode string) bool {
// 	/* This is used to check if a staffCode exists*/
// 	// the timeout which is used for accessing this route it must not be more than 100 seconds
// 	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// 	defer cancel()
// 	// Open the database connection and close it at the end
// 	db := database.DBConnection()
// 	defer database.CloseDB(db)

// 	var staff models.Staff
// 	err := db.WithContext(ctx).Model(&staff).Where("staff_code = ?", staffCode).First(&staff).Error

// 	// if the error is not found
// 	if err == gorm.ErrRecordNotFound {
// 		return false
// 	}
// 	return true

// }

// // MonthlyAttendance represents the attendance data for a specific month
// type MonthlyAttendance struct {
// 	Month   string `json:"month"`
// 	Present int64  `json:"present"`
// 	Absent  int64  `json:"absent"`
// 	Late    int64  `json:"late"`
// }

// // GetMonthlyAttendance returns the total present count for each month
// func (staffUtils *StaffUtils) GetMonthlyAttendance(ctx context.Context, db *gorm.DB, schoolID uuid.UUID) ([]MonthlyAttendance, error) {
// 	var monthlyAttendance []MonthlyAttendance

// 	query := fmt.Sprintf(`
// 		SELECT
// 			TO_CHAR(timestamp, 'Month YYYY') as month,
// 			SUM(CASE WHEN status = 'PRESENT' THEN 1 ELSE 0 END) as present,
// 			SUM(CASE WHEN status = 'ABSENT' THEN 1 ELSE 0 END) as absent,
// 			SUM(CASE WHEN status = 'LATE' THEN 1 ELSE 0 END) as late
// 		FROM student_attendances
// 		WHERE school_id = '%s' AND timestamp >= NOW() - INTERVAL '12 months'
// 		GROUP BY TO_CHAR(timestamp, 'Month YYYY')
// 		ORDER BY MIN(timestamp) ASC
// 	`, schoolID)

// 	err := db.Raw(query).Scan(&monthlyAttendance).Error

// 	if err != nil {
// 		return nil, err
// 	}

// 	return monthlyAttendance, nil
// }

// type StaffFeedbackPercentage struct {
// 	Rating     int     `json:"rating"`
// 	Percentage float64 `json:"percentage"`
// }

// // GetStaffFeedbackPercentage returns the percentage of each rating count (1 to 5) for a staff
// func (staffUtils *StaffUtils) GetStaffFeedbackPercentage(ctx context.Context, db *gorm.DB, staffID uuid.UUID) ([]StaffFeedbackPercentage, error) {
// 	var staffFeedbackPercentage []StaffFeedbackPercentage

// 	// Replace "your_table_name" with the actual name of your table
// 	err := db.Raw(`
// 		SELECT
// 			rating,
// 			ROUND(COUNT(*) * 100.0 / (SELECT COUNT(*) FROM staff_ratings WHERE staff_id = ?), 2) as percentage
// 		FROM staff_ratings
// 		WHERE staff_id = ?
// 		GROUP BY rating
// 		ORDER BY rating
// 	`, staffID, staffID).Scan(&staffFeedbackPercentage).Error

// 	if err != nil {
// 		return nil, err
// 	}

// 	return staffFeedbackPercentage, nil
// }
