package myDB

import (
	"log"
	"metonode-golang/task3/gorm_task/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var MyDB *gorm.DB

func initMyDB() error {
	dsn := "root:123456@tcp(127.0.0.1:3306)/gorm_test?charset=utf8mb4&parseTime=True&loc=Local"
	db, er := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if er != nil {
		log.Fatal("mysql connect err:", er)
		return er
	}
	log.Println("mysql connect success")
	if err := db.AutoMigrate(&models.User{}, &models.Post{}, &models.Comment{}); err != nil {
		log.Fatal("AutoMigrate err:", err)
	}
	log.Println("AutoMigrate success")
	MyDB = db
	return nil
}
