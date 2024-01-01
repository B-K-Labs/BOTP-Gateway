package router

import (
	authRouter "botp-gateway/router/auth"
	userRouter "botp-gateway/router/user"

	"github.com/gofiber/fiber/v2"
)

func New(app *fiber.App) {
	userRouter.CreateRouter(app)
	authRouter.CreateRouter(app)
}
