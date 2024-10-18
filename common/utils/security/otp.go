package security

import (
	"fmt"

	"api/common/constants"
	"api/config"

	"github.com/pquerna/otp/totp"
)

func GenerateOTP(userId int64, userName string, issuer string) (string, error) {
	// Generate OTP secret if not already generated
	secret, err := totp.Generate(totp.GenerateOpts{
		Issuer:      issuer,
		AccountName: userName,
	})
	if err != nil {
		return "", constants.HTTP_500_ERROR_MESSAGE("generate TOTP code")
	}

	// Save secret on redis
	config.SetRedisString(GetJWTCachedKey(userId, issuer), secret.Secret())

	// Construct the OTP URL for generating QR code
	otpURL := fmt.Sprintf("otpauth://totp/%s:%s?secret=%s&issuer=%s",
		config.Env.AppName,
		userName,
		secret.Secret(),
		issuer,
	)

	// Generate QR code
	// TODO contact storage api

	return otpURL, nil
}

// Validate the OTP code
func ValidateOTP(otpCode int, userId int64, issuer string) bool {
	userSecret, err := config.GetRedisString(GetJWTCachedKey(userId, issuer))
	if err != nil {
		return false
	}
	return totp.Validate(fmt.Sprintf("%d", otpCode), userSecret)
}
