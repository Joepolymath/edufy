package models

// import (
// 	"Learnium/logger"
// 	"context"
// 	"errors"
// 	"fmt"
// 	"github.com/google/uuid"
// 	"go.uber.org/zap"
// 	"gorm.io/gorm"
// )

// type Conversation struct {
// 	BaseModel
// 	School           *School    `json:"school,omitempty"  gorm:"constraint:OnDelete:CASCADE;ForeignKey:SchoolID"` // school the subject belongs to
// 	SchoolID         *uuid.UUID `json:"school_id" gorm:"not null;"`
// 	Name             *string    `json:"name" gorm:"not null;uniqueIndex"`
// 	ConversationType *string    `json:"conversation_type" gorm:"default:CHAT"` // CHAT OR ADMIN_TEACHERS OR ADMIN_ALL OR TEACHERS_STUDENTS
// 	LastMessage      *string    `json:"last_message" `
// }

// func (conversation *Conversation) GetOrConnectConversation(ctx context.Context, db *gorm.DB, schoolID, conversationUserID1, conversationUserID2 uuid.UUID) (Conversation, error) {
// 	// Check if ConversationUser records exist for the provided user IDs.
// 	var conversationUser1, conversationUser2 ConversationUser

// 	err := db.WithContext(ctx).Model(&ConversationUser{}).Where("id = ? AND school_id = ?", conversationUserID1, schoolID).
// 		First(&conversationUser1).Error
// 	if err != nil {
// 		return Conversation{}, err
// 	}

// 	err = db.WithContext(ctx).Model(&ConversationUser{}).Where("id = ? AND school_id = ?", conversationUserID2, schoolID).
// 		First(&conversationUser2).Error
// 	if err != nil {
// 		return Conversation{}, err
// 	}

// 	// Check if a conversation already exists with the given ConversationUser records.
// 	var existingConversation Conversation

// 	err = db.WithContext(ctx).
// 		Where("school_id = ?", schoolID).
// 		Where("name ILIKE ? AND name ILIKE ?", "%"+conversationUser1.ID.String()+"%", "%"+conversationUser2.ID.String()+"%").First(&existingConversation).
// 		Error

// 	if err != nil {
// 		// If the conversation doesn't exist, create a new one.
// 		if err != gorm.ErrRecordNotFound {
// 			return Conversation{}, err
// 		}

// 		name := fmt.Sprintf("%v__%v", conversationUserID1, conversationUserID2)
// 		conversationType := "CHAT"
// 		newConversation := Conversation{
// 			Name:             &name,
// 			SchoolID:         &schoolID,
// 			ConversationType: &conversationType,
// 		}

// 		// Create the new conversation and save it to the database.
// 		if err := db.WithContext(ctx).Create(&newConversation).Error; err != nil {
// 			return Conversation{}, err
// 		}

// 		return newConversation, nil
// 	}

// 	return existingConversation, nil
// }

// type Message struct {
// 	BaseModel
// 	School         *School           `json:"school,omitempty"  gorm:"constraint:OnDelete:CASCADE;ForeignKey:SchoolID"` // school the subject belongs to
// 	SchoolID       *uuid.UUID        `json:"school_id" gorm:"not null;"`
// 	Conversation   *Conversation     `json:"conversation,omitempty" gorm:"constraint:OnDelete:CASCADE;foreignKey:ConversationID;"`
// 	ConversationID *uuid.UUID        `json:"conversation_id" gorm:"not null;"`
// 	FromUser       *ConversationUser `json:"from_user,omitempty" gorm:"constraint:OnDelete:CASCADE;foreignKey:FromUserID;"`
// 	FromUserID     *uuid.UUID        `json:"from_user_id" gorm:"not null;"`
// 	ToUser         *ConversationUser `json:"to_user,omitempty" gorm:"constraint:OnDelete:CASCADE;foreignKey:ToUserID;"`
// 	ToUserID       *uuid.UUID        `json:"to_user_id" gorm:"not null;"`
// 	Content        *string           `json:"content"`
// 	File           *string           `json:"file"`
// 	Read           *bool             `json:"read"`
// }

