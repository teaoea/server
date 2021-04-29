package user

import (
	vars2 "Server/config/vars"
	"Server/models"
	"Server/tools/token"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func UpdateUser(c *gin.Context) {
	t := c.GetHeader("Authorization")
	parse := token.Parse(t)
	id := parse.(jwt.MapClaims)["id"]
	rows, _ := vars2.DB0.Table("user").Model(&models.User{}).Where("id = ?", id).Rows()
	for rows.Next() {
		var user models.User
		var u struct {
			models.User
		}
		_ = vars2.DB0.ScanRows(rows, &user)
		_ = c.ShouldBindJSON(&u)
		switch {

		case u.Avatar != "":
			vars2.DB0.Table("user").Model(&models.User{}).Where("id = ?", user.Id).Update("avatar", u.Avatar)

		case u.Gender != "":
			vars2.DB0.Table("user").Model(&models.User{}).Where("id = ?", user.Id).Update("gender", u.Gender)

		case u.Introduction != "":
			vars2.DB0.Table("user").Model(&models.User{}).Where("id = ?", user.Id).Update("introduction", u.Introduction)
		}
		c.SecureJSON(200, nil)
	}
}
