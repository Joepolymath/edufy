package models

// import (
// 	"context"
// 	"time"

// 	"github.com/google/uuid"
// 	"gorm.io/gorm"
// )

// // Employment this would be the job which is available to the public for teachers to apply
// type Employment struct {
// 	/*
// 		The employment is a job which also have a configuration that why we have a
// 		foreign key in here were we have created a configuration and set it to this particular job
// 	*/
// 	BaseModel
// 	School                    *School                  `json:"school,omitempty" gorm:"constraint:OnDelete:CASCADE;ForeignKey:SchoolID"`                                     // school the employment belongs to
// 	SchoolID                  *uuid.UUID               `json:"schoolID" gorm:"not null;"`                                                                                   // school the employment belongs to
// 	EmploymentConfiguration   *EmploymentConfiguration `json:"employment_configuration,omitempty" gorm:"constraint:OnDelete:CASCADE;foreignKey:EmploymentConfigurationID;"` // the employment configuration
// 	EmploymentConfigurationID *uuid.UUID               `json:"employment_configuration_id" gorm:"not null;" validate:"required"`                                            // the employment configuration
// 	// Role                      *Role                    `json:"role,omitempty" gorm:"constraint:OnDelete:SET NULL;ForeignKey:RoleID" `                                       // the role the employment belongs to
// 	RoleID         *uuid.UUID `json:"role_id" gorm:"not null;" validate:"required"`                                 // the role the employment belongs to
// 	Category       *Category  `json:"category,omitempty" gorm:"constraint:OnDelete:SET NULL;ForeignKey:CategoryID"` // the category the employment belongs to
// 	CategoryID     *uuid.UUID `json:"category_id" gorm:"not null;" validate:"required"`                             // the category the employment belongs to
// 	Description    *string    `json:"description" gorm:"not null;max=5000;" validate:"required"`
// 	EmploymentType *string    `json:"employment_type" gorm:"not null;max=250;" validate:"required"`
// 	Code           *string    `json:"code" gorm:"not null;max=250;" validate:"required"`
// }

// func (employment *Employment) Retrieve(ctx context.Context, db *gorm.DB, id uuid.UUID) (Employment, error) {
// 	// filter on gorm
// 	err := db.WithContext(ctx).Model(&employment).Where("id = ?", id).First(&employment).Error
// 	if err != nil {
// 		return *employment, err
// 	}

// 	return *employment, err
// }

// // EmploymentConfiguration /* the employment configuration is tied to the employment itself*/
// type EmploymentConfiguration struct {
// 	/*
// 		These are the default settings we set for the configuration of the project
// 	*/
// 	BaseModel

