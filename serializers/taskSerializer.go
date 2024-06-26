package serializers

// import (
// 	"Learnium/models"
// 	"github.com/google/uuid"
// 	"time"
// )

// type TaskCreateRequestSerializer struct {
// 	CourseID    *uuid.UUID `form:"course_id" validate:"required"`
// 	Title       *string    `form:"title" validate:"required"`       // title of the task
// 	Description *string    `form:"description" validate:"required"` // description of the task
// 	TotalPoint  *uint      `form:"total_point" validate:"required"` // total point of the task
// 	File        *string    `form:"file" `                           // file of the task
// 	DueDate     *time.Time `form:"due_date" validate:"required"`
// 	TaskType    *string    `form:"task_type" validate:"required,eq=EXAM|eq=ASSIGNMENT|eq=TEST|eq=QUIZ"`
// }

// type TaskUpdateRequestSerializer struct {
// 	Title       *string    `form:"title" `       // title of the task
// 	Description *string    `form:"description" ` // description of the task
// 	TotalPoint  *uint      `form:"total_point" ` // total point of the task
// 	File        *string    `form:"file" `        // file of the task
// 	DueDate     *time.Time `form:"due_date" `
// }

// type TaskQuestionCreateRequestSerializer struct {
// 	CourseTaskID      *uuid.UUID                         `form:"course_task_id" validate:"required" `                                                 // CourseTask the question belongs to ID
// 	Question          *string                            `form:"question"  validate:"required"`                                                       // question of the task
// 	File              *string                            `form:"file" `                                                                               // file of the task
// 	Point             *uint                              `form:"point"  validate:"required"`                                                          // point of the task
// 	QuestionType      *string                            `form:"question_type" validate:"required,eq=TEXT|eq=OPTION|eq=BOOLEAN|eq=FILL_IN_THE_BLANK"` // this could be TEXT or OPTION or BOOLEAN or FILL_IN_THE_BLANK
// 	FillInBlankAnswer *models.StringList                 `json:"fill_in_blank_answer"`                                                                // answer if its a fill in blank question
// 	BooleanAnswer     *bool                              `form:"boolean_answer"`
// 	Options           *[]QuestionOptionRequestSerializer `form:"options" `
// }

// type TaskQuestionUpdateRequestSerializer struct {
// 	Question *string `json:"question" ` // question of the task
// 	File     *string `json:"file" `     // file of the task
// 	Point    *uint   `json:"point" `    // point of the task
// }

// type QuestionOptionRequestSerializer struct {
// 	Option    *string `json:"option" validate:"required"` // option of the question
// 	IsCorrect *bool   `json:"is_correct" validate:"required"`
// }

// type QuestionOptionCreateRequestSerializer struct {
// 	QuestionID *uuid.UUID `json:"question_id" validate:"required"`
// 	Option     *string    `json:"option" validate:"required"` // option of the question
// 	IsCorrect  *bool      `json:"is_correct" validate:"required"`
// }

// type StudentAnswerRequestSerializer struct {
// 	QuestionID        *uuid.UUID         `json:"question_id" `
// 	TextAnswer        *string            `json:"text_answer" `
// 	BooleanAnswer     *bool              `json:"boolean_answer"`
// 	FillInBlankAnswer *models.StringList `json:"fill_in_blank_answer"`
// 	OptionID          *uuid.UUID         `json:"option_id"`
// 	AnswerType        *string            `json:"answer_type" validate:"required,eq=TEXT|eq=OPTION|eq=BOOLEAN|eq=FILL_IN_THE_BLANK"`
// }

// type CourseTaskDetailSerializer struct {
// 	ID           *uuid.UUID        `json:"id"`
// 	CourseID     *uuid.UUID        `json:"course_id" `   // course the task belongs to ID
// 	Title        *string           `json:"title" `       // title of the task
// 	Description  *string           `json:"description" ` // description of the task
// 	TotalPoint   *uint             `json:"total"`        // total point of the task
// 	File         *string           `json:"file" `        // file of the task
// 	ClassAverage *int              `json:"class_average"`
// 	TaskType     *string           `json:"task_type"` // EXAM or ASSIGNMENT or TEST
// 	TotalStudent *int              `json:"total_student"`
// 	DueDate      *time.Time        `json:"due_date" `
// 	Questions    []models.Question `json:"questions"`
// }

