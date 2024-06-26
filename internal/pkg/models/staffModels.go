package models

// import (
// 	"context"

// 	"github.com/google/uuid"
// 	"gorm.io/gorm"
// )

// type Staff struct {
// 	BaseModel
// 	School                  *School                `json:"school,omitempty" gorm:"constraint:OnDelete:CASCADE;ForeignKey:SchoolID"`
// 	SchoolID                *uuid.UUID             `json:"school_id" gorm:"not null;"`
// 	StaffCode               *string                `json:"staff_code" gorm:"not null"`
// 	User                    *User                  `json:"user,omitempty" gorm:"constraint:OnDelete:CASCADE;foreignKey:UserID;"` //	the one-to-one relationship
// 	UserID                  *uuid.UUID             `json:"user_id" gorm:"not null;"`
// 	EmploymentApplication   *EmploymentApplication `json:"employment_application,omitempty" gorm:"constraint:OnDelete:SET NULL;foreignKey:EmploymentApplicationID;"`
// 	EmploymentApplicationID *uuid.UUID             `json:"employment_application_id" ` // the application id
// 	Classes                 []*Class               `json:"classes,omitempty" gorm:"many2many:staff_classes;"`
// 	// Role                    *Role                  `json:"role,omitempty" gorm:"constraint:OnDelete:SET NULL;foreignKey:RoleID;"`
// 	RoleID        *uuid.UUID `json:"role_id" `
// 	Category      *Category  `json:"category,omitempty" gorm:"constraint:OnDelete:SET NULL;foreignKey:CategoryID;"`
// 	CategoryID    *uuid.UUID `json:"category_id" `
// 	Qualification *string    `json:"qualification" gorm:"max=250"`
// 	NetSalary     *float64   `json:"net_salary" ` // this would be updated anytime the salary is updated
// 	StaffType     *string    `json:"staff_type" gorm:"default:FULL_TIME" validate:"required, oneof=FULL_TIME PART_TIME"`
// 	Status        *string    `json:"status" gorm:"max=250" validate:"required, oneoff=ACTIVE INACTIVE"`
// }

// type StaffClass struct {
// 	BaseModel
// 	StaffID uuid.UUID `json:"staff_id" gorm:"type:uuid;primaryKey"`
// 	ClassID uuid.UUID `json:"class_id" gorm:"type:uuid;primaryKey"`
// }

// func (staff *Staff) Retrieve(ctx context.Context, db *gorm.DB, id uuid.UUID) (Staff, error) {
// 	// filter on gorm
// 	err := db.WithContext(ctx).Model(&staff).Where("id = ?", id).First(&staff).Error
// 	if err != nil {
// 		return *staff, err
// 	}

// 	return *staff, err
// }
// func (staff *Staff) RetrieveByIDAndSchool(ctx context.Context, db *gorm.DB, id uuid.UUID, schoolID uuid.UUID) (Staff, error) {
// 	// filter on gorm
// 	err := db.WithContext(ctx).Model(&staff).Where("id = ?", id).Where("school_id = ?", schoolID).First(&staff).Error
// 	if err != nil {
// 		return *staff, err
// 	}

// 	return *staff, err
// }

// func (staff *Staff) RetrieveByUserAndSchool(ctx context.Context, db *gorm.DB, schoolID uuid.UUID, userID uuid.UUID) (Staff, error) {
// 	// filter on gorm
// 	err := db.WithContext(ctx).Model(&staff).
// 		Where("school_id = ?", schoolID).
// 		Where("user_id = ?", userID).
// 		Where("status = ?", "ACTIVE").
// 		First(&staff).Error
// 	if err != nil {
// 		return *staff, err
// 	}

// 	return *staff, err
// }

// func (staff *Staff) CheckUserStaff(ctx context.Context, db *gorm.DB, user_id, school_id uuid.UUID) (Staff, error) {
// 	// filter on gorm
// 	err := db.WithContext(ctx).Model(&staff).
// 		Where("school_id = ?", school_id).
// 		Where("user_id = ?", user_id).First(&staff).Error
// 	if err != nil {
// 		return *staff, err
// 	}

// 	return *staff, err
// }

// type StaffRating struct {
// 	BaseModel
// 	School   *School    `json:"school,omitempty" gorm:"constraint:OnDelete:CASCADE;ForeignKey:SchoolID"`
// 	SchoolID *uuid.UUID `json:"school_id" gorm:"not null;"`
// 	Staff    *Staff     `json:"staff,omitempty" gorm:"constraint:OnDelete:CASCADE;ForeignKey:StaffID"`
// 	StaffID  *uuid.UUID `json:"staff_id" gorm:"not null;"`
// 	User     *User      `json:"user,omitempty" gorm:"constraint:OnDelete:CASCADE;ForeignKey:UserID"`
// 	UserID   *uuid.UUID `json:"user_id" `
// 	Message  *string    `json:"message"`
// 	Rating   *int       `json:"rating"`
// }
