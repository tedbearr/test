package dto

import "time"

type Auth struct {
	ID           int       `json:"id" form:"id"`
	Name         string    `json:"name" form:"name"`
	Username     string    `json:"username" form:"username"`
	Password     string    `json:"password" form:"password"`
	RefreshToken string    `json:"refresh_token" form:"refresh_token"`
	CreatedAt    time.Time `json:"created_at" form:"created_at"`
}

type Register struct {
	Name     string `json:"name" form:"name" validate:"required"`
	Username string `json:"username" form:"username" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
}

// Login model info
// @Description Response login
type Login struct {
	Username string `json:"username" form:"username" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
}

// ResponseToken model info
// @Description Response login
type ResponseToken struct {
	AccessToken  string `json:"accessToken" form:"accessToken"`
	RefreshToken string `json:"refreshToken" form:"refreshToken"`
}

type VerifyAccessToken struct {
	AccessToken string `json:"accessToken" form:"accessToken"`
}
