package tools

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/mail"
	"net/smtp"

	"server/config"
)

// SendMail
//example:
//  name = "recipient"
//	recipient = "recipient@example.com"
// 	subject = "subject"
//	content = "<h1>Hello</h1>"
//	return @false send failed @true send successfully
func SendMail(recipient, sub, content string) bool {
	var (
		conf config.Config
		c    = conf.Yaml()
	)

	from := mail.Address{Name: c.Mail.From, Address: c.Mail.User} // sender
	to := mail.Address{Name: "", Address: recipient}              // recipient
	headers := make(map[string]string)
	headers["From"] = from.String()
	headers["To"] = to.String()
	headers["Subject"] = sub
	headers["Content-Type"] = "text/html; charset=UTF-8"
	message := ""
	for key, value := range headers {
		message += fmt.Sprintf("%s: %s\r\n", key, value)
	}
	message += "\r\n" + content

	server := fmt.Sprintf("%s:%d", c.Mail.Smtp, c.Mail.Port)
	host, _, _ := net.SplitHostPort(server)
	auth := smtp.PlainAuth("", c.Mail.User, c.Mail.Password, host)

	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         host,
	}

	// When the program starts,use tcp to test the connectivity of the mail service
	conn, err := tls.Dial("tcp", server, tlsConfig)
	if err != nil {
		panic(err)
	}

	client, err := smtp.NewClient(conn, host)
	if err != nil {
		Err("tools/mail/send_mail.go->smtp.NewClient", err)
		return false
	}

	if err = client.Auth(auth); err != nil {
		Err("tools/mail/send_mail.go->client.Auth", err)
		return false
	}

	if err = client.Mail(from.Address); err != nil {
		Err("tools/mail/send_mail.go->client.Mail", err)
		return false
	}

	if err = client.Rcpt(to.Address); err != nil {
		Err("tools/mail/send_mail.go->client.Rcpt", err)
		return false
	}

	data, err := client.Data()
	if err != nil {
		Err("tools/mail/send_mail.go->client.Data", err)
		return false
	}

	_, err = data.Write([]byte(message))
	if err != nil {
		Err("tools/mail/send_mail.go->data.Write", err)
		return false
	}

	_ = client.Quit()
	return true
}
