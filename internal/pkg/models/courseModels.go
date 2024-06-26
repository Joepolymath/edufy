package models

// import (
// 	"context"
// 	"github.com/google/uuid"
// 	"gorm.io/gorm"
// 	"time"
// )

// // Course /*This contains the course which is either created by the owner of the school or the super admin of the project*/
// type Course struct {
// 	BaseModel
// 	School               *School    `json:"school,omitempty"  gorm:"constraint:OnDelete:CASCADE;ForeignKey:SchoolID"` // school the course belongs to
// 	SchoolID             *uuid.UUID `json:"school_id" gorm:"not null;"`
// 	Class                *Class     `json:"class,omitempty" gorm:"ForeignKey:ClassID"` // class the course belongs to
// 	ClassID              *uuid.UUID `json:"class_id" gorm:"not null;" validate:"required"`
// 	Staff                *Staff     `json:"staff,omitempty" gorm:"ForeignKey:StaffID"` // staff taking the course
// 	StaffID              *uuid.UUID `json:"staff_id" gorm:"not null;" validate:"required"`
// 	Name                 *string    `json:"name" gorm:"not null;max=250;" validate:"required,max=250"`
// 	Description          *string    `json:"description" gorm:"not null;" validate:"required"` // description of the course
// 	Image                *string    `json:"image" `
// 	Video                *string    `json:"video"`
// 	Performance          *int       `json:"performance"`
// 	Attendance           *int       `json:"attendance"`
// 	EnrolledStudentCount *int       `json:"enrolled_student_count"`
// }

// func (course *Course) Retrieve(ctx context.Context, db *gorm.DB, id uuid.UUID) (Course, error) {
// 	// filter on gorm
// 	err := db.WithContext(ctx).Model(&course).Where("id = ?", id).Find(&course).Error
// 	if err != nil {
// 		return *course, err
// 	}

// 	return *course, err
// }

// func (course *Course) RetrieveSchoolCourse(ctx context.Context, db *gorm.DB, id uuid.UUID, schoolID uuid.UUID) (Course, error) {
// 	// filter on gorm
// 	err := db.WithContext(ctx).Model(&course).Where("id = ?", id).Where("school_id = ?", schoolID).Find(&course).Error
// 	if err != nil {
// 		return *course, err
// 	}

// 	return *course, err
// }

// // CalculateCoursePerformance is used to calculate the performance of students in this course
// func (course *Course) CalculateCoursePerformance(ctx context.Context, db *gorm.DB, id uuid.UUID) error {
// 	var totalPoint int64
// 	var totalScore int64
// 	var totalStudent int64

// 	// Count the total number of students with a total score greater than 0
// 	err := db.WithContext(ctx).Model(&EnrolledCourse{}).
// 		Where("course_id = ?", id).
// 		Where("status ILIKE ?", "ENROLLED").
// 		Where("end_date > ?", time.Now().UTC()).
// 		Pluck("SUM(total)", &totalScore).Count(&totalStudent).Error
// 	if err != nil {
// 		return err
// 	}

// 	err = db.WithContext(ctx).Model(&CourseTask{}).
// 		Where("course_id = ?", id).
// 		Pluck("SUM(total_point)", &totalPoint).Error
// 	if err != nil {
// 		return err
// 	}

// 	// Calculate performance, handle division by zero
// 	var performance float64
// 	if totalPoint > 0 {
// 		performance = (float64(totalScore) / (float64(totalPoint) * float64(totalStudent))) * 100
// 	}

// 	// Update the course with the calculated performance
// 	err = db.Model(&Course{}).Where("id = ?", id).Update("performance", performance).Error
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// // CalculateCourseAttendance is used to calculate the attendance of students in this course
// func (course *Course) CalculateCourseAttendance(ctx context.Context, db *gorm.DB, id uuid.UUID) error {
// 	var totalAttendancePresent int64
// 	var totalStudent int64
// 	var totalStudentAttendance int64

// 	// Count the total number of students enrolled in the course
// 	err := db.WithContext(ctx).Model(&EnrolledCourse{}).Where("course_id = ?", id).Count(&totalStudent).Error
// 	if err != nil {
// 		return err
// 	}

