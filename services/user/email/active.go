package email

import (
	vars2 "Server/config/vars"
	"Server/models"
	"Server/tools/token"
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type active struct {
	Code string `json:"code"`
}

func Active(c *gin.Context) {

	active := active{}
	t := c.GetHeader("Authorization")
	parse := token.Parse(t)
	id := parse.(jwt.MapClaims)["id"]
	rows, _ := vars2.DB0.Table("user").Model(&models.User{}).Where("id = ?", id).Rows()

	for rows.Next() {
		var user models.User
		_ = vars2.DB0.ScanRows(rows, &user)
		_ = c.ShouldBindJSON(&active)

		value, _ := vars2.RedisCode.Get(context.Background(), user.Email).Result()

		if active.Code != value {
			c.SecureJSON(403, gin.H{
				"message": "验证码错误",
			})
			return
		}

		vars2.DB0.Table("user").Model(&models.User{}).Where("email = ?", user.Email).Update("email_active", true) // 邮箱激活
		vars2.DB0.Table("user").Model(&models.User{}).Where("email = ?", user.Email).Update("is_active", true)    // 账户激活
		c.SecureJSON(200, gin.H{
			"message": fmt.Sprintf("电子邮件地址%s已激活", user.Email),
		})
	}
}
