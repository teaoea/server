package mail

import (
	"crypto/tls"
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
