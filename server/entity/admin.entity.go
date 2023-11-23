package entity

type Admin struct {
	ID       uint64 `gorm:"primary_key:auto_increment" json:"id"`
	Name     string `gorm:"type:varchar(255)" json:"name"`
	Username string `gorm:"type:varchar(255)" json:"username"`
	Password string `gorm:"type:varchar(255)" json:"password"`
}
