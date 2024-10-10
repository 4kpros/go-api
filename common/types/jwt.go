package types

import (
	"github.com/golang-jwt/jwt/v5"
)

type JwtToken struct {
	jwt.RegisteredClaims
	UserId   int64  `json:"userId"`
	RoleId   int64  `json:"roleId"`
	Platform string `json:"platform"`
	Device   string `json:"device"`
	App      string `json:"app"`
	Code     int    `json:"code"`
}
