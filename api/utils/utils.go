package utils

import "encoding/base64"

func DecodeB64(data string) (string, error) {
	rawDecodedText, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return "", err
	}
	return string(rawDecodedText), nil
}
