package service

import (
	"errors"
	"server/dto"
	"server/helper"
	"server/repository"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type UserService interface {
	All(ctx echo.Context) helper.Response
	Insert(ctx echo.Context) helper.Response
	Find(ctx echo.Context) helper.Response
	Update(ctx echo.Context) helper.Response
	Delete(ctx echo.Context) helper.Response
}

type userService struct {
	connection repository.UserRepositoy
}

func NewUserService(repo repository.UserRepositoy) UserService {
	return &userService{
		connection: repo,
	}
}

func (repo *userService) All(ctx echo.Context) helper.Response {
	result := repo.connection.All()
	return helper.BuildResponse("00", "success", result)
}

func (repo *userService) Find(ctx echo.Context) helper.Response {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		return helper.BuildResponse("400", err.Error(), helper.EmptyObj{})
	}

	res, errFind := repo.connection.Find(id)
	if errFind != nil {
		return helper.BuildResponse("400", errFind.Error(), helper.EmptyObj{})
	}

	return helper.BuildResponse("00", "success", res)
}

func (repo *userService) Insert(ctx echo.Context) helper.Response {
	userInsert := new(dto.UserInsert)
	ctx.Bind(userInsert)

	validate := helper.Validate(userInsert)
	if validate != nil {
		res := helper.BuildResponse("500", validate.Error(), helper.EmptyObj{})
		return res
	}

	result := repo.connection.Insert(*userInsert)
	if result != nil {
		return helper.BuildResponse("500", result.Error(), helper.EmptyObj{})
	}

	return helper.BuildResponse("00", "success", helper.EmptyObj{})
}

func (repo *userService) Update(ctx echo.Context) helper.Response {
	userUpdate := new(dto.UserUpdate)

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return helper.BuildResponse("400", err.Error(), helper.EmptyObj{})
	}

	check := repo.connection.Check(id)
	errors.Is(check, gorm.ErrRecordNotFound)
	if check != nil {
		return helper.BuildResponse("400", check.Error(), helper.EmptyObj{})
	}

	userUpdate.UpdatedAt = time.Now()

	ctx.Bind(userUpdate)

	validate := helper.Validate(userUpdate)
	if validate != nil {
		res := helper.BuildResponse("500", validate.Error(), helper.EmptyObj{})
		return res
	}

	result := repo.connection.Update(*userUpdate, id)
	if result != nil {
		return helper.BuildResponse("400", result.Error(), helper.EmptyObj{})
	}

	return helper.BuildResponse("00", "success", helper.EmptyObj{})
}

func (repo *userService) Delete(ctx echo.Context) helper.Response {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return helper.BuildResponse("400", err.Error(), helper.EmptyObj{})
	}

	result := repo.connection.Delete(id)
	if result != nil {
		return helper.BuildResponse("00", result.Error(), helper.EmptyObj{})
	}

	return helper.BuildResponse("00", "success", helper.EmptyObj{})
}
