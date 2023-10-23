package middleware

import (
	"botp-gateway/gorm"
	model "botp-gateway/model"
	rsa "botp-gateway/utils/rsa"
	"fmt"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Extract API key, timestamp, nonce, and signature from the request headers.
		apiKeyHeader := c.Get("X-Botp-Api-Key")
		timestampHeader := c.Get("X-Botp-Api-Timestamp")
		nonceHeader := c.Get("X-Botp-Api-Nonce")
		signature := c.Get("X-Botp-Api-Signature")

		apiKey := model.ApiKeys{
			ApiKey: apiKeyHeader,
		}

		err := gorm.DB.First(&apiKey).Error
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized",
			})
		}
		timestampInt, _ := strconv.ParseInt(timestampHeader, 10, 64)

		currentTimestamp := time.Now().Unix()
		timeWindow := int64(300) // 5-minute time window

		if currentTimestamp-timestampInt > timeWindow {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized",
			})
		}

		// Retrieve the request body using Fiber's built-in parsing middleware.
		requestBody := string(c.Request().Body())

		// You can include other request data as needed for signature verification.
		message := []byte(fmt.Sprintf("%s%s%s%s%s%s", apiKey.ApiKey, timestampHeader, nonceHeader, c.Path(), c.Method(), requestBody))
		if !rsa.VerifyRSASignature(apiKey.PublicKeyPEM, signature, message) {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid signature",
			})
		}

		c.Locals("ClientID", apiKey.ClientID)
		// Authentication passed; continue with the next middleware or route handler.
		return c.Next()
	}
}
