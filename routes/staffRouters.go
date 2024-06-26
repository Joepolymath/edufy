package routes

// import (
// 	"Learnium/controllers"
// 	"Learnium/middleware"
// 	"github.com/gofiber/fiber/v2"
// )

// func StaffRouters(incomingRoutes *fiber.App) {
// 	//	Logged in route
// 	incomingRoutes.Get("api/v1/staffs/staff_list", middleware.AuthMiddleware, controllers.StaffListController)
// 	incomingRoutes.Get("api/v1/staffs/staff_list/:id", middleware.AuthMiddleware, controllers.StaffDetailController)
// 	incomingRoutes.Put("api/v1/staffs/staff_list/:id", middleware.AuthMiddleware, controllers.StaffUpdateController)

// 	// staff rating
// 	incomingRoutes.Post("api/v1/staffs/staff_rating", middleware.AuthMiddleware, controllers.StaffRatingCreateController)
// 	incomingRoutes.Get("api/v1/staffs/staffs_analytics", middleware.AuthMiddleware, controllers.StaffsAnalyticController)
// 	incomingRoutes.Get("api/v1/staffs/staff_analytics/:id", middleware.AuthMiddleware, controllers.StaffAnalyticsController)
// 	incomingRoutes.Get("api/v1/staffs/staff_rating_list/:id", middleware.AuthMiddleware, controllers.StaffRatingListController)
// }
