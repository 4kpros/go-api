package config

import (
	"github.com/4kpros/go-api/common/helpers"
	"go.uber.org/zap"
)

type Key struct {
	JwtPrivateKey string
	JwtPublicKey  string
}

var Keys = &Key{}

func LoadKeys() (err error) {
	var errTemp error
	// JWT private key
	Keys.JwtPrivateKey, errTemp = helpers.ReadFileContentToString("keys/jwt/private.pem")
	if errTemp != nil {
		helpers.Logger.Error(
			"Failed to load jwt/private.pem",
			zap.String("Error", errTemp.Error()),
		)
		err = errTemp
	} else {
		helpers.Logger.Info(
			"Key jwt/private.pem loaded!",
		)
	}

	// JWT public key
	Keys.JwtPublicKey, errTemp = helpers.ReadFileContentToString("keys/jwt/public.pem")
	if errTemp != nil {
		helpers.Logger.Error(
			"Failed to load jwt/public.pem",
			zap.String("Error", errTemp.Error()),
		)
		err = errTemp
	} else {
		helpers.Logger.Info(
			"Key jwt/public.pem loaded!",
		)
	}
	return
}
