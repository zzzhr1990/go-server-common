package strings

import "encoding/base64"

//URLSafeEncode encode bytes
func URLSafeEncode(data []byte) string {
	return base64.URLEncoding.EncodeToString(data)
}

//URLSafeEncodeString encode strings
func URLSafeEncodeString(str string) string {
	return base64.URLEncoding.EncodeToString([]byte(str))
}

//URLSafeDecodeString encode strings
func URLSafeDecodeString(str string) ([]byte, error) {
	return base64.URLEncoding.DecodeString(str)
}
