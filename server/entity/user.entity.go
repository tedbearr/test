package entity

type User struct {
	ID       uint64 `gorm:"primary_key:auto_increment" json:"id"`
	Name     string `gorm:"type:varchar(255)" json:"name"`
	Username string `gorm:"uniqueIndex;type:varchar(255)" json:"username"`
	AboutMe  string `gorm:"type:text" json:"about_me"`
	LinkedIn string `gorm:"type:varchar(255)" json:"linked_in"`
	Github   string `gorm:"type:varchar(255)" json:"github"`
}

func (User) TableName() string {
	return "users"
}
