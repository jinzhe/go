package base64_test

import (
	"testing"

	"github.com/jinzhe/go/assert"
	"github.com/jinzhe/go/base64"
)

func TestEncode(t *testing.T) {
	assert.Equal(t, base64.Encode("https://zee.kim"), "aHR0cHM6Ly96ZWUua2lt")
}

func TestDecode(t *testing.T) {
	assert.Equal(t, base64.Decode("aHR0cHM6Ly96ZWUua2lt"), "https://zee.kim")
}
