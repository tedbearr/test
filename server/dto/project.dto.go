package dto

import "time"

type ProjectAll struct {
	ID        uint64    `gorm:"primary_key:auto_increment" json:"id"`
	Name      string    `gorm:"type:varchar(255)" json:"name"`
	Image     string    `gorm:"type:varchar(255)" json:"image"`
	Link      string    `gorm:"type:varchar(255)" json:"link"`
	Tech      string    `gorm:"type:varchar(255)" json:"tech"`
	CreatedAt time.Time `gorm:"type:timestamp;default:now()" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:timestamp" json:"updated_at"`
	DeletedAt time.Time `gorm:"type:timestamp" json:"deleted_at"`
}

type ProjectCreate struct {
	Name  string `json:"name" form:"name" validate:"required"`
	Image string `json:"image" form:"image" validate:"required"`
	Link  string `json:"link" form:"link" validate:"required"`
	Tech  string `json:"tech" form:"tech" validate:"required"`
}

type ProjectUpdate struct {
	Name      string    `json:"name" form:"name" validate:"required"`
	Image     string    `json:"image" form:"image" validate:"required"`
	Link      string    `json:"link" form:"link" validate:"required"`
	Tech      string    `json:"tech" form:"tech" validate:"required"`
	UpdatedAt time.Time `json:"updated_at" form:"updated_at" validate:"required"`
}
