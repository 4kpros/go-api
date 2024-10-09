package utils

import (
	"crypto/ecdsa"
	"fmt"
	"runtime"
	"strconv"
	"time"

	"github.com/4kpros/go-api/common/types"
	"github.com/4kpros/go-api/config"
	"github.com/alexedwards/argon2id"
	"github.com/golang-jwt/jwt/v5"
)

var JwtIssuerSession = "session"
var JwtIssuerActivate = "activate"
var JwtIssuerResetCode = "resetCode"
var JwtIssuerResetNewPassword = "resetNewPassword"

func NewExpiresDateDefault() *time.Time {
	tempDate := time.Now().Add(time.Minute * time.Duration(config.Env.JwtExpiresDefault))
	return &tempDate
}

func NewExpiresDateSignIn(stayConnected bool) (date *time.Time) {
	if stayConnected {
		tempDate := time.Now().Add(time.Hour * time.Duration(24*config.Env.JwtExpiresSignInStayConnected))
		return &tempDate
	}
	tempDate := time.Now().Add(time.Minute * time.Duration(config.Env.JwtExpiresSignIn))
	return &tempDate
}

func GetCachedKey(jwtToken *types.JwtToken) string {
	return fmt.Sprintf("%d%s%s%d", jwtToken.UserId, jwtToken.Issuer, jwtToken.Device, jwtToken.Code)
}

func IsSameCachedKey(jwtToken1 *types.JwtToken, jwtToken2 *types.JwtToken) bool {
	if jwtToken1 == nil || jwtToken2 == nil {
		return false
	}
	if GetCachedKey(jwtToken1) != GetCachedKey(jwtToken2) {
		return false
	}

	return true
}

// Encrypt JWT
func EncryptJWTToken(jwtToken *types.JwtToken, privateKey string, loadCached bool) (newJwt *types.JwtToken, tokenStr string, err error) {
	// Check if there is some cached token
	if loadCached {
		tokenStr, err = config.GetRedisVal(GetCachedKey(jwtToken))
		if err == nil && len(tokenStr) > 0 {
			jwtDecrypted, errDecrypted := DecryptJWTToken(tokenStr, config.Keys.JwtPublicKey)
			if errDecrypted == nil && IsSameCachedKey(jwtToken, jwtDecrypted) {
				newJwt = jwtDecrypted
				return
			}
		}
	}

	// Otherwise generate new one. Also only add string on MapClaims, others types are not regonised
	token := jwt.NewWithClaims(jwt.SigningMethodES512, jwt.MapClaims{
		"iss":    jwtToken.Issuer,
		"userId": fmt.Sprintf("%d", jwtToken.UserId),
		"role":   jwtToken.Role,
		"device": jwtToken.Device,
		"exp":    jwt.NewNumericDate(jwtToken.Expires),
		"iat":    jwt.NewNumericDate(time.Now()),
		"code":   fmt.Sprintf("%d", jwtToken.Code),
	})
	var signedKey *ecdsa.PrivateKey
	signedKey, err = jwt.ParseECPrivateKeyFromPEM([]byte(privateKey))
	if err != nil {
		return
	}
	tokenStr, err = token.SignedString(signedKey)
	if err != nil {
		return
	}

	// Cache new token
	config.SetRedisVal(GetCachedKey(jwtToken), tokenStr)
	newJwt = jwtToken
	return
}

// Decrypt JWT
func DecryptJWTToken(tokenStr string, publicKey string) (*types.JwtToken, error) {
	// Parse the token and get claims
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (signedKey interface{}, err error) {
		// Validate the alg
		if _, ok := token.Method.(*jwt.SigningMethodECDSA); !ok {
			message := fmt.Sprintf("Unexpected signing method: %v", token.Header["alg"])
			return nil, fmt.Errorf("%s", message)
		}
		signedKey, err = jwt.ParseECPublicKeyFromPEM([]byte(publicKey))
		return
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)

	// Check if the token is valid
	if !ok || !token.Valid {
		message := "Invalid token or expired! Please enter valid information."
		return nil, fmt.Errorf("%s", message)
	}

	// Extract information
	iss := fmt.Sprintf("%s", claims["iss"])
	userId, _ := strconv.Atoi(fmt.Sprintf("%s", claims["userId"]))
	roleStr := fmt.Sprintf("%s", claims["role"])
	role, _ := strconv.Atoi(roleStr)
	codeStr := fmt.Sprintf("%s", claims["code"])
	code, _ := strconv.Atoi(codeStr)
	device := fmt.Sprintf("%s", claims["device"])

	// Return it
	return &types.JwtToken{
		UserId: uint(userId),
		Role:   role,
		Code:   code,
		Device: device,
		Issuer: iss,
	}, nil
}

// Verify JWT
func VerifyJWTToken(tokenStr string, publicKey string) (*types.JwtToken, bool) {
	// Decrypt the token
	jwtDecrypted, errDecrypted := DecryptJWTToken(tokenStr, publicKey)
	if errDecrypted != nil || jwtDecrypted == nil {
		return nil, false
	}

	// Check if the token is cached
	tokenStrCached, errCached := config.GetRedisVal(GetCachedKey(jwtDecrypted))
	if errCached != nil || len(tokenStrCached) <= 0 {
		return nil, false
	}
	if tokenStr != tokenStrCached {
		return nil, false
	}
	return jwtDecrypted, true
}

// Encrypt jwtToken using Argon2id
func EncryptWithArgon2id(password string) (hash string, err error) {
	params := &argon2id.Params{
		Memory:      uint32(config.Env.ArgonMemoryLeft * config.Env.ArgonMemoryRight),
		Iterations:  uint32(config.Env.ArgonIterations),
		Parallelism: uint8(runtime.NumCPU()),
		SaltLength:  uint32(config.Env.ArgonSaltLength),
		KeyLength:   uint32(config.Env.ArgonKeyLength),
	}
	tempHash, tempErr := argon2id.CreateHash(password, params)
	err = tempErr
	if err != nil {
		return
	}
	// Encode to base64
	hash = EncodeBase64(tempHash)
	return
}

// Verify if Argon2id password matches string
func CompareToArgon2id(password string, hashedPassword string) (match bool, err error) {
	initialHashedPassword, _ := DecodeBase64(hashedPassword)
	match, err = argon2id.ComparePasswordAndHash(password, initialHashedPassword)
	return
}
