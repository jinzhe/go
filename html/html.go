package html

import (
	"html"
	"regexp"
	"strings"
)

// Encode html chars to string
func Encode(value string) string {
	return html.EscapeString(value)
}

// Decode html string to html chars
func Decode(value string) string {
	return html.UnescapeString(value)
}

// Clean html tags
func Clean(source string) string {
	re := regexp.MustCompile(`(?s)<(?:style|script)[^<>]*>.*?</(?:style|script)>|</?[a-z][a-z0-9]*[^<>]*>|<!--.*?-->|&nbsp;|&amp;`)
	source = re.ReplaceAllString(source, "")
	re = regexp.MustCompile(`\s{2,}`)
	source = re.ReplaceAllString(source, "\n")
	return strings.TrimSpace(source)
}
