package controllers

import "database/sql"

type Config struct {
	DB *sql.DB
}
