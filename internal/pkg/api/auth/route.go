package auth

import (
	"Learnium/internal/database"
	"Learnium/internal/pkg/common"

	"Learnium/internal/pkg/adapters"

	"github.com/gofiber/fiber/v2"
)

var (
	authRepository   = NewUserRepository(database.DBConnection())
	authCache        = database.NewRedisDriver(database.RedisConnection())
	authService      = NewAuthService(authRepository, authCache)
	emailService     = adapters.NewSendGridEmailAdapter()
	authLogger       = common.NewLogger()
	requestValidator = common.NewValidator()
	controller       = NewAuthController(authService, emailService, requestValidator, authLogger)
	Middleware       = NewAuthMiddleware(authService)
)

// router function mapping routes to controller methods
func Router(app *fiber.App) {

	// auth group
	authGroup := app.Group("/api/v1/auth")
	authGroup.Post("/register", common.CustomHeaderMiddleware(), controller.handleRegistration)
	authGroup.Post("/login", common.CustomHeaderMiddleware(), controller.handleLearniumOSLogin)
	authGroup.Get("/register/resend_otp/:email", common.CustomHeaderMiddleware(), controller.resendRegistrationOTP)
	authGroup.Post("/register/confirm_otp", common.CustomHeaderMiddleware(), controller.handleRegistrationOTPValidation)
	authGroup.Post("/onboarding/personal_info", common.CustomHeaderMiddleware(),
		Middleware.GenericGuard(), controller.UpdateSuperAdminPersonalInfo)
}
