package route

import "github.com/labstack/echo/v4"

func AuthRoute(echo *echo.Group) {
	r := echo.Group("auth")
	r.POST("/login", authController.Login)
	r.POST("/register", authController.Register)
}
