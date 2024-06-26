package serializers

// import (
// 	"Learnium/models"
// 	"github.com/google/uuid"
// 	"time"
// )

// type StudentListSerializer struct {
// 	ID                    *uuid.UUID           `json:"id" `
// 	StudentCode           *string              `json:"student_code" `
// 	EnrollmentAdmissionID *uuid.UUID           `json:"enrollment_admission_id" `
// 	EnrollmentAdmission   *EnrollmentAdmission `json:"enrollment_admission" `
// 	SchoolID              *uuid.UUID           `json:"school_id" `
// 	User                  *User                `json:"user"` //	the one-to-one relationship
// 	UserID                *uuid.UUID           `json:"user_id" `
// 	Class                 *models.Class        `json:"class"` //	the one-to-one relationship
// 	ClassID               *uuid.UUID           `json:"class_id" `
// 	Department            *string              `json:"department" `
// 	Section               *string              `json:"section" `
// 	Status                *string              `json:"status" `
// }

// type NoteCreateSerializer struct {
// 	StudentID *uuid.UUID `json:"student_id" validate:"required"`
// 	NoteType  *string    `json:"note_type" validate:"required,eq=ACADEMIC|eq=BEHAVIOR|eq=EMOTION|eq=OTHERS"`                                  //  ACADEMIC , BEHAVIOR , EMOTION , OTHERS
// 	Share     *string    `json:"share"  validate:"required,eq=DONT_SHARE|eq=VISIBLE_TO_TEACHERS|eq=VISIBLE_TO_STUDENTS|eq=WITH_PARENTS_ONLY"` // DONT_SHARE, VISIBLE_TO_TEACHERS, VISIBLE_TO_STUDENTS, WITH_PARENTS_ONLY , ADMIN
// 	Title     *string    `json:"title" validate:"required"`
// 	Note      *string    `json:"note" validate:"required"`
// }

// type NoteUpdateSerializer struct {
// 	NoteType *string `json:"note_type,omitempty" validate:"eq=ACADEMIC|eq=BEHAVIOR|eq=EMOTION|eq=OTHERS"`                                  //  ACADEMIC , BEHAVIOR , EMOTION , OTHERS
// 	Share    *string `json:"share,omitempty"  validate:"eq=DONT_SHARE|eq=VISIBLE_TO_TEACHERS|eq=VISIBLE_TO_STUDENTS|eq=WITH_PARENTS_ONLY"` // DONT_SHARE, VISIBLE_TO_TEACHERS, VISIBLE_TO_STUDENTS, WITH_PARENTS_ONLY , ADMIN
// 	Title    *string `json:"title" validate:""`
// 	Note     *string `json:"note" validate:""`
// }

// type ClinicVisitationCreateRequestSerializer struct {
// 	StudentID      *uuid.UUID `json:"student_id"   validate:"required" `    // added user_id to the serializer so you can create by ur self
// 	Name           *string    `json:"name"    validate:"required"`          // name of the clinic
// 	Description    *string    `json:"description"   validate:"required"`    // what the user went to do in the clinic
// 	DoctorsName    *string    `json:"doctors_name"    validate:"required" ` // The name of the doctor
// 	VisitationTime *time.Time `json:"visitation_time"  validate:"required"`
// 	Status         *string    `json:"status"   validate:"required"` // the status of the visit if its for something very important
// }
// type ClinicVisitationUpdateRequestSerializer struct {
// 	Name           *string    `json:"name"  `          // name of the clinic
// 	Description    *string    `json:"description" `    // what the user went to do in the clinic
// 	DoctorsName    *string    `json:"doctors_name"   ` // The name of the doctor
// 	VisitationTime *time.Time `json:"visitation_time" `
// 	Status         *string    `json:"status" ` // the status of the visit if its for something very important
// }
