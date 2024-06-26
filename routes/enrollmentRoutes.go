package routes

// import (
// 	"Learnium/controllers"
// 	"Learnium/middleware"
// 	"github.com/gofiber/fiber/v2"
// )

// func EnrollmentRouters(incomingRoutes *fiber.App) {
// 	//	Employment Configuration
// 	incomingRoutes.Post("api/v1/enrollments/enrollment_config_create", middleware.AuthMiddleware, controllers.EnrollmentConfigurationCreateController)
// 	incomingRoutes.Get("api/v1/enrollments/enrollment_config_list", controllers.EnrollmentConfigurationListController)
// 	incomingRoutes.Get("api/v1/enrollments/enrollment_config_list/:id", controllers.EnrollmentConfigurationDetailController)
// 	incomingRoutes.Put("api/v1/enrollments/enrollment_config_list/:id", middleware.AuthMiddleware, controllers.EnrollmentConfigurationUpdateController)
// 	incomingRoutes.Delete("api/v1/enrollments/enrollment_config_list/:id", middleware.AuthMiddleware, controllers.EnrollmentConfigurationDeleteController)

// 	// Applicant Routes
// 	incomingRoutes.Post("api/v1/enrollments/enrollment_applicant", controllers.EnrollmentApplicationController)
// 	incomingRoutes.Get("api/v1/enrollments/enrollment_applicants", middleware.AuthMiddleware, controllers.EnrollmentApplicantListController)
// 	incomingRoutes.Get("api/v1/enrollments/enrollment_applicants/:id", middleware.AuthMiddleware, controllers.EnrollmentApplicantDetailController)

// 	// Accept Applicant
// 	incomingRoutes.Post("api/v1/enrollments/enrollment_applicant_accept", middleware.AuthMiddleware, controllers.EnrollmentApplicantAcceptController)
// 	incomingRoutes.Get("api/v1/enrollments/enrollment_analytics", middleware.AuthMiddleware, controllers.EnrollmentAnalyticsController)

// }
