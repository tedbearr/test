package route

import "github.com/labstack/echo/v4"

func UserRoute(echo *echo.Group) {
	r := echo.Group("user")
	r.GET("/", userController.All)
	r.GET("/:id", userController.Find)
	r.POST("/insert", userController.Insert)
	r.POST("/update/:id", userController.Update)
	r.POST("delete", userController.Delete)
}
