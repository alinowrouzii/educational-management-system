// main.go

package main

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type config struct {
	AppDbUsername string `env:"APP_DB_USERNAME,file"`
	AppDbPassword string `env:"APP_DB_PASSWORD,file"`
	AppDbName     string `env:"APP_DB_NAME,file"`
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	a := App{}
	wantsToDropDatabase := os.Getenv("WANTS_TO_DROP_DATABASE")
	wantsToWriteData := os.Getenv("WANTS_TO_WRITE_DATA")
	wantsToMakeMigrations := os.Getenv("WANTS_TO_MAKE_MIGRATIONS")

	wantsToDropDatabaseBool, err := strconv.ParseBool(wantsToDropDatabase)
	if err != nil {
		log.Fatal(err)
	}

	wantsToWriteDataBool, err := strconv.ParseBool(wantsToWriteData)
	if err != nil {
		log.Fatal(err)
	}
	wantsToMakeMigrationsBool, err := strconv.ParseBool(wantsToMakeMigrations)
	if err != nil {
		log.Fatal(err)
	}

	a.wantsToDropDatabase = wantsToDropDatabaseBool

	a.Initialize(
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"),
		os.Getenv("SECRET_KEY"))

	a.Run(":8010", wantsToWriteDataBool, wantsToMakeMigrationsBool)
}
