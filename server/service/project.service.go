package service

import (
	"fmt"
	"io"
	"os"
	"server/dto"
	"server/helper"
	"server/repository"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ProjectService interface {
	All(ctx echo.Context) helper.Response
	Find(ctx echo.Context) helper.Response
	Insert(ctx echo.Context) helper.Response
}

type projectService struct {
	connection repository.ProjectRepository
}

func NewProjectService(repo repository.ProjectRepository) ProjectService {
	return &projectService{
		connection: repo,
	}
}

func (repository *projectService) All(ctx echo.Context) helper.Response {
	result := repository.connection.All()
	return helper.BuildResponse("00", "success", result)

}

func (repository *projectService) Find(ctx echo.Context) helper.Response {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		return helper.BuildResponse("400", err.Error(), helper.EmptyObj{})
	}

	res, errFind := repository.connection.Find(id)
	if errFind != nil {
		return helper.BuildResponse("400", errFind.Error(), helper.EmptyObj{})
	}

	return helper.BuildResponse("00", "success", res)
}

func (repo *projectService) Insert(ctx echo.Context) helper.Response {
	projectInsert := new(dto.ProjectCreate)
	ctx.Bind(projectInsert)
	fmt.Println(projectInsert)
	file, err := ctx.FormFile("file")
	fmt.Println(file)
	if err != nil {
		fmt.Println("qwqw", err.Error())
		return helper.BuildResponse("500", err.Error(), helper.EmptyObj{})
	}

	src, err := file.Open()
	if err != nil {
		return helper.BuildResponse("500", err.Error(), helper.EmptyObj{})
	}
	defer src.Close()

	if err := os.Mkdir("upload/image", 0777); err != nil {
		return helper.BuildResponse("500", err.Error(), helper.EmptyObj{})
	}

	dst, err := os.Create("upload/image/" + file.Filename)
	if err != nil {
		return helper.BuildResponse("500", err.Error(), helper.EmptyObj{})
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return helper.BuildResponse("500", err.Error(), helper.EmptyObj{})
	}

	projectInsert.Image = file.Filename

	validate := helper.Validate(projectInsert)
	if validate != nil {
		res := helper.BuildResponse("500", validate.Error(), helper.EmptyObj{})
		return res
	}

	result := repo.connection.Insert(*projectInsert)
	if result != nil {
		return helper.BuildResponse("500", result.Error(), helper.EmptyObj{})
	}

	return helper.BuildResponse("00", "success", helper.EmptyObj{})
}
