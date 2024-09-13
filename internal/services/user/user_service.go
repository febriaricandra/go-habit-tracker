package services

import (
	"errors"

	"github.com/febriaricandra/go-habit-tracker/internal/database"
	"github.com/febriaricandra/go-habit-tracker/internal/utils"
)

type UserService struct{}

func (s *UserService) GetUsers() ([]database.User, error) {
	var users []database.User
	err := database.Instance.Find(&users).Error
	return users, err
}

func (s *UserService) CreateUser(username, password string) (database.User, error) {
	type Parameter struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var p Parameter

	p.Username = username
	p.Password = password

	if p.Username == "" || p.Password == "" {
		return database.User{}, errors.New("Username or Password is missing")
	}

	// Check if the user exists
	var user database.User
	database.Instance.Select("id", "username", "password").Where("username = ?", p.Username).First(&user)
	if user.ID != 0 {
		return database.User{}, errors.New("User already exists")
	}

	// Hash the password
	hash, err := utils.HashPassword(p.Password)
	if err != nil {
		return database.User{}, err
	}

	// Create the user
	user = database.User{
		Username: p.Username,
		Password: hash,
	}

	err = database.Instance.Create(&user).Error
	return user, err
}
