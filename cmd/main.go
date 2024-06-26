package main

import (
	// lmLogger "Learnium/internal/pkg/logger"
	"Learnium/internal/config"
	"Learnium/internal/pkg/common"
	"Learnium/setup"

	// "Learnium/setup"
	"log"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {

	app := fiber.New(fiber.Config{
		BodyLimit: 20 * 1024 * 1024, // Set the body limit to 20MB
	})
	// Use the logger middleware
	app.Use(logger.New())

	// app.Use(cache.New(cache.Config{
	// 	Next: func(c *fiber.Ctx) bool {
	// 		return strings.Contains(c.Route().Path, "/ws")
	// 	},
	// }))

	//app.Use(cors.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
		AllowHeaders:     "Origin,Content-Type,Authorization,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin,LEARNIUM_SK_HEADER",
		AllowCredentials: true,
	}))

	app.Use("/ws", func(c *fiber.Ctx) error {
		// IsWebSocketUpgrade returns true if the client
		// requested upgrade to the WebSocket protocol.
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	var cfg = config.Config

	// Make migrations
	setup.MigrateDatabase()

	// Serve static files from the "uploads" directory
	app.Static("/uploads", "./uploads")

	app.Use(common.CustomHeaderMiddleware())
	app.Use(common.LimitMiddleware)

	// Register user routes
	setup.HttpRoutes(app)
	setup.WebSocketRouters(app)

	// Start the server
	port := cfg.PORT

	log.Printf("Server listening on port %s", port)
	log.Fatal(app.Listen(":" + port))

}
