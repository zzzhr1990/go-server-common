package strings

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
)

//SHA256String Calc String sha256
func SHA256String(source string) string {
	h := sha256.New()
	h.Write([]byte(source))
	return hex.EncodeToString(h.Sum(nil))
}

//MD5String Calc String md5
func MD5String(source string) string {
	h := md5.New()
	h.Write([]byte(source))
	return hex.EncodeToString(h.Sum(nil))
}
