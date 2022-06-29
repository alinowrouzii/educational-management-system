package routers

import (
	"github.com/gorilla/mux"
)

func InitRouter(r *mux.Router) {
	InitStudentRouter(r)
	// r.Handle("/student", studentRouter)
	// fmt.Println((studentRouter))
}
