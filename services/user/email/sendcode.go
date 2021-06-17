package email

import (
	"context"
	"fmt"
	"time"

	"server/config/vars"
	"server/models"
	"server/tools"
	"server/tools/mail"

	"github.com/gin-gonic/gin"
)

func SendCode(c *gin.Context) {

	value := c.GetHeader("Authorization")
	rows, _ := vars.DB0.Table("user").Model(&models.User{}).Where("id = ?", tools.Parse(value)).Rows()

	for rows.Next() {
		var user models.User
		_ = vars.DB0.ScanRows(rows, &user)

		emailCheck := vars.DB0.Table("user").Where(&models.User{Email: user.Email}, "email").Find(&user).RowsAffected
		switch {
		case emailCheck == 1:
			to := user.Email                       //recipient email address
			subject := "Verify your email address" // email subject
			code := tools.RandomDig(7)             // verification code
			// email content, you can use html, verification code use tag with "strong", accessibility service
			html := fmt.Sprintf("<h1>Hello,%s,verify your email address</h1>\n", user.Username) +
				fmt.Sprintf("<h3>You use %s to sign in an account. To verify that this email address belongs to you, please enter the verification code below in the verification code input box. The verification code is valid for 5 minutes!!!\n", user.Email) +
				fmt.Sprintf("<h2><strong>%s</strong></h2>", code) +
				"<h2><strong>The reason you received this email:</strong></h2>" +
				fmt.Sprintf("<h3>Someone uses this %s email address to register an account with <a href=\"https://www.teaoea.com\"> teaoea </a>. If you have not registered an account, please ignore this email.</h3>\n", user.Email)
			err := mail.SendMail(to, subject, html)
			if !err {
				c.SecureJSON(200, gin.H{
					"message": 1011,
				})
			} else {
				_ = vars.RedisCode.Set(context.TODO(), user.Email, code, time.Minute*5)
				c.SecureJSON(200, nil)
			}
		default:
			c.SecureJSON(200, gin.H{
				"message": 1012,
			})
		}
	}
}
