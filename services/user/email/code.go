package email

import (
	"Server/models"
	"Server/tools"
	"Server/tools/token"
	"Server/vars"
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

func SendCode(c *gin.Context) {

	t := c.GetHeader("Authorization")
	parse := token.Parse(t)
	id := parse.(jwt.MapClaims)["id"]
	rows, _ := vars.PDB0.Table("user").Model(&models.User{}).Where("id = ?", id).Rows()

	for rows.Next() {
		var user models.User
		_ = vars.PDB0.ScanRows(rows, &user)

		emailCheck := vars.PDB0.Table("user").Where(&models.User{Email: user.Email}, "email").Find(&user).RowsAffected
		switch {
		case emailCheck == 1:
			to := user.Email        //收件人邮箱
			subject := "验证你的电子邮件地址" // 邮件标题
			const a = "https://www.teaoea.com"
			code := tools.RandomDig(7) // 验证码
			html := []byte(fmt.Sprintf("<h1>你好,%s,验证你的电子邮件地址</h1>", user.Name) +
				fmt.Sprintf("<h3>你使用%s注册账户,为验证此电子邮件地址属于你,请在验证码输入框输入下方验证码,验证码有效期5分钟!!!", user.Email) +
				fmt.Sprintf("<h2><strong>%s</strong></h2>", code) +
				"<h2><strong>你收到此邮件原因:</strong></h2>" +
				fmt.Sprintf("<h3>有人使用此%s电子邮件地址,在<a href=%s>teaoea</a>注册了账户,如果你未注册账户,请忽视此邮件.</h3>", user.Email, a),
			) // 邮件正文,可以使用html语法,验证码使用<strong>标签,无障碍服务
			_ = tools.SendMail(strings.Fields(to), subject, html, "")
			vars.RDBCODE.Set(context.Background(), user.Email, code, time.Second*300)
			c.SecureJSON(200, gin.H{
				"message": fmt.Sprintf("验证码已发送到电子邮件地址%s", user.Email),
			})
		default:
			c.SecureJSON(403, gin.H{
				"message": "验证码发送失败",
			})
		}
	}
}
