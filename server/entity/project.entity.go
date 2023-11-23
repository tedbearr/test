package entity

import "time"

type Project struct {
	ID        uint64    `gorm:"primary_key:auto_increment" json:"id"`
	Name      string    `gorm:"type:varchar(255)" json:"name"`
	Image     string    `gorm:"type:varchar(255)" json:"image"`
	Link      string    `gorm:"type:varchar(255)" json:"link"`
	Tech      string    `gorm:"type:varchar(255)" json:"tech"`
	CreatedAt time.Time `gorm:"type:timestamp;default:now()" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:timestamp" json:"updated_at"`
	DeletedAt time.Time `gorm:"type:timestamp" json:"deleted_at"`
}

func (Project) TableName() string {
	return "project"
}
