package server

import (
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDatabase() *gorm.DB {
	var db *gorm.DB
	var err error

	for i := 0; i < 5; i++ { // 设置最大重试次数
		db, err = gorm.Open(mysql.Open("root:123456@tcp(db:3306)/dev_go?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
		if err == nil {
			return db
		}
		log.Printf("Failed to connect to database, attempt %d: %v\n", i+1, err)
		time.Sleep(5 * time.Second) // 等待5秒后重试
	}

	log.Fatalf("failed to connect database after 5 attempts: %v", err)

	return db
}
