package serializers

import (
	"time"

	"github.com/google/uuid"
)

type SchoolTypeCreateRequestSerializer struct {
	Name  *string `form:"name" json:"name" validate:"required,max=250"`
	Image *string `form:"image" json:"image" validate:"required,max=250"`
}

type SchoolTypeUpdateRequestSerializer struct {
	Name  *string `form:"name" `
	Image *string `form:"image" `
}

type SchoolInviteCreateRequestSerializer struct {
	// this contains custom functions that other models could use and fields like the id and timestamp
	UserType *string `json:"user_type,omitempty" validate:"omitempty,eq=ADMINISTRATOR|eq=TEACHING_STAFF|eq=NON_TEACHING_STAFF"`
	Email    *string `json:"email"`
}

type SchoolCreateRequestSerializer struct {
	/* This is used in creating the school*/
	// SchoolTypeID *uuid.UUID `form:"school_type_id" validate:"required"`
	Name *string `form:"name" validate:"required,max=250"`
	// Address      *string    `form:"address" validate:"required,max=250"`
	// Mail         *string    `form:"mail" validate:"required,max=250"`
	// Phone        *string    `form:"phone" validate:"required,max=250"`
	// BrandColor   *string    `form:"brand_color" `
	Logo     *string `form:"logo"`
	Document *string `form:"document"`
}

type SchoolUpdateRequestSerializer struct {
	/* This is used to update the school*/
	SchoolTypeID *uuid.UUID `form:"school_type_id" `
	Name         *string    `form:"name" `
	Address      *string    `form:"address" `
	Mail         *string    `form:"mail" `
	Phone        *string    `form:"phone" `
	BrandColor   *string    `form:"brand_color" `
	Logo         *string    `form:"logo"`
	Document     *string    `form:"document"`
}

type ClassCreateUpdateSerializer struct {
	/* this is used to create class */
	Name    *string `json:"name" gorm:"not null;"`
	Faculty *string `json:"faculty"`
}

type SchoolAdminCreateRequestSerializer struct {
	UserID *uuid.UUID `json:"user_id"`
}

type SchoolAdminUpdateRequestSerializer struct {
	Status *string `json:"status"  validate:"required,eq=ACTIVE|eq=INACTIVE"`
}

type AnnouncementConfigurationRequestSerializer struct {
	ToRecipientPortal *bool   `json:"to_recipient_portal" `
	ToEmail           *bool   `json:"to_email" `
	ReceiveCopy       *bool   `json:"receive_copy"`
	SendFromOrder     *string `json:"send_from_order"`
	AnnouncementTag   *string `json:"announcement_tag"`
	NotificationType  *string `json:"notification_type"`
	Layout            *string `json:"layout"`
}
type NotificationConfigurationRequestSerializer struct {
	ToRecipientPortal *bool   `json:"to_recipient_portal" `
	ToEmail           *bool   `json:"to_email" `
	ReceiveCopy       *bool   `json:"receive_copy"`
	SendFromOrder     *string `json:"send_from_order"`
	AnnouncementTag   *string `json:"announcement_tag"`
	NotificationType  *string `json:"notification_type"`
	Layout            *string `json:"layout"`
	AttendanceStatus  *string `json:"attendance_status"`
	Description       *string `json:"description"`
}

type EventCreateRequestSerializer struct {
	Title              *string      `json:"title" `
	Description        *string      `json:"description"`
	Time               *time.Time   `json:"time" gorm:"not null;"`       // the time the event is to take place
	EventType          *string      `json:"event_type" gorm:"not null;"` // the type of event
	EventsParticipants []*uuid.UUID `json:"events_participants" gorm:"many2many:event_participants;"`
}
