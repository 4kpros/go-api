package security

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

// EncodeHMAC_SHA256 Generates an HMAC-SHA256 signature for the given message using the provided private key.
// Returns the hex-encoded signature as a string and any error encountered.
func EncodeHMAC_SHA256(message string, privateKey string) (string, error) {
	mac := hmac.New(sha256.New, []byte(privateKey))
	_, err := mac.Write([]byte(message))
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(mac.Sum(nil)), nil
}

// VerifyHmacSha256 Verifies the HMAC-SHA256 signature of a message using the provided private key.
// Returns true if the signature is valid, false otherwise.
func VerifyHmacSha256(message string, privateKey string, hash string) (bool, error) {
	sig, err := hex.DecodeString(hash)
	if err != nil {
		return false, err
	}
	mac := hmac.New(sha256.New, []byte(privateKey))
	mac.Write([]byte(message))

	return hmac.Equal(sig, mac.Sum(nil)), nil
}
