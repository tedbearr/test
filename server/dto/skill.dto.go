package dto

import "time"

type Skill struct {
	ID        uint64    `json:"id" form:"id"`
	Name      string    `json:"name" form:"name"`
	Skill     string    `gorm:"type:varchar(255)" json:"skill"`
	CreatedAt time.Time `json:"created_at" form:"created_at"`
	UpdatedAt time.Time `json:"updated_at" form:"updated_at"`
	DeletedAt time.Time `json:"deleted_at" form:"deleted_at"`
}

type SkillInsert struct {
	Name  string `json:"name" form:"name" validate:"required"`
	Skill string `json:"skill" form:"skill" validate:"required"`
}

type SkillUpdate struct {
	Name      string    `json:"name" form:"name" validate:"required"`
	Skill     string    `json:"skill" form:"skill" validate:"required"`
	UpdatedAt time.Time `json:"updated_at" form:"updated_at" validate:"required"`
}
