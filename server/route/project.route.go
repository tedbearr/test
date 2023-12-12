package route

import "github.com/labstack/echo/v4"

func ProjectRoute(echo *echo.Group) {
	r := echo.Group("project")
	r.GET("/", projectController.All)
	r.GET("/:id", projectController.Find)
	r.POST("/insert", projectController.Insert)
	// r.POST("/update/:id", skillController.Update)
	// r.POST("/delete/:id", skillController.Delete)
}
