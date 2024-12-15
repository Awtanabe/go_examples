package db

import (
	"fmt"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)


func NewDB() *gorm.DB {
	var db *gorm.DB
	var err error

	USER := os.Getenv("DB_USER")
	PASS := os.Getenv("DB_PASS")
	DB_HOST := os.Getenv("DB_HOST")
	if DB_HOST == "" {
		DB_HOST = "db"
	}
  // PROTOCOL := fmt.Sprintf("tcp(%s:3306)", DB_HOST)
	DB_NAME := os.Getenv("DB_NAME")

  dsn := USER + ":" + PASS + "@" + "tcp(db:3306)" + "/" + DB_NAME + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"


	for i := 0; i < 10; i++ { // 最大10回リトライ
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err == nil {
			break
		}
		fmt.Println("Retrying database connection...")
		time.Sleep(2 * time.Second) // 2秒間隔でリトライ
	}

	if err != nil {
		panic("failed to connect database")
	}


	return db
}