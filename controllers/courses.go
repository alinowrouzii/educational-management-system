package controllers

import (
	"fmt"
	"net/http"

	"github.com/alinowrouzii/educational-management-system/models"
	"github.com/alinowrouzii/educational-management-system/token"
)

func (cfg *Config) GetCoursesHandler(w http.ResponseWriter, r *http.Request) {

	user := r.Context().Value("payload")
	fmt.Println("==========")
	token := user.(*token.Payload)
	username := token.Username
	fmt.Println(username)

	courses, err := models.GetUserCourses(cfg.DB, username)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	RespondWithJSON(w, http.StatusOK, courses)
}
