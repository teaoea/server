package user

import (
	"Server/models"
	"Server/tools/token"
	"Server/vars"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type passwordUpdate struct {
	Old       string `json:"old"`
	Password1 string `json:"password1"`
	Password2 string `json:"password2"`
}

func ChangePassword(c *gin.Context) {
	pwd := passwordUpdate{}
	_ = c.ShouldBindJSON(&pwd)

	t := c.GetHeader("Authorization")
	parse := token.Parse(t)
	id := parse.(jwt.MapClaims)["id"]
	rows, _ := vars.PDB0.Table("user").Model(&models.User{}).Where("id = ? ", id).Rows()

	for rows.Next() {
		var user models.User
		_ = vars.PDB0.ScanRows(rows, &user)

		decodePWD := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pwd.Old))

		switch {

		case decodePWD != nil || pwd.Password1 != pwd.Password2:
			c.SecureJSON(403, gin.H{
				"message": "密码错误",
			})

		case len(pwd.Password2) < 8 && len(pwd.Password2) > 16:
			c.SecureJSON(403, gin.H{
				"message": "密码需要大于8位,小于16位",
			})

		default:
			hash, _ := bcrypt.GenerateFromPassword([]byte(pwd.Password2), bcrypt.DefaultCost) //加密处理
			encodePWD := string(hash)
			vars.PDB0.Table("user").Model(&models.User{}).Where("id = ?", user.Id).Update("password", encodePWD)
			c.SecureJSON(200, gin.H{
				"message": "密码已修改完成",
			})
		}
	}
}
