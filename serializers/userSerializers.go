package serializers

import (
	"github.com/google/uuid"
	"time"
)

// User /*
// the user serializer
type User struct {
	FirstName   *string    `json:"first_name" `
	LastName    *string    `json:"last_name" `
	Email       *string    `json:"email"  `
	ID          *uuid.UUID `json:"id"`
	PhoneNumber *string    `json:"phone_number" `
	LastLogin   *time.Time `json:"last_login"`
}

type UserPermissionUpdateRequestSerializer struct {
	Email       *string `json:"email" validate:"required"`
	IsStaff     *bool   `json:"is_staff"`
	IsSuperUser *bool   `json:"is_super_user"`
}

// UserAndProfileUpdateRequestSerializer /* This is used for getting the request from the frontend and also updating the model struct*/
type UserAndProfileUpdateRequestSerializer struct {
	FirstName    *string    `form:"first_name" `
	LastName     *string    `form:"last_name" `
	Email        *string    `form:"email"  `
	IsVerified   *bool      `form:"is_verified" `
	PhoneNumber  *string    `form:"phone_number" validate:"omitempty,max=15,min=5" `
	ProfileImage *string    `form:"profile_image" `
	Gender       *string    `form:"gender"`
	Address      *string    `form:"address"`
	Country      *string    `form:"country"`
	BirthDay     *time.Time `form:"birth_day"`
	PostalCode   *string    `form:"postal_code"`
	TwoFactor    *bool      `form:"two_factor"`
}

// HealthUpdateRequestSerializer /* This is used to update the health info of the user */
type HealthUpdateRequestSerializer struct {
	BloodGroup      *string `json:"blood_group"  validate:"max=250"`
	Genotype        *string `json:"genotype"  validate:"max=250"`
	Allergy         *string `json:"allergy"  validate:"max=250"`
	Visual          *string `json:"visual"  validate:"max=250"`
	BloodPressure   *string `json:"blood_pressure"  validate:"max=250"`
	Disability      *string `json:"disability"  validate:"max=250"`
	ParentWord      *string `json:"parent_word"  validate:"max=250"`
	State           *string `json:"state"  validate:"max=250"`
	Ulcer           *bool   `json:"ulcer" `
	Covid19         *bool   `json:"covid19"  `
	Addiction       *string `json:"addiction"  validate:"max=250"`
	SugarLevel      *string `json:"sugar_level"  validate:"max=250"`
	ParentsPhone    *string `json:"parents_phone"  validate:"max=15"`
	DoctorsName     *string `json:"doctors_name"  validate:"max=50"`
	DoctorsHospital *string `json:"doctors_hospital"  validate:"max=50"`
	DoctorsPhone    *string `json:"doctors_phone"  validate:"max=50"`
}

// UserInfoUpdateRequestSerializer All user Info Update Serializer
type UserInfoUpdateRequestSerializer struct {
	FirstName       *string    `json:"first_name" ` // User part
	LastName        *string    `json:"last_name" `
	Email           *string    `json:"email"  `
	IsVerified      *bool      `json:"is_verified" `
	PhoneNumber     *string    `json:"phone_number" validate:"omitempty,max=15,min=5" ` // profile part
	ProfileImage    *string    `json:"profile_image" `
	Gender          *string    `json:"gender"`
	Address         *string    `json:"address"`
	Country         *string    `json:"country"`
	BirthDay        *time.Time `json:"birth_day"`
	PostalCode      *string    `json:"postal_code"`
	TwoFactor       *bool      `json:"two_factor"`
	BloodGroup      *string    `json:"blood_group"  validate:"max=250"` // health part
	Genotype        *string    `json:"genotype"  validate:"max=250"`
	Allergy         *string    `json:"allergy"  validate:"max=250"`
	Visual          *string    `json:"visual"  validate:"max=250"`
	BloodPressure   *string    `json:"blood_pressure"  validate:"max=250"`
	Disability      *string    `json:"disability"  validate:"max=250"`
	ParentWord      *string    `json:"parent_word"  validate:"max=250"`
	State           *string    `json:"state"  validate:"max=250"`
	Ulcer           *bool      `json:"ulcer" `
	Covid19         *bool      `json:"covid19"  `
	Addiction       *string    `json:"addiction"  validate:"max=250"`
	SugarLevel      *string    `json:"sugar_level"  validate:"max=250"`
	ParentsPhone    *string    `json:"parents_phone"  validate:"max=15"`
	DoctorsName     *string    `json:"doctors_name"  validate:"max=50"`
	DoctorsHospital *string    `json:"doctors_hospital"  validate:"max=50"`
	DoctorsPhone    *string    `json:"doctors_phone"  validate:"max=50"`
	StaffCode       *string    `json:"staff_code"`
	StudentCode     *string    `json:"student_code"`
}
