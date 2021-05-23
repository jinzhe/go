package base64

import (
	"encoding/base64"
)

// Encode
func Encode(options ...interface{}) string {
	if len(options) > 1 {
		return base64.RawStdEncoding.EncodeToString([]byte(options[0].(string)))
	}
	return base64.StdEncoding.EncodeToString([]byte(options[0].(string)))
}

// Decode
func Decode(options ...interface{}) string {
	var (
		s []byte
		e error
	)
	if len(options) > 1 {
		s, e = base64.RawStdEncoding.DecodeString(options[0].(string))
	} else {
		s, e = base64.StdEncoding.DecodeString(options[0].(string))
	}
	if e != nil {
		return ""
	}
	return string(s)
}
