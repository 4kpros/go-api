package config

import (
	"api/common/constants"
	"api/common/helpers"
	"api/common/utils"
	"go.uber.org/zap"
)

type Key struct {
	JwtPrivateKey *string
	JwtPublicKey  *string
}

var Keys = &Key{}

// Loads the necessary cryptographic keys.
func LoadKeys() error {
	var err error
	var errRead error
	// JWT private key
	Keys.JwtPrivateKey, errRead = utils.ReadFileToString(constants.AssetKeysPath + "/jwt/private.pem")
	if errRead != nil {
		err = errRead
		helpers.Logger.Error(
			"Failed to load jwt/private.pem",
			zap.String("Error", errRead.Error()),
		)
	} else {
		helpers.Logger.Info("Key jwt/private.pem loaded!")
	}

	// JWT public key
	Keys.JwtPublicKey, errRead = utils.ReadFileToString(constants.AssetKeysPath + "/jwt/public.pem")
	if errRead != nil {
		err = errRead
		helpers.Logger.Error(
			"Failed to load jwt/public.pem",
			zap.String("Error", errRead.Error()),
		)
	} else {
		helpers.Logger.Info("Key jwt/public.pem loaded!")
	}
	return err
}
