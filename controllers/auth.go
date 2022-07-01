package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/alinowrouzii/educational-management-system/models"
)

var changeUserPassword = `SELECT change_student_password(?, ?, ?) as shit`

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

	payload, err := user.Login(cfg.JWT, cfg.DB)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	RespondWithJSON(w, http.StatusCreated, map[string]interface{}{"result": payload})
}

func (cfg *Config) LogoutHandler(w http.ResponseWriter, r *http.Request) {

	token := map[string]interface{}{"token": ""}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&token); err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	fmt.Println("here is token`", token["token"] == "", "`", "`")

	// payload, err := user.Login(cfg.JWT, cfg.DB)
	if token["token"] == "" {
		RespondWithError(w, http.StatusBadRequest, "required token")
		return
	}

	res, err := models.Logout(cfg.JWT, token["token"].(string))

	if err != nil {
		RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	RespondWithJSON(w, http.StatusCreated, res)
}

func (cfg *Config) ChangePasswordHandler(w http.ResponseWriter, r *http.Request) {
	var u models.UserChangePassword
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&u); err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := Validator.Struct(u); err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	defer r.Body.Close()

	if err := u.ChangePassword(cfg.DB); err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondWithJSON(w, http.StatusCreated, u)
}
