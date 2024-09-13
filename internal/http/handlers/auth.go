package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/febriaricandra/go-habit-tracker/config"
	"github.com/febriaricandra/go-habit-tracker/internal/database"
	"github.com/febriaricandra/go-habit-tracker/internal/utils"
)

// HomeHandler is the handler for the home route
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	type Parameter struct {
		Message string `json:"message"`
	}
	mess := &Parameter{
		Message: "Welcome to the Habit Tracker API",
	}
	utils.RespondWithJSON(w, http.StatusOK, mess)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	type Parameter struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	type Response struct {
		Token string `json:"token"`
	}

	var p Parameter
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&p)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if p.Username == "" || p.Password == "" {
		utils.RespondWithError(w, http.StatusBadRequest, "Username or Password is missing")
		return
	}

	// Check if the user exists
	var user database.User
	database.Instance.Select("id", "username", "password").Where("username = ?", p.Username).First(&user)
	if user.ID == 0 {
		utils.RespondWithError(w, http.StatusUnauthorized, "Invalid Username or Password")
		return
	}

	// Check if the password is correct
	if !utils.CheckPasswordHash(p.Password, user.Password) {
		err := errors.New("invalid Username or Password")
		utils.RespondWithError(w, http.StatusUnauthorized, err.Error())
		return
	}

	// Generate a JWT token
	token, err := utils.MakeJWT(int(user.ID), config.GetSecret(), 3600)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to generate token")
		return
	}

	// Respond with the token
	utils.RespondWithJSON(w, http.StatusOK, Response{Token: token})
}
