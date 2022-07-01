package routers

import (
	"database/sql"

	"github.com/alinowrouzii/educational-management-system/controllers"
	"github.com/alinowrouzii/educational-management-system/token"
	"github.com/gorilla/mux"
)

func InitRouter(r *mux.Router, db *sql.DB, jwt *token.JWTMaker) {
	cfg := &controllers.Config{
		DB:  db,
		JWT: jwt,
	}

	InitStudentRouter(r, cfg)
	InitAuthRouter(r, cfg)
	InitCoursesRouter(r, cfg)

}
