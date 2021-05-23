package uri

import (
	"net/url"
)

// URIEncode url encode string, is + not %20
func Encode(value string) string {
	return url.QueryEscape(value)
}

// URIDecode url decode string
func Decode(value string) string {
	s, e := url.QueryUnescape(value)
	if e != nil {
		return ""
	}
	return s
}
