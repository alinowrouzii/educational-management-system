package routers

import (
	"database/sql"
	"fmt"

	"github.com/alinowrouzii/educational-management-system/controllers"
	"github.com/gorilla/mux"
)

func InitStudentRouter(r *mux.Router, db *sql.DB) {
	fmt.Println("Initialize student route...")
	cfg := controllers.Config{
		DB: db,
	}

	// studentRouter := mux.NewRouter()
	r.PathPrefix("/student").Subrouter().HandleFunc("/student/test", cfg.TestHandler).Methods("GET")
	r.PathPrefix("/student").Subrouter().HandleFunc("/student", cfg.GetStudentHandler).Methods("GET")

}
