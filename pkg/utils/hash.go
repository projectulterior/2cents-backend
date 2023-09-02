package utils

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
)

// json.Marshal then sha256
func SHA256Base64(content interface{}) (string, error) {
	b, err := json.Marshal(content)
	if err != nil {
		return "", err
	}

	hash := sha256.Sum256(b)
	return base64.RawURLEncoding.EncodeToString(hash[:]), nil
}
