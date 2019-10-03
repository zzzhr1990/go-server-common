package strings

import "encoding/base64"

// Base64URLSafeEncode encode bytes
func Base64URLSafeEncode(data []byte) string {
	return base64.URLEncoding.EncodeToString(data)
}

// Base64URLSafeEncodeString encode strings
func Base64URLSafeEncodeString(str string) string {
	return base64.URLEncoding.EncodeToString([]byte(str))
}

// Base64URLSafeDecodeString encode strings
func Base64URLSafeDecodeString(str string) ([]byte, error) {
	return base64.URLEncoding.DecodeString(str)
}
