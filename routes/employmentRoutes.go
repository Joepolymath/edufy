package routes

// import (
// 	"Learnium/controllers"
// 	"Learnium/middleware"
// 	"github.com/gofiber/fiber/v2"
// )

// func EmploymentRouters(incomingRoutes *fiber.App) {
// 	//	Employment Configuration
// 	incomingRoutes.Post("api/v1/employments/employment_config_create", middleware.AuthMiddleware, controllers.EmploymentConfigurationCreateController)
// 	incomingRoutes.Get("api/v1/employments/employment_config_list", middleware.AuthMiddleware, controllers.EmploymentConfigurationListController)
// 	incomingRoutes.Get("api/v1/employments/employment_config_list/:id", middleware.AuthMiddleware, controllers.EmploymentConfigurationDetailController)
// 	incomingRoutes.Put("api/v1/employments/employment_config_list/:id", middleware.AuthMiddleware, controllers.EmploymentConfigurationUpdateController)
// 	incomingRoutes.Delete("api/v1/employments/employment_config_list/:id", middleware.AuthMiddleware, controllers.EmploymentConfigurationDeleteController)

// 	//  Employment routes (Jobs)
// 	incomingRoutes.Post("api/v1/employments/employment_create", middleware.AuthMiddleware, controllers.EmploymentCreateController)
// 	incomingRoutes.Get("api/v1/employments/employment_list", controllers.EmploymentListController)
// 	incomingRoutes.Get("api/v1/employments/employment_list/:id", controllers.EmploymentDetailController)
// 	incomingRoutes.Put("api/v1/employments/employment_list/:id", middleware.AuthMiddleware, controllers.EmploymentUpdateController)
// 	incomingRoutes.Delete("api/v1/employments/employment_list/:id", middleware.AuthMiddleware, controllers.EmploymentDeleteController)

// 	// Applicant Routes
// 	incomingRoutes.Post("api/v1/employments/employment_applicant/:id", controllers.EmploymentApplicationController)
// 	incomingRoutes.Get("api/v1/employments/employment_applicant/:id", middleware.AuthMiddleware, controllers.EmploymentApplicantListController)

// 	// accept applicant route
// 	incomingRoutes.Post("api/v1/employments/employment_applicant_accept", middleware.AuthMiddleware, controllers.EmploymentApplicantAcceptController)

// }
