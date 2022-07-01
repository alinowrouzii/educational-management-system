package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/alinowrouzii/educational-management-system/models"
	"github.com/go-playground/validator/v10"
)

var Validator = validator.New()

type testStruct struct {
	Test string `json:"test"`
}

func (cfg *Config) TestHandler(w http.ResponseWriter, _ *http.Request) {

	RespondWithJSON(w, http.StatusOK, testStruct{
		Test: "hello world",
	})
}

// func (cfg *Config) CreateStudentHandler(w http.ResponseWriter, r *http.Request) {

// 	var s models.Student
// 	decoder := json.NewDecoder(r.Body)
// 	if err := decoder.Decode(&s); err != nil {
// 		RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
// 		return
// 	}
// 	if err := Validator.Struct(s); err != nil {
// 		RespondWithError(w, http.StatusInternalServerError, err.Error())
// 		return
// 	}

// 	defer r.Body.Close()

// 	if err := s.CreateStudent(cfg.DB); err != nil {
// 		RespondWithError(w, http.StatusInternalServerError, err.Error())
// 		return
// 	}
// 	respondWithJSON(w, http.StatusCreated, s)
// }

// func (cfg *Config) GetStudentHandler(w http.ResponseWriter, r *http.Request) {

// 	vars := mux.Vars(r)
// 	studentName, ok := vars["studentName"]
// 	if !ok {
// 		RespondWithError(w, http.StatusBadRequest, "name is required")
// 		return
// 	}
// 	log.Println("student name is", studentName)

// 	student := models.Student{Name: &studentName}
// 	if err := student.GetStudentByName(cfg.DB); err != nil {
// 		switch err {
// 		case sql.ErrNoRows:
// 			RespondWithError(w, http.StatusNotFound, "Student not found")
// 		default:
// 			RespondWithError(w, http.StatusInternalServerError, err.Error())
// 		}
// 		return
// 	}

// 	respondWithJSON(w, http.StatusOK, student)
// }

func (cfg *Config) ChangeStudentPasswordHandler(w http.ResponseWriter, r *http.Request) {
	var s models.Student
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&s); err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := Validator.Struct(s); err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	defer r.Body.Close()

	if err := s.ChangeStudentPassword(cfg.DB); err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondWithJSON(w, http.StatusCreated, s)
}

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	fmt.Println(payload)
	response, _ := json.Marshal(payload)
	// fmt.Println(response)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func RespondWithError(w http.ResponseWriter, code int, message string) {
	RespondWithJSON(w, code, map[string]string{"error": message})
}
