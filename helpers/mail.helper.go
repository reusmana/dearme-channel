package helpers

import (
	"crypto/tls"
	"fmt"

	"gopkg.in/gomail.v2"
)

func KirimEmail(subject string, pesan string, path string) {

	mailer := gomail.NewMessage()
	mailer.SetHeader("From", "wms-noreply@ramayana.co.id")
	mailer.SetHeader("To", "eggy.nuko@ramayana.co.id")
	//mailer.SetAddressHeader("Cc", "tralalala@gmail.com", "CEKCEK")
	mailer.SetHeader("Subject", subject)
	// mailer.SetBody("text/html", "Hello, <b>have a nice day</b>")
	mailer.SetBody("text/html", pesan)
	// mailer.Attach("./sample.png")

	if path != "" {
		mailer.Attach(path)
	}

	dialer := &gomail.Dialer{Host: "172.16.1.111", Port: 25}
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	err := dialer.DialAndSend(mailer)
	if err != nil {
		fmt.Println("send mail error")
		fmt.Println(err.Error())
	}

}
