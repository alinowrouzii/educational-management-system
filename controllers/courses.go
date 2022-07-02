package controllers

import (
	"encoding/json"
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

	exams, err := models.GetCourseExams(cfg.DB, username, courseID)

	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondWithJSON(w, http.StatusOK, exams)
}

func (cfg *Config) CreateCourseExamHandler(w http.ResponseWriter, r *http.Request) {

	var exam models.Exam
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&exam); err != nil {
		fmt.Println(err)
		RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := Validator.Struct(exam); err != nil {
		RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	user := r.Context().Value("payload")
	fmt.Println("==========")
	token := user.(*token.Payload)
	username := token.Username
	fmt.Println(username)

	err := exam.CreateCourseExam(cfg.DB, username)

	if err != nil {
		RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	RespondWithJSON(w, http.StatusCreated, exam)
}

func (cfg *Config) AddExamQuestionHandler(w http.ResponseWriter, r *http.Request) {

	var question models.ExamQuestion
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&question); err != nil {
		fmt.Println(err)
		RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := Validator.Struct(question); err != nil {
		RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	user := r.Context().Value("payload")
	fmt.Println("==========")
	token := user.(*token.Payload)
	username := token.Username
	fmt.Println(username)

	err := question.AddExamQuestion(cfg.DB, username)

	if err != nil {
		RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	RespondWithJSON(w, http.StatusCreated, question)
}

func (cfg *Config) SubmitExamAnswerHandler(w http.ResponseWriter, r *http.Request) {

	var answer models.ExamAnswer
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&answer); err != nil {
		fmt.Println(err)
		RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := Validator.Struct(answer); err != nil {
		RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	user := r.Context().Value("payload")
	fmt.Println("==========")
	token := user.(*token.Payload)
	username := token.Username
	fmt.Println(username)

	answer.StudentNO = username
	err := answer.SubmitExamAnswer(cfg.DB)

	if err != nil {
		RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	RespondWithJSON(w, http.StatusCreated, answer)
}

func (cfg *Config) GetExamScoreHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println(mux.Vars(r)["exam_id"])
	examID := mux.Vars(r)["exam_id"]
	if examID == "" {
		RespondWithError(w, http.StatusInternalServerError, "exam_id is required!")
		return
	}

	user := r.Context().Value("payload")
	fmt.Println("==========")
	token := user.(*token.Payload)
	username := token.Username
	fmt.Println(username)

	score, err := models.GetStudentExamScore(cfg.DB, username, examID)

	if err != nil {
		RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	RespondWithJSON(w, http.StatusCreated, map[string]interface{}{"score": score})
}
