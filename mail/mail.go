package mail

import (
	"fmt"
	"net/smtp"
	"strings"
)

// 发送邮件
func Send(params ...string) error {
	to := params[0]
	subject := params[1]
	body := params[2]
	from := params[3]
	host := params[4]
	port := params[5]
	user := params[6]
	password := params[7]
	auth := smtp.PlainAuth("", user, password, host)
	mailto := strings.Split(to, ";")
	data := []byte("To: " + to + "\r\nFrom: " + from + "\r\nSubject: " + subject + "\r\nContent-Type: text/html; charset=UTF-8\r\n\r\n" + body)
	e := smtp.SendMail(host+":"+port, auth, from, mailto, data)
	if e != nil {
		fmt.Println(data, e)
	} else {
		fmt.Println("Send OK!")
	}
	return e
}
