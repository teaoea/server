package user

import (
	"fmt"
	"regexp"
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
	usernameMatchString, _ := regexp.MatchString("^[a-zA-Z]", register.Username)
	numberMatchString, _ := regexp.MatchString("^[0-9]", register.Number)

	switch {
	case register.Password != register.Password2:
		c.SecureJSON(460, gin.H{
			"message": "The two passwords entered are inconsistent",
		})

	case !tools.CheckPassword(register.Password2):
		c.SecureJSON(461, gin.H{
			"message": "The password isn't secure enough",
		})

	case !mail.SuffixCheck(register.Email):
		c.SecureJSON(462, gin.H{
			"message": "Email address suffix cannot be used for registration",
		})

	case nameCheck != 0 || register.Username == "":
		c.SecureJSON(463, gin.H{
			"message": "Username has been signed up",
		})

	case !usernameMatchString:
		c.SecureJSON(464, gin.H{
			"message": "Username can only be English characters",
		})

	case emailCheck != 0 || register.Email == "":
		c.SecureJSON(465, gin.H{
			"message": "Email address has been signed up",
		})

	case numberCheck != 0 || register.Number == "":
		c.SecureJSON(466, gin.H{
			"message": "Phone number has been signed up",
		})
	case !numberMatchString:
		c.SecureJSON(467, gin.H{
			"message": "The phone number can only be a number",
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
		c.SecureJSON(200, gin.H{
			"message": "Signed up successfully",
		})
	}
}
