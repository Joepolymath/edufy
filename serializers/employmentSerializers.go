package serializers

// import (
// 	"Learnium/models"
// 	"github.com/google/uuid"
// 	"time"
// )

// type AcceptEmploymentSerializer struct {
// 	EmploymentID *uuid.UUID `json:"employment_id" validate:"required"`
// 	ApplicantID  *uuid.UUID `json:"applicant_id" validate:"required"`
// 	ClassID      *uuid.UUID `json:"class_id" `
// 	Status       *string    `json:"status" validate:"eq=APPROVE|eq=DECLINE"`
// }

// // EmploymentApplicationRequestSerializer /* this is the application for the employment*/
// type EmploymentApplicationRequestSerializer struct {
// 	/*
// 		This is the application form where teacher input his or her details
// 	*/
// 	SchoolID                     *uuid.UUID                        `form:"schoolID" `
// 	EmploymentID                 *uuid.UUID                        `form:"employment_id" `
// 	FirstName                    *string                           `form:"first_name" `
// 	LastName                     *string                           `form:"last_name" `
// 	Email                        *string                           `form:"email" `
// 	Phone                        *string                           `form:"phone" `
// 	Location                     *string                           `form:"location" `
// 	Gender                       *string                           `form:"gender"  validate:"eq=MALE|eq=FEMALE"`
// 	DateOfBirth                  *time.Time                        `form:"date_of_birth" `
// 	HomeAddress                  *string                           `form:"home_address" `
// 	State                        *string                           `form:"state" `
// 	NextOfKin                    *string                           `form:"next_of_kin" `
// 	NextOfKinPhone               *string                           `form:"next_of_kin_phone" `
// 	MaritalStatus                *string                           `form:"marital_status,omitempty"  validate:"oneof=SINGLE MARRIED RELATIONSHIP"`
// 	Religion                     *string                           `form:"religion" `
// 	YearsOfExperience            *string                           `form:"years_of_experience" `
// 	Degree                       *string                           `form:"degree" `
// 	YearOfGraduation             *string                           `form:"year_of_graduation" `
// 	PreviousRole                 *string                           `form:"previous_role" `
// 	SchoolOfGraduation           *string                           `form:"school_of_graduation" `
// 	ExpectedSalary               *string                           `form:"expected_salary" `
// 	OtherQualification           *string                           `form:"other_qualification" `
// 	CurriculumVideo              *string                           `form:"curriculum_video" `
// 	Resume                       *string                           `form:"resume" `
// 	CoverLetter                  *string                           `form:"cover_letter" `
// 	References                   *string                           `form:"references" `
// 	Transcript                   *string                           `form:"transcript" `
// 	TeachingCertification        *string                           `form:"teaching_certification" `
// 	LetterOfRecommendation       *string                           `form:"letter_of_recommendation" `
// 	EmploymentCustomFieldAnswers *models.CustomFieldAnswers        `form:"employment_custom_field_answers" json:"employment_custom_field_answers" `
// 	AdditionalQuestionsAnswer    *models.AdditionalQuestionAnswers `form:"additional_questions_answer" json:"additional_questions_answer" `
// }

// type EnrollmentAdmission struct {
// 	ID          *uuid.UUID `json:"id"`
// 	ParentWard  *string    `json:"parent_ward"`
// 	ParentPhone *string    `json:"parent_phone"`
// }
