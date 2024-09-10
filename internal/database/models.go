package database

import (
	"time"
)

type User struct {
	ID        int    `gorm:"primaryKey"`
	Username  string `gorm:"unique"`
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Habit struct {
	ID        int `gorm:"primaryKey"`
	UserID    int
	Name      string
	Frequency string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type HabitLog struct {
	ID        int `gorm:"primaryKey"`
	HabitID   int
	Completed bool
	Date      time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserHabit struct {
	ID        int `gorm:"primaryKey"`
	UserID    int
	HabitID   int
	CreatedAt time.Time
	UpdatedAt time.Time
}
