package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	Server ServerConfig `mapstructure:"server"`
	Mysql  MysqlConfig  `mapstructure:"mysql"`
	Jwt    JwtConfig    `mapstructure:"jwt"`
}

type ServerConfig struct {
	Port string `mapstructure:"port"`
}

type MysqlConfig struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"dbname"`
}

type JwtConfig struct {
	Secret string `mapstructure:"secret"`
	Expire int    `mapstructure:"expire"`
}

var AppConfig Config

func LoadConfig() error {
	// 设置 viper
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	// 尝试多种路径查找配置文件
	// 1. 当前目录
	viper.AddConfigPath(".")
	// 2. config 目录
	viper.AddConfigPath("personal_blog2/config")
	//// 3. 上级目录
	//viper.AddConfigPath("../")
	//// 4. 上级目录的 config 子目录
	//viper.AddConfigPath("../config")

	viper.AutomaticEnv()

	// 读取配置文件
	err := viper.ReadInConfig()
	if err != nil {
		// 打印当前工作目录和尝试的配置路径，方便调试
		cwd, _ := os.Getwd()
		return fmt.Errorf("读取配置文件失败（当前工作目录：%s）：%w", cwd, err)
	}

	// 解析配置赋值到config结构体
	err = viper.Unmarshal(&AppConfig)
	if err != nil {
		return fmt.Errorf("解析配置失败：%w", err)
	}

	return nil
}
