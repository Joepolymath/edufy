package models

// import (
// 	"Learnium/adapters"
// 	"Learnium/logger"
// 	"context"
// 	"errors"
// 	"github.com/google/uuid"
// 	"go.uber.org/zap"
// 	"gorm.io/gorm"
// 	"time"
// )

// // CourseTask /* This is the task of a course*/
// type CourseTask struct {
// 	BaseModel
// 	School       *School    `json:"school,omitempty" gorm:"constraint:OnDelete:CASCADE;ForeignKey:SchoolID"`
// 	SchoolID     *uuid.UUID `json:"school_id" gorm:"not null;"`
// 	Course       *Course    `json:"course,omitempty" gorm:"constraint:OnDelete:CASCADE;ForeignKey:CourseID"` // course the task belongs to
// 	CourseID     *uuid.UUID `json:"course_id" gorm:"not null;"`                                              // course the task belongs to ID
// 	Title        *string    `json:"title" gorm:"not null;max=250;"`                                          // title of the task
// 	Description  *string    `json:"description" `                                                            // description of the task
// 	TotalPoint   *uint      `json:"total_point" `                                                            // total point of the task
// 	File         *string    `json:"file" `                                                                   // file of the task
// 	TaskType     *string    `json:"task_type"`                                                               // EXAM or ASSIGNMENT or TEST or QUIZ
// 	ClassAverage *int       `json:"class_average"`
// 	TotalStudent *int       `json:"total_student"`
// 	DueDate      *time.Time `json:"due_date" `
// }

// func (courseTask *CourseTask) CalculateClassAverage(ctx context.Context, db *gorm.DB, id uuid.UUID, taskPoint uint) error {
// 	var totalPoint int64
// 	var numberOfStudents int64

// 	// Fetch all the student tasks for the course task
// 	err := db.WithContext(ctx).Model(&StudentTask{}).
// 		Where("course_task_id = ?", id).
// 		Pluck("SUM(total_point)", &totalPoint).Count(&numberOfStudents).Error
// 	if err != nil {
// 		return err
// 	}

// 	// Calculate class average, handle division by zero
// 	var classAverage float64
// 	if numberOfStudents > 0 {
// 		classAverage = (float64(totalPoint) / (float64(numberOfStudents) * float64(taskPoint))) * 100
// 	}

// 	// Update the course task with the calculated class average
// 	err = db.WithContext(ctx).Model(&CourseTask{}).
// 		Where("id = ?", id).
// 		Update("class_average", classAverage).Error
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// // Question /* This is the question with the foreign key related */
// type Question struct {
// 	BaseModel
// 	School            *School           `json:"school,omitempty" gorm:"constraint:OnDelete:CASCADE;ForeignKey:SchoolID"`
// 	SchoolID          *uuid.UUID        `json:"school_id" gorm:"not null;"`
// 	CourseTask        *CourseTask       `json:"course_task,omitempty" gorm:"constraint:OnDelete:CASCADE;ForeignKey:CourseTaskID"` // CourseTask the question belongs to
// 	CourseTaskID      *uuid.UUID        `json:"course_task_id" gorm:"not null;"`                                                  // CourseTask the question belongs to ID
// 	Question          *string           `json:"question" gorm:"not null;max=250;"`                                                // question of the task
// 	QuestionType      *string           `json:"question_type"`                                                                    // this could be TEXT or OPTION or BOOLEAN  OR FILL_IN_THE_BLANK
// 	File              *string           `json:"file" `                                                                            // file of the task
// 	Point             *uint             `json:"point" gorm:"not null;"`                                                           // point of the task
// 	FillInBlankAnswer *StringList       `json:"fill_in_blank_answer"`                                                             // answer if its a fill in blank question
// 	BooleanAnswer     *bool             `json:"boolean_answer"`                                                                   // answer if its a boolean question
// 	Options           []*QuestionOption `json:"options,omitempty" gorm:"many2many:question_question_options;"`                    // Define many-to-many relationship
// }

// type QuestionOption struct {
// 	BaseModel
// 	School    *School    `json:"school,omitempty" gorm:"constraint:OnDelete:CASCADE;ForeignKey:SchoolID"`
// 	SchoolID  *uuid.UUID `json:"school_id" gorm:"not null;"`
// 	Option    *string    `json:"option" gorm:"not null;max=250;"` // option of the question
// 	IsCorrect *bool      `json:"is_correct" gorm:"not null;"`
// }

