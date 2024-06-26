package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type (
	RoleType string // role category for users
)

// users are account types with definite association with schools
type User struct {
	ID             string         `json:"id" gorm:"primaryKey;type:uuid;"`
	FirstName      *string        `json:"first_name" validate:"required,max=250,min=2" gorm:"size:250;"`
	LastName       *string        `json:"last_name" validate:"required,max=250,min=2" gorm:"size:250;"`
	Email          string         `json:"email" validate:"max=100,min=3"  gorm:"unique;"`
	EmailVerified  bool           `json:"email_verified" gorm:"default:false"`
	PhoneNumber    *string        `json:"phone_number" validate:"required,max=50,min=5" gorm:"size:50"`
	PhoneVerified  bool           `json:"phone_verified" gorm:"default:false"`
	IsSuperAdmin   bool           `json:"is_super_admin" gorm:"default:false"`
	IsSchoolMember bool           `json:"is_school_member" gorm:"default:false"`
	IsStudent      bool           `json:"is_student" gorm:"default:false"`
	IsParent       bool           `json:"is_parent" gorm:"default:false"`
	LastLogin      *time.Time     `json:"last_login"`
	Password       *string        `json:"-" validate:"required,max=250,min=5"`
	Avatar         *string        `json:"avatar"`
	Gender         *string        `json:"gender"`
	Address        *string        `json:"address"`
	Country        *string        `json:"country"`
	BirthDay       *time.Time     `json:"birth_day"`
	PostalCode     *string        `json:"postal_code"`
	TwoFactor      bool           `json:"two_factor" gorm:"default:false;type:bool;"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	// Generate a new UUID for the ID field
	if u.ID == "" {
		u.ID = uuid.New().String()
	}
	return nil
}

func (u *User) ExportDetails() User {
	u.Password = nil
	return *u
}

const (
	// Admin role for Learnium OS
	ADMIN RoleType = "School Admin"
	// Non Teaching Staff role for Learnium OS
	NON_TEACHING_STAFF RoleType = "Non Teaching Staff"
	// Teaching Staff role for Learnium OS
	TEACHING_STAFF RoleType = "Teaching Staff"
)
