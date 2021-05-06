package tools

import (
	"context"
	"fmt"
	"github.com/jordan-wright/email"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"net/smtp"
	"server/config/vars"
	"strings"
)

// SendMail example:
// 	subject = "标题"
//	to = "recipient@example.com"
//	body = []byte("<h1>hello</h1>")
//	filename = "./file.txt" // 如果不需要发送附件,使用""
//  SendMail(strings.Fields(to),subject,body,filename)
func SendMail(to []string, subject string, body []byte, filename string) *email.Pool {

	mail := email.NewEmail()
	mail.To = to
	mail.From = vars.MailForm
	mail.Subject = subject
	mail.HTML = body
	_, _ = mail.AttachFile(filename)
	auth := smtp.PlainAuth("", vars.MailUser, vars.MailPassword, vars.MailSmtp)
	pool, _ := email.NewPool(fmt.Sprintf("%s:%s", vars.MailSmtp, vars.MailPort), 4, auth)

	return pool
}

// SuffixCheck
// email: 邮箱地址
// doc: 根据参数email截取邮箱后缀,在数据库MongoDB,数据库名conf,集合suffixes查找后缀是否存在
// return 后缀存在返回true,不存在返回false
func SuffixCheck(email string) bool {
	var suffix struct {
		Suffix string
	}

	addr := strings.Split(email, "@") // 字符串分割
	suf := "@" + addr[1]              // 截取邮箱后缀
	filter := bson.D{
		bson.E{Key: "suffix", Value: suf},
	}
	val := vars.MongoSuffix.FindOne(context.TODO(), filter).Decode(&suffix)
	return val != mongo.ErrNoDocuments
}
