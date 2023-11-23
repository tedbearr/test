package dto

import "time"

type User struct {
	ID        int       `json:"id" form:"id"`
	Name      string    `json:"name" form:"name"`
	AboutMe   string    `json:"about_me" form:"about_me"`
	LinkedIn  string    `json:"linked_in" form:"linked_in"`
	Github    string    `json:"github" form:"github"`
	CreatedAt time.Time `json:"created_at" form:"created_at"`
	UpdatedAt time.Time `json:"updated_at" form:"updated_at"`
	DeletedAt time.Time `json:"deleted_at" form:"deleted_at"`
}

type UserInsert struct {
	Name      string    `json:"name" form:"name" validate:"required"`
	AboutMe   string    `json:"about_me" form:"about_me" validate:"required"`
	LinkedIn  string    `json:"linked_in" form:"linked_in" validate:"required"`
	Github    string    `json:"github" form:"github" validate:"required"`
	CreatedAt time.Time `json:"created_at" form:"created_at"`
}

type UserUpdate struct {
	Name      string    `json:"name" form:"name" validate:"required"`
	AboutMe   string    `json:"about_me" form:"about_me" validate:"required"`
	LinkedIn  string    `json:"linked_in" form:"linked_in" validate:"required"`
	Github    string    `json:"github" form:"github" validate:"required"`
	UpdatedAt time.Time `json:"updated_at" form:"updated_at" validate:"required"`
}
