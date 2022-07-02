package controllers

import (
	"fmt"
	"net/http"

	"github.com/alinowrouzii/educational-management-system/models"
	"github.com/alinowrouzii/educational-management-system/token"
	"github.com/gorilla/mux"
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

func (cfg *Config) GetCourseStudentHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(mux.Vars(r)["course_id"])
	courseID := mux.Vars(r)["course_id"]
	if courseID == "" {
		RespondWithError(w, http.StatusInternalServerError, "course_id is required!")
		return
	}

	user := r.Context().Value("payload")
	fmt.Println("==========")
	token := user.(*token.Payload)
	username := token.Username
	fmt.Println(username)
	students, err := models.GetCourseStudents(cfg.DB, username, courseID)

	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	RespondWithJSON(w, http.StatusOK, students)
}

func (cfg *Config) GetCourseHWExamHandler(w http.ResponseWriter, r *http.Request) {
	// TODO
}
