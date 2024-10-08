package types

import "time"

type JwtToken struct {
	UserId  uint
	Role    int
	Expires time.Time
	Code    int
	Device  string
	Issuer  string
}
