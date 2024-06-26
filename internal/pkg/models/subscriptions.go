package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type NewsletterSubscription struct {
	ID         string `json:"id" gorm:"primaryKey;type:uuid;"`
	Email      string `gorm:"unique;not null"`
	Subscribed bool   `gorm:"default:true"`
	FirstName  *string
	LastName   *string
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

func (n *NewsletterSubscription) BeforeCreate(tx *gorm.DB) (err error) {
	// Generate a new UUID for the ID field
	if n.ID == "" {
		n.ID = uuid.New().String()
	}
	return nil
}
