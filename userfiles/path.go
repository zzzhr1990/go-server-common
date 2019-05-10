package userfiles

import (
	"path/filepath"
	"strings"

	// github.com/zzzhr1990/go-server-common/userfiles
	gString "github.com/zzzhr1990/go-server-common/strings"
)

// FormatPath to format windows path like unix path but
func FormatPath(path string) string {
	cleanStr := filepath.Clean(strings.Replace(strings.Trim(path, " \r\n\t"), "\\", "/", -1))
	if cleanStr == "." {
		return "/"
	}
	if strings.HasPrefix(cleanStr, "..") {
		return FormatPath("/" + cleanStr)
	}
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
