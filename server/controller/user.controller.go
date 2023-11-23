package controller

import (
	"net/http"
	"server/service"

	"github.com/labstack/echo/v4"
)

type UserController interface {
	All(echo.Context) error
	Find(echo.Context) error
	Insert(echo.Context) error
	Update(echo.Context) error
	Delete(echo.Context) error
}

type userController struct {
	connection service.UserService
}

func NewUserController(service service.UserService) UserController {
	return &userController{
		connection: service,
	}
}

func (service *userController) All(ctx echo.Context) error {
	result := service.connection.All(ctx)
	return ctx.JSON(http.StatusOK, result)
}

func (service *userController) Find(ctx echo.Context) error {
	result := service.connection.Find(ctx)
	return ctx.JSON(http.StatusOK, result)
}

func (service *userController) Insert(ctx echo.Context) error {
	result := service.connection.Insert(ctx)
	return ctx.JSON(http.StatusOK, result)
}

func (service *userController) Update(ctx echo.Context) error {
	result := service.connection.Update(ctx)
	return ctx.JSON(http.StatusOK, result)
}

func (service *userController) Delete(ctx echo.Context) error {
	result := service.connection.Delete(ctx)
	return ctx.JSON(http.StatusOK, result)

}
