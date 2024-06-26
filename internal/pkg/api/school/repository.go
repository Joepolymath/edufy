package school

import (
	"Learnium/internal/pkg/common"
	"Learnium/internal/pkg/models"
	"context"
	"fmt"
	"strings"

	"gorm.io/gorm"
)

// describe behaviour of school repository
type ISchoolRepository interface {
	// School Repository Methods
	CreateSchool(ctx context.Context, sch models.School) (models.School, error)
	CreateSession(ctx context.Context, sess models.Session) (models.Session, error)
	CreateSchoolRole(ctx context.Context, role models.SchoolRole) (models.SchoolRole, error)
	FilterSchoolRole(ctx context.Context, conditions map[string]interface{}) ([]models.SchoolRole, error)
	CreateSchoolRoles(ctx context.Context, roles []models.SchoolRole) error
	GetSchoolByID(ctx context.Context, id string) (models.School, error)
	GetParentSchoolByOwnerID(ctx context.Context, id string) (models.School, error)
	GetSchoolBranches(ctx context.Context, parentSchoolId string) ([]models.School, error)
	GetSchoolByEmail(ctx context.Context, email string) (models.School, error)
}

// school repository connects with the school table in the database for data manipulation
type SchoolRepository struct {
	db *gorm.DB
}

// create new instance of school repository
func NewSchoolRepository(db *gorm.DB) ISchoolRepository {
	schRepo := &SchoolRepository{
		db,
	}
	repository := ISchoolRepository(schRepo)
	return repository
}

// add new school to school subscriptions table
func (sr *SchoolRepository) CreateSchool(ctx context.Context,
	sch models.School) (models.School, error) {
	db := sr.db.WithContext(ctx).Model(&models.School{}).Create(&sch)
	if db.Error != nil {
		if strings.Contains(db.Error.Error(), "duplicate key value") {
			return models.School{}, common.ErrDuplicate
		}
		return models.School{}, db.Error
	}
	return sch, nil
}

// filter school role search by attribute
func (sr *SchoolRepository) FilterSchoolRole(ctx context.Context,
	conditions map[string]interface{}) ([]models.SchoolRole, error) {
	var roles []models.SchoolRole

	query := sr.db.WithContext(ctx).Model(&roles)

	for key, value := range conditions {
		query = query.Where(fmt.Sprintf("%s = ?", key), value)
	}

	if err := query.Find(&roles).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return roles, common.ErrRecordNotFound
		}
		return roles, err
	}

	return roles, nil
}

func (sr *SchoolRepository) CreateSession(ctx context.Context,
	sess models.Session) (models.Session, error) {
	db := sr.db.WithContext(ctx).Model(&models.Session{}).Create(&sess)
	if db.Error != nil {
		if strings.Contains(db.Error.Error(), "duplicate key value") {
			return models.Session{}, common.ErrDuplicate
		}
		return models.Session{}, db.Error
	}
	return sess, nil
}

// create school based role
func (sr *SchoolRepository) CreateSchoolRole(ctx context.Context,
	role models.SchoolRole) (models.SchoolRole, error) {
	db := sr.db.WithContext(ctx).Model(&models.SchoolRole{}).Create(&role)
	if db.Error != nil {
		if strings.Contains(db.Error.Error(), "duplicate key value") {
			return models.SchoolRole{}, common.ErrDuplicate
		}
		return models.SchoolRole{}, db.Error
	}
	return role, nil
}

func (sr *SchoolRepository) CreateSchoolRoles(ctx context.Context,
	roles []models.SchoolRole) error {
	if err := sr.db.Create(&roles).Error; err != nil {
		if strings.Contains(err.Error(), "duplicate key value") {
			return common.ErrDuplicate
		}
		return err
	}

	return nil
}

func (sr *SchoolRepository) GetSchoolByID(ctx context.Context, id string) (models.School, error) {
	var sch models.School
	db := sr.db.WithContext(ctx).Where("id = ?", id).First(&sch)
	if db.Error != nil || strings.EqualFold(sch.ID, "") {
		return sch, common.ErrRecordNotFound
	}
	return sch, nil
}

func (sr *SchoolRepository) GetSchoolBranches(ctx context.Context, id string) ([]models.School, error) {
	var schs []models.School
	db := sr.db.WithContext(ctx).Where("parent_school_id = ?", id).Find(&schs)
	if db.Error != nil {
		if db.Error == gorm.ErrRecordNotFound {
			return nil, common.ErrRecordNotFound
		}
		return nil, db.Error
	}
	return schs, nil
}

func (sr *SchoolRepository) GetParentSchoolByOwnerID(ctx context.Context, id string) (models.School, error) {
	var sch models.School
	db := sr.db.WithContext(ctx).Where("owner_id = ? AND is_parent_school = ?", id, true).First(&sch)
	if db.Error != nil || sch.ID == "" {
		return sch, common.ErrRecordNotFound
	}
	return sch, nil
}

// get school by email
func (sr *SchoolRepository) GetSchoolByEmail(ctx context.Context, email string) (models.School, error) {
	var sch models.School
	db := sr.db.WithContext(ctx).Where("email = ?", email).First(&sch)
	if db.Error != nil || strings.EqualFold(sch.ID, "") {
		return sch, common.ErrRecordNotFound
	}
	return sch, nil
}
