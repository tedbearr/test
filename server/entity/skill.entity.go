package entity

import "time"

type Skill struct {
	ID uint64 `gorm:"primary_key:auto_increment" json:"id"`
	// UserID int8   `gorm:"not null" json:"-"`
	// User   User   `gorm:"foreignkey:UserID" json:"user"`
	// SkillTypeID uint64 `gorm:"not null" json:"-"`
	Type      string    `gorm:"type:varchar(255)" json:"type"`
	Skill     string    `gorm:"type:varchar(255)" json:"skill"`
	CreatedAt time.Time `gorm:"type:timestamp;default:now()" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:timestamp" json:"updated_at"`
	DeletedAt time.Time `gorm:"type:timestamp" json:"deleted_at"`
}

func (Skill) TableName() string {
	return "skill"
}
