package database

import (
	"log"
	"metonode-golang/personal_blog/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var MySqlDB *gorm.DB

func InitMySqlDB() {
	dsn := "root:Mysql5735@tcp(127.0.0.1:13306)/personal_blog_task?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	MySqlDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("mysql connect fail :", err)
		return
	}

	// 自动创建或更新表
	if err := MySqlDB.AutoMigrate(&models.User{}, &models.Post{}, &models.Comment{}); err != nil {
		log.Fatal("mysql migrate fail :", err)
	}
}