// 	// contact info
// 	School                 *School    `json:"school,omitempty" gorm:"constraint:OnDelete:CASCADE;ForeignKey:SchoolID"`
// 	SchoolID               *uuid.UUID `json:"schoolID" gorm:"not null;"`
// 	FirstName              *bool      `json:"first_name" gorm:"default:true;"`
// 	FirstNameRequired      *bool      `json:"first_name_required" gorm:"default:true;"`
// 	LastName               *bool      `json:"last_name" gorm:"default:true;"`
// 	LastNameRequired       *bool      `json:"last_name_required" gorm:"default:true;"`
// 	Email                  *bool      `json:"email" gorm:"not null;default:true;"`
// 	EmailRequired          *bool      `json:"email_required" gorm:"not null;default:true;"`
// 	Phone                  *bool      `json:"phone" gorm:"not null;default:true;"`
// 	PhoneRequired          *bool      `json:"phone_required" gorm:"not null;default:true;"`
// 	Location               *bool      `json:"location" gorm:"not null;default:true;"`
// 	LocationRequired       *bool      `json:"location_required" gorm:"not null;default:true;"`
// 	DateOfBirth            *bool      `json:"date_of_birth" gorm:"not null;default:true;"`
// 	DateOfBirthRequired    *bool      `json:"date_of_birth_required" gorm:"not null;default:true;"`
// 	HomeAddress            *bool      `json:"home_address" gorm:"not null;default:true;"`
// 	HomeAddressRequired    *bool      `json:"home_address_required" gorm:"not null;default:true;"`
// 	State                  *bool      `json:"state" gorm:"not null;default:true;"`
// 	StateRequired          *bool      `json:"state_required" gorm:"not null;default:true;"`
// 	NextOfKin              *bool      `json:"next_of_kin" gorm:"not null;default:false;"`
// 	NextOfKinRequired      *bool      `json:"next_of_kin_required" gorm:"not null;default:true;"`
// 	NextOfKinPhone         *bool      `json:"next_of_kin_phone" gorm:"not null;default:false;"`
// 	NextOfKinPhoneRequired *bool      `json:"next_of_kin_phone_required" gorm:"not null;default:true;"`
// 	MaritalStatus          *bool      `json:"marital_status" gorm:"not null;default:false;"`
// 	MaritalStatusRequired  *bool      `json:"marital_status_required" gorm:"not null;default:true;"`
// 	Gender                 *bool      `json:"gender" gorm:"not null;default:true;"`
// 	GenderRequired         *bool      `json:"gender_required" gorm:"not null;default:true;"`
// 	Religion               *bool      `json:"religion" gorm:"not null;default:true;"`
// 	ReligionRequired       *bool      `json:"religion_required" gorm:"not null;default:true;"`
// 	// experience
// 	YearsOfExperience          *bool `json:"years_of_experience" gorm:"not null;default:true;"`
// 	YearsOfExperienceRequired  *bool `json:"years_of_experience_required" gorm:"not null;default:true;"`
// 	Degree                     *bool `json:"degree" gorm:"not null;default:true;"`
// 	DegreeRequired             *bool `json:"degree_required" gorm:"not null;default:true;"`
// 	YearOfGraduation           *bool `json:"year_of_graduation" gorm:"not null;default:true;"`
// 	YearOfGraduationRequired   *bool `json:"year_of_graduation_required" gorm:"not null;default:true;"`
// 	PreviousRole               *bool `json:"previous_role" gorm:"not null;default:true;"`
// 	PreviousRoleRequired       *bool `json:"previous_role_required" gorm:"not null;default:true;"`
// 	SchoolOfGraduation         *bool `json:"school_of_graduation" gorm:"not null;default:true;"`
// 	SchoolOfGraduationRequired *bool `json:"school_of_graduation_required" gorm:"not null;default:true;"`
// 	ExpectedSalary             *bool `json:"expected_salary" gorm:"not null;default:true;"`
// 	ExpectedSalaryRequired     *bool `json:"expected_salary_required" gorm:"not null;default:true;"`
// 	OtherQualification         *bool `json:"other_qualification" gorm:"not null;default:true;"`
// 	OtherQualificationRequired *bool `json:"other_qualification_required" gorm:"not null;default:true;"`

// 	//	document Upload
// 	CurriculumVideo                *bool `json:"curriculum_video" gorm:"not null;default:true;"`
// 	CurriculumVideoRequired        *bool `json:"curriculum_video_required" gorm:"not null;default:true;"`
// 	Resume                         *bool `json:"resume" gorm:"not null;default:false;"`
// 	ResumeRequired                 *bool `json:"resume_required" gorm:"not null;default:true;"`
// 	CoverLetter                    *bool `json:"cover_letter" gorm:"not null;default:true;"`
// 	CoverLetterRequired            *bool `json:"cover_letter_required" gorm:"not null;default:true;"`
// 	References                     *bool `json:"references" gorm:"not null;default:false;"`
// 	ReferencesRequired             *bool `json:"references_required" gorm:"not null;default:true;"`
// 	Transcript                     *bool `json:"transcript" gorm:"not null;default:false;"`
// 	TranscriptRequired             *bool `json:"transcript_required" gorm:"not null;default:true;"`
// 	TeachingCertification          *bool `json:"teaching_certification" gorm:"not null;default:false;"`
// 	TeachingCertificationRequired  *bool `json:"teaching_certification_required" gorm:"not null;default:true;"`
// 	LetterOfRecommendation         *bool `json:"letter_of_recommendation" gorm:"not null;default:false;"`
// 	LetterOfRecommendationRequired *bool `json:"letter_of_recommendation_required" gorm:"not null;default:true;"`

// 	// Additional Questions
// 	EmploymentCustomFields *CustomFields        `json:"employment_custom_fields" gorm:"type:jsonb;"`
// 	AdditionalQuestions    *AdditionalQuestions `json:"additional_questions,omitempty" gorm:"type:jsonb"`
// }

