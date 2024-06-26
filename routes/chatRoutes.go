package routes

// import (
// 	"Learnium/adapters"
// 	"Learnium/controllers"
// 	"Learnium/middleware"
// 	"github.com/gofiber/contrib/websocket"
// 	"github.com/gofiber/fiber/v2"
// )

// func ChatWebSocketRouters(incomingRoutes *fiber.App, wsConfig websocket.Config, hub adapters.HubInterface) {

// 	// WebSocket route for conversation
// 	incomingRoutes.Get("/ws/user/", middleware.AuthMiddleware, websocket.New(func(c *websocket.Conn) {
// 		controllers.ConversationUserListWebSocket(c, hub)
// 	}, wsConfig))
// 	incomingRoutes.Get("/ws/message/:name", middleware.AuthMiddleware, websocket.New(func(c *websocket.Conn) {
// 		controllers.ConversationMessageUserWebsocket(c, hub)
// 	}, wsConfig))

// 	incomingRoutes.Get("/ws/user/notifications", middleware.AuthMiddleware, websocket.New(func(c *websocket.Conn) {
// 		controllers.ConversationUserMessageNotificationWebsocket(c, hub)
// 	}, wsConfig))

// }
