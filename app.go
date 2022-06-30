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
	Router              *mux.Router
	DB                  *sql.DB
	wantsToDropDatabase bool
}

func (a *App) dropAndCreateDatabase(connectionString, dbName string) {

	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec("DROP DATABASE final_db")
	if err != nil {
		panic(err)
	}
	createDBStmt := fmt.Sprintf("CREATE DATABASE %s DEFAULT CHARACTER SET = 'utf8mb4'", dbName)
	_, err = db.Exec(createDBStmt)
	if err != nil {
		panic(err)
	}

	fmt.Println("Database created successfully")
}

func (a *App) connectDB(user, password, dbName string) {

	connectionString := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/", user, password)
	if a.wantsToDropDatabase {
		fmt.Println("hereee")
		a.dropAndCreateDatabase(connectionString, dbName)
	}
	connectionString += dbName

	dbConn, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	a.DB = dbConn
}

func (a *App) Initialize(user, password, dbname string) {
	a.connectDB(user, password, dbname)

	a.Router = mux.NewRouter()
	routers.InitRouter(a.Router, a.DB)
}

func (a *App) Run(addr string, wantsToWriteData bool) {
	if a.wantsToDropDatabase {
		MakeMigrations(a.DB)
	}

	if wantsToWriteData {
		WriteDataToDatabsae(a.DB)
	}

	log.Fatal(http.ListenAndServe(addr, a.Router))

}
