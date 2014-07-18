package securerandom

import (
	"crypto/rand"
	"encoding/base64"
	"strings"
)


func UrlSafeBase64Padded(n int) (string, error) {
	bytes := make([]byte, n)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	result := base64.StdEncoding.EncodeToString(bytes)
	return strings.Replace(result, "\n", "", -1), nil
}

func UrlSafeBase64(n int) (string, error) {
	result, err := UrlSafeBase64Padded(n)
	return strings.Replace(result, "=", "", -1), err
}
