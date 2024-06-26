package utils

// import (
// 	"Learnium/adapters"
// 	"Learnium/models"
// 	"context"
// 	"github.com/google/uuid"
// 	"gorm.io/gorm"
// 	"gorm.io/gorm/clause"
// 	"time"
// )

// type CourseUtilsInterface interface {
// 	GetHighestPerformingStudentAnalytics(ctx context.Context, db *gorm.DB, courseID uuid.UUID) ([]HighestPerformingStudent, error)
// 	GetCourseCurriculumAnalytics(ctx context.Context, db *gorm.DB, courseID uuid.UUID) ([]CourseCurriculumInfo, float64, error)
// 	GetStudentPerformanceAnalytics(ctx context.Context, db *gorm.DB, courseID uuid.UUID) ([]StudentPerformanceAnalytics, error)
// }

// type CourseUtils struct {
// }

// func NewCourseUtils() CourseUtilsInterface {
// 	return &CourseUtils{}
// }

// type HighestPerformingStudent struct {
// 	Name        *string `json:"name"`
// 	Performance *uint   `json:"performance"`
// }

// func (courseUtils *CourseUtils) GetHighestPerformingStudentAnalytics(ctx context.Context, db *gorm.DB, courseID uuid.UUID) ([]HighestPerformingStudent, error) {
// 	var enrolledCourses []models.EnrolledCourse
// 	var highestPerformingStudents []HighestPerformingStudent

// 	err := db.WithContext(ctx).Model(&models.EnrolledCourse{}).
// 		Where("course_id =?", courseID).
// 		Where("status = ?", "ENROLLED").
// 		Where("end_date > ?", time.Now().UTC()).
// 		Preload("Student.User").
// 		Order(clause.OrderByColumn{Column: clause.Column{Name: "total"}, Desc: true}).
// 		Find(&enrolledCourses).Error
// 	if err != nil {
// 		return nil, err
// 	}

// 	for i, enrolledCourse := range enrolledCourses {
// 		highestPerformingStudent := HighestPerformingStudent{
// 			Name:        enrolledCourse.Student.User.FirstName,
// 			Performance: enrolledCourse.Total,
// 		}
// 		highestPerformingStudents = append(highestPerformingStudents, highestPerformingStudent)
// 		if i == 10 {
// 			break
// 		}
// 	}

// 	return highestPerformingStudents, err

// }

// type CourseCurriculumInfo struct {
// 	ID               *uuid.UUID `json:"id"`
// 	Subject          *string    `json:"subject"`
// 	PercentCompleted *float64   `json:"percent_completed"`
// }

// func (courseUtils *CourseUtils) GetCourseCurriculumAnalytics(ctx context.Context, db *gorm.DB, courseID uuid.UUID) ([]CourseCurriculumInfo, float64, error) {
// 	var curriculums []models.Curriculum
// 	var courseCurriculumInfos []CourseCurriculumInfo
// 	var curriculumProgress float64
// 	var totalCurriculumProgress float64
// 	var totalCurriculum int64

// 	pointerAdapters := adapters.NewPointer()

// 	err := db.WithContext(ctx).Model(&models.Curriculum{}).Where("course_id =?", courseID).
// 		Order(clause.OrderByColumn{Column: clause.Column{Name: "order"}, Desc: false}).Find(&curriculums).Count(&totalCurriculum).Error
// 	if err != nil {
// 		return nil, 0, err
// 	}

// 	for _, curriculum := range curriculums {
// 		var totalLessons int64
// 		var completedLessons int64
// 		err := db.WithContext(ctx).Model(&models.Lesson{}).
// 			Where("curriculum_id =?", curriculum.ID).Count(&totalLessons).Error // total lessons in a curriculum
// 		if err != nil {
// 			return nil, 0, err
// 		}
// 		err = db.WithContext(ctx).Model(&models.Lesson{}).Where("curriculum_id =?", curriculum.ID).
// 			Where("completed =?", true).Count(&completedLessons).Error // completed lesson
// 		if err != nil {
// 			return nil, 0, err
// 		}

