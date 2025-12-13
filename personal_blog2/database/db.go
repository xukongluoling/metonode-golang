package database

import (
	"fmt"
	"log"
	"metonode-golang/personal_blog2/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var MysqlDB *gorm.DB

func InitMysqlDB() {
	var err error

	mysqlConfig := config.AppConfig.Mysql
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		mysqlConfig.User,
		mysqlConfig.Password,
		mysqlConfig.Host,
		mysqlConfig.Port,
		mysqlConfig.DBName)

	MysqlDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("mysql connect fail: %v", err)
		return
	}
	log.Printf("mysql connect success")

	// 创建或更新表

}
