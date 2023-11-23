package repository

import (
	"server/dto"

	"gorm.io/gorm"
)

type AuthRepository interface {
	Insert(authData *dto.Auth) error
	CheckUsername(username string) (dto.Auth, error)
	UpdateRefreshToken(authData *dto.Auth, username string) error
	// CheckEmail(email string) (dto.Auth, error)
}

type authRepository struct {
	connection *gorm.DB
}

func NewAuthRepository(dbConn *gorm.DB) AuthRepository {
	return &authRepository{
		connection: dbConn,
	}
}

func (db *authRepository) Insert(authData *dto.Auth) error {
	err := db.connection.Table("admin").
		Create(&authData).Error
	return err
}

func (db *authRepository) CheckUsername(username string) (dto.Auth, error) {
	var auth dto.Auth
	check := db.connection.Table("admin").
		Where("username = ?", username).
		First(&auth).Error
	return auth, check
}

// func (db *authRepository) CheckEmail(email string) (dto.Auth, error) {
// 	var auth dto.Auth
// 	check := db.connection.Table("admin").
// 		Where("email = ?", email).
// 		First(&auth).Error
// 	return auth, check
// }

func (db *authRepository) UpdateRefreshToken(authData *dto.Auth, username string) error {
	update := db.connection.Table("admin").
		Where("username = ?", username).
		Updates(&authData).Error
	return update
}
