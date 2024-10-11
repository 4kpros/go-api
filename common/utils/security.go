package utils

import (
	base64 "encoding/base64"
	"fmt"
	"runtime"
	"time"

	"github.com/4kpros/go-api/common/constants"
	"github.com/4kpros/go-api/common/types"
	"github.com/4kpros/go-api/config"
	"github.com/alexedwards/argon2id"
	"github.com/golang-jwt/jwt/v5"
)

// Returns the default expiration time for JWT. It's 10 min.
func NewExpiresDateDefault() *time.Time {
	tempDate := time.Now().Add(time.Minute * time.Duration(config.Env.JwtExpiresDefault))
	return &tempDate
}

// Returns the login expiration time for JWT.
// With default ENV var, the JWT login expiration time is 24 hours
// if stayConnected is true, otherwise it's 1 hour.
func NewExpiresDateSignIn(stayConnected bool) (date *time.Time) {
	if stayConnected {
		tempDate := time.Now().Add(time.Hour * time.Duration(24*config.Env.JwtExpiresSignInStayConnected))
		return &tempDate
	}
	tempDate := time.Now().Add(time.Minute * time.Duration(config.Env.JwtExpiresSignIn))
	return &tempDate
}

// Returns the key of Redis entry(combining userId and Issuer)
func GetJWTCachedKey(jwtToken *types.JwtToken) (key string) {
	if jwtToken != nil {
		return fmt.Sprintf("%d%s", jwtToken.UserId, jwtToken.Issuer)
	}
	return ""
}

// Encode the JWT token and store it in the cache. This returns a signed string token.
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

// Decode the JWT token using the signed string token and public key.
// This returns a JWT token object.
func DecodeJWTToken(token string, publicKey string) (*types.JwtToken, error) {
	// Parse the token and get claims
	jwtToken, err := jwt.ParseWithClaims(token, &types.JwtToken{}, func(token *jwt.Token) (signedKey interface{}, err error) {
		if _, ok := token.Method.(*jwt.SigningMethodECDSA); !ok {
			return nil, fmt.Errorf("%s", fmt.Sprintf("Unexpected signing method: %v", token.Header["alg"]))
		}
		signedKey, err = jwt.ParseECPublicKeyFromPEM([]byte(publicKey))
		return
	})
	if err != nil {
		return nil, err
	} else if claims, ok := jwtToken.Claims.(*types.JwtToken); ok && jwtToken.Valid {
		return claims, nil
	}
	return nil, constants.HTTP_401_ERROR_MESSAGE()
}

// Validate the token by checking if it is cached.
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

// Apply the Argon2id hashing algorithm to the password and return the resulting hashed string.
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

// Verify if the Argon2id password matches the string.
func CompareArgon2id(password string, hashedPassword string) (bool, error) {
	initialHashedPassword, _ := DecodeBase64(hashedPassword)
	var match, err = argon2id.ComparePasswordAndHash(password, initialHashedPassword)
	return match, err
}

// Encodes the input string into Base64 format.
func EncodeBase64(data string) string {
	return base64.StdEncoding.EncodeToString([]byte(data))
}

// Decodes a Base64-encoded string and returns an error if the input is invalid.
func DecodeBase64(data string) (string, error) {
	var base64Text = make([]byte, base64.StdEncoding.DecodedLen(len(data)))
	var n, err = base64.StdEncoding.Decode(base64Text, []byte(data))
	var output = string(base64Text[:n])
	return output, err
}