// func (q *QuestionOption) FindQuestionOptionByQuestionAndOptionID(ctx context.Context, db *gorm.DB, questionID uuid.UUID, optionID uuid.UUID) (*QuestionOption, error) {
// 	var questionOption QuestionOption

// 	if err := db.Joins("JOIN question_question_options ON question_question_options.question_option_id = question_options.id").
// 		Where("question_question_options.question_id = ? AND question_question_options.question_option_id = ?", questionID, optionID).
// 		First(&questionOption).Error; err != nil {
// 		if errors.Is(err, gorm.ErrRecordNotFound) {
// 			return nil, errors.New("QuestionOption not found for the provided QuestionID and OptionID")
// 		}
// 		logger.Error(ctx, "Error filtering many to many field", zap.Error(err))
// 		return nil, err
// 	}

// 	return &questionOption, nil
// }

// type QuestionQuestionOption struct {
// 	BaseModel
// 	QuestionID       uuid.UUID `json:"question_id" gorm:"type:uuid;primaryKey"`
// 	QuestionOptionID uuid.UUID `json:"question_option_id" gorm:"type:uuid;primaryKey"`
// }

// // StudentTask /* This is an task model for a student which is tied to the Course task */
// type StudentTask struct {
// 	BaseModel
// 	School         *School     `json:"school,omitempty" gorm:"constraint:OnDelete:CASCADE;ForeignKey:SchoolID"`
// 	SchoolID       *uuid.UUID  `json:"school_id" gorm:"not null;"`
// 	Student        *Student    `json:"student,omitempty" gorm:"constraint:OnDelete:CASCADE;ForeignKey:StudentID"` // the student task belongs to
// 	StudentID      *uuid.UUID  `json:"student_id" gorm:"not null;"`                                               // the student task belongs to ID
// 	CourseTask     *CourseTask `json:"course_task,omitempty" gorm:"ForeignKey:CourseTaskID"`                      // task the student belongs to
// 	CourseTaskID   *uuid.UUID  `json:"task_id" gorm:"not null;"`                                                  // task the student belongs to ID
// 	StartTime      *time.Time  `json:"start_time" gorm:"not null"`                                                // The time the student starts taking the task
// 	DueDate        *time.Time  `json:"due_date" `                                                                 // The time the student is due to submit the task
// 	SubmissionTime *time.Time  `json:"submission_time" `                                                          // The time the student submits the task
// 	TotalPoint     *uint       `json:"total_point" `                                                              // The total point the student gets for the task
// 	Status         *string     `json:"status" gorm:"default:PENDING"`                                             // this could be PENDING, SUBMITTED
// }

// // CreateStudentTask to create the task that is attached to the StudentFor that Course
// func (studentTask *StudentTask) CreateStudentTask(ctx context.Context, db *gorm.DB, schoolID uuid.UUID, courseID uuid.UUID, studentID uuid.UUID) (err error) {

// 	var courseTasks []CourseTask

// 	// create the assessment related to this course for this student
// 	err = db.Model(&CourseTask{}).Where("course_id = ?", courseID).Find(&courseTasks).Error
// 	if err != nil {
// 		return err
// 	}

// 	// loop through all the course tasks and create the student task
// 	for _, courseTask := range courseTasks {
// 		var courseTaskCount int64

// 		err := db.WithContext(ctx).Model(&studentTask).Where("id = ? AND student_id =? and status =?", courseTask.ID, studentID, "PENDING").
// 			Count(&courseTaskCount).Error
// 		if err != nil {
// 			logger.Error(ctx, "Error counting student task", zap.Error(err))
// 			return err
// 		}
// 		if courseTaskCount > 0 {
// 			continue // check if the task exists on the student and continues
// 		}

// 		startTime := time.Now()
// 		status := "PENDING"
// 		studentTask := StudentTask{
// 			SchoolID:       &schoolID,
// 			StudentID:      &studentID,
// 			CourseTaskID:   &courseTask.ID,
// 			StartTime:      &startTime,
// 			DueDate:        courseTask.DueDate,
// 			SubmissionTime: nil,
// 			Status:         &status,
// 		}
// 		logger.Info(ctx, "The student task", zap.Any("studentTask", studentTask))

