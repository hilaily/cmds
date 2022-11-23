package main

import (
	"fmt"

	"gopkg.in/gomail.v2"
)

/*
	发送邮件逻辑
*/

func SendMail(configPath, subject, msg string, to, attachs []string) {
	config := loadConfig(configPath)

	m := gomail.NewMessage()
	m.SetHeader("From", config.Mail)
	m.SetHeader("To", to...)
	//m.SetAddressHeader("Cc", "dan@example.com", "Dan")
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", msg)
	//m.SetBody("text/html", "Hello <b>Bob</b> and <i>Cora</i>!")

	for _, v := range attachs {
		m.Attach(v)
	}

	fmt.Printf("%+v\n", config)
	fmt.Println(subject, msg, to, attaches)

	d := gomail.NewDialer(config.SmtpHost, config.SmtpPort, config.Mail, config.Password)
	d.SSL = config.UseSsl

	err := d.DialAndSend(m)
	if err != nil {
		fmt.Println("send mail failed: ", err)
		return
	}

	fmt.Println("send mail succ")
}
