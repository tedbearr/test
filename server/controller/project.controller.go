package controller

import (
	"net/http"
	"server/service"

	"github.com/labstack/echo/v4"
)

type ProjectController interface {
	All(ctx echo.Context) error
	Find(ctx echo.Context) error
	Insert(ctx echo.Context) error
}

type projectController struct {
	connection service.ProjectService
}

func NewProjectController(serv service.ProjectService) ProjectController {
	return &projectController{
		connection: serv,
	}
}

func (service *projectController) All(ctx echo.Context) error {
	res := service.connection.All(ctx)
	return ctx.JSON(http.StatusOK, res)
}

func (service *projectController) Find(ctx echo.Context) error {
	res := service.connection.Find(ctx)
	return ctx.JSON(http.StatusOK, res)
}

func (service *projectController) Insert(ctx echo.Context) error {
	res := service.connection.Insert(ctx)
	return ctx.JSON(http.StatusOK, res)
}
