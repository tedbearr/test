package entity

type Skill struct {
	ID uint64 `gorm:"primary_key:auto_increment" json:"id"`
	// UserID int8   `gorm:"not null" json:"-"`
	// User   User   `gorm:"foreignkey:UserID" json:"user"`
	// SkillTypeID uint64 `gorm:"not null" json:"-"`
	Type  string `gorm:"type:varchar(255)" json:"type"`
	Skill string `gorm:"type:varchar(255)" json:"skill"`
}
