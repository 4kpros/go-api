package types

import (
	"github.com/golang-jwt/jwt/v5"
)

type JwtToken struct {
	UserId   uint   `json:"userId"`
	RoleId   int    `json:"roleId"`
	Platform string `json:"platform"`
	Device   string `json:"device"`
	App      string `json:"app"`
	Code     int    `json:"code"`
	jwt.RegisteredClaims
}
