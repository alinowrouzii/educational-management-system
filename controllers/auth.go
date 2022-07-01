package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/alinowrouzii/educational-management-system/models"
)

func (cfg *Config) LoginHandler(w http.ResponseWriter, r *http.Request) {

	var user models.UserLogin
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if err := Validator.Struct(user); err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	token, err := user.Login(cfg.JWT)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	RespondWithJSON(w, http.StatusCreated, map[string]interface{}{"token": token})
}
