package route

import "github.com/labstack/echo/v4"

func AdminRoute(echo *echo.Group) {
	r := echo.Group("admin")
	r.GET("/", adminController.All)
	r.GET("/:id", adminController.Find)
	r.POST("/insert", adminController.Insert)
	r.POST("/update/:id", adminController.Update)
	r.POST("delete", adminController.Delete)
}
