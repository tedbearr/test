package controller

import (
	"net/http"
	"server/dto"
	"server/helper"
	"server/service"
	"sync"

	"github.com/labstack/echo/v4"
	"golang.org/x/exp/slog"
)

type AuthController interface {
	Login(context echo.Context) error
	Register(context echo.Context) error
}

type authController struct {
	authService service.AuthService
}

func NewAuthController(service service.AuthService) AuthController {
	return &authController{
		authService: service,
	}
}

func (service *authController) Login(ctx echo.Context) error {
	var wg sync.WaitGroup
	var User dto.Login
	// a := context.Request()
	// fmt.Println(a)
	uniqueCode := helper.UniqueCode()

	ctx.Bind(&User)

	// if err := context.BodyParser(&User); err != nil {
	// 	res := helper.BuildResponse("400", err.Error(), helper.EmptyObj{})
	// 	slog.Info(uniqueCode+" Login response ", res)
	// 	return ctx.JSON(http.StatusOK, res)
	// }

	slog.Info(uniqueCode+" Login request ", User)

	// validate := helper.Validate(User)
	// if validate != nil {
	// 	res := helper.BuildResponse("500", validate.Error(), helper.EmptyObj{})
	// 	slog.Info(uniqueCode+" Login response ", res)
	// 	return ctx.JSON(http.StatusOK, res)
	// }
	wg.Add(1)
	result, err := service.authService.Login(User, uniqueCode, &wg)
	wg.Wait()
	if err != nil {
		res := helper.BuildResponse("400", err.Error(), helper.EmptyObj{})
		slog.Info(uniqueCode+" Login response ", res)
		return ctx.JSON(http.StatusOK, res)
	}

	res := helper.BuildResponse("00", "success", result)
	slog.Info(uniqueCode+" Login response ", res)
	return ctx.JSON(http.StatusOK, res)
}

func (service *authController) Register(ctx echo.Context) error {
	result, err := service.authService.Register(ctx)
	if err != nil {
		return ctx.JSON(http.StatusOK, err.Error())
	} else {
		return ctx.JSON(http.StatusOK, result)
	}

}
