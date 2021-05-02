package user

import (
	"Server/config/vars"
	"Server/models"
	"Server/services/user/auth"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Me(c *gin.Context) {

	t := c.GetHeader("Authorization")
	parse := auth.Parse(t)
	id := parse.(jwt.MapClaims)["id"]
	rows, _ := vars.DB0.Table("user").Model(&models.User{}).Where("id = ?", id).Rows()

	for rows.Next() {
		var user models.User
		_ = vars.DB0.ScanRows(rows, &user)

		c.SecureJSON(200, gin.H{
			"name":          user.Name,
			"email":         user.Email,
			"email_active":  user.EmailActive,
			"country":       user.Country,
			"number":        user.Number,
			"number_active": user.NumberActive,
			"avatar":        user.Avatar,
			"gender":        user.Gender,
		})
	}
}