// 	// Sum up the attendance counts for all attendance records
// 	err = db.WithContext(ctx).Model(&Attendance{}).
// 		Where("course_id = ?", id).
// 		Pluck("SUM(present + late)", &totalAttendancePresent).
// 		Error
// 	if err != nil {
// 		return err
// 	}

// 	// Count the total number of students with attendance records for the given course
// 	err = db.WithContext(ctx).Model(&StudentAttendance{}).
// 		Joins("JOIN attendances ON student_attendances.attendance_id = attendances.id").
// 		Where("attendances.course_id = ?", id).
// 		Count(&totalStudentAttendance).Error

// 	// Calculate overall attendance percentage, handle division by zero
// 	var overallAttendancePercentage float64
// 	if totalStudentAttendance > 0 {
// 		overallAttendancePercentage = (float64(totalAttendancePresent) / float64(totalStudentAttendance)) * 100
// 	}

// 	// Update the course with the calculated overall attendance percentage
// 	err = db.Model(&Course{}).Where("id = ?", id).
// 		Update("attendance", overallAttendancePercentage).Error
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// type Curriculum struct {
// 	BaseModel
// 	School          *School    `json:"school,omitempty"  gorm:"constraint:OnDelete:CASCADE;ForeignKey:SchoolID"` // school the course belongs to
// 	SchoolID        *uuid.UUID `json:"school_id" gorm:"not null;"`
// 	Course          *Course    `json:"course,omitempty" gorm:"constraint:OnDelete:CASCADE;ForeignKey:CourseID"` // course the lesson belongs to
// 	CourseID        *uuid.UUID `json:"course_id" gorm:"not null;" validate:"required"`
// 	Staff           *Staff     `json:"staff,omitempty" gorm:"constraint:OnDelete:SET NULL;ForeignKey:StaffID"` // course the lesson belongs to
// 	StaffID         *uuid.UUID `json:"staff_id"  validate:"required"`
// 	Name            *string    `json:"name" gorm:"not null;max=250;" validate:"required,max=250"`
// 	Description     *string    `json:"description" gorm:"not null;" validate:"required,max=50000"` // description of the lesson
// 	File            *string    `json:"file"  validate:"required"`                                  // file of the lesson (The file could be pdf, video, audio, etc)
// 	Order           *int       `json:"order" gorm:"not null;" validate:"required"`                 // order of the lesson within the course
// 	TotalLesson     *int       `json:"total_lesson"`
// 	CompletedLesson *int       `json:"completed_lesson"`
// }

// // Lesson /* This is the lesson of a course*/
// type Lesson struct {
// 	BaseModel
// 	School       *School     `json:"school,omitempty"  gorm:"constraint:OnDelete:CASCADE;ForeignKey:SchoolID"` // school the course belongs to
// 	SchoolID     *uuid.UUID  `json:"school_id" gorm:"not null;"`
// 	Curriculum   *Curriculum `json:"curriculum,omitempty" gorm:"constraint:OnDelete:CASCADE;ForeignKey:CurriculumID"` // curriculum the lesson belongs to
// 	CurriculumID *uuid.UUID  `json:"curriculum_id" gorm:"not null;" validate:"required"`
// 	Staff        *Staff      `json:"staff,omitempty" gorm:"constraint:OnDelete:SET NULL;ForeignKey:StaffID"` // course the lesson belongs to
// 	StaffID      *uuid.UUID  `json:"staff_id"  validate:"required"`
// 	Name         *string     `json:"name" gorm:"not null;max=250;" validate:"required,max=250"`
// 	Description  *string     `json:"description" gorm:"not null;" validate:"required,max=50000"` // description of the lesson
// 	File         *string     `json:"file" `                                                      // file of the lesson (The file could be pdf, video, audio, etc)
// 	Order        *int        `json:"order" `                                                     // order of the lesson within the course
// 	Completed    *bool       `json:"completed" gorm:"default:false"`
// }

// func (l *Lesson) UpdateCurriculumLessonCount(db *gorm.DB) (err error) {
// 	var curriculumLessonsCount int64
// 	var curriculum Curriculum

