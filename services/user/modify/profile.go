package modify

import (
	"server/config/vars"
	"server/models"
	"server/tools"

	"github.com/gin-gonic/gin"
)

func Profile(c *gin.Context) {

	value := c.GetHeader("Authorization")
	rows, _ := vars.DB0.Table("user").Model(&models.User{}).Where("id = ?", tools.Parse(value)).Rows()

	for rows.Next() {
		var (
			user models.User
			u    struct {
				models.User
			}
		)
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
