package userRouter

import (
	authService "botp-gateway/service/auth"

	"github.com/gofiber/fiber/v2"
)

func CreateRouter(app *fiber.App) {
	authRouter := app.Group("/authen")
	{
		authRouter.Post("/createAccount", authService.CreateAccount)
	}
	dashBoardRouter := app.Group("/dashboardAuthen")
	{
		dashBoardRouter.Post("/signIn", authService.DashboardSignIn)
		dashBoardRouter.Get("/info", authService.DashboardUserInfo)
	}
}
