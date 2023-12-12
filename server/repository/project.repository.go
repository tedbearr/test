package repository

import (
	"server/dto"

	"gorm.io/gorm"
)

type ProjectRepository interface {
	All() []dto.ProjectAll
	Find(id int) (dto.ProjectAll, error)
	Insert(data dto.ProjectCreate) error
	Update(data dto.ProjectUpdate, id int) error
	Delete(id int) error
}

type projectRepository struct {
	connection *gorm.DB
}

func NewProjectRepository(db *gorm.DB) ProjectRepository {
	return &projectRepository{
		connection: db,
	}
}

func (db *projectRepository) All() []dto.ProjectAll {
	var project []dto.ProjectAll
	db.connection.Table("project").Order("id DESC").Find(&project)
	return project
}

func (db *projectRepository) Find(id int) (dto.ProjectAll, error) {
	var project dto.ProjectAll
	res := db.connection.Table("project").Where("id = ?", id).First(&project)
	return project, res.Error
}

func (db *projectRepository) Insert(data dto.ProjectCreate) error {
	res := db.connection.Table("project").Create(&data).Error
	return res
}

func (db *projectRepository) Update(data dto.ProjectUpdate, id int) error {
	res := db.connection.Table("project").Where("id = ?", id).Updates(&data).Error
	return res
}

func (db *projectRepository) Delete(id int) error {
	var project dto.ProjectAll
	res := db.connection.Table("project").Where("id = ?", id).Delete(&project).Error
	return res
}
