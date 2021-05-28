package modify

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"server/config/vars"
	"server/models"
	"server/tools"
)

type modifyPassword struct {
	Old       string `json:"old"`
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

		decodePWD := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password.Old))

		switch {

		case decodePWD != nil || password.Password1 != password.Password2:
			c.SecureJSON(403, gin.H{
				"message": "密码错误",
			})

		case len(password.Password2) < 8 && len(password.Password2) > 16:
			c.SecureJSON(403, gin.H{
				"message": "密码需要大于8位,小于16位",
			})

		default:
			hash, _ := bcrypt.GenerateFromPassword([]byte(password.Password2), bcrypt.DefaultCost) //加密处理
			encodePWD := string(hash)
			vars.DB0.Table("user").Model(&models.User{}).Where("id = ?", user.Id).Update("password", encodePWD)
			c.SecureJSON(200, gin.H{
				"message": "密码已修改完成",
			})
		}
	}
}
