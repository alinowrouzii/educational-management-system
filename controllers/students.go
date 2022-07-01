package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/alinowrouzii/educational-management-system/token"
	"github.com/go-playground/validator/v10"
)

var Validator = validator.New()

type testStruct struct {
	Test string `json:"test"`
}

func (cfg *Config) TestHandler(w http.ResponseWriter, r *http.Request) {

	user := r.Context().Value("payload")
	fmt.Println("==========")
	token := user.(*token.Payload)
	username := token.Username
	fmt.Println(username)

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
