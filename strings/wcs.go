package strings

import "strings"

// GetWcsKey wcs
func GetWcsKey(str string) (string, string) {
	idx := strings.Index(str, ":")
	if idx > -1 {
		bucket := str[:idx]
		key := str[idx+1:]
		return bucket, key
	}
	return "", str
}
