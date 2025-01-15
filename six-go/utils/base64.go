package utils

import "encoding/base64"

func Base64Encode(value string) string {
	return base64.StdEncoding.EncodeToString([]byte(value))
}

func Base64Decode(b64 string) string {
	decodeString, err := base64.StdEncoding.DecodeString(b64)
	if err != nil {
		return ""
	}
	return string(decodeString)
}
