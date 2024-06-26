package routes

// import (
// 	"Learnium/controllers"
// 	"Learnium/middleware"
// 	"github.com/gofiber/fiber/v2"
// )

// func TaskRouters(incomingRoutes *fiber.App) {
// 	// course ( staff)
// 	incomingRoutes.Post("api/v1/tasks/course_task_create/", middleware.AuthMiddleware, controllers.CourseTaskCreateController)
// 	incomingRoutes.Get("api/v1/tasks/course_task_list/:id", middleware.AuthMiddleware, controllers.CourseTaskListController)
// 	incomingRoutes.Get("api/v1/tasks/course_task_list/:course_id/:id", middleware.AuthMiddleware, controllers.CourseTaskDetailController)
// 	incomingRoutes.Put("api/v1/tasks/course_task_list/:course_id/:id", middleware.AuthMiddleware, controllers.CourseTaskUpdateController)
// 	incomingRoutes.Delete("api/v1/tasks/course_task_list/:course_id/:id", middleware.AuthMiddleware, controllers.CourseTaskDeleteController)

// 	// questions
// 	incomingRoutes.Post("api/v1/tasks/course_task_question_create/", middleware.AuthMiddleware, controllers.CourseTaskQuestionCreateController)
// 	incomingRoutes.Put("api/v1/tasks/course_task_question/:id", middleware.AuthMiddleware, controllers.CourseTaskQuestionUpdateController)
// 	incomingRoutes.Delete("api/v1/tasks/course_task_question/:id", middleware.AuthMiddleware, controllers.CourseTaskQuestionDeleteController)
// 	incomingRoutes.Get("api/v1/tasks/course_task_question/:id", middleware.AuthMiddleware, controllers.CourseTaskQuestionDetailController)

// 	// Options routes
// 	incomingRoutes.Put("api/v1/tasks/course_task_option/:id", middleware.AuthMiddleware, controllers.CourseTaskOptionUpdateController)
// 	incomingRoutes.Delete("api/v1/tasks/course_task_option/:id", middleware.AuthMiddleware, controllers.CourseTaskOptionDeleteController)
// 	incomingRoutes.Get("api/v1/tasks/course_task_option/:id", middleware.AuthMiddleware, controllers.CourseTaskOptionDetailController)
// 	incomingRoutes.Post("api/v1/tasks/course_task_option_create", middleware.AuthMiddleware, controllers.CourseTaskOptionCreateController)

// 	// student
// 	incomingRoutes.Get("api/v1/tasks/student_tasks/", middleware.AuthMiddleware, controllers.StudentAllCourseAssigmentListController)
// 	incomingRoutes.Get("api/v1/tasks/student_tasks/:id", middleware.AuthMiddleware, controllers.StudentAnswerTaskDetailController)
// 	incomingRoutes.Post("api/v1/tasks/student_answer_tasks/:id", middleware.AuthMiddleware, controllers.StudentTaskAnswerQuestionController)
// 	incomingRoutes.Post("api/v1/tasks/student_submit_task", middleware.AuthMiddleware, controllers.StudentSubmitTaskController)

// 	incomingRoutes.Get("api/v1/tasks/student_course_tasks/:id", middleware.AuthMiddleware, controllers.StudentCourseAssigmentDetailController)

// }
