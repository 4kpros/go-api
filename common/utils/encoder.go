package utils

import (
	base64 "encoding/base64"
)

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
