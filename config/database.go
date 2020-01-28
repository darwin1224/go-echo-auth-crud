package config

import (
	"github.com/darwin1224/go-echo-auth-crud/models"
	"github.com/jinzhu/gorm"
)

func InitDB() *gorm.DB {
	db, err := gorm.Open("mysql", "root:rahasia@/go_test")
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(new(models.User))
	return db
}
