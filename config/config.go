package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
}

type AppConfig struct {
	Name string `json:"name"`
	Mode string `json:"mode"`
	Port string `json:"port"`
}

type DatabaseConfig struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Database string `json:"database"`
}

var appConfig AppConfig
var databaseConfig DatabaseConfig

func InitViper() {
	viper.SetConfigName("dev")
	viper.SetConfigType("toml")
	viper.AddConfigPath("./config")

	// 读取配置文件
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w \n", err))
	}

	// 将配置映射到结构体
	err = viper.UnmarshalKey("app", &appConfig)
	if err != nil {
		panic(fmt.Errorf("unable to decode into struct: %v", err))
	}

	err = viper.UnmarshalKey("database", &databaseConfig)
	if err != nil {
		panic(fmt.Errorf("unable to decode database config: %v", err))
	}
}

// 获取应用配置
func GetAppConfig() AppConfig {
	return appConfig
}

// 获取数据库配置
func GetDatabaseConfig() DatabaseConfig {
	return databaseConfig
}
