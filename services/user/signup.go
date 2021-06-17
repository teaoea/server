package user

import (
	"fmt"
	"time"

	"server/config/vars"
	"server/models"
	"server/tools"
	"server/tools/mail"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *gin.Context) {

	var (
		user     models.User
		register struct {
			models.User
			Password2 string `json:"password2"`
			Country   string `json:"country"`
		}
	)
	_ = c.ShouldBindJSON(&register)

	nameCheck := vars.DB0.Table("user").Where(&models.User{Username: register.Username}, "username").Find(&user).RowsAffected

	emailCheck := vars.DB0.Table("user").Where(&models.User{Email: register.Email}, "email").Find(&user).RowsAffected

	numberCheck := vars.DB0.Table("user").Where(&models.User{Number: register.Number}, "number").Find(&user).RowsAffected

	switch {

	case register.Password != register.Password2:
		c.SecureJSON(200, gin.H{
			"message": "1003",
		})

	case len(register.Password2) < 8 || len(register.Password2) > 32:
		c.SecureJSON(200, gin.H{
			"message": 1004,
		})

	case !mail.SuffixCheck(register.Email):
		// addr := strings.Split(register.Email, "@")
		// suffix := "@" + addr[1]
		c.SecureJSON(200, gin.H{
			"message": 1005,
		})

	case nameCheck != 0 || register.Username == "":
		c.SecureJSON(200, gin.H{
			"message": 1006,
		})

	case emailCheck != 0 || register.Email == "":
		c.SecureJSON(200, gin.H{
			"message": 1007,
		})

	case numberCheck != 0 || register.Number == "":
		c.SecureJSON(200, gin.H{
			"message": 1008,
		})

	default:
		hash, _ := bcrypt.GenerateFromPassword([]byte(register.Password2), bcrypt.DefaultCost) // 加密密码
		encodePWD := string(hash)

		user = models.User{
			Id:        tools.NewId(),
			Username:  register.Username,
			Password:  encodePWD,
			Email:     register.Email,
			Number:    fmt.Sprintf("%s-%s", register.Country, register.Number),
			CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
		}

		vars.DB0.Table("user").Create(&user)
		c.SecureJSON(200, nil)
	}
}
