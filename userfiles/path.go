package userfiles

import (
	// "path/filepath"
	"strings"

	// github.com/zzzhr1990/go-server-common/userfiles
	gString "github.com/zzzhr1990/go-server-common/strings"
)

// FormatPath to format windows path like unix path but
func FormatPath(path string) string {
	// cleanStr := strings.Replace(path, ".", "", -1)
	cleanStr := strings.Replace(strings.Trim(path, " \r\n\t"), "\\", "/", -1)

	// split every
	s := strings.Split(cleanStr, "/")
	var sb strings.Builder

	for _, ss := range s {
		// sb.WriteString("/")
		clear := strings.Trim(ss, " \r\n\t")
		if clear != ".." && clear != "." && len(clear) > 0 {
			sb.WriteString("/")
			sb.WriteString(clear)
		}
	}

	cleanStr = sb.String()

	/*
		if strings.HasSuffix(cleanStr, "/") {
			cleanStr = cleanStr[:len(cleanStr)-1]
		}
	*/

	if !strings.HasPrefix(cleanStr, "/") {
		return "/" + cleanStr
	}

	return cleanStr
}

// getIdentity calc a file's uuid.
func getIdentity(path string) string {

	return gString.MD5String(path)
}

// GetFormatedIdentity calc a file's identity.
func GetFormatedIdentity(path string) (string, string) {
	formated := FormatPath(path)
	return getIdentity(formated), formated
}
