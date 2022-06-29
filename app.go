package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/alinowrouzii/educational-management-system/routers"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (a *App) Initialize(user, password, dbname string) {
	fmt.Println("here")
	connectionString := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, password, dbname)

	fmt.Println((connectionString))
	var err error
	a.DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("here is err", err)
	fmt.Println(a.DB)

	a.Router = mux.NewRouter()
	routers.InitRouter(a.Router, a.DB)
}

func (a *App) Run(addr string) {
	MakeMigrations(a.DB)
	// log.Fatal(http.ListenAndServe(":8010", a.Router))
}
