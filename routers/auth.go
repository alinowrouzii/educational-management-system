package routers

import (
	"fmt"

	"github.com/alinowrouzii/educational-management-system/controllers"
	"github.com/gorilla/mux"
)

func InitAuthRouter(r *mux.Router, cfg *controllers.Config) {
	fmt.Println("Initialize auth route...")

	r.PathPrefix("/auth").Subrouter().HandleFunc("/login", cfg.LoginHandler).Methods("POST")
}
