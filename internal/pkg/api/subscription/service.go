package subscription

import (
	"Learnium/internal/pkg/models"
	"context"
	// "github.com/patrickmn/go-cache"
)

// describe service to be injected by subscription controller to accomplish subscription related tasks
type ISubscriptionService interface {
	SubscribeNewsletter(ctx context.Context, email string) (models.NewsletterSubscription, error)
}

type SubscriptionService struct {
	repository ISubscriptionRepository
}

// create new subscription service instance
func NewSubscriptionService(repository ISubscriptionRepository) ISubscriptionService {
	Srv := &SubscriptionService{
		repository,
	}
	service := ISubscriptionService(Srv)
	return service
}

// subscribe user to newsletter based on email
func (srv *SubscriptionService) SubscribeNewsletter(ctx context.Context, email string) (models.NewsletterSubscription, error) {
	sub := &models.NewsletterSubscription{
		Email:      email,
		Subscribed: true,
	}
	return srv.repository.CreateNewsletterSubscription(ctx, *sub)

}
