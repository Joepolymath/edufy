package utils

// import (
// 	"Learnium/adapters"
// 	"Learnium/logger"
// 	"Learnium/models"
// 	"context"
// 	"crypto/rand"
// 	"errors"
// 	"github.com/google/uuid"
// 	"go.uber.org/zap"
// 	"gorm.io/gorm"
// 	"math/big"
// )

// type UserUtilsInterface interface {
// 	CreateUser(db *gorm.DB, ctx context.Context, user models.User, sendEmail bool) (models.User, error)
// 	GenerateRandomPassword(length int) (string, error)
// }

// type UserUtils struct {
// }

// func NewUserUtils() UserUtilsInterface {
// 	return &UserUtils{}
// }

// func (u *UserUtils) CreateUser(db *gorm.DB, ctx context.Context, user models.User, sendEmail bool) (models.User, error) {

// 	emailAdapters := adapters.NewEmailAdapter()

// 	// hash the password
// 	password := HashPassword(*user.Password)
// 	user.Password = &password
// 	user.ID = uuid.New()

// 	// create the user
// 	err := db.WithContext(ctx).Model(&user).Create(&user).Error
// 	if err != nil {
// 		return user, errors.New(err.Error())
// 	}

// 	// create the user profile and health
// 	_, err = user.CreateAssociation(ctx, db)
// 	if err != nil {
// 		// delete the user if the profile and health creation fails
// 		db.WithContext(ctx).Model(&user).Delete(&user)
// 		return models.User{}, err
// 	}

// 	if sendEmail {
// 		_, err := emailAdapters.SendAccountCreateEmail(ctx, *user.FirstName, *user.LastName, *user.Email, password)
// 		if err != nil {
// 			db.WithContext(ctx).Model(&user).Delete(&user)
// 			return user, errors.New(err.Error())
// 		}
// 	}

// 	return user, nil
// }

// // GenerateRandomPassword generates a random password of the specified length
// func (u *UserUtils) GenerateRandomPassword(length int) (string, error) {
// 	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789@#$&*"
// 	password := make([]byte, length)

// 	for i := range password {
// 		// Generate a random index within the charset
// 		idx, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
// 		if err != nil {
// 			return "", err
// 		}

// 		// Use the index to select a character from the charset
// 		password[i] = charset[idx.Int64()]
// 	}
// 	logger.Info(context.Background(), "The password Generated is: ", zap.String("password", string(password)))

// 	return string(password), nil
// }