// type ConversationUser struct {
// 	BaseModel
// 	School    *School    `json:"school,omitempty"  gorm:"constraint:OnDelete:CASCADE;ForeignKey:SchoolID"` // school the subject belongs to
// 	SchoolID  *uuid.UUID `json:"school_id" gorm:"not null;"`
// 	User      *User      `json:"user,omitempty" gorm:"constraint:OnDelete:CASCADE;ForeignKey:UserID" `
// 	UserID    *uuid.UUID `json:"user_id"`
// 	UserType  *string    `json:"user_type"` // STAFF OR STUDENT
// 	Student   *Student   `json:"student,omitempty" gorm:"constraint:OnDelete:CASCADE;ForeignKey:StudentID" `
// 	StudentID *uuid.UUID `json:"student_id"`
// 	Staff     *Staff     `json:"staff,omitempty" gorm:"constraint:OnDelete:CASCADE;ForeignKey:StaffID" `
// 	StaffID   *uuid.UUID `json:"staff_id"`
// 	IsOnline  *bool      `json:"is_online"`
// }

// func (conversationUser *ConversationUser) GetConversationUserWithID(ctx context.Context, db *gorm.DB, conversationUserID uuid.UUID) (ConversationUser, error) {
// 	var conversationUsr ConversationUser

// 	err := db.WithContext(ctx).Model(&ConversationUser{}).Where("id =?", conversationUserID).Preload("User").First(&conversationUsr).Error
// 	if err != nil {
// 		if err != gorm.ErrRecordNotFound {
// 			logger.Error(ctx, fmt.Sprintf("Error getting  the conversation user for %p", conversationUsr.UserType), zap.Error(err))
// 		}
// 		return conversationUsr, err
// 	}
// 	return conversationUsr, nil
// }

// func (conversationUser *ConversationUser) CreateConversationUser(ctx context.Context, db *gorm.DB) error {
// 	err := db.WithContext(ctx).Model(&ConversationUser{}).Create(&conversationUser).Preload("User").First(&conversationUser).Error
// 	if err != nil {
// 		logger.Error(ctx, fmt.Sprintf("Error creating the conversation user for %p", conversationUser.UserType), zap.Error(err))
// 		return err
// 	}
// 	return nil
// }

// func (conversationUser *ConversationUser) GetOrCreateConversationUser(ctx context.Context, db *gorm.DB, userID uuid.UUID, schoolID uuid.UUID, schoolCode string) (ConversationUser, error) {
// 	var school School
// 	var student Student
// 	var staff Staff

// 	err := db.WithContext(ctx).Model(&ConversationUser{}).Where("school_id = ? AND user_id =? ", schoolID, userID).Preload("User").First(&conversationUser).Error
// 	if err != nil {
// 		if err != gorm.ErrRecordNotFound {
// 			logger.Error(ctx, "error filtering conversation user")
// 			return *conversationUser, err
// 		}
// 		isAdmin := school.IsSchoolAdminOrOwner(ctx, db, schoolCode, userID)
// 		if !isAdmin {
// 			logger.Error(ctx, "User exists who is not an admin user and still have no conversation user",
// 				zap.String("error", "User not and admin user and no conversation user"))
// 			return *conversationUser, errors.New("not an admin user")
// 		}

// 		studentExist := true
// 		student, err = student.RetrieveStudentByUserIDAndSchool(db, ctx, userID, schoolID)
// 		if err != nil {
// 			studentExist = false
// 		}

// 		staffExist := true
// 		staff, err = staff.RetrieveByUserAndSchool(ctx, db, schoolID, userID)
// 		if err != nil {
// 			staffExist = false
// 		}

