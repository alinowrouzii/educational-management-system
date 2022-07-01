package models

import (
	"time"

	"github.com/alinowrouzii/educational-management-system/token"
)

type UserLogin struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (u *UserLogin) Login(jwt *token.JWTMaker) (string, error) {

	// 10000 seconds for expiration time
	var d time.Duration = 10000000000000

	payload, err := jwt.CreateToken(u.Username, d)
	return payload, err
}
