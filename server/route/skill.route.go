package route

import "github.com/labstack/echo/v4"

func SkillRoute(echo *echo.Group) {
	r := echo.Group("skill")
	r.GET("/", skillController.All)
	r.GET("/:id", skillController.Find)
	r.POST("/insert", skillController.Insert)
	r.POST("/update/:id", skillController.Update)
	r.POST("delete", skillController.Delete)
}
