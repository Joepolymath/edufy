package models

import (
	"github.com/google/uuid"
	"time"
)

// Schedule /* The is the alert time which is sent to students' base on a new lecture availability */
type Schedule struct {
	BaseModel
	School   *School    `gorm:"ForeignKey:SchoolID"`        // the school schedule belongs to
	SchoolID *uuid.UUID `json:"school_id" gorm:"not null;"` // the school schedule belongs to
	Time     *time.Time `json:"time" gorm:"not null;"`      // time of the schedule
}

type Lecture struct {
}
