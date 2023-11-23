package dto

import "time"

type Admin struct {
	ID           uint64    `json:"id" form:"id"`
	Name         string    `json:"name" form:"name"`
	Username     string    `json:"username" form:"username"`
	Password     string    `json:"password" form:"password"`
	RefreshToken string    `json:"refresh_token" form:"refresh_token"`
	CreatedAt    time.Time `json:"created_at" form:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" form:"updated_at"`
	DeletedAt    time.Time `json:"deleted_at" form:"deleted_at"`
}

type AdminInsert struct {
	Name     string `json:"name" form:"name" validate:"required"`
	Username string `json:"username" form:"username" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
}

type AdminUpdate struct {
	Name      string    `json:"name" form:"name" validate:"required"`
	Username  string    `json:"username" form:"username" validate:"required"`
	Password  string    `json:"password" form:"password" validate:"required"`
	UpdatedAt time.Time `json:"updated_at" form:"updated_at" validate:"required"`
}
