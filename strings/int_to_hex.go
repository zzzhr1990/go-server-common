package strings

import "strconv"

//IntToHex get int hex
func IntToHex(n int64) []byte {
	return []byte(strconv.FormatInt(n, 16))
}

//IntToHexString get hex string
func IntToHexString(n int64) string {
	return strconv.FormatInt(n, 16)
}

//string()
