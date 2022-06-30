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
	a.Initialize(
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"))
	wantsToMigrate := os.Getenv("WANTS_TO_MIGRATE")
	wantsToWriteData := os.Getenv("WANTS_TO_WRITE_DATA")

	wantsToMigrateBool, err := strconv.ParseBool(wantsToMigrate)
	if err != nil {
		log.Fatal(err)
	}

	wantsToWriteDataBool, err := strconv.ParseBool(wantsToWriteData)
	if err != nil {
		log.Fatal(err)
	}

	a.Run(":8010", wantsToMigrateBool, wantsToWriteDataBool)
}
