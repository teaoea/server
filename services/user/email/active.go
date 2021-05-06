package email

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"server/config/vars"
	"server/models"
	"server/services/user/auth"
)

type active struct {
	Code string `json:"code"`
}

func Active(c *gin.Context) {

	active := active{}
	token := c.GetHeader("Authorization")
	parse := auth.Parse(token)
	id := parse.(jwt.MapClaims)["id"]
	rows, _ := vars.DB0.Table("user").Model(&models.User{}).Where("id = ?", id).Rows()

	for rows.Next() {
		var user models.User
		_ = vars.DB0.ScanRows(rows, &user)
		_ = c.ShouldBindJSON(&active)

		value, _ := vars.RedisCode.Get(context.Background(), user.Email).Result()

		if active.Code != value {
			c.SecureJSON(403, gin.H{
				"message": "验证码错误",
			})
			return
		}

		vars.DB0.Table("user").Model(&models.User{}).Where("email = ?", user.Email).Update("email_active", true) // 邮箱激活
		vars.DB0.Table("user").Model(&models.User{}).Where("email = ?", user.Email).Update("is_active", true)    // 账户激活
		c.SecureJSON(200, gin.H{
			"message": fmt.Sprintf("电子邮件地址%s已激活", user.Email),
		})
	}
}