// 	err = db.Model(&Lesson{}).Where("curriculum_id = ?", l.CurriculumID).Count(&curriculumLessonsCount).Error
// 	if err != nil {
// 		return err
// 	}
// 	// update the curriculum lesson count and add 1 because
// 	err = db.Model(&Curriculum{}).Where("id = ?", l.CurriculumID).Update("total_lesson", curriculumLessonsCount).First(&curriculum).Error
// 	if err != nil {
// 		return err
// 	}
// 	return
// }

// func (l *Lesson) UpdateCurriculumCompletedCount(db *gorm.DB) (err error) {

// 	var lessonCompletedCount int64

// 	err = db.Model(&Lesson{}).Where("curriculum_id = ?", l.CurriculumID).Where("completed", true).Count(&lessonCompletedCount).Error
// 	if err != nil {
// 		return err
// 	}

// 	// update the curriculum lesson count
// 	err = db.Model(&Curriculum{}).Where("id = ?", l.CurriculumID).Update("completed_lesson", lessonCompletedCount).Error
// 	if err != nil {
// 		return err
// 	}

// 	return
// }

// type EnrolledCourse struct {
// 	/* This is the join models that enable us to connect the payment link to the transaction*/
// 	// this contains the required fields
// 	BaseModel
// 	School          *School    `json:"school,omitempty"  gorm:"constraint:OnDelete:CASCADE;ForeignKey:SchoolID"` // school the course belongs to
// 	SchoolID        *uuid.UUID `json:"school_id" gorm:"not null;"`
// 	Student         *Student   `json:"student,omitempty"  gorm:"constraint:OnDelete:CASCADE;ForeignKey:StudentID"`
// 	StudentID       *uuid.UUID `json:"student_id"  gorm:"not null;"`
// 	Course          *Course    `json:"course"  gorm:"constraint:OnDelete:CASCADE;ForeignKey:CourseID"`
// 	CourseID        *uuid.UUID `json:"course_id"  gorm:"not null;"`
// 	Status          *string    `json:"status" gorm:"not null;default:'ENROLLED'"` //  could be ENROLLED, FAILED, COMPLETED
// 	Lesson          *Lesson    `json:"lesson" gorm:"constraint:OnDelete:SET NULL; ForeignKey:LessonID"`
// 	LessonID        *uuid.UUID `json:"lesson_id"`
// 	StartDate       *time.Time `json:"start_date"`
// 	AssignmentScore *uint      `json:"assignment_score"`
// 	TestScore       *uint      `json:"test_score"`
// 	ExamScore       *uint      `json:"exam_score"`
// 	Total           *uint      `json:"total"`
// 	EndDate         *time.Time `json:"end_date"`
// }

// // BeforeCreate hook to update the Course's EnrolledStudentCount
// func (enrolledCourse *EnrolledCourse) BeforeCreate(db *gorm.DB) (err error) {
// 	// Increment the Course's EnrolledStudentCount
// 	enrolledCourse.ID = uuid.New()

// 	var totalEnrolledStudentCount int64
// 	err = db.Model(&EnrolledCourse{}).
// 		Where("course_id = ?", enrolledCourse.CourseID).
// 		Where("status = ?", "ENROLLED").
// 		Where("end_date > ?", time.Now().UTC()).
// 		Count(&totalEnrolledStudentCount).
// 		Error
// 	if err != nil {
// 		return err
// 	}

// 	// update the course
// 	err = db.Model(&Course{}).Where("id =?", enrolledCourse.CourseID).Update("enrolled_student_count", totalEnrolledStudentCount+1).Error
// 	if err != nil {
// 		return err
// 	}

// 	return
// }

// type CourseSchedule struct {
// 	BaseModel
// 	School       *School    `json:"school,omitempty"  gorm:"constraint:OnDelete:CASCADE;ForeignKey:SchoolID"` // school the course belongs to
// 	SchoolID     *uuid.UUID `json:"school_id" gorm:"not null;"`
// 	Staff        *Staff     `json:"staff,omitempty" gorm:"constraint:OnDelete:CASCADE;ForeignKey:StaffID"`
// 	StaffID      *uuid.UUID `json:"staff_id" gorm:"not null"`
// 	Course       *Course    `json:"course,omitempty" gorm:"constraint:OnDelete:CASCADE;ForeignKey:CourseID"`
// 	CourseID     *uuid.UUID `json:"course_id" gorm:"not null"`
// 	ScheduleDate *time.Time `json:"schedule_date"`
// 	FromTime     *string    `json:"from_time"`
// 	ToTime       *string    `json:"to_time"`
// }

