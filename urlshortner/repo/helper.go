package repo

import (
	"crypto/md5"
	b64 "encoding/base64"
	"hash"
)

var md5Controller hash.Hash = md5.New()

func CreateShortUrl(url string, length int) string {
	encodedString := b64.StdEncoding.EncodeToString(md5Controller.Sum([]byte(url)))
	var result string = ""
	for i := 1; i < 64; i = 2 * i {
		result += string(encodedString[i])
	}
	return result
}
