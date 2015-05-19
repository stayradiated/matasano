package matasano

import (
	"encoding/base64"
	"encoding/hex"
)

func ConvertHexToBase64(input string) (string, error) {
	bytes, err := hex.DecodeString(input)
	if err != nil {
		return "", err
	}
	output := base64.StdEncoding.EncodeToString(bytes)
	return output, nil
}
