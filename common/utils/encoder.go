package utils

import (
	base64 "encoding/base64"
)

func EncodeBase64(data string) (output string) {
	output = base64.StdEncoding.EncodeToString([]byte(data))
	return
}

func DecodeBase64(data string) (output string, err error) {
	base64Text := make([]byte, base64.StdEncoding.DecodedLen(len(data)))
	n, tempErr := base64.StdEncoding.Decode(base64Text, []byte(data))
	output = string(base64Text[:n])
	err = tempErr
	return
}
