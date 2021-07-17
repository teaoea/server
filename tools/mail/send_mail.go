package mail

import (
	"crypto/tls"
	"fmt"
	"net/smtp"
	"time"

	"server/config/vars"

	"github.com/xhit/go-simple-mail/v2"
)

// SendMail
//example:
//	to = "recipient@example.com"
// 	subject = "subject"
//	html = <h1>Hello World</h1>
//	return @false send failed @true send successfully
func SendMail(to, subject, html string) bool {
	e := mail.NewSMTPClient()
	e.Host = vars.MailSmtp
	e.Port = vars.MailPort
	e.Username = vars.MailUser
	e.Password = vars.MailPassword
	e.Encryption = mail.EncryptionSTARTTLS
	e.KeepAlive = false
	e.ConnectTimeout = time.Second * 10
	e.SendTimeout = time.Second * 10
	e.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	smtpClient, err := e.Connect()
	if err != nil {
		return false
	}
	send := mail.NewMSG()
	send.SetFrom(vars.MailForm).
		AddTo(to).
		SetSubject(subject).
		SetBody(mail.TextHTML, html)
	err = send.Send(smtpClient)
	return err == nil
}

func SendMail1(subject, content, to string) (bool, error) {
	sub := fmt.Sprintf("%s\r\n", subject)
	form := vars.MailForm
	contentType := "Content-Type: text/html; charset=UTF-8\r\n\r\n"
	message := []byte(sub + form + to + contentType + content)
	addr := fmt.Sprintf("%s:%d", vars.MailSmtp, vars.MailPort)
	auth := smtp.PlainAuth("", vars.MailUser, vars.MailPassword, vars.MailSmtp)
	send := smtp.SendMail(addr, auth, form, []string{to}, message)
	if send != nil {
		return false, send
	}
	return true, nil
}
