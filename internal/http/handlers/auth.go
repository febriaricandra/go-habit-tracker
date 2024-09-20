package handlers

import (
	"encoding/json"
	"net/http"

	auth "github.com/febriaricandra/go-habit-tracker/internal/services/auth"
	"github.com/febriaricandra/go-habit-tracker/internal/utils"
)

var authService = auth.AuthService{}

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

	// Check if the user exists
	token, err := authService.Login(p.Username, p.Password)
	if err != nil {
		utils.RespondWithError(w, http.StatusUnauthorized, err.Error())
		return
	}

	// Respond with the token
	utils.RespondWithJSON(w, http.StatusOK, Response{Token: token})
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
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

	// Check if the user exists
	token, err := authService.Register(p.Username, p.Password)
	if err != nil {
		utils.RespondWithError(w, http.StatusUnauthorized, err.Error())
		return
	}

	// Respond with the token
	utils.RespondWithJSON(w, http.StatusOK, Response{Token: token})
}
