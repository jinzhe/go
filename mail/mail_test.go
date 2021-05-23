package mail_test

import (
	"testing"

	"github.com/jinzhe/go/mail"
)

func TestSend(t *testing.T) {
	e := mail.Send(map[string]string{
		"to":       "zee.kim@qq.com",
		"subject":  "Subject",
		"body":     "Body",
		"from":     "zee.kim@qq.com",
		"host":     "ssl://smtp.exmail.qq.com",
		"port":     "465",
		"user":     "smtp@zee.kim",
		"password": "*****",
	})
	if e != nil {
		t.Skipped()
	}
}
