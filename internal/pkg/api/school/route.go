package school

import (
	"Learnium/internal/database"
	"Learnium/internal/pkg/common"

	"Learnium/internal/pkg/adapters"

	"Learnium/internal/pkg/api/auth"

	"github.com/gofiber/fiber/v2"
)

var (
	repository       = NewSchoolRepository(database.DBConnection())
	schoolService    = NewSchoolService(repository)
	emailService     = adapters.NewSendGridEmailAdapter()
	storageService   = adapters.NewFileUploadAdapter()
	logger           = common.NewLogger()
	requestValidator = common.NewValidator()
	controller       = NewSchoolController(schoolService, emailService, storageService, requestValidator, logger)
)

// router function mapping routes to controller methods
func Router(app *fiber.App) {

	schoolGroup := app.Group("/api/v1/school")

	schoolGroup.Post("/setup", common.CustomHeaderMiddleware(),
		auth.Middleware.SuperAdminGuard(), controller.handleSchoolSetUp)
	schoolGroup.Post("/branch", common.CustomHeaderMiddleware(),
		auth.Middleware.SuperAdminGuard(), controller.handleBranchAddition)
	schoolGroup.Get("/", common.CustomHeaderMiddleware(),
		auth.Middleware.SuperAdminGuard(), controller.GetSchoolsOwned)
	schoolGroup.Post("/session", common.CustomHeaderMiddleware(),
		auth.Middleware.SuperAdminGuard(), controller.handleSessionManagement)
	schoolGroup.Get("/roles", common.CustomHeaderMiddleware(), controller.handleRolesFetching)
}
