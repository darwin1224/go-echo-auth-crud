package models

type (
	User struct {
		ID       uint   `gorm:"primary_key"`
		Name     string `gorm:"column:name;type:varchar(50);not null" json:"name"`
		Username string `gorm:"column:username;type:varchar(100);not null" json:"username"`
		Password string `gorm:"column:password;type:text;not null" json:"password"`
		RoleUser string `gorm:"column:role_user;type:varchar(30);not null" json:"role_user"`
	}
)
