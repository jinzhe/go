package uri_test

import (
	"testing"

	"github.com/jinzhe/go/assert"
	"github.com/jinzhe/go/uri"
)

func TestEncode(t *testing.T) {
	assert.Equal(t, uri.Encode("https://zee.kim"), "https%3A%2F%2Fzee.kim")

}

func TestDecode(t *testing.T) {
	assert.Equal(t, uri.Decode("https%3A%2F%2Fzee.kim"), "https://zee.kim")
}
