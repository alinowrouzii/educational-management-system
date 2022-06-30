package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/alinowrouzii/educational-management-system/routers"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (a *App) Initialize(user, password, dbname string) {
	connectionString := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, password, dbname)

	var err error
	a.DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	a.Router = mux.NewRouter()
	routers.InitRouter(a.Router, a.DB)
}

func (a *App) Run(addr string, wantsToMigrate bool) {
	if wantsToMigrate {
		MakeMigrations(a.DB)
	} else {
		log.Fatal(http.ListenAndServe(addr, a.Router))
	}
}
