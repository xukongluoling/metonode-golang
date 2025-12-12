package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

// MySQLConfig 定义MySQL配置结构体
type MySQLConfig struct {
	Host      string `mapstructure:"host"`
	Port      string `mapstructure:"port"`
	User      string `mapstructure:"user"`
	Password  string `mapstructure:"password"`
	DBName    string `mapstructure:"dbname"`
	Charset   string `mapstructure:"charset"`
	ParseTime bool   `mapstructure:"parseTime"`
	Loc       string `mapstructure:"loc"`
}

type ServerConfig struct {
	Port string `mapstructure:"port"`
}

// JWTConfig 定义JWT配置结构体
type JWTConfig struct {
	Secret      string `mapstructure:"secret"`
	ExpireHours int    `mapstructure:"expireHours"`
}

// Config 定义全局配置结构体
type Config struct {
	Server ServerConfig `mapstructure:"server"`
	MySQL  MySQLConfig  `mapstructure:"mysql"`
	JWT    JWTConfig    `mapstructure:"jwt"`
}

// 全局配置变量
var AppConfig Config

// LoadConfig 加载配置文件
func LoadConfig() error {
	// 设置配置文件路径
	configPath := "config"

	// 检查配置文件是否存在
	_, err := os.Stat(configPath)
	if os.IsNotExist(err) {
		// 尝试从当前目录的上一级查找
		configPath = "./config"
		_, err = os.Stat(configPath)
		if os.IsNotExist(err) {
			return fmt.Errorf("配置文件目录不存在")
		}
	}

	// 设置viper配置
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(configPath)
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	// 读取配置文件
	err = viper.ReadInConfig()
	if err != nil {
		return fmt.Errorf("读取配置文件失败: %w", err)
	}

	// 解析配置到结构体
	err = viper.Unmarshal(&AppConfig)
	if err != nil {
		return fmt.Errorf("解析配置失败: %w", err)
	}

	return nil
}
