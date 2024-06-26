package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

// BaseModel / This is actually used to create most used fields like timestamp, uuid and do some custom process **/
type BaseModel struct {
	ID        uuid.UUID  `json:"id" gorm:"primaryKey;type:uuid;"`
	Timestamp *time.Time `json:"timestamp" gorm:"default:CURRENT_TIMESTAMP;"`
}

// BeforeCreate setting the uuid of the value
func (m *BaseModel) BeforeCreate(db *gorm.DB) (err error) {
	m.ID = uuid.New()
	return
}

// AdditionalQuestion /* This is used to ask additional question */
type AdditionalQuestion struct {
	Question   string `json:"question"  validate:"required"`
	IsRequired bool   `json:"isRequired"  `
}

type AdditionalQuestions []AdditionalQuestion

// Scan implements the sql.Scanner interface
func (aq *AdditionalQuestions) Scan(src interface{}) error {
	switch src := src.(type) {
	case []byte:
		return json.Unmarshal(src, aq)
	case string:
		return json.Unmarshal([]byte(src), aq)
	default:
		return errors.New("incompatible type for AdditionalQuestions")
	}
}

// Value implements the driver.Valuer interface
func (aq AdditionalQuestions) Value() (driver.Value, error) {
	return json.Marshal(aq)
}

// AdditionalQuestionAnswer /* This is used for answers*/
type AdditionalQuestionAnswer struct {
	Question   string `json:"question"  validate:"required"`
	IsRequired bool   `json:"isRequired"  validate:"required"`
	Answer     string `json:"answer"  validate:"required"`
}

type AdditionalQuestionAnswers []AdditionalQuestionAnswer

// Scan implements the sql.Scanner interface
func (aq *AdditionalQuestionAnswers) Scan(src interface{}) error {
	switch src := src.(type) {
	case []byte:
		return json.Unmarshal(src, aq)
	case string:
		return json.Unmarshal([]byte(src), aq)
	default:
		return errors.New("incompatible type for AdditionalQuestionAnswers")
	}
}

// Value implements the driver.Valuer interface
func (aq AdditionalQuestionAnswers) Value() (driver.Value, error) {
	return json.Marshal(aq)
}

// CustomField /* This is used for answers*/
type CustomField struct {
	FieldName  string `json:"field_name" validate:"required"`
	IsRequired bool   `json:"isRequired"`
	FieldPath  string `json:"field_path"  validate:"required,uppercase,oneof=EXPERIENCE DOCUMENT_UPLOAD CONTACT_INFO STUDENT_INFO PARENT_INFO"` // this is where the field could be found in the employment object (ContactInfo, Experience, etc)
}

type CustomFields []CustomField

// Scan implements the sql.Scanner interface
func (aq *CustomFields) Scan(src interface{}) error {
	switch src := src.(type) {
	case []byte:
		return json.Unmarshal(src, aq)
	case string:
		return json.Unmarshal([]byte(src), aq)
	default:
		return errors.New("incompatible type for CustomFields")
	}
}

// Value implements the driver.Valuer interface
func (aq CustomFields) Value() (driver.Value, error) {
	return json.Marshal(aq)
}

// CustomFieldAnswer /* This is used for answers*/
type CustomFieldAnswer struct {
	FieldName string `json:"field_name"  validate:"required"`
	FieldPath string `json:"field_path"  validate:"required"` // this is where the field could be found in the employment object (ContactInfo, Experience, etc)
	Answer    string `json:"answer"  validate:"required"`
}

type CustomFieldAnswers []CustomFieldAnswer

// Scan implements the sql.Scanner interface
func (aq *CustomFieldAnswers) Scan(src interface{}) error {
	switch src := src.(type) {
	case []byte:
		return json.Unmarshal(src, aq)
	case string:
		return json.Unmarshal([]byte(src), aq)
	default:
		return errors.New("incompatible type for CustomFieldAnswers")
	}
}

// Value implements the driver.Valuer interface
func (aq CustomFieldAnswers) Value() (driver.Value, error) {
	return json.Marshal(aq)
}

type StringList []string

// Scan implements the sql.Scanner interface
func (aq *StringList) Scan(src interface{}) error {
	switch src := src.(type) {
	case []byte:
		return json.Unmarshal(src, aq)
	case string:
		return json.Unmarshal([]byte(src), aq)
	default:
		return errors.New("incompatible type for StringList")
	}
}

// Value implements the driver.Valuer interface
func (aq StringList) Value() (driver.Value, error) {
	return json.Marshal(aq)
}
