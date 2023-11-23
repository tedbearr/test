package repository

import (
	"server/dto"

	"gorm.io/gorm"
)

type UserRepositoy interface {
	All() []dto.User
	Find(id int) (dto.User, error)
	Insert(user dto.UserInsert) error
	Update(user dto.UserUpdate, id int) error
	Delete(id int) error
	Check(id int) error
}

type userRepository struct {
	connection *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepositoy {
	return &userRepository{
		connection: db,
	}
}

func (db *userRepository) All() []dto.User {
	var user []dto.User
	db.connection.Table("users").Find(&user)
	return user
}

func (db *userRepository) Find(id int) (dto.User, error) {
	var user dto.User
	res := db.connection.Table("users").First(&user).Error
	return user, res
}

func (db *userRepository) Insert(user dto.UserInsert) error {
	res := db.connection.Table("users").Create(&user).Error
	return res
}

func (db *userRepository) Update(user dto.UserUpdate, id int) error {
	res := db.connection.Table("users").Where("id = ?", id).Updates(&user).Error
	return res
}

func (db *userRepository) Delete(id int) error {
	var user dto.User
	res := db.connection.Table("users").Where("id = ?", id).Delete(&user).Error
	return res
}

func (db *userRepository) Check(id int) error {
	var user dto.User
	res := db.connection.Table("users").Where("id = ?", id).First(&user).Error
	return res
}
