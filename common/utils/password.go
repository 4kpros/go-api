package utils

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890!@#$%^&*()-=_+"
const letterBytesCode = "1234567890"

var src = rand.NewSource(time.Now().UnixNano())

const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

// Generates a random numeric code of specified length.
// Returns the generated code and an error if any.
func GenerateRandomCode(length int) (int, error) {
	safeLength := length
	if safeLength <= 0 {
		safeLength = 1
	}
	return strconv.Atoi(generateRandomValue(letterBytesCode, length))
}

// Returns a random password of the specified length.
func GenerateRandomPassword(length int) string {
	safeLength := length
	if safeLength <= 0 {
		safeLength = 1
	}
	return generateRandomValue(letterBytes, safeLength)
}

// Returns a random string of specified length, using provided characters.
// It's useful to generate passwords, OTP code and various other things
func generateRandomValue(letters string, length int) string {
	sb := strings.Builder{}
	sb.Grow(length)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := length-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letters) {
			sb.WriteByte(letters[idx])
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return sb.String()
}
