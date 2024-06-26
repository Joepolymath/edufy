package routes

// import (
// 	"Learnium/controllers"
// 	"Learnium/middleware"
// 	"github.com/gofiber/fiber/v2"
// )

// func StudentRouters(incomingRoutes *fiber.App) {
// 	////	Logged in route
// 	incomingRoutes.Get("api/v1/students/student_list", middleware.AuthMiddleware, controllers.StudentListController)
// 	//incomingRoutes.Get("api/v1/students/student_list/:id", middleware.AuthMiddleware, controllers.StudentDetailController)
// 	incomingRoutes.Post("api/v1/students/student_note_create/", middleware.AuthMiddleware, controllers.StudentNoteCreateController)
// 	incomingRoutes.Get("api/v1/students/student_notes/:id", middleware.AuthMiddleware, controllers.StudentNoteListController)
// 	incomingRoutes.Put("api/v1/students/student_note_list/:id", middleware.AuthMiddleware, controllers.StudentNoteUpdateController)
// 	incomingRoutes.Delete("api/v1/students/student_note_list/:id", middleware.AuthMiddleware, controllers.StudentNoteDeleteController)
// 	incomingRoutes.Get("api/v1/students/student_note_list/:id", middleware.AuthMiddleware, controllers.StudentNoteDetailController)

// 	// clinic visitation

// 	incomingRoutes.Get("api/v1/students/clinic_visitations/:id", middleware.AuthMiddleware, controllers.ClinicVisitationListController)
// 	incomingRoutes.Post("api/v1/students/clinic_visitation_list/", middleware.AuthMiddleware, controllers.ClinicVisitationCreateController)
// 	incomingRoutes.Put("api/v1/students/clinic_visitation_list/:id", middleware.AuthMiddleware, controllers.ClinicVisitationUpdateController)
// 	incomingRoutes.Get("api/v1/students/clinic_visitation_list/:id", middleware.AuthMiddleware, controllers.ClinicVisitationDetailController)
// 	incomingRoutes.Delete("api/v1/students/clinic_visitation_list/:id", middleware.AuthMiddleware, controllers.ClinicVisitationDeleteController)

// }