// func (employmentConfig *EmploymentConfiguration) Retrieve(ctx context.Context, db *gorm.DB, id uuid.UUID) (EmploymentConfiguration, error) {
// 	// filter on gorm
// 	err := db.WithContext(ctx).Model(&employmentConfig).Where("id = ?", id).First(&employmentConfig).Error
// 	if err != nil {
// 		return *employmentConfig, err
// 	}

// 	return *employmentConfig, err
// }

// // EmploymentApplication /* this is the application for the employment*/
// type EmploymentApplication struct {
// 	/*
// 		This is the application form where teacher input his or her details
// 	*/
// 	BaseModel
// 	ApplicantID                  *string                    `json:"applicant_id"`
// 	Status                       *string                    `json:"status" gorm:"default:IN_PROGRESS"`                                       // IN_PROGRESS / APPROVE / DECLINE
// 	School                       *School                    `json:"school,omitempty" gorm:"constraint:OnDelete:CASCADE;ForeignKey:SchoolID"` // school the employment belongs to
// 	SchoolID                     *uuid.UUID                 `json:"schoolID" gorm:"not null;"`
// 	Employment                   *Employment                `json:"employment,omitempty" gorm:"constraint:OnDelete:CASCADE;foreignKey:EmploymentID;"` // the employment application belongs to
// 	EmploymentID                 *uuid.UUID                 `json:"employment_id" gorm:"not null;"`
// 	FirstName                    *string                    `json:"first_name" gorm:"max=250;"`
// 	LastName                     *string                    `json:"last_name" gorm:"max=250;"`
// 	Email                        *string                    `json:"email" gorm:"max=250;"`
// 	Phone                        *string                    `json:"phone" gorm:"max=250;"`
// 	Location                     *string                    `json:"location" gorm:"max=250;"`
// 	Gender                       *string                    `json:"gender" gorm:"max=250;" validate:"required,uppercase,oneof=MALE FEMALE"`
// 	DateOfBirth                  *time.Time                 `json:"date_of_birth" gorm:"max=250;"`
// 	HomeAddress                  *string                    `json:"home_address" gorm:"max=250;"`
// 	State                        *string                    `json:"state" gorm:"max=250;"`
// 	NextOfKin                    *string                    `json:"next_of_kin" gorm:"max=250;"`
// 	NextOfKinPhone               *string                    `json:"next_of_kin_phone" gorm:"max=250;"`
// 	MaritalStatus                *string                    `json:"marital_status,omitempty" gorm:"max=250;" validate:"oneof=SINGLE MARRIED RELATIONSHIP"`
// 	Religion                     *string                    `json:"religion" gorm:"max=250;"`
// 	YearsOfExperience            *string                    `json:"years_of_experience" gorm:"max=250;"`
// 	Degree                       *string                    `json:"degree" gorm:"max=250;"`
// 	YearOfGraduation             *string                    `json:"year_of_graduation" gorm:"max=250;"`
// 	PreviousRole                 *string                    `json:"previous_role" gorm:"max=250;"`
// 	SchoolOfGraduation           *string                    `json:"school_of_graduation" gorm:"max=250;"`
// 	ExpectedSalary               *string                    `json:"expected_salary" gorm:"max=250;"`
// 	OtherQualification           *string                    `json:"other_qualification" gorm:"max=250;"`
// 	CurriculumVideo              *string                    `json:"curriculum_video" `
// 	Resume                       *string                    `json:"resume" `
// 	CoverLetter                  *string                    `json:"cover_letter" `
// 	References                   *string                    `json:"references" `
// 	Transcript                   *string                    `json:"transcript" `
// 	TeachingCertification        *string                    `json:"teaching_certification" `
// 	LetterOfRecommendation       *string                    `json:"letter_of_recommendation" `
// 	EmploymentCustomFieldAnswers *CustomFieldAnswers        `json:"employment_custom_field_answers" gorm:"type:jsonb;"`
// 	AdditionalQuestionsAnswer    *AdditionalQuestionAnswers `json:"additional_questions_answer" gorm:"type:jsonb;"`
// }

// func (employmentApplication *EmploymentApplication) Retrieve(ctx context.Context, db *gorm.DB, id uuid.UUID) (EmploymentApplication, error) {
// 	// filter on gorm
// 	err := db.WithContext(ctx).Model(&employmentApplication).Where("id = ?", id).First(&employmentApplication).Error
// 	if err != nil {
// 		return *employmentApplication, err
// 	}

// 	return *employmentApplication, err
// }
