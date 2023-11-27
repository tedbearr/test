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

type SkillService interface {
	All(ctx echo.Context) helper.Response
	Insert(ctx echo.Context) helper.Response
	Find(ctx echo.Context) helper.Response
	Update(ctx echo.Context) helper.Response
	Delete(ctx echo.Context) helper.Response
}

type skillService struct {
	connection repository.SkillRepository
}

func NewSkillService(repo repository.SkillRepository) SkillService {
	return &skillService{
		connection: repo,
	}
}

func (repo *skillService) All(ctx echo.Context) helper.Response {
	result := repo.connection.All()
	return helper.BuildResponse("00", "success", result)
}

func (repo *skillService) Find(ctx echo.Context) helper.Response {
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

func (repo *skillService) Insert(ctx echo.Context) helper.Response {
	SkillInsert := new(dto.SkillInsert)
	ctx.Bind(SkillInsert)

	validate := helper.Validate(SkillInsert)
	if validate != nil {
		res := helper.BuildResponse("500", validate.Error(), helper.EmptyObj{})
		return res
	}

	result := repo.connection.Insert(*SkillInsert)
	if result != nil {
		return helper.BuildResponse("500", result.Error(), helper.EmptyObj{})
	}

	return helper.BuildResponse("00", "success", helper.EmptyObj{})
}

func (repo *skillService) Update(ctx echo.Context) helper.Response {
	SkillUpdate := new(dto.SkillUpdate)

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return helper.BuildResponse("400", err.Error(), helper.EmptyObj{})
	}

	check := repo.connection.Check(id)
	errors.Is(check, gorm.ErrRecordNotFound)
	if check != nil {
		return helper.BuildResponse("400", check.Error(), helper.EmptyObj{})
	}

	SkillUpdate.UpdatedAt = time.Now()

	ctx.Bind(SkillUpdate)

	validate := helper.Validate(SkillUpdate)
	if validate != nil {
		res := helper.BuildResponse("500", validate.Error(), helper.EmptyObj{})
		return res
	}

	result := repo.connection.Update(*SkillUpdate, id)
	if result != nil {
		return helper.BuildResponse("400", result.Error(), helper.EmptyObj{})
	}

	return helper.BuildResponse("00", "success", helper.EmptyObj{})
}

func (repo *skillService) Delete(ctx echo.Context) helper.Response {
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
