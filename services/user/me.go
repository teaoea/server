package user

import (
	vars2 "Server/config/vars"
	"Server/models"
	"Server/tools/token"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Me(c *gin.Context) {

	t := c.GetHeader("Authorization")
	parse := token.Parse(t)
	id := parse.(jwt.MapClaims)["id"]
	rows, _ := vars2.DB0.Table("user").Model(&models.User{}).Where("id = ?", id).Rows()

	for rows.Next() {
		var user models.User
		_ = vars2.DB0.ScanRows(rows, &user)

		c.SecureJSON(200, gin.H{
			"name":          user.Name,
			"email":         user.Email,
			"email_active":  user.EmailActive,
			"country":       user.Country,
			"number":        user.Number,
			"number_active": user.NumberActive,
			"avatar":        user.Avatar,
			"gender":        user.Gender,
			"introduction":  user.Introduction,
		})
	}
}
