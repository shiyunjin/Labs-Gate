package util

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
)

func HmacSha1(data string, key string) string {

	// hmac ,use sha1
	mac := hmac.New(sha1.New, []byte(key))
	mac.Write([]byte(data))

	return hex.EncodeToString(mac.Sum(nil))
}