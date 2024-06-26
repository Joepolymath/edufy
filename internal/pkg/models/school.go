package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

var (
	HRDesc = `Manages all aspects of human resources, including recruitment, onboarding, staff development, and employee relations.
 Ensures compliance with employment laws, maintains personnel records, and supports overall staff well-being`

	AODesc = ` Manages the student admission process, including application reviews, interviews, and enrollment. 
Coordinates with prospective students and their families.`

	CounselDesc = `Provides counseling services to students. Addresses academic, personal, and emotional concerns, 
and collaborates with teachers and parents to support student well-being.`

	HWCDesc = `Promotes the health and well-being of students and staff. Coordinates health programs, organizes wellness activities, and manages health-related records.`

	ACCDesc = `Manages financial transactions, budgeting, and payroll for the school. Responsible for keeping accurate financial records.`

	ECDesc = `Manages the planning and execution of examinations. Coordinates with teachers, ensures proper exam administration, and oversees the grading process.`

	CDevDesc = `Designs and develops the school curriculum, ensuring alignment with educational standards and goals. Collaborates with other educators to enhance teaching materials.`

	CTDesc = ` Manages a specific class or group of students. Acts as the primary point of contact for parents and is responsible for monitoring the overall well-being and academic progress of the students.`

	STDesc = `Teaches a specific subject or multiple subjects. Responsible for creating lesson plans, assessing student performance, and providing feedback.`

	HODDesc = `Leads a specific academic department, overseeing curriculum, lesson planning, and the performance of teachers within the department.`

	SchoolAdminDesc = ` Manages the overall administration of the school, including student enrollment, staff management, and academic settings.
 Has the authority to configure system settings related to the school's policies.`
)

type SchoolType string

const (
	K12        SchoolType = "K-12"
	RELIGIOUS  SchoolType = "Faith-Based"
	TERTIARY   SchoolType = "Tertiary"
	SUMMER     SchoolType = "Summer"
	GOV        SchoolType = "Government"
	VOCATIONAL SchoolType = "Vocational"
)

type SchoolMemberProfile struct {
	ID        string      `json:"id" gorm:"primaryKey;type:uuid;"`
	MemberId  string      `json:"owner_id" gorm:"not null;"`
	Member    *User       `json:"owner,omitempty" gorm:"constraint:OnDelete:CASCADE;foreignKey:MemberId;"`
	SchoolID  *string     `json:"school_id,omitempty" gorm:"type:uuid;"`
	School    *School     `json:"school,omitempty" gorm:"constraint:OnDelete:CASCADE;foreignKey:SchoolID;"`
	Avatar    *string     `json:"avatar"`
	RoleID    *string     `json:"role_id"`
	FirstName *string     `json:"first_name" validate:"required,max=250,min=2" gorm:"size:250;"`
	LastName  *string     `json:"last_name" validate:"required,max=250,min=2" gorm:"size:250;"`
	Role      *SchoolRole `json:"role"`
	WorkEmail *string     `json:"work_email"`
}

// Role // this is the role of the user associated with a particular school branch
type SchoolRole struct {
	ID          string  `json:"id" gorm:"primaryKey;type:uuid;"`
	SchoolID    *string `json:"school_id,omitempty" gorm:"type:uuid;"`
	School      *School
	Permissions []SchoolPermission `gorm:"many2many:role_permissions;"`
	Name        string             `json:"name" gorm:"not null;"`
	RoleType    RoleType           `json:"role"`
	CreatedAt   time.Time          `json:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at"`
	DeletedAt   gorm.DeletedAt     `json:"deleted_at" gorm:"index"`
	Description *string            `json:"description" gorm:"type:text"`
}

