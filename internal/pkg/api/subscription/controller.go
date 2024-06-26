package subscription

import (
	"Learnium/internal/pkg/adapters"
	"Learnium/internal/pkg/common"
	"context"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type SubscriptionController struct {
	service      ISubscriptionService
	mailProvider adapters.IEmailAdapter
	serializer   common.IValidator
	logger       common.ILogger
}

// create new instance of Subscription controller
func NewSubscriptionController(srv ISubscriptionService, mail adapters.IEmailAdapter,
	validator common.IValidator, logger common.ILogger) SubscriptionController {
	// instantiate subscription controller with dependencies
	controller := &SubscriptionController{
		service:      srv,
		mailProvider: mail,
		serializer:   validator,
		logger:       logger,
	}
	return *controller
}

func (cntrl *SubscriptionController) handleNewsletterSubscription(c *fiber.Ctx) error {

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	var requestBody SignupNewsletterDto

	// validate request body
	if err := cntrl.serializer.ValidateRequestBody(c, &requestBody); err != nil {
		return c.Status(422).JSON(common.APIResponse{Message: "Invalid Request Body", Success: false, Detail: err.Error()})
	}

	// subscribe to newsletter
	savedSub, err := cntrl.service.SubscribeNewsletter(ctx, requestBody.Email)
	if err != nil && err == common.ErrDuplicate {
		return c.Status(409).JSON(common.APIResponse{Message: "Subscription Failed", Success: false, Detail: "User with Email subscribed already"})
	} else if err != nil {
		cntrl.logger.Error(ctx, "Error while subscribing to newsletter",
			zap.String("additional_info", err.Error()))
		return c.Status(500).JSON(common.APIResponse{Message: "Subscription Failed", Success: false, Detail: "Internal Server Error"})
	}

	mailBody := fmt.Sprintf("You have subscribed to Learnium Newsletters")
	// send otp mail to user
	err = cntrl.mailProvider.SendEmail(savedSub.Email, "Newsletter Subscription", mailBody)
	if err != nil {
		cntrl.logger.Error(ctx, "error during registration otp mail sending",
			zap.String("additional_info", err.Error()))
	}

	return c.Status(201).JSON(common.APIResponse{
		Message: "User Subscribed to Newsletter!",
		Success: true})
}
