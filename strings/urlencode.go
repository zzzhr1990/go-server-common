package strings

import (
	"net/url"
	"strings"
)

//EncodeURI encode strings to url safe
func EncodeURI(source string) string {

	u, err := url.Parse(source)
	if err != nil {
		return source
	}

	return u.String()

}

// EncodeURIComponent encode url
func EncodeURIComponent(source string) string {
	return url.QueryEscape(source)
}

// EncodeURIComponentCap like js.
func EncodeURIComponentCap(source string) string {
	r := url.QueryEscape(source)
	return strings.Replace(r, "+", "%20", -1)
}

/*
func EncodeURI() string {
	u, err := url.Parse(source)

}
*/
