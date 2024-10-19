package security

import (
	"api/common/constants"
	"api/common/types"
	"api/config"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Returns the default expiration time for JWT. It's 10 min.
func NewExpiresDateDefault() *time.Time {
	tempDate := time.Now().Add(time.Minute * time.Duration(config.Env.JwtExpiresDefault))
	return &tempDate
}

// Returns the login expiration time for JWT.
// With default ENV var, the JWT login expiration time is
// 30 days if stayConnected is true, otherwise it's 1 hour.
func NewExpiresDateLogin(stayConnected bool) (date *time.Time) {
	if stayConnected {
		tempDate := time.Now().Add(time.Hour * time.Duration(24*config.Env.JwtExpiresLoginStayConnected))
		return &tempDate
	}
	tempDate := time.Now().Add(time.Minute * time.Duration(config.Env.JwtExpiresLogin))
	return &tempDate
}

// Returns the key of Redis entry(combining userId and Issuer)
func GetJWTCachedKey(userId int64, issuer string) (key string) {
	return fmt.Sprintf("%d_%s", userId, issuer)
}

// Encode the JWT token and store it in the cache. This returns a signed string token.
func EncodeJWTToken(jwtToken *types.JwtToken, issuer string, expires *time.Time, privateKey *string, cacheFunc func(string, string) error) (*types.JwtToken, string, error) {
	// Encode token with claims
	jwtToken.Issuer = issuer
	jwtToken.ExpiresAt = jwt.NewNumericDate(*expires)
	jwtToken.IssuedAt = jwt.NewNumericDate(time.Now())
	jwtTokenClaimed := jwt.NewWithClaims(jwt.SigningMethodES512, *jwtToken)
	signedKey, errParse := jwt.ParseECPrivateKeyFromPEM([]byte(*privateKey))
	if errParse != nil {
		return nil, "", errParse
	}
	token, errSigning := jwtTokenClaimed.SignedString(signedKey)
	if errSigning != nil {
		return nil, "", errSigning
	}

	// Cache new token
	errCache := cacheFunc(GetJWTCachedKey(jwtToken.UserId, jwtToken.Issuer), token)
	if errCache != nil {
		return nil, "", errCache
	}
	return jwtToken, token, nil
}

// Decode the JWT token using the signed string token and public key.
// This returns a JWT token object.
func DecodeJWTToken(token string, publicKey *string) (*types.JwtToken, error) {
	// Parse the token and get claims
	jwtToken, err := jwt.ParseWithClaims(token, &types.JwtToken{}, func(token *jwt.Token) (signedKey interface{}, err error) {
		if _, ok := token.Method.(*jwt.SigningMethodECDSA); !ok {
			return nil, fmt.Errorf("%s", fmt.Sprintf("Unexpected signing method: %v", token.Header["alg"]))
		}
		signedKey, err = jwt.ParseECPublicKeyFromPEM([]byte(*publicKey))
		return
	})
	if err != nil {
		return nil, err
	} else if claims, ok := jwtToken.Claims.(*types.JwtToken); ok && jwtToken.Valid {
		return claims, nil
	}
	return nil, constants.Http401InvalidTokenErrorMessage()
}

// Validate the token by checking if it is cached.
func ValidateJWTToken(token string, jwtToken *types.JwtToken, loadCachedFunc func(string) (string, error)) bool {
	tokenCached, errCached := loadCachedFunc(GetJWTCachedKey(jwtToken.UserId, jwtToken.Issuer))
	if errCached != nil || len(tokenCached) <= 0 {
		return false
	}
	if token != tokenCached {
		return false
	}
	return true
}
