package routes

// import (
// 	"Learnium/controllers"
// 	"Learnium/middleware"
// 	"github.com/gofiber/fiber/v2"
// )

// func CourseRouters(incomingRoutes *fiber.App) {
// 	//	Course
// 	incomingRoutes.Post("api/v1/courses/course_create", middleware.AuthMiddleware, controllers.AdminCreateCourseController)
// 	incomingRoutes.Get("api/v1/courses/course_admin_list", middleware.AuthMiddleware, controllers.AdminCourseListAllController)
// 	incomingRoutes.Get("api/v1/courses/course_list/:id", middleware.AuthMiddleware, controllers.StaffCourseDetailController)
// 	incomingRoutes.Get("api/v1/courses/course_analytics/:id", middleware.AuthMiddleware, controllers.StaffCourseAnalyticsController)

// 	// curriculum
// 	incomingRoutes.Post("api/v1/courses/curriculum_create", middleware.AuthMiddleware, controllers.StaffCurriculumCreateController)
// 	incomingRoutes.Get("api/v1/courses/staff_curriculum_list/:id", middleware.AuthMiddleware, controllers.StaffCurriculumListController)
// 	incomingRoutes.Get("api/v1/courses/curriculum_detail/:id", middleware.AuthMiddleware, controllers.StaffCurriculumDetailController)

// 	// lesson
// 	incomingRoutes.Post("api/v1/courses/lesson_create", middleware.AuthMiddleware, controllers.StaffLessonCreateController)
// 	incomingRoutes.Put("api/v1/courses/lesson_update/:id", middleware.AuthMiddleware, controllers.StaffLessonUpdateController)
// 	incomingRoutes.Get("api/v1/courses/lesson_list/:id", middleware.AuthMiddleware, controllers.StaffLessonListController)

// 	// enrolled course
// 	incomingRoutes.Post("api/v1/courses/enroll_course_create", middleware.AuthMiddleware, controllers.StaffEnrollCourseCreateController)
// 	incomingRoutes.Get("api/v1/courses/list_enroll_student/:id", middleware.AuthMiddleware, controllers.StaffListEnrolledStudentsInCourseController)
// 	incomingRoutes.Post("api/v1/courses/staff_register_attendance", middleware.AuthMiddleware, controllers.StaffRegisterStudentAttendanceController)
// 	incomingRoutes.Get("api/v1/courses/enroll_course_list", middleware.AuthMiddleware, controllers.StudentEnrollCourseListController)

// 	// staff
// 	incomingRoutes.Get("api/v1/courses/staff_student_enrolled/:id", middleware.AuthMiddleware, controllers.StaffOrAdminCourseEnrolledStudentListController)
// 	incomingRoutes.Get("api/v1/courses/staff_course_list/:id", middleware.AuthMiddleware, controllers.StaffCoursesListController)

// 	// schedule
// 	incomingRoutes.Post("api/v1/courses/course_schedule_create", middleware.AuthMiddleware, controllers.CourseScheduleCreateController)
// 	incomingRoutes.Get("api/v1/courses/staff_course_schedules/:id", middleware.AuthMiddleware, controllers.StaffScheduleListController)
// }
