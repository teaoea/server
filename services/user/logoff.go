package user

import (
	"context"
	"fmt"
	"os"
	"time"

	"server/config/vars"
	"server/models"
	"server/tools"
	"server/tools/mail"

	"github.com/gin-gonic/gin"
)

func SendEmail(c *gin.Context) {
	value := c.GetHeader("Authorization")
	rows, _ := vars.DB0.Table("user").Model(&models.User{}).Where("id = ?", tools.Parse(value)).Rows()

	for rows.Next() {
		var user models.User
		_ = vars.DB0.ScanRows(rows, &user)

		emailCheck := vars.DB0.Table("user").Where(&models.User{Email: user.Email}, "email").Find(&user).RowsAffected
		switch {
		case emailCheck == 1:
			to := user.Email                 //recipient email address
			subject := "Delete your account" // email subject
			code := tools.RandomDig(7)       // verification code
			// email content, you can use html, verification code use tag with "strong", accessibility service
			html := "<h1 style=\"text-align: center\">Delete your account</h1>" +
				fmt.Sprintf("<h3><strong style=\"text-align: left\">verification code:%s</strong></h3>", code) +
				"<h3><strong>The verification code is valid for 5 minutes!!!</strong></h3>"

			err := mail.SendMail(to, subject, html)
			if !err {
				c.SecureJSON(403, gin.H{
					"message": fmt.Sprintf("failed to send mail to email address \"%s\"", user.Email),
				})
			} else {
				vars.RedisLogoff.Set(context.TODO(), user.Email, code, time.Minute*5)
				c.SecureJSON(200, gin.H{
					"message": fmt.Sprintf("the verification code has been sent to the email address \"%s\", please check the email!!!", user.Email),
				})
			}
		default:
			c.SecureJSON(403, gin.H{
				"message": "failed to send verification code",
			})
		}
	}
}
func Logoff(c *gin.Context) {

	value := c.GetHeader("Authorization")
	rows, _ := vars.DB0.Table("user").Model(&models.User{}).Where("id = ?", tools.Parse(value)).Rows()

	for rows.Next() {
		var (
			user   models.User
			logoff struct {
				models.User
				Code string `json:"code"`
			}
		)
		_ = vars.DB0.ScanRows(rows, &user)
		_ = c.ShouldBindJSON(&logoff)
		value, _ := vars.RedisLogoff.Get(context.TODO(), user.Email).Result()
		if logoff.Code != value {
			c.SecureJSON(403, gin.H{
				"message": "verification code error",
			})
		} else {
			// delete account
			vars.DB0.Table("user").Delete(&models.User{}, user.Id)
			// remove github account binding
			vars.DB0.Table("github").Delete(&models.Github{}, user.Id)
			// delete draft article
			vars.DB0.Table("draft").Where("author = ?", user.Username).Delete(&models.Article{})
			// delete file
			_ = os.RemoveAll(fmt.Sprintf("./static/article/draft/%d", user.Id))

			vars.DB0.Table("article").Where("author = ?", user.Username).Updates(map[string]interface{}{"author": "account deleted"})
			vars.DB0.Table("comment").Where("user = ?", user.Username).Updates(map[string]interface{}{"user": "account deleted"})
			vars.DB0.Table("reply").Where("user = ?", user.Username).Updates(map[string]interface{}{"user": "account deleted"})
			c.SecureJSON(200, nil)
		}
	}
}
