package rateLimitMiddleware

import (
	"botp-gateway/common/constants"
	"botp-gateway/utils/http"
	jsonWebToken "botp-gateway/utils/jwt"

	"botp-gateway/validator"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func RateLimit() fiber.Handler {
	return limiter.New(limiter.Config{
		Max:        30,
		Expiration: constants.RATE_LIMITER_TIMEOUT * time.Second,
		KeyGenerator: func(c *fiber.Ctx) string {
			reqToken := http.GetBearerTokenFromHeader(c)
			if len(reqToken) > 0 {
				id, err := jsonWebToken.ParseTokenToUserId(reqToken)
				if (err == nil) && (id != "") {
					return id
				}
			}
			return c.IP()
		},
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"success": false,
				"data":    "Too many requests",
			})
		},
	})
}

func RateLimitByEmail() fiber.Handler {
	return limiter.New(limiter.Config{
		Max:        1,
		Expiration: constants.RATE_LIMITER_TIMEOUT * time.Second,
		KeyGenerator: func(c *fiber.Ctx) string {

			payload := validator.WithEmail{}

			if err := c.BodyParser(&payload); err != nil {
				return ""
			}

			return payload.Email
		},
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"success": false,
				"data":    "Too many requests",
			})
		},
	})
}
