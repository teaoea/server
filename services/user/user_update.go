package user

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"server/config/vars"
	"server/models"
	"server/tools"
)

func UpdateUser(c *gin.Context) {
	token := c.GetHeader("Authorization")
	parse := tools.Parse(token)
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
			fallthrough

		case u.Gender != "":
			vars.DB0.Table("user").Model(&models.User{}).Where("id = ?", user.Id).Update("gender", u.Gender)

			c.SecureJSON(200, nil)
		}
	}
}
