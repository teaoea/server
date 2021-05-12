package user

import (
	"fmt"
	"server/config/vars"
	"server/models"
	"server/tools"
	"server/tools/mail"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	var register struct {
		models.User
		Password2 string `json:"password2"`
		Country   string `json:"country"`
	}
	var user models.User
	_ = c.ShouldBindJSON(&register)

	nameCheck := vars.DB0.Table("user").Where(&models.User{Name: register.Name}, "name").Find(&user).RowsAffected

	emailCheck := vars.DB0.Table("user").Where(&models.User{Email: register.Email}, "email").Find(&user).RowsAffected

	numberCheck := vars.DB0.Table("user").Where(&models.User{Number: register.Number}, "number").Find(&user).RowsAffected

	switch {

	case register.Password != register.Password2: // 校验密码是否一致
		c.SecureJSON(403, gin.H{
			"message": "两次输入的密码不一致",
		})

	case len(register.Password2) < 8 || len(register.Password2) > 32: // 校验密码是否安全
		c.SecureJSON(403, gin.H{
			"message": "密码需要大于8位,小于16位",
		})

	case !mail.SuffixCheck(register.Email): // 校验电子邮件服务商是否运行注册
		addr := strings.Split(register.Email, "@") // 字符串分割
		suffix := "@" + addr[1]                    // 截取邮箱后缀
		c.SecureJSON(403, gin.H{
			"message": fmt.Sprintf("邮箱后缀%s无法用于注册", suffix),
		})

	case nameCheck != 0 || register.Name == "": // 校验用户名是否已被注册
		c.SecureJSON(403, gin.H{
			"message": fmt.Sprintf("用户名%s已被注册", register.Name),
		})

	case emailCheck != 0 || register.Email == "": // 校验电子邮件地址是否已被注册
		c.SecureJSON(403, gin.H{
			"message": fmt.Sprintf("电子邮件地址%s已被注册", register.Email),
		})

	case numberCheck != 0 || register.Number == "": // 校验手机号是否已被注册
		c.SecureJSON(403, gin.H{
			"message": fmt.Sprintf("手机号%s已被注册", register.Number),
		})

	default:
		hash, _ := bcrypt.GenerateFromPassword([]byte(register.Password2), bcrypt.DefaultCost) // 加密密码
		encodePWD := string(hash)

		user = models.User{
			Id:        tools.NewId(),
			Name:      register.Name,
			Password:  encodePWD,
			Email:     register.Email,
			Number:    fmt.Sprintf("%s-%s", register.Country, register.Number),
			CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
		}

		vars.DB0.Table("user").Create(&user)
		c.SecureJSON(200, gin.H{
			"message": fmt.Sprintf("%s已完成注册", register.Name),
		})
	}
}
