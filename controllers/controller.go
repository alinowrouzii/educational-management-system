package controllers

import (
	"database/sql"

	"github.com/alinowrouzii/educational-management-system/token"
)

type Config struct {
	DB  *sql.DB
	JWT *token.JWTMaker
}
