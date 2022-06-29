package routers

import (
	"database/sql"

	"github.com/gorilla/mux"
)

func InitRouter(r *mux.Router, db *sql.DB) {
	InitStudentRouter(r, db)

}
