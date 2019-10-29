package strings

import (
	"errors"
	sysString "strings"
)

// ConvertToCheckString conv
func ConvertToCheckString(source string, key string) string {
	check := MD5String(source)
	return Base64URLSafeEncodeString(check + ":_c_:" + source)
}

// RecoveryCheckString conv
func RecoveryCheckString(sourceB64 string) (string, error) {
	sourceByte, err := Base64URLSafeDecodeString(sourceB64)
	if err != nil {
		return "", err
	}
	source := string(sourceByte)
	sp := sysString.Index(source, ":_c_:")
	if sp < 0 {
		return "", errors.New("Cannot decode string")
	}
	prev := source[:sp]
	suf := source[sp+5:]
	check := MD5String(suf)
	if prev != check {
		return "", errors.New("Source check failed")
	}
	return suf, nil
}
