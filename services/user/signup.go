package user

import (
	"regexp"
	"time"

	"server/config/vars"
	"server/models"
	"server/tools"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *gin.Context) {

	var (
		user   models.User
		signup struct {
			models.User
			Password2 string `json:"password2"`
		}
	)
	_ = c.ShouldBindJSON(&signup)

	nameCheck := vars.DB0.Table("user").Where(&models.User{Username: signup.Username}, "username").Find(&user).RowsAffected
	emailCheck := vars.DB0.Table("user").Where(&models.User{Email: signup.Email}, "email").Find(&user).RowsAffected
	numberCheck := vars.DB0.Table("user").Where(&models.User{Phone: signup.Phone}, "number").Find(&user).RowsAffected
	usernameMatchString, _ := regexp.MatchString("^[a-zA-Z]", signup.Username)
	numberMatchString, _ := regexp.MatchString("^[0-9]", signup.Phone)

	switch {
	case signup.Password != signup.Password2:
		c.SecureJSON(460, gin.H{
			"message": "The two passwords entered are inconsistent",
		})

	case !tools.CheckPassword(signup.Password2):
		c.SecureJSON(461, gin.H{
			"message": "The password isn't secure enough",
		})

	case !tools.SuffixCheck(signup.Email):
		c.SecureJSON(462, gin.H{
			"message": "Email address suffix cannot be used for registration",
		})

	case nameCheck != 0 || signup.Username == "":
		c.SecureJSON(463, gin.H{
			"message": "Username has been signed up",
		})

	case !usernameMatchString:
		c.SecureJSON(464, gin.H{
			"message": "Username can only be English characters",
		})

	case emailCheck != 0 || signup.Email == "":
		c.SecureJSON(465, gin.H{
			"message": "Email address has been signed up",
		})

	case numberCheck != 0 || signup.Phone == "":
		c.SecureJSON(466, gin.H{
			"message": "Phone number has been signed up",
		})
	case !numberMatchString:
		c.SecureJSON(467, gin.H{
			"message": "The phone number can only be a number",
		})
	default:
		hash, _ := bcrypt.GenerateFromPassword([]byte(signup.Password2), bcrypt.DefaultCost) // 加密密码
		encodePWD := string(hash)

		user = models.User{
			Id:        tools.NewId(),
			Username:  signup.Username,
			Password:  encodePWD,
			Email:     signup.Email,
			Prefix:    signup.Prefix,
			Phone:     signup.Phone,
			CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
		}

		vars.DB0.Table("user").Create(&user)
		c.SecureJSON(200, gin.H{
			"message": "Signed up successfully",
		})
	}
}
