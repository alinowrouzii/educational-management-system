package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/alinowrouzii/educational-management-system/routers"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (a *App) Initialize(user, password, dbname string) {
	connectionString := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s", user, password, dbname)

	var err error
	a.DB, err = sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	a.Router = mux.NewRouter()
	routers.InitRouter(a.Router, a.DB)
}

func (a *App) Run(addr string, wantsToMigrate bool, wantsToWriteData bool) {
	if wantsToMigrate {
		MakeMigrations(a.DB)
	} else if wantsToWriteData {
		WriteDataToDatabsae(a.DB)
	} else {
		log.Fatal(http.ListenAndServe(addr, a.Router))
	}
}