// 		var percentCompleted float64
// 		if totalLessons > 0 {
// 			percentCompleted = (float64(completedLessons) / float64(totalLessons)) * 100
// 			curriculumProgress += percentCompleted
// 		}

// 		var curriculumInfo CourseCurriculumInfo
// 		curriculumInfo = CourseCurriculumInfo{
// 			ID:               pointerAdapters.UUIDPointer(curriculum.ID),
// 			Subject:          curriculum.Name,
// 			PercentCompleted: pointerAdapters.Float64Pointer(percentCompleted),
// 		}
// 		courseCurriculumInfos = append(courseCurriculumInfos, curriculumInfo)
// 	}

// 	totalCurriculumProgress = (curriculumProgress / (float64(totalCurriculum) * 100)) * 100
// 	return courseCurriculumInfos, totalCurriculumProgress, err
// }

// type StudentPerformanceAnalytics struct {
// 	ScoreRange string  `json:"score_range"`
// 	Percentage float64 `json:"percentage"`
// }

// func (courseUtils *CourseUtils) GetStudentPerformanceAnalytics(ctx context.Context, db *gorm.DB, courseID uuid.UUID) ([]StudentPerformanceAnalytics, error) {
// 	var enrolledCourses []models.EnrolledCourse
// 	var studentPerformanceAnalytics []StudentPerformanceAnalytics

// 	err := db.WithContext(ctx).Model(&models.EnrolledCourse{}).
// 		Where("course_id =?", courseID).
// 		Where("status = ?", "ENROLLED").
// 		Where("end_date > ?", time.Now().UTC()).
// 		Preload("Student.User").
// 		Find(&enrolledCourses).Error
// 	if err != nil {
// 		return studentPerformanceAnalytics, err
// 	}

// 	// Initialize score range counters
// 	var count70to100, count50to69, count25to49, count0to24 int

// 	// Count students in each score range
// 	for _, enrolledCourse := range enrolledCourses {
// 		if enrolledCourse.Total != nil {
// 			score := *enrolledCourse.Total
// 			switch {
// 			case score >= 70:
// 				count70to100++
// 			case score >= 50:
// 				count50to69++
// 			case score >= 25:
// 				count25to49++
// 			default:
// 				count0to24++
// 			}
// 		}
// 	}

// 	// Calculate percentages
// 	totalStudents := len(enrolledCourses)
// 	if totalStudents > 0 {
// 		studentPerformanceAnalytics = append(studentPerformanceAnalytics, StudentPerformanceAnalytics{
// 			ScoreRange: "70-100",
// 			Percentage: float64(count70to100) / float64(totalStudents) * 100,
// 		}, StudentPerformanceAnalytics{
// 			ScoreRange: "50-69",
// 			Percentage: float64(count50to69) / float64(totalStudents) * 100,
// 		}, StudentPerformanceAnalytics{
// 			ScoreRange: "25-49",
// 			Percentage: float64(count25to49) / float64(totalStudents) * 100,
// 		}, StudentPerformanceAnalytics{
// 			ScoreRange: "0-24",
// 			Percentage: float64(count0to24) / float64(totalStudents) * 100,
// 		})
// 	}

// 	return studentPerformanceAnalytics, nil
// }

// type CourseCompletionPercentage struct {
// 	ID                  uuid.UUID `json:"id"`
// 	Name                *string   `json:"name"`
// 	CompletedCurriculum *int      `json:"completed_curriculum"`
// 	TotalCurriculum     *int      `json:"total_curriculum"`
// }

// func (courseUtils *CourseUtils) GetCourseCompletionPercentage(ctx context.Context, db *gorm.DB, staffID uuid.UUID) ([]CourseCompletionPercentage, error) {
// 	var courses []models.Course
// 	var courseCompletionPercentages []CourseCompletionPercentage

// 	err := db.WithContext(ctx).Model(&models.Course{}).Where("staff_id = ?", staffID).Find(&courses).Error
// 	if err != nil {
// 		return courseCompletionPercentages, err
// 	}

// 	// fixme : fix this
// 	return courseCompletionPercentages, err
// }
