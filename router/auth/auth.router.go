package userRouter

import (
	authService "botp-gateway/service/auth"

	"github.com/gofiber/fiber/v2"
)

func CreateRouter(app *fiber.App) {
	r := app.Group("/authen")
	{
		r.Post("/createAccount", authService.CreateAccount)
	}
}
