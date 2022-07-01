package routers

import (
	"fmt"
	"net/http"

	"github.com/alinowrouzii/educational-management-system/controllers"
	"github.com/alinowrouzii/educational-management-system/middleware"
	"github.com/gorilla/mux"
)

func InitStudentRouter(r *mux.Router, cfg *controllers.Config) {
	fmt.Println("Initialize student route...")
	// studentRouter := mux.NewRouter()
	r.PathPrefix("/student").Subrouter().Handle("/test", middleware.TokenMiddleware(http.HandlerFunc(cfg.TestHandler), cfg.JWT)).Methods("GET")
	// r.PathPrefix("/student").Subrouter().HandleFunc("/", cfg.CreateStudentHandler).Methods("POST")
	// r.PathPrefix("/student").Subrouter().HandleFunc("/{studentName}", cfg.GetStudentHandler).Methods("GET")
}