// 		if studentExist {
// 			userType := "STUDENT"
// 			conversationUser := ConversationUser{
// 				SchoolID:  &schoolID,
// 				UserID:    &userID,
// 				UserType:  &userType,
// 				StudentID: &student.ID,
// 			}
// 			err := conversationUser.CreateConversationUser(ctx, db)
// 			if err != nil {
// 				return ConversationUser{}, err
// 			}
// 		}

// 		if staffExist {
// 			userType := "STAFF"
// 			conversationUser := ConversationUser{
// 				SchoolID: &schoolID,
// 				UserID:   &userID,
// 				UserType: &userType,
// 				StaffID:  &staff.ID,
// 			}
// 			err := conversationUser.CreateConversationUser(ctx, db)
// 			if err != nil {
// 				return ConversationUser{}, err
// 			}
// 		}

// 		if isAdmin {
// 			userType := "STAFF"
// 			conversationUser := ConversationUser{
// 				SchoolID: &schoolID,
// 				UserID:   &userID,
// 				UserType: &userType,
// 			}
// 			err := conversationUser.CreateConversationUser(ctx, db)
// 			if err != nil {
// 				return ConversationUser{}, err
// 			}
// 		}

// 	}
// 	return *conversationUser, nil
// }

// func (conversationUser *ConversationUser) CreateConversationUserForAllUsers(ctx context.Context, db *gorm.DB, schoolID uuid.UUID) error {
// 	var students []Student
// 	var staffs []Staff

// 	err := db.WithContext(ctx).Model(&Student{}).Where("school_id =?", schoolID).Find(&students).Error
// 	if err != nil {
// 		logger.Error(ctx, fmt.Sprintf("Error getting all students"), zap.Error(err))
// 		return err
// 	}

// 	err = db.WithContext(ctx).Model(&Staff{}).Where("school_id =?", schoolID).Find(&staffs).Error
// 	if err != nil {
// 		logger.Error(ctx, fmt.Sprintf("Error getting all staffs"), zap.Error(err))
// 		return err
// 	}

// 	for _, student := range students {
// 		userType := "STUDENT"
// 		IsOnline := false
// 		conversationUser := ConversationUser{
// 			SchoolID:  &schoolID,
// 			UserID:    student.UserID,
// 			UserType:  &userType,
// 			StudentID: &student.ID,
// 			IsOnline:  &IsOnline,
// 		}
// 		err = db.WithContext(ctx).Model(&ConversationUser{}).Where("school_id = ? AND user_id =? ", schoolID, student.UserID).First(&conversationUser).Error
// 		if err != nil {
// 			if err != gorm.ErrRecordNotFound {
// 				logger.Error(ctx, "error filtering conversation user")
// 				return err
// 			}
// 			err := db.WithContext(ctx).Model(&ConversationUser{}).Create(&conversationUser).Error
// 			if err != nil {
// 				logger.Error(ctx, "Error creating student conversation user", zap.Error(err))
// 				return err
// 			}
// 		}

// 	}

// 	for _, staff := range staffs {
// 		userType := "STAFF"
// 		IsOnline := false
// 		conversationUser := ConversationUser{
// 			SchoolID: &schoolID,
// 			UserID:   staff.UserID,
// 			UserType: &userType,
// 			StaffID:  &staff.ID,
// 			IsOnline: &IsOnline,
// 		}
// 		err = db.WithContext(ctx).Model(&ConversationUser{}).Where("school_id = ? AND user_id =? ", schoolID, staff.UserID).First(&conversationUser).Error
// 		if err != nil {
// 			if err != gorm.ErrRecordNotFound {
// 				logger.Error(ctx, "error filtering conversation user for staff")
// 				return err
// 			}
// 			err := db.WithContext(ctx).Model(&ConversationUser{}).Create(&conversationUser).Error
// 			if err != nil {
// 				logger.Error(ctx, "Error creating staff conversation user", zap.Error(err))
// 				return err
// 			}
// 		}

// 	}

// 	return nil
// }
