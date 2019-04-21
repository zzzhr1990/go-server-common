package strings

import "net/url"

//EncodeURI encode strings to url safe
func EncodeURI(source string) string {

	u, err := url.Parse(source)
	if err != nil {
		return source
	}

	return u.String()

}

/*
func EncodeURI() string {
	u, err := url.Parse(source)

}
*/
