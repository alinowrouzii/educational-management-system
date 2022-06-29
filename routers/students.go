package routers

import (
	"fmt"

	"github.com/alinowrouzii/educational-management-system/controllers"
	"github.com/gorilla/mux"
)

func InitStudentRouter(r *mux.Router) {
	fmt.Println("Initialize student route...")

	// studentRouter := mux.NewRouter()
	r.PathPrefix("/student").Subrouter().HandleFunc("/student/test", controllers.TestHandler).Methods("GET")
}
