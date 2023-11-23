package repository

import (
	"server/dto"

	"gorm.io/gorm"
)

type AdminRepositoy interface {
	All() []dto.Admin
	Find(id int) (dto.Admin, error)
	Insert(admin dto.AdminInsert) error
	Update(admin dto.AdminUpdate, id int) error
	Delete(id int) error
	Check(id int) error
}

type adminRepository struct {
	connection *gorm.DB
}

func NewAdminRepository(db *gorm.DB) AdminRepositoy {
	return &adminRepository{
		connection: db,
	}
}

func (db *adminRepository) All() []dto.Admin {
	var admin []dto.Admin
	db.connection.Table("admin").Find(&admin)
	return admin
}

func (db *adminRepository) Find(id int) (dto.Admin, error) {
	var admin dto.Admin
	res := db.connection.Table("admin").First(&admin).Error
	return admin, res
}

func (db *adminRepository) Insert(admin dto.AdminInsert) error {
	res := db.connection.Table("admin").Create(&admin).Error
	return res
}

func (db *adminRepository) Update(admin dto.AdminUpdate, id int) error {
	res := db.connection.Table("admin").Where("id = ?", id).Updates(&admin).Error
	return res
}

func (db *adminRepository) Delete(id int) error {
	var admin dto.Admin
	res := db.connection.Table("admin").Where("id = ?", id).Delete(&admin).Error
	return res
}

func (db *adminRepository) Check(id int) error {
	var admin dto.Admin
	res := db.connection.Table("admin").Where("id = ?", id).First(&admin).Error
	return res
}
