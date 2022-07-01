package routers

import (
	"fmt"

	"github.com/alinowrouzii/educational-management-system/controllers"
	"github.com/gorilla/mux"
)

func InitStudentRouter(r *mux.Router, cfg *controllers.Config) {
	fmt.Println("Initialize student route...")
	// studentRouter := mux.NewRouter()
	r.PathPrefix("/student").Subrouter().HandleFunc("/test", cfg.TestHandler).Methods("GET")
	// r.PathPrefix("/student").Subrouter().HandleFunc("/", cfg.CreateStudentHandler).Methods("POST")
	// r.PathPrefix("/student").Subrouter().HandleFunc("/{studentName}", cfg.GetStudentHandler).Methods("GET")
	r.PathPrefix("/student").Subrouter().HandleFunc("/changePassword", cfg.ChangeStudentPasswordHandler).Methods("PATCH")
}
