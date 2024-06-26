package serializers

// import (
// 	"Learnium/models"
// 	"github.com/google/uuid"
// 	"time"
// )

// type AcceptEnrollmentSerializer struct {
// 	ApplicantID *uuid.UUID `json:"applicant_id" validate:"required"`
// 	ClassID     *uuid.UUID `json:"class_id" `
// 	Status      *string    `json:"status" validate:"oneof=APPROVE DECLINE "`
// }

// type EnrollmentApplicationRequestSerializer struct {
// 	SchoolID                     *uuid.UUID                        `form:"schoolID" `
// 	EnrollmentConfigurationID    *uuid.UUID                        `form:"enrollment_configuration_id" `
// 	FullName                     *string                           `form:"full_name" `
// 	DateOfBirth                  *time.Time                        `form:"date_of_birth"`
// 	HomeAddress                  *string                           `form:"home_address" `
// 	Allergies                    *string                           `form:"allergies" `
// 	CurrentMedication            *string                           `form:"current_medication" `
// 	NextOfKin                    *string                           `form:"next_of_kin" `
// 	NextOfKinPhone               *string                           `form:"next_of_kin_phone" `
// 	PreviousSchool               *string                           `form:"previous_school" `
// 	CurrentGradeLevel            *string                           `form:"current_grade_level" `
// 	Gender                       *string                           `form:"gender"  validate:"oneof=MALE FEMALE OTHERS"`
// 	Email                        *string                           `form:"email" `
// 	Phone                        *string                           `form:"phone" `
// 	Religion                     *string                           `form:"religion" `
// 	BirthCertificate             *string                           `form:"birth_certificate" `
// 	ImmunizationRecord           *string                           `form:"immunization_record" `
// 	ProofOfAddress               *string                           `form:"proof_of_address" `
// 	ReportCard                   *string                           `form:"report_card" `
// 	OtherSupportingDocument      *string                           `form:"other_supporting_document" `
// 	EnrollmentCustomFieldAnswers *models.CustomFieldAnswers        `form:"enrollment_custom_field_answers" `
// 	AdditionalQuestionsAnswer    *models.AdditionalQuestionAnswers `form:"additional_questions_answer" `
// }

// type EnrollmentAnalyticsResponseSerializer struct {
// 	AllApplicants      int64 `json:"all_applicants"`
// 	ApprovedApplicants int64 `json:"approved_applicants"`
// 	DeclinedApplicants int64 `json:"declined_applicants"`
// 	PendingApplicants  int64 `json:"pending_applicants"`
// }
