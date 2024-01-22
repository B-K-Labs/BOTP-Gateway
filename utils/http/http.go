package http

import (
	"strings"

	"github.com/gofiber/fiber/v2"
)

func GetBearerTokenFromHeader(c *fiber.Ctx) string {
	authHeader := c.Get("Authorization")
	splitToken := strings.Split(authHeader, "Bearer ")
	if len(splitToken) < 2 {
		return ""
	}
	return splitToken[1]
}
