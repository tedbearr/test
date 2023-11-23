package entity

import "time"

type Experience struct {
	ID          uint64    `gorm:"primary_key:auto_increment" json:"id"`
	Name        string    `gorm:"type:varchar(255)" json:"name"`
	Description string    `gorm:"type:varchar(255)" json:"description"`
	Date        string    `gorm:"type:date" json:"date"`
	CreatedAt   time.Time `gorm:"type:timestamp;default:now()" json:"created_at"`
	UpdatedAt   time.Time `gorm:"type:timestamp" json:"updated_at"`
	DeletedAt   time.Time `gorm:"type:timestamp" json:"deleted_at"`
}

func (Experience) TableName() string {
	return "experience"
}
