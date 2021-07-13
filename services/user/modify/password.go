package modify

import (
	"regexp"

	"server/config/vars"
	"server/models"
	"server/tools"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
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
		passwordMatchString, _ := regexp.MatchString("(?=.*[a-z])(?=.*[A-Z])(?=.*[0-9]).{8,32}", password.Password2)
		switch {

		case decodePWD != nil || password.Password1 != password.Password2:
			c.SecureJSON(460, gin.H{
				"message": "Mistake password",
			})

		case !passwordMatchString:
			c.SecureJSON(461, gin.H{
				"message": "The password isn't secure enough",
			})

		default:
			hash, _ := bcrypt.GenerateFromPassword([]byte(password.Password2), bcrypt.DefaultCost) //加密处理
			encodePWD := string(hash)
			vars.DB0.Table("user").Model(&models.User{}).Where("id = ?", user.Id).Update("password", encodePWD)
			c.SecureJSON(200, gin.H{
				"message": "Modify the password successfully",
			})
		}
	}
}
