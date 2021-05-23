package request_test

import (
	"testing"

	"github.com/jinzhe/go/request"
)

func TestGet(t *testing.T) {
	data := request.Get("https://baidu.com")
	if data == "" {
		t.Error("Cant't get data!")
	}
}

func TestPost(t *testing.T) {
	data := request.Post("https://baidu.com", map[string]string{
		"test": "9",
	})
	if data == "" {
		t.Error("Cant't get data!")
	}
}
