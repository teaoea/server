package mail

import (
	"fmt"
	"net/smtp"
	"time"

	"server/config"

	"github.com/jordan-wright/email"
)

// SendMail
//example:
//	to = "recipient@example.com"
// 	subject = "subject"
//	content = []byte("<h1>Hello</h1>")
//	return @false send failed @true send successfully
func SendMail(to, subject string, content []byte) error {
	var (
		conf config.Config
		c    = conf.Yaml()
		ch   <-chan *email.Email
	)
	e := email.NewEmail()
	e.From = c.Mail.From
	e.To = []string{to}
	e.Subject = subject
	e.HTML = content
	addr := fmt.Sprintf("%s:%d", c.Mail.Smtp, c.Mail.Port)
	auth := smtp.PlainAuth("", c.Mail.User, c.Mail.Password, c.Mail.Smtp)

	pool, err := email.NewPool(addr, 4, auth)
	for i := 0; i < 4; i++ {
		go func() {
			for e := range ch {
				_ = pool.Send(e, 10*time.Second)
				return
			}
		}()
	}
	if err != nil {
		return err
	}
	return nil
}
