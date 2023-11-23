package entity

type Experience struct {
	ID          uint64 `gorm:"primary_key:auto_increment" json:"id"`
	Name        string `gorm:"type:varchar(255)" json:"name"`
	Description string `gorm:"type:varchar(255)" json:"description"`
	Date        string `gorm:"type:date" json:"date"`
}