// type Attendance struct {
// 	BaseModel
// 	School            *School    `json:"school,omitempty"  gorm:"constraint:OnDelete:CASCADE;ForeignKey:SchoolID"`
// 	SchoolID          *uuid.UUID `json:"school_id" gorm:"not null;"`
// 	Course            *Course    `json:"course" gorm:"constraint:OnDelete:CASCADE;ForeignKey:CourseID"`
// 	CourseID          *uuid.UUID `json:"course_id" gorm:"not null;"`
// 	Present           *int64     `json:"present"`
// 	Absent            *int64     `json:"absent"`
// 	Late              *int64     `json:"late"`
// 	AttendancePercent *float64   `json:"attendance_percent"`
// }

// // CalculateAttendancePercent is used to calculate the attendance percentage for an attendance record
// func (attendance *Attendance) CalculateAttendancePercent(ctx context.Context, db *gorm.DB, id uuid.UUID) error {
// 	var totalStudent int64
// 	var presentCount int64
// 	var absentCount int64
// 	var lateCount int64

// 	err := db.WithContext(ctx).Model(&Attendance{}).
// 		Where("id = ?", id).
// 		Error
// 	if err != nil {
// 		return err
// 	}

// 	// Count the total number of students enrolled in the course
// 	err = db.WithContext(ctx).Model(&StudentAttendance{}).
// 		Where("attendance_id = ?", id).
// 		Count(&totalStudent).Error
// 	if err != nil {
// 		return err
// 	}

// 	// Count the total number of students with attendance marked as 'PRESENT'
// 	err = db.WithContext(ctx).Model(&StudentAttendance{}).
// 		Where("attendance_id = ?", id).
// 		Where("status = ?", "PRESENT").
// 		Count(&presentCount).Error
// 	if err != nil {
// 		return err
// 	}
// 	// Count the total number of students with attendance marked as 'PRESENT'
// 	err = db.WithContext(ctx).Model(&StudentAttendance{}).
// 		Where("attendance_id = ?", id).
// 		Where("status = ?", "ABSENT").
// 		Count(&absentCount).Error
// 	if err != nil {
// 		return err
// 	}
// 	// Count the total number of students with attendance marked as 'PRESENT'
// 	err = db.WithContext(ctx).Model(&StudentAttendance{}).
// 		Where("attendance_id = ?", id).
// 		Where("status = ?", "LATE").
// 		Count(&lateCount).Error
// 	if err != nil {
// 		return err
// 	}

// 	// Calculate attendance percentage, handle division by zero
// 	var attendancePercentage float64
// 	if totalStudent > 0 {
// 		attendancePercentage = ((float64(presentCount) + float64(lateCount)) / float64(totalStudent)) * 100
// 	}

// 	// Update the Attendance model with the calculated attendance percentage
// 	err = db.Model(&Attendance{}).
// 		Where("id = ?", id).
// 		Updates(map[string]interface{}{
// 			"attendance_percent": attendancePercentage,
// 			"present":            presentCount,
// 			"absent":             absentCount,
// 			"late":               lateCount,
// 		}).Error
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// type StudentAttendance struct {
// 	BaseModel
// 	School       *School     `json:"school,omitempty" gorm:"constraint:OnDelete:CASCADE;ForeignKey:SchoolID"`
// 	SchoolID     *uuid.UUID  `json:"school_id" gorm:"not null;"`
// 	Student      *Student    `json:"student,omitempty" gorm:"constraint:OnDelete:CASCADE;ForeignKey:StudentID"`
// 	StudentID    *uuid.UUID  `json:"student_id"`
// 	Attendance   *Attendance `json:"attendance,omitempty" gorm:"constraint:OnDelete:CASCADE;ForeignKey:AttendanceID"`
// 	AttendanceID *uuid.UUID  `json:"attendance_id"`
// 	Status       *string     `json:"status" gorm:"not null;"` // could be PRESENT, ABSENT, LATE
// }
