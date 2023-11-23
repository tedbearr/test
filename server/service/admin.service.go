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

type AdminService interface {
	All(ctx echo.Context) helper.Response
	Insert(ctx echo.Context) helper.Response
	Find(ctx echo.Context) helper.Response
	Update(ctx echo.Context) helper.Response
	Delete(ctx echo.Context) helper.Response
}

type adminService struct {
	connection repository.AdminRepositoy
}

func NewAdminService(repo repository.AdminRepositoy) AdminService {
	return &adminService{
		connection: repo,
	}
}

func (repo *adminService) All(ctx echo.Context) helper.Response {
	result := repo.connection.All()
	return helper.BuildResponse("00", "success", result)
}

func (repo *adminService) Find(ctx echo.Context) helper.Response {
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

func (repo *adminService) Insert(ctx echo.Context) helper.Response {
	adminInsert := new(dto.AdminInsert)
	ctx.Bind(adminInsert)

	validate := helper.Validate(adminInsert)
	if validate != nil {
		res := helper.BuildResponse("500", validate.Error(), helper.EmptyObj{})
		return res
	}

	result := repo.connection.Insert(*adminInsert)
	if result != nil {
		return helper.BuildResponse("500", result.Error(), helper.EmptyObj{})
	}

	return helper.BuildResponse("00", "success", helper.EmptyObj{})
}

func (repo *adminService) Update(ctx echo.Context) helper.Response {
	adminUpdate := new(dto.AdminUpdate)

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return helper.BuildResponse("400", err.Error(), helper.EmptyObj{})
	}

	check := repo.connection.Check(id)
	errors.Is(check, gorm.ErrRecordNotFound)
	if check != nil {
		return helper.BuildResponse("400", check.Error(), helper.EmptyObj{})
	}

	adminUpdate.UpdatedAt = time.Now()

	ctx.Bind(adminUpdate)

	validate := helper.Validate(adminUpdate)
	if validate != nil {
		res := helper.BuildResponse("500", validate.Error(), helper.EmptyObj{})
		return res
	}

	result := repo.connection.Update(*adminUpdate, id)
	if result != nil {
		return helper.BuildResponse("400", result.Error(), helper.EmptyObj{})
	}

	return helper.BuildResponse("00", "success", helper.EmptyObj{})
}

func (repo *adminService) Delete(ctx echo.Context) helper.Response {
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
