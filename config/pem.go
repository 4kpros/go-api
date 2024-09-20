package config

import (
	"github.com/4kpros/go-api/common/helpers"
	"go.uber.org/zap"
)

type Pem struct {
	JwtPrivateKey string
	JwtPublicKey  string
}

var AppPem = &Pem{}

func LoadPem() (err error) {
	var err1, err2 error
	// JWT private key
	AppPem.JwtPrivateKey, err1 = helpers.ReadFileContentToString("keys/jwt_private.pem")
	if err1 != nil {
		helpers.Logger.Warn(
			"Failed to load jwt_private.pem",
			zap.String("Error", err1.Error()),
		)
		err = err1
	} else {
		helpers.Logger.Warn(
			"Pem jwt_private.pem loaded!",
		)
	}

	// JWT public key
	AppPem.JwtPublicKey, err2 = helpers.ReadFileContentToString("keys/jwt_public.pem")
	if err2 != nil {
		helpers.Logger.Warn(
			"Failed to load jwt_public.pem",
			zap.String("Error", err2.Error()),
		)
		err = err2
	} else {
		helpers.Logger.Warn(
			"Pem jwt_public.pem loaded!",
		)
	}
	return
}
