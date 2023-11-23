package entity

import "time"

type Admin struct {
	ID           uint64    `gorm:"primary_key:auto_increment" json:"id"`
	Name         string    `gorm:"type:varchar(255)" json:"name"`
	Username     string    `gorm:"type:varchar(255)" json:"username"`
	Password     string    `gorm:"type:varchar(255)" json:"password"`
	RefreshToken string    `gorm:"type:varchar(255)" json:"refresh_token"`
	CreatedAt    time.Time `gorm:"type:timestamp;default:now()" json:"created_at"`
	UpdatedAt    time.Time `gorm:"type:timestamp" json:"updated_at"`
	DeletedAt    time.Time `gorm:"type:timestamp" json:"deleted_at"`
}

func (Admin) TableName() string {
	return "admin"
}