// type StudentTaskDetailSerializer struct {
// 	ID             *uuid.UUID             `json:"id"`
// 	Student        *models.Student        `json:"student,omitempty"`      // the student task belongs to
// 	StudentID      *uuid.UUID             `json:"student_id" `            // the student task belongs to ID
// 	CourseTask     *models.CourseTask     `json:"course_task,omitempty" ` // task the student belongs to
// 	CourseTaskID   *uuid.UUID             `json:"task_id" `               // task the student belongs to ID
// 	StartTime      *time.Time             `json:"start_time" `            // The time the student starts taking the task
// 	DueDate        *time.Time             `json:"due_date" `              // The time the student is due to submit the task
// 	SubmissionTime *time.Time             `json:"submission_time" `       // The time the student submits the task
// 	TotalPoint     *uint                  `json:"total_point" `           // The total point the student gets for the task
// 	Status         *string                `json:"status" `
// 	StudentAnswers []models.StudentAnswer `json:"student_answer"`
// }

// // StudentCourseTaskDetailSerializer /*This is used to get the detail of an assigment by the student*/
// type StudentCourseTaskDetailSerializer struct {
// 	ID             *uuid.UUID                               `json:"id"`
// 	Student        *models.Student                          `json:"student"`          // the student task belongs to
// 	StudentID      *uuid.UUID                               `json:"student_id" `      // the student task belongs to ID
// 	CourseTask     *StudentCourseTaskCourseDetailSerializer `json:"course_task" `     // task the student belongs to
// 	CourseTaskID   *uuid.UUID                               `json:"task_id" `         // task the student belongs to ID
// 	StartTime      *time.Time                               `json:"start_time" `      // The time the student starts taking the task
// 	DueDate        *time.Time                               `json:"due_date" `        // The time the student is due to submit the task
// 	SubmissionTime *time.Time                               `json:"submission_time" ` // The time the student submits the task
// 	TotalPoint     *uint                                    `json:"total_point" `     // The total point the student gets for the task
// 	Status         *string                                  `json:"status" `
// 	StudentAnswers []models.StudentAnswer                   `json:"student_answer"`
// }

// // StudentCourseTaskCourseDetailSerializer this shows all the detail of a course task excluding the answer of an option
// type StudentCourseTaskCourseDetailSerializer struct {
// 	ID           *uuid.UUID         `json:"id"`
// 	SchoolID     *uuid.UUID         `json:"school_id" `
// 	CourseID     *uuid.UUID         `json:"course_id" `   // course the task belongs to ID
// 	Title        *string            `json:"title" `       // title of the task
// 	Description  *string            `json:"description" ` // description of the task
// 	TotalPoint   *uint              `json:"total_point" ` // total point of the task
// 	File         *string            `json:"file" `        // file of the task
// 	ClassAverage *int               `json:"class_average"`
// 	DueDate      *time.Time         `json:"due_date" `
// 	Question     []*StudentQuestion `json:"question"`
// }

// type StudentQuestion struct {
// 	ID           *uuid.UUID               `json:"id"`
// 	SchoolID     *uuid.UUID               `json:"school_id" `
// 	Question     *string                  `json:"question" `          // question of the task
// 	QuestionType *string                  `json:"question_type"`      // this could be TEXT or OPTION
// 	File         *string                  `json:"file" `              // file of the task
// 	Point        *uint                    `json:"point" `             // point of the task
// 	Options      []*StudentQuestionOption `json:"options,omitempty" ` // Define many-to-many relationship
// }

// type StudentQuestionOption struct {
// 	ID       *uuid.UUID `json:"id"`
// 	SchoolID *uuid.UUID `json:"school_id" gorm:"not null;"`
// 	Option   *string    `json:"option" gorm:"not null;max=250;"` // option of the question
// }

// type StudentTaskSubmitRequestSerializer struct {
// 	StudentTaskID *uuid.UUID `json:"student_task_id" validate:"required"`
// }
