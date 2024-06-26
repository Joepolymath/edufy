package setup

import (
	"Learnium/internal/pkg/adapters"
	"Learnium/internal/pkg/api/auth"
	"Learnium/internal/pkg/api/school"
	"Learnium/internal/pkg/api/subscription"

	// "Learnium/internal/pkg/logger"
	// "context"

	// "github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	// "go.uber.org/zap"
)

/*This contains all the routes on the user-services Combined
 */

func HttpRoutes(app *fiber.App) {

	auth.Router(app)         // mount auth routes
	subscription.Router(app) // mount subscription routes
	school.Router(app)       // mount school routes

}

func WebSocketRouters(app *fiber.App) {
	// websocket Config
	// websocketConfig := websocket.Config{
	// 	RecoverHandler: func(conn *websocket.Conn) {
	// 		if err := recover(); err != nil {
	// 			err := conn.WriteJSON(fiber.Map{"customError": "error occurred"})
	// 			if err != nil {
	// 				logger.Error(context.Background(), "error making recover handler  in websocket routes,", zap.Error(err))
	// 				return
	// 			}
	// 		}
	// 	},
	// }

	// Create a new instance of the Hub
	hub := adapters.NewHub()
	go hub.Run()

	// ChatWebSocketRouters(app, websocketConfig, hub)
}
