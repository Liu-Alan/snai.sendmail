package main

import (
	"fmt"
	"net/smtp"

	"github.com/jordan-wright/email"
)

type MailContext struct {
	From       string
	To         []string
	Subject    string
	Content    []byte
	SMTPServer string
	Port       int
	Pwd        string
}

func SendMail(mail *MailContext) error {
	em := email.NewEmail()

	em.From = mail.From
	em.To = mail.To
	em.Subject = mail.Subject
	em.Text = mail.Content
	addr := fmt.Sprintf("%s:%d", mail.SMTPServer, mail.Port)

	err := em.Send(addr, smtp.PlainAuth("", mail.From, mail.Pwd, mail.SMTPServer))
	if err != nil {
		return err
	}
	return nil
}

func main() {

	mail := MailContext{
		From:       "xxx@qq.com",                             // 发送者
		To:         []string{"yyy@qq.com"},                   // 接收者，可以多个
		Subject:    "golang 发送邮件",                            // 邮件主题
		Content:    []byte("Hello world， 咱们用 golang 发个邮件！！"), // 邮件内容
		SMTPServer: "smtp.qq.com",                            // 邮件服务器
		Port:       587,                                      // 端口
		Pwd:        "aaabbbccc",                              // 密码或授权码
	}

	err := SendMail(&mail)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("发送成功")
	}

	fmt.Scanln()
}
