package routers

import (
	"fmt"
	"net/http"

	"github.com/alinowrouzii/educational-management-system/controllers"
	"github.com/alinowrouzii/educational-management-system/middleware"
	"github.com/gorilla/mux"
)

func InitCoursesRouter(r *mux.Router, cfg *controllers.Config) {
	fmt.Println("Initialize courses route...")

	r.PathPrefix("/courses").Subrouter().Handle("/", middleware.TokenMiddleware(http.HandlerFunc(cfg.GetCoursesHandler), cfg.JWT)).Methods("GET")
	r.PathPrefix("/courses").Subrouter().Handle("/{course_id}/students/", middleware.TokenMiddleware(http.HandlerFunc(cfg.GetCourseStudentHandler), cfg.JWT)).Methods("GET")
}
