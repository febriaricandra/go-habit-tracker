package database

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        int    `gorm:"primaryKey"`
	Username  string `gorm:"unique"`
	Password  string `json:"-"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Habit struct {
	gorm.Model
	ID        int `gorm:"primaryKey"`
	UserID    int
	Name      string
	Frequency string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Activity struct {
	gorm.Model
	ID        int `gorm:"primaryKey"`
	HabitID   int
	Completed bool
	Note      string
	Date      time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}
