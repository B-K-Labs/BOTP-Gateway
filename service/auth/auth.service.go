package auth

import (
	"botp-gateway/common/http_client"

	"github.com/gofiber/fiber/v2"
)

// Api create account
func CreateAccount(c *fiber.Ctx) error {
	return http_client.MakeRequestFromFiberWithSamePath(c)
}

// Api create account
func DashboardSignIn(c *fiber.Ctx) error {
	return http_client.MakeRequestFromFiberWithSamePath(c)
}

func DashboardUserInfo(c *fiber.Ctx) error {
	return http_client.MakeRequestFromFiberWithSamePath(c)
}
