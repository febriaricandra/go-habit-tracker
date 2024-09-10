package database

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"fmt"
	"os"
)

var Instance *gorm.DB
var err error

func Connect() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	Instance, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	log.Println("Connected to database")
}

func Migrate() {
	Instance.AutoMigrate(&User{}, &Habit{}, &HabitLog{}, &UserHabit{})
	log.Println("Database Migration Completed...")
}
