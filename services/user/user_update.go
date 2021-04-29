package user

import (
	"Server/config/vars"
	"Server/models"
	"Server/services/user/auth"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func UpdateUser(c *gin.Context) {
	token := c.GetHeader("Authorization")
	parse := auth.Parse(token)
	id := parse.(jwt.MapClaims)["id"]
	rows, _ := vars.DB0.Table("user").Model(&models.User{}).Where("id = ?", id).Rows()
	for rows.Next() {
		var user models.User
		var u struct {
			models.User
		}
		_ = vars.DB0.ScanRows(rows, &user)
		_ = c.ShouldBindJSON(&u)
		switch {

		case u.Avatar != "":
			vars.DB0.Table("user").Model(&models.User{}).Where("id = ?", user.Id).Update("avatar", u.Avatar)

		case u.Gender != "":
			vars.DB0.Table("user").Model(&models.User{}).Where("id = ?", user.Id).Update("gender", u.Gender)

		case u.Introduction != "":
			vars.DB0.Table("user").Model(&models.User{}).Where("id = ?", user.Id).Update("introduction", u.Introduction)
		}
		c.SecureJSON(200, nil)
	}
}