type SchoolPermission struct {
	ID        string  `json:"id" gorm:"primaryKey;type:uuid;"`
	SchoolID  *string `json:"school_id,omitempty" gorm:"type:uuid;"`
	School    *School
	Name      string
	Roles     []SchoolRole `gorm:"many2many:role_permissions;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// type SchoolType struct {
// 	// this contains custom functions that other models could use and fields like the id and timestamp
// 	ID   string  `json:"id" gorm:"primaryKey;type:uuid;"`
// 	Type *string `json:"name"`
// 	// Image *string `json:"image"`
// }

// School /* This contains the info when creating a new school. the name*/
type School struct {
	ID             string     `json:"id" gorm:"primaryKey;type:uuid;"`
	ParentSchoolID *string    `json:"parent_school_id,omitempty" gorm:"type:uuid;"`
	ParentSchool   *School    `json:"parent_school,omitempty" gorm:"foreignKey:ParentSchoolID;constraint:OnDelete:SET NULL"`
	OwnerID        string     `json:"owner_id" gorm:"not null;"`
	Owner          *User      `json:"owner,omitempty" gorm:"constraint:OnDelete:CASCADE;foreignKey:OwnerID;"` //	the one-to-one relationship
	Type           SchoolType `json:"school_type,omitempty"`
	IsParentSchool bool       `json:"is_parent_school"`
	IsBranchSchool bool       `json:"is_branch_school"`
	// SchoolCode           *string `json:"school_code" gorm:"not null;uniqueIndex;size:250;"`
	Name                 string  `json:"name" gorm:"not null;size:250;"`
	Address              *string `json:"address" gorm:"size:250;"`
	Mail                 string  `json:"mail" gorm:"size:250;"`
	Phone                string  `json:"phone" gorm:"size:250"`
	BrandColor           *string `json:"brand_color" gorm:"size:20"`
	Logo                 *string `json:"logo"`
	RegistrationDocument *string `json:"registrationDocument"`
	CreatedAt            time.Time
	UpdatedAt            time.Time
	DeletedAt            gorm.DeletedAt `gorm:"index"`
}

type Session struct {
	ID                   string  `json:"id" gorm:"primaryKey;type:uuid;"`
	SchoolID             *string `json:"school_id,omitempty" gorm:"type:uuid;"`
	School               *School `json:"school,omitempty" gorm:"foreignKey:ParentSchoolID;constraint:OnDelete:SET NULL"`
	Type                 *string `json:"session_type,omitempty"`
	Name                 string  `json:"name" gorm:"not null;size:250;"`
	StartDate            *time.Time
	EndDate              *time.Time
	RegistrationDocument *string `json:"registrationDocument"`
	CreatedAt            time.Time
	UpdatedAt            time.Time
	DeletedAt            gorm.DeletedAt `gorm:"index"`
}

func (s *Session) BeforeCreate(tx *gorm.DB) (err error) {
	// Generate a new UUID for the ID field
	if s.ID == "" {
		s.ID = uuid.New().String()
	}
	return nil
}

func (s *School) BeforeCreate(tx *gorm.DB) (err error) {
	// Generate a new UUID for the ID field
	if s.ID == "" {
		s.ID = uuid.New().String()
	}
	return nil
}

func (s *SchoolPermission) BeforeCreate(tx *gorm.DB) (err error) {
	// Generate a new UUID for the ID field
	if s.ID == "" {
		s.ID = uuid.New().String()
	}
	return nil
}

func (s *SchoolRole) BeforeCreate(tx *gorm.DB) (err error) {
	// Generate a new UUID for the ID field
	if s.ID == "" {
		s.ID = uuid.New().String()
	}
	return nil
}

// func (school *School) Retrieve(ctx context.Context, db *gorm.DB, id uuid.UUID) (School, error) {
// 	// filter on gorm
// 	err := db.WithContext(ctx).Model(&school).Where("id = ?", id).First(&school).Error
// 	if err != nil {
// 		return *school, err
// 	}

// 	return *school, err
// }

// func (school *School) RetrieveBySchoolCode(ctx context.Context, db *gorm.DB, schoolCode string) (School, error) {

// 	if err := db.WithContext(ctx).Model(&school).Where("school_code = ?", schoolCode).First(&school).Error; err != nil {
// 		// Handle the error, e.g., return false or log the error
// 		return *school, errors.New("school with this code does not exists")
// 	}
// 	return *school, nil
// }

// func (school *School) IsSchoolAdminOrOwner(ctx context.Context, db *gorm.DB, schoolCode string, currentUserID uuid.UUID) bool {
// 	// Retrieve the school by school code
// 	if err := db.WithContext(ctx).Model(&school).Where("school_code = ?", schoolCode).First(&school).Error; err != nil {
// 		// Handle the error, e.g., return false or log the error
// 		return false
// 	}
// 	// Check if the user is the owner of the school
// 	if school.OwnerID != nil && *school.OwnerID == currentUserID {
// 		return true
// 	}

// 	// Check if the user is an admin of the school
// 	var adminCount int64
// 	if err := db.WithContext(ctx).Model(&Admin{}).Where("school_id = ? AND user_id = ? and status =?", school.ID, currentUserID, "ACTIVE").Count(&adminCount).Error; err != nil {
// 		// Handle the error, e.g., return false or log the error
// 		return false
// 	}
// 	return adminCount > 0
// }

// func (school *School) IsSchoolAdminOrOwnerOrStaff(ctx context.Context, db *gorm.DB, schoolCode string, currentUserID uuid.UUID) bool {
// 	var staff Staff
// 	var adminCount int64
// 	// Retrieve the school by school code
// 	if err := db.WithContext(ctx).Model(&school).Where("school_code = ?", schoolCode).First(&school).Error; err != nil {
// 		// Handle the error, e.g., return false or log the error
// 		return false
// 	}
// 	// Check if the user is the owner of the school
// 	if school.OwnerID != nil && *school.OwnerID == currentUserID {
// 		return true
// 	}

// 	if err := db.WithContext(ctx).Model(&Staff{}).Where("school_id = ?", school.ID).Where("user_id =?", currentUserID).First(&staff); err != nil {
// 		// Check if the user is an admin of the school
// 		if err := db.WithContext(ctx).Model(&Admin{}).Where("school_id = ? AND user_id = ? and status =?", school.ID, currentUserID, "ACTIVE").Count(&adminCount).Error; err != nil {
// 			// Handle the error, e.g., return false or log the error
// 			return false
// 		}
// 	}

// 	return adminCount > 0
// }

// type SchoolInvite struct {
// 	// this contains custom functions that other models could use and fields like the id and timestamp
// 	BaseModel
// 	School   *School    `json:"school,omitempty" gorm:"constraint:OnDelete:CASCADE;ForeignKey:SchoolID"`
// 	SchoolID *uuid.UUID `json:"school_id" gorm:"not null;"`
// 	UserType *string    `json:"user_type" gorm:"type:varchar(20);check:user_type IN ('ADMINISTRATOR', 'TEACHING_STAFF','NON_TEACHING_STAFF');default:NULL"`
// 	Status   *string    `json:"status" gorm:"type:varchar(20);check:status IN ('ACCEPTED', 'PENDING')"`
// 	Email    *string    `json:"email"`
// }

// // Admin these admins enables us to have multiple users to be able to access the application dashboard for maintenance
// type Admin struct {
// 	BaseModel
// 	School   *School    `json:"school,omitempty" gorm:"constraint:OnDelete:CASCADE;ForeignKey:SchoolID"`
// 	SchoolID *uuid.UUID `json:"school_id" gorm:"not null;"`
// 	User     *User      `gorm:"constraint:OnDelete:CASCADE;ForeignKey:UserID"`
// 	UserID   *uuid.UUID `json:"user_id" gorm:"not null;"`
// 	Status   *string    `json:"status" gorm:"type:varchar(20);check:status IN ('ACTIVE', 'INACTIVE')"`
// }

// type Class struct {
// 	// this contains custom functions that other models could use and fields like the id and timestamp
// 	BaseModel
// 	School        *School    `json:"school,omitempty" gorm:"constraint:OnDelete:CASCADE;ForeignKey:SchoolID"`
// 	SchoolID      *uuid.UUID `json:"school_id" gorm:"not null;"`
// 	Name          *string    `json:"name" gorm:"not null;"`
// 	Population    *int       `json:"population"`
// 	Subjects      *string    `json:"subjects"`
// 	MaleAndFemale *string    `json:"male_and_female"`
// 	Faculty       *string    `json:"faculty"`
// }

// func (class *Class) Retrieve(ctx context.Context, db *gorm.DB, id uuid.UUID) (Class, error) {
// 	// filter on gorm
// 	err := db.WithContext(ctx).Model(&class).Where("id = ?", id).First(&class).Error
// 	if err != nil {
// 		return *class, err
// 	}

// 	return *class, err
// }

// // Category // this is the category of the user in the school
// type Category struct {
// 	BaseModel
// 	School   *School    `json:"school,omitempty" gorm:"constraint:OnDelete:CASCADE;ForeignKey:SchoolID"`
// 	SchoolID *uuid.UUID `json:"school_id" gorm:"not null;"`
// 	Name     *string    `json:"name" gorm:"not null;"`
// }

// // Event /* This enables creating an event to the parent or teachers about an upcoming meeting*/
// type Event struct {
// 	// this contains custom functions that other models could use and fields like the id and timestamp
// 	BaseModel
// 	School             *School    `gorm:"constraint:OnDelete:CASCADE;ForeignKey:SchoolID"`
// 	SchoolID           *uuid.UUID `json:"school_id" gorm:"not null;"`
// 	Time               *time.Time `json:"time" gorm:"not null;"` // the time the event is to take place
// 	Title              *string    `json:"title" `
// 	Description        *string    `json:"description"`
// 	EventType          *string    `json:"event_type" gorm:"not null;"` // the type of event
// 	EventsParticipants []*Staff   `json:"events_participants" gorm:"many2many:event_participants;"`
// }

// type EventParticipant struct {
// 	EventID *uuid.UUID `json:"event_id" gorm:"type:uuid;primaryKey"`
// 	StaffID *uuid.UUID `json:"staff_id" gorm:"type:uuid;primaryKey"`
// }

// // TableName Add this function to define the composite primary key
// func (EventParticipant) TableName() string {
// 	return "event_participants"
// }

// type NotificationConfiguration struct {
// 	BaseModel
// 	School            *School    `json:"school,omitempty" gorm:"constraint:OnDelete:CASCADE;ForeignKey:SchoolID"`
// 	SchoolID          *uuid.UUID `json:"school_id" gorm:"not null;"`
// 	ToRecipientPortal *bool      `json:"to_recipient_portal" `
// 	ToEmail           *bool      `json:"to_email" `
// 	ReceiveCopy       *bool      `json:"receive_copy"`
// 	SendFromOrder     *string    `json:"send_from_order"`
// 	AnnouncementTag   *string    `json:"announcement_tag"`
// 	NotificationType  *string    `json:"notification_type"`
// 	Layout            *string    `json:"layout"`
// 	AttendanceStatus  *string    `json:"attendance_status"`
// 	Description       *string    `json:"description"`
// }

// type AnnouncementConfiguration struct {
// 	BaseModel
// 	School            *School    `json:"school,omitempty" gorm:"constraint:OnDelete:CASCADE;ForeignKey:SchoolID"`
// 	SchoolID          *uuid.UUID `json:"school_id" gorm:"not null;"`
// 	ToRecipientPortal *bool      `json:"to_recipient_portal" `
// 	ToEmail           *bool      `json:"to_email" `
// 	ReceiveCopy       *bool      `json:"receive_copy"`
// 	SendFromOrder     *string    `json:"send_from_order"`
// 	AnnouncementTag   *string    `json:"announcement_tag"`
// 	NotificationType  *string    `json:"notification_type"`
// 	Layout            *string    `json:"layout"`
// }
