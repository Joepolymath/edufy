package auth

import (
	"context"
	"strings"

	"gorm.io/gorm"

	"Learnium/internal/pkg/models"
)

// describe behaviour of user repository
type IUserRepository interface {
	// User Repository Methods
	CreateUser(ctx context.Context, user models.User) (models.User, error)
	GetUserByID(ctx context.Context, id string) (models.User, error)
	GetUserByEmail(ctx context.Context, email string) (models.User, error)
	UpdateUser(ctx context.Context, user models.User) (models.User, error)
}

// user repository connects with the users table in the database
// for user data manipulation
type UserRepository struct {
	db *gorm.DB
	// store *cache.Cache
}

// create new instance of user repository
func NewUserRepository(db *gorm.DB) IUserRepository {
	userRepo := &UserRepository{
		db,
	}
	repository := IUserRepository(userRepo)
	return repository
}

// add new user to users table based on user model
func (u *UserRepository) CreateUser(ctx context.Context, user models.User) (models.User, error) {
	db := u.db.WithContext(ctx).Model(&models.User{}).Create(&user)
	if db.Error != nil {
		if strings.Contains(db.Error.Error(), "duplicate key value") {
			return models.User{}, ErrDuplicate
		}
		return models.User{}, db.Error
	}
	return user, nil
}

// get user with unique ID
func (u *UserRepository) GetUserByID(ctx context.Context, id string) (models.User, error) {
	var user models.User
	db := u.db.WithContext(ctx).Where("id = ?", id).First(&user)
	if db.Error != nil || strings.EqualFold(user.ID, "") {
		return user, ErrRecordNotFound
	}
	return user, nil
}

// get user with email
func (u *UserRepository) GetUserByEmail(ctx context.Context, email string) (models.User, error) {
	var user models.User
	db := u.db.WithContext(ctx).Where("email = ?", email).First(&user)
	if db.Error != nil || strings.EqualFold(user.ID, "") {
		return user, ErrRecordNotFound
	}
	return user, nil
}

// update user in database
func (u *UserRepository) UpdateUser(ctx context.Context, user models.User) (models.User, error) {
	// Ensure the user has a valid ID
	if user.ID == "" {
		return models.User{}, ErrID
	}

	// Use `Updates` instead of `Save` for updating specific fields
	db := u.db.WithContext(ctx).Model(&models.User{}).Where("id = ?", user.ID).Updates(&user)
	if db.Error != nil {
		return models.User{}, db.Error
	}

	// Check if the record was found and updated
	if db.RowsAffected == 0 {
		return models.User{}, ErrRecordNotFound
	}

	return user, nil

}

//get roles by roletype
// func (u *UserRepository) GetDefaultRolesByType(ctx context.Context, roleType models.RoleType) (models.User, error) {
// 	var roles []models.Role
// 	db := u.db.WithContext(ctx).Where("email = ?", email).First(&user)
// 	if db.Error != nil || strings.EqualFold(user.ID, "") {
// 		return user, ErrRecordNotFound
// 	}
// 	return user, nil
// }
