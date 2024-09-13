package database

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"fmt"

	"github.com/febriaricandra/go-habit-tracker/config"
)

var Instance *gorm.DB
var err error

func Connect(cfg *config.AppConfig) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DB_USER, cfg.DB_PASS, cfg.DB_HOST, cfg.DB_PORT, cfg.DB_NAME,
	)

	Instance, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	log.Println("Connected to database")
}

func Migrate() {
	Instance.AutoMigrate(&User{}, &Habit{}, &Activity{})
	log.Println("Database Migration Completed...")
}