// 		err = db.WithContext(ctx).Model(&StudentTask{}).Create(&studentTask).Error
// 		if err != nil {
// 			logger.Error(ctx, "Error creating student task", zap.Error(err))
// 			return err
// 		}
// 		// update a registered student in the course task
// 		var totalStudent int
// 		if courseTask.TotalStudent != nil {
// 			totalStudent = *courseTask.TotalStudent + 1
// 		} else {
// 			totalStudent = 1
// 		}
// 		err = db.Model(&CourseTask{}).Where("course_id = ?", courseID).Update("total_student", totalStudent).Error
// 		if err != nil {
// 			logger.Error(ctx, "error updating course task", zap.Error(err))
// 			return err
// 		}
// 	}
// 	return nil
// }

// // CalculateTaskPoint This is used to calculate the point gotten by the student for the task
// func (studentTask *StudentTask) CalculateTaskPoint(ctx context.Context, db *gorm.DB, studentID uuid.UUID, courseID uuid.UUID, taskType string) error {
// 	var enrolledCourse EnrolledCourse

// 	pointerAdapters := adapters.NewPointer()

// 	// add all the student task type point together
// 	err := db.WithContext(ctx).Model(&EnrolledCourse{}).
// 		Where("student_id = ? AND  course_id =? ", studentID, courseID).First(&enrolledCourse).Error
// 	if err != nil {
// 		return err
// 	}

// 	if enrolledCourse.ExamScore == nil {
// 		enrolledCourse.ExamScore = pointerAdapters.UIntPointer(0)
// 	}
// 	if enrolledCourse.TestScore == nil {
// 		enrolledCourse.TestScore = pointerAdapters.UIntPointer(0)
// 	}
// 	if enrolledCourse.TestScore == nil {
// 		enrolledCourse.TestScore = pointerAdapters.UIntPointer(0)
// 	}
// 	if enrolledCourse.AssignmentScore == nil {
// 		enrolledCourse.AssignmentScore = pointerAdapters.UIntPointer(0)
// 	}

// 	if enrolledCourse.Total == nil {
// 		enrolledCourse.Total = pointerAdapters.UIntPointer(0)
// 	}

// 	switch taskType {
// 	case "EXAM":
// 		*enrolledCourse.ExamScore += *studentTask.TotalPoint
// 	case "TEST":
// 		*enrolledCourse.TestScore += *studentTask.TotalPoint
// 	case "ASSIGNMENT":
// 		*enrolledCourse.AssignmentScore += *studentTask.TotalPoint
// 	}

// 	*enrolledCourse.Total = *enrolledCourse.ExamScore + *enrolledCourse.TestScore + *enrolledCourse.AssignmentScore

// 	// update the enroll course
// 	err = db.WithContext(ctx).Model(&EnrolledCourse{}).
// 		Where("student_id = ? AND  course_id =? ", studentID, courseID).Updates(&enrolledCourse).Error
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// // StudentAnswer /* This is the answer of a student to a question */
// type StudentAnswer struct {
// 	BaseModel
// 	School        *School      `json:"school,omitempty" gorm:"constraint:OnDelete:CASCADE;ForeignKey:SchoolID"`
// 	SchoolID      *uuid.UUID   `json:"school_id" gorm:"not null;"`
// 	StudentTask   *StudentTask `json:"student_task,omitempty" gorm:"constraint:OnDelete:CASCADE;ForeignKey:StudentTaskID"` // student task the answer belongs to
// 	StudentTaskID *uuid.UUID   `json:"student_task_id" gorm:"not null;"`                                                   // student task the answer belongs to ID
// 	Question      *Question    `json:"question,omitempty" gorm:"constraint:OnDelete:CASCADE;ForeignKey:QuestionID"`        // the question the answer belongs to
// 	QuestionID    *uuid.UUID   `json:"question_id" gorm:"not null;"`                                                       // question the answer belongs to ID
// 	Answer        *string      `json:"answer" gorm:"not null;"`                                                            // answer of the student
// 	Point         *uint        `json:"point" gorm:"not null;"`                                                             // point of the answer or the point given to the student for the answer
// }
