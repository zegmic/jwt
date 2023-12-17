package jwt

import (
	"encoding/base64"
	"strings"
)

func encodeBase64URL(content []byte) string {
	enc := base64.URLEncoding.EncodeToString(content)
	enc = strings.Replace(enc, "+", "-", -1)
	enc = strings.Replace(enc, "/", "_", -1)
	enc = strings.Replace(enc, "=", "", -1)

	return enc
}
