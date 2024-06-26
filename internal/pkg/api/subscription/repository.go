package subscription

import (
	"Learnium/internal/pkg/common"
	"Learnium/internal/pkg/models"
	"context"
	"strings"

	"gorm.io/gorm"
)

// describe behaviour of subscription repository
type ISubscriptionRepository interface {
	// User Repository Methods
	CreateNewsletterSubscription(ctx context.Context, sub models.NewsletterSubscription) (models.NewsletterSubscription, error)
	GetNewsletterSubByID(ctx context.Context, id string) (models.NewsletterSubscription, error)
	GetNewsletterSubByEmail(ctx context.Context, email string) (models.NewsletterSubscription, error)
}

// subscription repository connects with the subscriptions table in the database for data manipulation
type SubscriptionRepository struct {
	db *gorm.DB
}

// create new instance of user repository
func NewSubscriptionRepository(db *gorm.DB) ISubscriptionRepository {
	subRepo := &SubscriptionRepository{
		db,
	}
	repository := ISubscriptionRepository(subRepo)
	return repository
}

// add new newsletter subscription to newsletter subscriptions table
func (sr *SubscriptionRepository) CreateNewsletterSubscription(ctx context.Context,
	sub models.NewsletterSubscription) (models.NewsletterSubscription, error) {
	db := sr.db.WithContext(ctx).Model(&models.NewsletterSubscription{}).Create(&sub)
	if db.Error != nil {
		if strings.Contains(db.Error.Error(), "duplicate key value") {
			return models.NewsletterSubscription{}, common.ErrDuplicate
		}
		return models.NewsletterSubscription{}, db.Error
	}
	return sub, nil
}

// get newsletter subscription by id
func (sr *SubscriptionRepository) GetNewsletterSubByID(ctx context.Context, id string) (models.NewsletterSubscription, error) {
	var sub models.NewsletterSubscription
	db := sr.db.WithContext(ctx).Where("id = ?", id).First(&sub)
	if db.Error != nil || strings.EqualFold(sub.ID, "") {
		return sub, common.ErrRecordNotFound
	}
	return sub, nil
}

// get newsletter subscription by email
func (sr *SubscriptionRepository) GetNewsletterSubByEmail(ctx context.Context, email string) (models.NewsletterSubscription, error) {
	var sub models.NewsletterSubscription
	db := sr.db.WithContext(ctx).Where("email = ?", email).First(&sub)
	if db.Error != nil || strings.EqualFold(sub.ID, "") {
		return sub, common.ErrRecordNotFound
	}
	return sub, nil
}
