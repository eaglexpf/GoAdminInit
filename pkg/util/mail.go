package util

import (
	"fmt"
	"net/smtp"
	"strings"

	"github.com/eaglexpf/GoAdminInit/pkg/setting"
)

func SendMail(to []string, subject, body string) bool {
	sec, err := setting.Cfg.GetSection("mail")
	if err != nil {
		fmt.Println("Fail to get section 'mail': %v", err)
	}
	username := sec.Key("UserName").String()
	password := sec.Key("Password").String()
	host := sec.Key("Host").String()
	port := sec.Key("Port").String()
	nickname := sec.Key("NickName").String()
	auth := smtp.PlainAuth("", username, password, host)
	content_type := "Content-Type: text/plain; charset=UTF-8"
	msg := []byte("To: " + strings.Join(to, ",") + "\r\nFrom: " + nickname +
		"<" + username + ">\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
	err = smtp.SendMail(host+":"+port, auth, username, to, msg)
	if err != nil {
		fmt.Printf("send mail error: %v", err)
		return false
	}
	return true
}
