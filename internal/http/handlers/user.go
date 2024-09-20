package handlers

import (
	"net/http"

	"github.com/febriaricandra/go-habit-tracker/config"
	"github.com/febriaricandra/go-habit-tracker/internal/database"
	services "github.com/febriaricandra/go-habit-tracker/internal/services/user"
	"github.com/febriaricandra/go-habit-tracker/internal/utils"
)

var userService = services.UserService{}

func GetUserByName(w http.ResponseWriter, r *http.Request) {
	type Response struct {
		User database.User `json:"user"`
	}

	username := r.URL.Query().Get("username")
	if username == "" {
		// Get token from authorization header
		token, err := utils.GetBearerToken(r.Header)
		if err != nil {
			utils.RespondWithError(w, http.StatusUnauthorized, "Unauthorized")
			return
		}

		// Validate the token
		claims, err := utils.ValidateJWT(token, config.GetSecret())
		if err != nil {
			utils.RespondWithError(w, http.StatusUnauthorized, "Unauthorized")
			return
		}

		user, err := userService.GetUserByID(claims)

		if err != nil {
			utils.RespondWithError(w, http.StatusNotFound, "User not found")
			return
		}

		utils.RespondWithJSON(w, http.StatusOK, Response{User: user})
	} else {
		user, err := userService.GetUser(username)
		if err != nil {
			utils.RespondWithError(w, http.StatusNotFound, "User not found")
			return
		}

		utils.RespondWithJSON(w, http.StatusOK, Response{User: user})
	}
}
