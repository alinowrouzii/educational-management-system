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
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if err := Validator.Struct(user); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	token, err := user.Login(cfg.JWT)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, map[string]interface{}{token: token})
}
