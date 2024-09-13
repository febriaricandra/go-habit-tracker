package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/febriaricandra/go-habit-tracker/internal/database"
	"github.com/febriaricandra/go-habit-tracker/internal/utils"
)

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	type Response struct {
		Users []database.User `json:"users"`
	}

	var users []database.User
	database.Instance.Find(&users)

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

	hashedPassword, err := utils.HashPassword(p.Password)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to hash password")
		return
	}

	user := database.User{
		Username: p.Username,
		Password: hashedPassword,
	}

	database.Instance.Create(&user)

	utils.RespondWithJSON(w, http.StatusCreated, Response{User: user})
}
