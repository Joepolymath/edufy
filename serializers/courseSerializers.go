package serializers

// import (
// 	"Learnium/models"
// 	"Learnium/utils"
// 	"github.com/google/uuid"
// 	"time"
// )

// type CourseCreateRequestSerializer struct {
// 	SchoolID             *uuid.UUID `form:"school_id" gorm:"not null;"`
// 	ClassID              *uuid.UUID `form:"class_id" gorm:"not null;" validate:"required"`
// 	StaffID              *uuid.UUID `form:"staff_id" gorm:"not null;" validate:"required"`
// 	Name                 *string    `form:"name" gorm:"not null;max=250;" validate:"required,max=250"`
// 	Description          *string    `form:"description" gorm:"not null;" validate:"required"` // description of the course
// 	Image                *string    `form:"image" `
// 	Video                *string    `form:"video"`
// 	EnrolledStudentCount *int       `form:"enrolled_student_count"`
// }

// type CourseDetailSerializer struct {
// 	ID                   *uuid.UUID                   `json:"id"`
// 	School               *models.School               `json:"course,omitempty"  ` // school the course belongs to
// 	SchoolID             *uuid.UUID                   `json:"school_id" `
// 	Class                *models.Class                `json:"class"` // class the course belongs to
// 	ClassID              *uuid.UUID                   `json:"class_id" `
// 	Staff                *models.Staff                `json:"staff,omitempty"` // staff taking the course
// 	StaffID              *uuid.UUID                   `json:"staff_id" validate:"required"`
// 	Name                 *string                      `json:"name"  validate:"required,max=250"`
// 	Description          *string                      `json:"description" validate:"required"` // description of the course
// 	Image                *string                      `json:"image" `
// 	Video                *string                      `json:"video"`
// 	Curriculums          []CurriculumDetailSerializer `json:"curriculums" `
// 	Performance          *int                         `json:"performance"`
// 	Attendance           *int                         `json:"attendance"`
// 	EnrolledStudentCount *int                         `json:"enrolled_student_count"`
// 	Timestamp            *time.Time                   `json:"timestamp"`
// }
// type CurriculumDetailSerializer struct {
// 	ID          *uuid.UUID      `json:"id"`
// 	School      *models.School  `json:"course,omitempty"  ` // school the course belongs to
// 	SchoolID    *uuid.UUID      `json:"school_id" `
// 	Staff       *models.Staff   `json:"staff,omitempty"` // staff taking the course
// 	StaffID     *uuid.UUID      `json:"staff_id" validate:"required"`
// 	Name        *string         `json:"name"  validate:"required,max=250"`
// 	Description *string         `json:"description" validate:"required"` // description of the course
// 	Order       *int            `json:"order" `
// 	Lessons     []models.Lesson `json:"lessons"` // lesson would cause invalid foreign key sha
// 	Timestamp   *time.Time      `json:"timestamp"`
// }

// type CurriculumCreateSerializer struct {
// 	CourseID    *uuid.UUID `form:"course_id"  validate:"required"`
// 	Name        *string    `form:"name"  validate:"required,max=250"`
// 	Description *string    `form:"description"  ` // description of the lesson
// 	File        *string    `form:"file"  `        // file of the lesson (The file could be pdf, video, audio, etc)
// 	Order       *int       `form:"order"  `
// }

// type LessonCreateSerializer struct {
// 	CurriculumID *uuid.UUID `form:"curriculum_id"  validate:"required"`
// 	Name         *string    `form:"name"  validate:"required,max=250"`
// 	Description  *string    `form:"description"  validate:"required,max=50000"` // description of the lesson
// 	File         *string    `form:"file"  `                                     // file of the lesson (The file could be pdf, video, audio, etc)
// 	Order        *int       `form:"order" `
// }
// type LessonUpdateSerializer struct {
// 	CurriculumID *uuid.UUID `form:"curriculum_id"  `
// 	Name         *string    `form:"name" `
// 	Description  *string    `form:"description" `
// 	File         *string    `form:"file"  `
// 	Order        *int       `form:"order" `
// 	Completed    *bool      `form:"completed"  `
// }

// type EnrollCourseRequestSerializer struct {
// 	CourseID  *uuid.UUID `json:"course_id" validate:"required"`
// 	StudentID *uuid.UUID `json:"student_id" validate:"required"`
// 	StartDate *time.Time `json:"start_date" validate:"required"`
// 	EndDate   *time.Time `json:"end_date" validate:"required"`
// }

// type CourseScheduleRequestSerializer struct {
// 	CourseID     *uuid.UUID `json:"course_id"  validate:"required"`
// 	StaffID      *uuid.UUID `json:"staff_id"`
// 	ScheduleDate *time.Time `json:"schedule_date" validate:"required"`
// 	FromTime     *string    `json:"from_time" validate:"required"`
// 	ToTime       *string    `json:"to_time" validate:"required"`
// }

// type RegisterStudentPresenceRequestSerializer struct {
// 	CourseID         *uuid.UUID                          `json:"course_id" validate:"required"`
// 	StudentPresences []*StudentPresenceRequestSerializer `json:"student_presences"`
// }

// type StudentPresenceRequestSerializer struct {
// 	StudentID *uuid.UUID `json:"student_id" validate:"required"`
// 	Status    *string    `json:"status" validate:"oneoff=PRESENT ABSENT LATE"`
// }

// type CourseAnalyticsSerializer struct {
// 	RegisteredStudents          *int                                 `json:"registered_students"`
// 	TeachingHours               *uint                                `json:"teaching_hours"`
// 	Attendance                  *int                                 `json:"attendance"`
// 	CurriculumProgress          *float64                             `json:"curriculum_progress"`
// 	CourseCurriculumInfo        *[]utils.CourseCurriculumInfo        `json:"CourseCurriculumInfo"`
// 	HighestPerformingStudents   *[]utils.HighestPerformingStudent    `json:"highest_performing_students"`
// 	StudentPerformanceAnalytics *[]utils.StudentPerformanceAnalytics `json:"student_performance_analytics"`
// }
