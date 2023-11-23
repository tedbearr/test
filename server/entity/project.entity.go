package entity

type Project struct {
	ID    uint64 `gorm:"primary_key:auto_increment" json:"id"`
	Name  string `gorm:"type:varchar(255)" json:"name"`
	Image string `gorm:"type:varchar(255)" json:"image"`
	Link  string `gorm:"type:varchar(255)" json:"link"`
	Tech  string `gorm:"type:varchar(255)" json:"tech"`
}
