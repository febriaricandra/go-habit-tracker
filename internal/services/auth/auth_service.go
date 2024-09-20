package auth

import (
	"errors"

	"github.com/febriaricandra/go-habit-tracker/config"
	"github.com/febriaricandra/go-habit-tracker/internal/database"
	"github.com/febriaricandra/go-habit-tracker/internal/utils"
)

type AuthService struct{}

func (a *AuthService) Login(username, password string) (string, error) {
	type Parameter struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var p Parameter

	p.Username = username
	p.Password = password

	if p.Username == "" || p.Password == "" {
		return "", errors.New("username or Password is missing")
	}

	// Check if the user exists
	var user database.User
	database.Instance.Select("id", "username", "password").Where("username = ?", p.Username).First(&user)
	if user.ID == 0 {
		return "", errors.New("invalid Username or Password")
	}

	// Check if the password is correct
	if !utils.CheckPasswordHash(p.Password, user.Password) {
		err := errors.New("invalid Username or Password")
		return "", err
	}

	// Generate a JWT token
	token, err := utils.MakeJWT(int(user.ID), config.GetSecret(), 3600)
	if err != nil {
		return "", errors.New("failed to generate token")
	}

	return token, nil
}

func (a *AuthService) Register(username, password string) (string, error) {
	type Parameter struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var p Parameter

	p.Username = username
	p.Password = password

	if p.Username == "" || p.Password == "" {
		return "", errors.New("username or Password is missing")
	}

	// Check if the user exists
	var user database.User
	database.Instance.Select("id", "username", "password").Where("username = ?", p.Username).First(&user)
	if user.ID != 0 {
		return "", errors.New("user already exists")
	}

	// Hash the password
	hash, err := utils.HashPassword(p.Password)
	if err != nil {
		return "", err
	}

	// Create the user
	user = database.User{
		Username: p.Username,
		Password: hash,
	}

	err = database.Instance.Create(&user).Error
	if err != nil {
		return "", err
	}

	// Generate a JWT token
	token, err := utils.MakeJWT(int(user.ID), config.GetSecret(), 3600)
	if err != nil {
		return "", errors.New("failed to generate token")
	}

	return token, nil
}
