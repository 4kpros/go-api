package utils

import (
	"fmt"
	"runtime"
	"time"

	"github.com/4kpros/go-api/common/types"
	"github.com/4kpros/go-api/config"
	"github.com/alexedwards/argon2id"
	"github.com/golang-jwt/jwt/v5"
)

const JWT_ISSUER_SESSION = "JWT_ISSUER_SESSION"
const JWT_ISSUER_SESSION_GENERATED = "JWT_ISSUER_SESSION_GENERATED"
const JWT_ISSUER_ACTIVATE = "JWT_ISSUER_ACTIVATE"
const JWT_ISSUER_RESET_CODE = "JWT_ISSUER_RESET_CODE"
const JWT_ISSUER_RESET_PASSWORD = "JWT_ISSUER_RESET_PASSWORD"

// Return default time for JWT expires
func NewExpiresDateDefault() *time.Time {
	tempDate := time.Now().Add(time.Minute * time.Duration(config.Env.JwtExpiresDefault))
	return &tempDate
}

// Return time for JWT expires when user try to login
func NewExpiresDateSignIn(stayConnected bool) (date *time.Time) {
	if stayConnected {
		tempDate := time.Now().Add(time.Hour * time.Duration(24*config.Env.JwtExpiresSignInStayConnected))
		return &tempDate
	}
	tempDate := time.Now().Add(time.Minute * time.Duration(config.Env.JwtExpiresSignIn))
	return &tempDate
}

// Return the key of Redis entry(combining userId and Issuer)
func GetJWTCachedKey(jwtToken *types.JwtToken) (key string) {
	if jwtToken != nil {
		return fmt.Sprintf("%d%s", jwtToken.UserId, jwtToken.Issuer)
	}
	return ""
}

// Encode JWT token and add to cache. This return signed string token
func EncodeJWTToken(jwtToken *types.JwtToken, issuer string, expires *time.Time, privateKey string, cacheFunc func(string, string) error) (*types.JwtToken, string, error) {
	// Encode token with claims
	jwtToken.Issuer = issuer
	jwtToken.ExpiresAt = jwt.NewNumericDate(*expires)
	jwtToken.IssuedAt = jwt.NewNumericDate(time.Now())
	var jwtTokenClaimed = jwt.NewWithClaims(jwt.SigningMethodES512, *jwtToken)
	var signedKey, errParse = jwt.ParseECPrivateKeyFromPEM([]byte(privateKey))
	if errParse != nil {
		return nil, "", errParse
	}
	var token, errSigning = jwtTokenClaimed.SignedString(signedKey)
	if errSigning != nil {
		return nil, "", errSigning
	}

	// Cache new token
	var errCache = cacheFunc(GetJWTCachedKey(jwtToken), token)
	if errCache != nil {
		return nil, "", errCache
	}
	return jwtToken, token, nil
}

// Decode JWT token by providing signed string token and public key. This return JWT token object
func DecodeJWTToken(token string, publicKey string) (*types.JwtToken, error) {
	// Parse the token and get claims
	var errMessage string
	jwtToken, err := jwt.ParseWithClaims(token, &types.JwtToken{}, func(token *jwt.Token) (signedKey interface{}, err error) {
		if _, ok := token.Method.(*jwt.SigningMethodECDSA); !ok {
			errMessage = fmt.Sprintf("Unexpected signing method: %v", token.Header["alg"])
			return nil, fmt.Errorf("%s", errMessage)
		}
		signedKey, err = jwt.ParseECPublicKeyFromPEM([]byte(publicKey))
		return
	})
	if err != nil {
		return nil, err
	} else if claims, ok := jwtToken.Claims.(*types.JwtToken); ok && jwtToken.Valid {
		return claims, nil
	}
	errMessage = "Invalid token or expired! Please enter valid information."
	return nil, fmt.Errorf("%s", errMessage)
}

// Validate the token by checking if the token it's cached
func ValidateJWTToken(token string, jwtToken *types.JwtToken, loadCachedFunc func(string) (string, error)) bool {
	tokenCached, errCached := loadCachedFunc(GetJWTCachedKey(jwtToken))
	if errCached != nil || len(tokenCached) <= 0 {
		return false
	}
	if token != tokenCached {
		return false
	}
	return true
}

// Encode password using Argon2id and returns new hashed password
func EncodeArgon2id(password string) (string, error) {
	params := &argon2id.Params{
		Memory:      uint32(config.Env.ArgonMemoryLeft * config.Env.ArgonMemoryRight),
		Iterations:  uint32(config.Env.ArgonIterations),
		Parallelism: uint8(runtime.NumCPU()),
		SaltLength:  uint32(config.Env.ArgonSaltLength),
		KeyLength:   uint32(config.Env.ArgonKeyLength),
	}
	tempHash, tempErr := argon2id.CreateHash(password, params)
	if tempErr != nil {
		return "", tempErr
	}
	// Encode to base64
	var hash = EncodeBase64(tempHash)
	return hash, nil
}

// Verify if Argon2id password matches string
func CompareArgon2id(password string, hashedPassword string) (bool, error) {
	initialHashedPassword, _ := DecodeBase64(hashedPassword)
	var match, err = argon2id.ComparePasswordAndHash(password, initialHashedPassword)
	return match, err
}
