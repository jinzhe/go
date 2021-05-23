package zip_test

import (
	"testing"

	"github.com/jinzhe/go/zip"
)

func TestZip(t *testing.T) {
	zip.In("./", "t.zip")
	zip.Out("t.zip")
}
