package models

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/alinowrouzii/educational-management-system/token"
)

var insertNewToken = "INSERT into token (id, username, issue_at, expired_at) VALUES (?, ?, ?, ?)"
var loginUser = "SELECT login_user(username, password) as shit"

type UserLogin struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (u *UserLogin) Login(jwt *token.JWTMaker, db *sql.DB) (map[string]interface{}, error) {

	fmt.Println("here is user", u)
	checkLoginSuccessfull := "FAIL"
	err := db.QueryRow(loginUser, u.Username, u.Password).Scan(&checkLoginSuccessfull)

	if err != nil {
		log.Fatal("Error after login", err)
	}

	if checkLoginSuccessfull == "FAIL" {
		return nil, errors.New("There is no user with provided credentials")
	}

	// elsewhere create new token

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

func Logout(jwt *token.JWTMaker, token string) (map[string]interface{}, error) {
	res, err := jwt.RevokeToken(token)

	return res, err
}
