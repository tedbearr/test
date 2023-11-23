package controller

import (
	"net/http"
	"server/service"

	"github.com/labstack/echo/v4"
)

type AdminController interface {
	All(echo.Context) error
	Find(echo.Context) error
	Insert(echo.Context) error
	Update(echo.Context) error
	Delete(echo.Context) error
}

type adminController struct {
	connection service.AdminService
}

func NewAdminController(service service.AdminService) AdminController {
	return &adminController{
		connection: service,
	}
}

func (service *adminController) All(ctx echo.Context) error {
	result := service.connection.All(ctx)
	return ctx.JSON(http.StatusOK, result)
}

func (service *adminController) Find(ctx echo.Context) error {
	result := service.connection.Find(ctx)
	return ctx.JSON(http.StatusOK, result)
}

func (service *adminController) Insert(ctx echo.Context) error {
	result := service.connection.Insert(ctx)
	return ctx.JSON(http.StatusOK, result)
}

func (service *adminController) Update(ctx echo.Context) error {
	result := service.connection.Update(ctx)
	return ctx.JSON(http.StatusOK, result)
}

func (service *adminController) Delete(ctx echo.Context) error {
	result := service.connection.Delete(ctx)
	return ctx.JSON(http.StatusOK, result)

}
