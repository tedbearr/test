package repository

import (
	"server/dto"

	"gorm.io/gorm"
)

type SkillRepository interface {
	All() []dto.Skill
	Find(id int) (dto.Skill, error)
	Insert(skill dto.SkillInsert) error
	Update(skill dto.SkillUpdate, id int) error
	Delete(id int) error
	Check(id int) error
}

type skillRepository struct {
	connection *gorm.DB
}

func NewSkillRepository(db *gorm.DB) SkillRepository {
	return &skillRepository{
		connection: db,
	}
}

func (db *skillRepository) All() []dto.Skill {
	var skill []dto.Skill
	db.connection.Table("skill").Order("id DESC").Find(&skill)
	return skill
}

func (db *skillRepository) Find(id int) (dto.Skill, error) {
	var skill dto.Skill
	res := db.connection.Table("skill").Where("id = ?", id).First(&skill).Error
	return skill, res
}

func (db *skillRepository) Insert(skill dto.SkillInsert) error {
	res := db.connection.Table("skill").Create(&skill).Error
	return res
}

func (db *skillRepository) Update(skill dto.SkillUpdate, id int) error {
	res := db.connection.Table("skill").Where("id = ?", id).Updates(&skill).Error
	return res
}

func (db *skillRepository) Delete(id int) error {
	var skill dto.Skill
	res := db.connection.Table("skill").Where("id = ?", id).Delete(&skill).Error
	return res
}

func (db *skillRepository) Check(id int) error {
	var skill dto.Skill
	res := db.connection.Table("skill").Where("id = ?", id).First(&skill).Error
	return res
}
