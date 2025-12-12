package database

import (
	"fmt"
	"log"
	"metonode-golang/personal_blog/config"
	"metonode-golang/personal_blog/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var MySqlDB *gorm.DB

func InitMySqlDB() {
	// 从配置文件获取MySQL连接信息
	mysqlConfig := config.AppConfig.MySQL
	
	// 构建DSN
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%t&loc=%s",
		mysqlConfig.User,
		mysqlConfig.Password,
		mysqlConfig.Host,
		mysqlConfig.Port,
		mysqlConfig.DBName,
		mysqlConfig.Charset,
		mysqlConfig.ParseTime,
		mysqlConfig.Loc,
	)
	
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
