package models

import (
	"database/sql"
	"time"

	"github.com/alinowrouzii/educational-management-system/token"
)

var insertNewToken = "INSERT into token (id, username, issue_at, expired_at) VALUES (?, ?, ?, ?)"

type UserLogin struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (u *UserLogin) Login(jwt *token.JWTMaker, db *sql.DB) (map[string]interface{}, error) {

	// 10000 seconds for expiration time
	var d time.Duration = 10000000000000
	payload, token, err := jwt.CreateToken(u.Username, d)
	res := map[string]interface{}{
		"payload": payload,
		"token":   token,
	}
	db.QueryRow(insertNewToken, payload.ID, payload.Username, payload.IssuedAt, payload.ExpiredAt)
	return res, err
}
