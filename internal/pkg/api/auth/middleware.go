package auth

import (
	// "Learnium/internal/pkg/models"
	"context"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

type AuthMiddleware struct {
	srv IAuthService
}

func NewAuthMiddleware(srv IAuthService) *AuthMiddleware {
	return &AuthMiddleware{
		srv,
	}
}

func (md *AuthMiddleware) GenericGuard() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Check if the user is authenticated
		// You can use any authentication mechanism you prefer, such as JWT or session cookies
		// For example, you can check if the "Authorization" header contains a valid JWT token
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Authorization Bearer was not passed"})
		}
		// Verify the JWT token and return true if it's valid
		clientToken := strings.TrimPrefix(authHeader, "Bearer ")
		if clientToken == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Authorization Bearer was not passed"})
		}

		// get the claims and the token of the user with the token
		claims, err := md.srv.ValidateJwtToken(clientToken)
		if err != "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "token not valid"})
		}

		user, err2 := md.srv.GetUser(ctx, claims.ID)
		if err2 != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Error getting user with claims . Please try  logging in again."})
		}
		// Check if the user  is verified
		if user.EmailVerified == false {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "User is not verified. Please verify email"})

		}

		// Set the user in the Locals
		c.Locals("user", user)

		return c.Next()

	}
}

func (md *AuthMiddleware) SuperAdminGuard() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Check if the user is authenticated
		// You can use any authentication mechanism you prefer, such as JWT or session cookies
		// For example, you can check if the "Authorization" header contains a valid JWT token
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Authorization Bearer was not passed"})
		}
		// Verify the JWT token and return true if it's valid
		clientToken := strings.TrimPrefix(authHeader, "Bearer ")
		if clientToken == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Authorization Bearer was not passed"})
		}

		// get the claims and the token of the user with the token
		claims, err := md.srv.ValidateJwtToken(clientToken)
		if err != "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "token not valid"})
		}

		user, err2 := md.srv.GetUser(ctx, claims.ID)
		if err2 != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Error getting user with claims . Please try  logging in again."})
		}
		// Check if the user  is verified
		if !user.EmailVerified {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "User is not verified. Please verify email"})
		} else if !user.IsSuperAdmin {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "User has no permission to use this service"})
		}

		// Set the user ID in the Locals
		c.Locals("user", user)

		return c.Next()

	}
}
