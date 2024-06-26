package models

import (
	"context"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type EnrollmentConfiguration struct {
	BaseModel
	School                          *School              `json:"school,omitempty" gorm:"constraint:OnDelete:CASCADE;ForeignKey:SchoolID"`
	SchoolID                        *uuid.UUID           `json:"schoolID" gorm:"not null;"`
	FullName                        *bool                `json:"full_name" gorm:"default:true;"`
	FullNameRequired                *bool                `json:"full_name_required" gorm:"default:true;"`
	DateOfBirth                     *bool                `json:"date_of_birth" gorm:"default:true;"`
	DateOfBirthRequired             *bool                `json:"date_of_birth_required" gorm:"default:true;"`
	HomeAddress                     *bool                `json:"home_address" gorm:"default:true;"`
	HomeAddressRequired             *bool                `json:"home_address_required" gorm:"default:true;"`
	Allergies                       *bool                `json:"allergies" gorm:"default:true;"`
	AllergiesRequired               *bool                `json:"allergies_required" gorm:"default:true;"`
	CurrentMedication               *bool                `json:"current_medication" gorm:"default:true"`
	CurrentMedicationRequired       *bool                `json:"current_medication_required" gorm:"default:true"`
	NextOfKin                       *bool                `json:"next_of_kin" gorm:"default:false"`
	NextOfKinRequired               *bool                `json:"next_of_kin_required" gorm:"default:false"`
	NextOfKinPhone                  *bool                `json:"next_of_kin_phone" gorm:"default:false"`
	NextOfKinPhoneRequired          *bool                `json:"next_of_kin_phone_required" gorm:"default:false"`
	PreviousSchool                  *bool                `json:"previous_school" gorm:"default:false"`
	PreviousSchoolRequired          *bool                `json:"previous_school_required" gorm:"default:false"`
	CurrentGradeLevel               *bool                `json:"current_grade_level" gorm:"default:false"`
	CurrentGradeLevelRequired       *bool                `json:"current_grade_level_required" gorm:"default:false"`
	Gender                          *bool                `json:"gender" gorm:"default:true"`
	GenderRequired                  *bool                `json:"gender_required" gorm:"default:true"`
	Email                           *bool                `json:"email" gorm:"default:true;"`
	EmailRequired                   *bool                `json:"email_required" gorm:"default:true;"`
	Phone                           *bool                `json:"phone" gorm:"default:true;"`
	PhoneRequired                   *bool                `json:"phone_required" gorm:"default:true;"`
	ParentPhone                     *bool                `json:"parent_phone" gorm:"default:true;"`
	ParentPhoneRequired             *bool                `json:"parent_phone_required" gorm:"default:true;"`
	ParentWard                      *bool                `json:"parent_ward" gorm:"default:true;"`
	ParentWardRequired              *bool                `json:"parent_ward_required" gorm:"default:true;"`
	Religion                        *bool                `json:"religion" gorm:"default:true;"`
	ReligionRequired                *bool                `json:"religion_required" gorm:"default:true;"`
	BirthCertificate                *bool                `json:"birth_certificate" gorm:"default:true"`
	BirthCertificateRequired        *bool                `json:"birth_certificate_required" gorm:"default:true"`
	ImmunizationRecord              *bool                `json:"immunization_record" gorm:"default:false"`
	ImmunizationRecordRequired      *bool                `json:"immunization_record_required" gorm:"default:false"`
	ProofOfAddress                  *bool                `json:"proof_of_address" gorm:"default:false"`
	ProofOfAddressRequired          *bool                `json:"proof_of_address_required" gorm:"default:false"`
	ReportCard                      *bool                `json:"report_card" gorm:"default:false"`
	ReportCardRequired              *bool                `json:"report_card_required" gorm:"default:false"`
	OtherSupportingDocument         *bool                `json:"other_supporting_document" gorm:"default:false"`
	OtherSupportingDocumentRequired *bool                `json:"other_supporting_document_required" gorm:"default:false"`
	EnrollmentCustomFields          *CustomFields        `json:"enrollment_custom_fields" gorm:"type:jsonb;"`
	AdditionalQuestions             *AdditionalQuestions `json:"additional_questions" gorm:"type:jsonb;"`
}

func (course *EnrollmentConfiguration) Retrieve(ctx context.Context, db *gorm.DB, id uuid.UUID) (EnrollmentConfiguration, error) {
	// filter on gorm
	err := db.WithContext(ctx).Model(&course).Where("id = ?", id).Find(&course).Error
	if err != nil {
		return *course, err
	}

	return *course, err
}

type EnrollmentAdmission struct {
	BaseModel
	School                       *School                    `json:"school,omitempty" gorm:"constraint:OnDelete:CASCADE;ForeignKey:SchoolID"`
	SchoolID                     *uuid.UUID                 `json:"schoolID" gorm:"not null;"`
	EnrollmentConfiguration      *EnrollmentConfiguration   `json:"enrollment_configuration,omitempty" gorm:"constraint:OnDelete:CASCADE;ForeignKey:EnrollmentConfigurationID"`
	EnrollmentConfigurationID    *uuid.UUID                 `json:"enrollment_configuration_id" gorm:"not null;"`
	FullName                     *string                    `json:"full_name" gorm:"max=500;"`
	Status                       *string                    `json:"status" gorm:"default:PENDING;"` // APPROVED / DECLINED/ PENDING
	DateOfBirth                  *time.Time                 `json:"date_of_birth"`
	HomeAddress                  *string                    `json:"home_address" gorm:"max=500;"`
	Allergies                    *string                    `json:"allergies" gorm:"max=500;"`
	CurrentMedication            *string                    `json:"current_medication" gorm:"max=1000;"`
	NextOfKin                    *string                    `json:"next_of_kin" gorm:"max=500;"`
	NextOfKinPhone               *string                    `json:"next_of_kin_phone" gorm:"max=500;"`
	PreviousSchool               *string                    `json:"previous_school" gorm:"max=500;"`
	CurrentGradeLevel            *string                    `json:"current_grade_level" gorm:"max=500;"`
	Gender                       *string                    `json:"gender" gorm:"max=500;" validate:"oneof=MALE FEMALE OTHERS"`
	Email                        *string                    `json:"email" gorm:"max=500;"`
	Phone                        *string                    `json:"phone" gorm:"max=500;"`
	ParentPhone                  *string                    `json:"parent_phone" gorm:"max=500;"`
	ParentWard                   *string                    `json:"parent_ward"`
	Religion                     *string                    `json:"religion" gorm:"max=500;"`
	BirthCertificate             *string                    `json:"birth_certificate" `
	ImmunizationRecord           *string                    `json:"immunization_record" `
	ProofOfAddress               *string                    `json:"proof_of_address" gorm:"max=500;"`
	ReportCard                   *string                    `json:"report_card" `
	OtherSupportingDocument      *string                    `json:"other_supporting_document" `
	EnrollmentCustomFieldAnswers *CustomFieldAnswers        `json:"enrollment_custom_field_answers" gorm:"type:jsonb;"`
	AdditionalQuestionsAnswer    *AdditionalQuestionAnswers `json:"additional_questions_answer" gorm:"type:jsonb;"`
}

func (enrollmentAdmission *EnrollmentAdmission) Retrieve(ctx context.Context, db *gorm.DB, id uuid.UUID) (EnrollmentAdmission, error) {
	// filter on gorm
	err := db.WithContext(ctx).Model(&enrollmentAdmission).Where("id = ?", id).First(&enrollmentAdmission).Error
	if err != nil {
		return *enrollmentAdmission, err
	}

	return *enrollmentAdmission, err
}
