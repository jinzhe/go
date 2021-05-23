package mail

import (
	"errors"
	"net/smtp"
	"strings"
)

// 发送邮件
func Send(options map[string]string) error {
	auth := smtp.PlainAuth("", options["user"], options["password"], options["host"])
	to := strings.Split(options["to"], ";")
	headers := make([]string, 0)
	headers = append(headers, "To: "+options["to"])
	headers = append(headers, "From: "+options["from"])
	headers = append(headers, "Subject: "+options["subject"])
	headers = append(headers, "Content-Type: text/html; charset=UTF-8")
	headers = append(headers, "")
	headers = append(headers, options["body"])

	data := []byte(strings.Join(headers, "\r\n"))
	e := smtp.SendMail(options["host"]+":"+options["port"], auth, options["from"], to, data)
	if e != nil {
		return errors.New("failed to send")
	}
	return e
}
