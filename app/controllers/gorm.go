package controllers

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/masato-kataoka/studygo/app/models"
)

var DB *gorm.DB

func InitDB() {
	db, err := gorm.Open("mysql", "root:@/studygo2?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		fmt.Println(err)
	}

	db.DB()
	db.AutoMigrate(&models.Article{})

	DB = db
}
