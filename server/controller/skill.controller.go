package controller

import (
	"net/http"
	"server/service"

	"github.com/labstack/echo/v4"
)

type SkillController interface {
	All(echo.Context) error
	Find(echo.Context) error
	Insert(echo.Context) error
	Update(echo.Context) error
	Delete(echo.Context) error
}

type skillController struct {
	connection service.SkillService
}

func NewSkillController(service service.SkillService) SkillController {
	return &skillController{
		connection: service,
	}
}

func (service *skillController) All(ctx echo.Context) error {
	result := service.connection.All(ctx)
	return ctx.JSON(http.StatusOK, result)
}

func (service *skillController) Find(ctx echo.Context) error {
	result := service.connection.Find(ctx)
	return ctx.JSON(http.StatusOK, result)
}

func (service *skillController) Insert(ctx echo.Context) error {
	result := service.connection.Insert(ctx)
	return ctx.JSON(http.StatusOK, result)
}

func (service *skillController) Update(ctx echo.Context) error {
	result := service.connection.Update(ctx)
	return ctx.JSON(http.StatusOK, result)
}

func (service *skillController) Delete(ctx echo.Context) error {
	result := service.connection.Delete(ctx)
	return ctx.JSON(http.StatusOK, result)

}
