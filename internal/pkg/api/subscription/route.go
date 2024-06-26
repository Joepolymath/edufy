package subscription

import (
	"Learnium/internal/database"
	"Learnium/internal/pkg/common"

	"Learnium/internal/pkg/adapters"

	"github.com/gofiber/fiber/v2"
)

var (
	repository       = NewSubscriptionRepository(database.DBConnection())
	subService       = NewSubscriptionService(repository)
	emailService     = adapters.NewSendGridEmailAdapter()
	authLogger       = common.NewLogger()
	requestValidator = common.NewValidator()
	controller       = NewSubscriptionController(subService, emailService, requestValidator, authLogger)
)

// router function mapping routes to controller methods
func Router(app *fiber.App) {

	// subscriptions group
	subGroup := app.Group("/api/v1/subscription")
	subGroup.Post("/newsletter", common.CustomHeaderMiddleware(), controller.handleNewsletterSubscription)

}
