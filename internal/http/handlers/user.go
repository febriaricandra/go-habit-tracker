package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/febriaricandra/go-habit-tracker/internal/database"
	services "github.com/febriaricandra/go-habit-tracker/internal/services/user"
	"github.com/febriaricandra/go-habit-tracker/internal/utils"
)

var userService = services.UserService{}

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	type Response struct {
		Users []database.User `json:"users"`
	}

	users, err := userService.GetUsers()
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to fetch users")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, Response{Users: users})
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	type Parameter struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	type Response struct {
		database.User
	}

	//decode the request body
	var p Parameter
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&p)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	//create the user
	user, err := userService.CreateUser(p.Username, p.Password)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, Response{User: user})
}
