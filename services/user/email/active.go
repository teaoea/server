package email

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"server/config/vars"
	"server/models"
	"server/tools"
)

type active struct {
	Code string `json:"code"`
}

func Active(c *gin.Context) {
	value := c.GetHeader("Authorization")
	rows, _ := vars.DB0.Table("user").Model(&models.User{}).Where("id = ?", tools.Parse(value)).Rows()

	for rows.Next() {
		var (
			user   models.User
			active active
		)

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
