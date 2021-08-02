package modify

import (
	"context"

	"server/config/vars"
	"server/models"
	"server/tools"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type modifyPassword struct {
	Code      string `json:"Code"`
	Password1 string `json:"password1"`
	Password2 string `json:"password2"`
}

func Password(c *gin.Context) {

	value := c.GetHeader("Authorization")
	rows, _ := vars.DB0.Table("user").Model(&models.User{}).Where("id = ?", tools.Parse(value)).Rows()

	for rows.Next() {
		var (
			user     models.User
			password modifyPassword
		)

		_ = vars.DB0.ScanRows(rows, &user)
		_ = c.ShouldBindJSON(&password)

		value, _ := vars.RedisPasswordCode.Get(context.Background(), user.Email).Result()
		if password.Code != value {
			c.SecureJSON(460, gin.H{
				"message": "Mistake verification code",
			})
			return
		}

		if !tools.CheckPassword(password.Password2) {
			c.SecureJSON(461, gin.H{
				"message": "The password isn't secure enough",
			})
			return
		}

		hash, _ := bcrypt.GenerateFromPassword([]byte(password.Password2), bcrypt.DefaultCost) //加密处理
		encodePWD := string(hash)
		vars.DB0.Table("user").Model(&models.User{}).Where("id = ?", user.Id).Update("password", encodePWD)
		c.SecureJSON(200, gin.H{
			"message": "Modify the password successfully",
		})
	}
}
