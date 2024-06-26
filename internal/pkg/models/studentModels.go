package models

// import (
// 	"context"
// 	"github.com/google/uuid"
// 	"gorm.io/gorm"
// 	"time"
// )

// type Student struct {
// 	BaseModel
// 	School                *School              `json:"school,omitempty" gorm:"constraint:OnDelete:CASCADE;ForeignKey:SchoolID"`
// 	SchoolID              *uuid.UUID           `json:"school_id" gorm:"not null;"`
// 	StudentCode           *string              `json:"student_code" gorm:"not null"`
// 	User                  *User                `json:"user,omitempty" gorm:"constraint:OnDelete:CASCADE;foreignKey:UserID;"` //	the one-to-one relationship
// 	UserID                *uuid.UUID           `json:"user_id" gorm:"not null;"`
// 	Class                 *Class               `json:"class,omitempty" gorm:"constraint:OnDelete:SET NULL;foreignKey:ClassID;"` //	the one-to-one relationship
// 	ClassID               *uuid.UUID           `json:"class_id" `
// 	EnrollmentAdmission   *EnrollmentAdmission `json:"enrollment_admission,omitempty" gorm:"constraint:OnDelete:SET NULL;foreignKey:EnrollmentAdmissionID;"`
// 	EnrollmentAdmissionID *uuid.UUID           `json:"enrollment_admission_id" `
// 	Department            *string              `json:"department" gorm:"max=500"`
// 	Section               *string              `json:"section" gorm:"max=500"`
// 	Status                *string              `json:"status" gorm:"max=500"`
// 	Gender                *string              `json:"gender"`
// }

// func (student *Student) RetrieveStudentByUserIDAndSchool(db *gorm.DB, ctx context.Context, userID, schoolID uuid.UUID) (Student, error) {

// 	err := db.WithContext(ctx).Model(&Student{}).
// 		Where("user_id = ?", userID).
// 		Where("school_id = ?", schoolID).
// 		First(&student).Error
// 	if err != nil {
// 		return *student, err
// 	}
// 	return *student, nil
// }

// func (student *Student) RetrieveStudentByIDAndSchool(ctx context.Context, db *gorm.DB, studentID, schoolID uuid.UUID) (Student, error) {

// 	err := db.WithContext(ctx).Model(&Student{}).
// 		Where("id = ?", studentID).
// 		Where("school_id = ?", schoolID).
// 		First(&student).Error
// 	if err != nil {
// 		return *student, err
// 	}
// 	return *student, nil
// }

// // ClinicVisitation /* This contains the history of how the user has been visiting the clinic*/
// type ClinicVisitation struct {
// 	// this contains custom functions that other models could use and fields like the id and timestamp
// 	BaseModel
// 	Student        *Student   `json:"student,omitempty" gorm:"constraint:OnDelete:CASCADE;ForeignKey:StudentID"`
// 	StudentID      *uuid.UUID `json:"student_id" gorm:"not null;"`
// 	School         *School    `json:"school,omitempty"  gorm:"constraint:OnDelete:CASCADE;ForeignKey:SchoolID"` // school the course belongs to
// 	SchoolID       *uuid.UUID `json:"school_id" gorm:"not null;"`
// 	Name           *string    `json:"name" gorm:"max=50" `          // name of the clinic
// 	Description    *string    `json:"description" gorm:"max=250"`   // what the user went to do in the clinic
// 	DoctorsName    *string    `json:"doctors_name"  gorm:"max=50" ` // The name of the doctor
// 	VisitationTime *time.Time `json:"visitation_time" gorm:"default:CURRENT_TIMESTAMP;"`
// 	Status         *string    `json:"status" ` // the status of the visit if its for something very important
// }

// type Note struct {
// 	BaseModel
// 	School    *School    `json:"school,omitempty" gorm:"constraint:OnDelete:CASCADE;ForeignKey:SchoolID"`
// 	SchoolID  *uuid.UUID `json:"school_id" gorm:"not null;"`
// 	Student   *Student   `json:"student,omitempty" gorm:"constraint:OnDelete:CASCADE;ForeignKey:StudentID"`
// 	StudentID *uuid.UUID `json:"student_id" gorm:"not null;"`
// 	Staff     *Staff     `json:"staff,omitempty" gorm:"constraint:OnDelete:CASCADE;ForeignKey:StaffID"`
// 	StaffID   *uuid.UUID `json:"staff_id" gorm:"not null;"`
// 	NoteType  *string    `json:"note_type"` //  ACADEMIC , BEHAVIOR , EMOTION , OTHERS
// 	Share     *string    `json:"share"`     // DONT_SHARE, VISIBLE_TO_TEACHERS, VISIBLE_TO_STUDENTS, WITH_PARENTS_ONLY , ADMIN
// 	Title     *string    `json:"title"`
// 	Note      *string    `json:"note"`
// }
