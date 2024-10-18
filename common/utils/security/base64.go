package security

import "encoding/base64"

// Encodes the input string into Base64 format.
func EncodeBase64(data string) string {
	return base64.StdEncoding.EncodeToString([]byte(data))
}

// Decodes a Base64-encoded string and returns an error if the input is invalid.
func DecodeBase64(data string) (string, error) {
	base64Text := make([]byte, base64.StdEncoding.DecodedLen(len(data)))
	n, err := base64.StdEncoding.Decode(base64Text, []byte(data))
	return string(base64Text[:n]), err
}
