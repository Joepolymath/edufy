package routes

// import (
// 	"Learnium/controllers"
// 	"Learnium/middleware"

// 	"github.com/gofiber/fiber/v2"
// )

// func SchoolRouters(incomingRoutes *fiber.App) {

// 	// school type routes
// 	incomingRoutes.Post("api/v1/schools/school_type_create", middleware.AuthMiddleware, controllers.SchoolTypeCreateController)
// 	incomingRoutes.Get("api/v1/schools/school_type_list", controllers.SchoolTypeListController)
// 	incomingRoutes.Put("api/v1/schools/school_type_list/:id", middleware.AuthMiddleware, controllers.SchoolTypeUpdateController)
// 	incomingRoutes.Delete("api/v1/schools/school_type_list/:id", middleware.AuthMiddleware, controllers.SchoolTypeDeleteController)

// 	//	school routes
// 	incomingRoutes.Get("api/v1/schools/schools_admin_list", middleware.AuthMiddleware, controllers.SchoolOwnerListCreatedController)
// 	incomingRoutes.Put("api/v1/schools/school_update/:id", middleware.AuthMiddleware, controllers.SchoolUpdateController)
// 	// incomingRoutes.Post("api/v1/schools/school_create", middleware.AuthMiddleware, controllers.SchoolCreateController)
// 	incomingRoutes.Post("api/v1/schools/school_create", controllers.SchoolCreateController)

// 	// User invite routes
// 	incomingRoutes.Post("api/v1/schools/school_invite_create", middleware.AuthMiddleware, controllers.SchoolInviteCreateController)
// 	incomingRoutes.Get("api/v1/schools/school_invite_list", middleware.AuthMiddleware, controllers.SchoolInviteListController)

// 	//	class routes
// 	incomingRoutes.Post("api/v1/schools/class_create", middleware.AuthMiddleware, controllers.ClassCreateController)
// 	incomingRoutes.Put("api/v1/schools/class_update/:id", middleware.AuthMiddleware, controllers.ClassUpdateController)
// 	incomingRoutes.Get("api/v1/schools/class_list", middleware.AuthMiddleware, controllers.ClassListController)
// 	incomingRoutes.Delete("api/v1/schools/class_delete/:id", middleware.AuthMiddleware, controllers.ClassDeleteController)

// 	// role routes
// 	incomingRoutes.Post("api/v1/schools/role_create", middleware.AuthMiddleware, controllers.RoleCreateController)
// 	incomingRoutes.Get("api/v1/schools/role_list", controllers.RoleListController)
// 	incomingRoutes.Get("api/v1/schools/role_list/:id", controllers.RoleDetailController)
// 	incomingRoutes.Put("api/v1/schools/role_list/:id", middleware.AuthMiddleware, controllers.RoleUpdateController)
// 	incomingRoutes.Delete("api/v1/schools/role_list/:id", middleware.AuthMiddleware, controllers.RoleDeleteController)

// 	// category routes
// 	incomingRoutes.Post("api/v1/schools/category_create", middleware.AuthMiddleware, controllers.CategoryCreateController)
// 	incomingRoutes.Get("api/v1/schools/category_list", controllers.CategoryListController)
// 	incomingRoutes.Get("api/v1/schools/category_list/:id", controllers.CategoryDetailController)
// 	incomingRoutes.Put("api/v1/schools/category_list/:id", middleware.AuthMiddleware, controllers.CategoryUpdateController)
// 	incomingRoutes.Delete("api/v1/schools/category_list/:id", middleware.AuthMiddleware, controllers.CategoryDeleteController)

// 	// school admin
// 	incomingRoutes.Get("api/v1/schools/school_admin_list", middleware.AuthMiddleware, controllers.SchoolAdminListController)
// 	incomingRoutes.Put("api/v1/schools/school_admin_list/:id", middleware.AuthMiddleware, controllers.SchoolAdminUpdateController)
// 	incomingRoutes.Get("api/v1/schools/school_admin_list/:id", middleware.AuthMiddleware, controllers.SchoolAdminDetailController)

// 	// event
// 	incomingRoutes.Post("api/v1/schools/event_create", middleware.AuthMiddleware, controllers.EventCreateController)
// 	incomingRoutes.Post("api/v1/schools/event_list", middleware.AuthMiddleware, controllers.EventListController)

// 	// announcement
// 	incomingRoutes.Put("api/v1/schools/announcement_update", middleware.AuthMiddleware, controllers.UpdateAnnouncementConfigurationController)

// 	// notification
// 	incomingRoutes.Put("api/v1/schools/notification_update", middleware.AuthMiddleware, controllers.UpdateNotificationConfigurationController)

// 	// add admin
// 	incomingRoutes.Post("api/v1/schools/school_admin_add", middleware.AuthMiddleware, controllers.SchoolAdminAddController)

// }
