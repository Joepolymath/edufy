package routes

// import (
// 	"Learnium/controllers"
// 	"Learnium/middleware"
// 	"github.com/gofiber/fiber/v2"
// )

// func UserRouters(incomingRoutes *fiber.App) {
// 	//	Logged in route
// 	incomingRoutes.Get("api/v1/users/user_detail", middleware.AuthMiddleware, controllers.UserDetailController)
// 	incomingRoutes.Put("api/v1/users/health_update", middleware.AuthMiddleware, controllers.HealthInfoUpdateController)
// 	incomingRoutes.Put("api/v1/users/user_update", middleware.AuthMiddleware, controllers.UserAndProfileUpdateInfoController)
// 	incomingRoutes.Get("api/v1/users/health_detail", middleware.AuthMiddleware, controllers.HealthInfoDetailController)
// 	incomingRoutes.Put("api/v1/users/user_info_update/:id", middleware.AuthMiddleware, controllers.UserInfoUpdateController)
// 	incomingRoutes.Put("api/v1/users/update_user_permission/", middleware.AuthMiddleware, controllers.LearniumUpdateUserPermission)
// }
