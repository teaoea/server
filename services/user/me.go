package user

import (
	"fmt"

	"server/config/vars"
	"server/models"
	"server/tools"

	"github.com/gin-gonic/gin"
)

func Me(c *gin.Context) {
	value := c.GetHeader("Authorization")
	rows, _ := vars.DB0.Table("user").Model(&models.User{}).Where("id = ?", tools.Parse(value)).Rows()

	for rows.Next() {
		var user models.User
		_ = vars.DB0.ScanRows(rows, &user)

		c.SecureJSON(200, gin.H{
			"username":     user.Username,
			"email":        user.Email,
			"email_active": user.EmailActive,
			"phone":        fmt.Sprintf("%s-%s", user.Prefix, user.Phone),
			"phone_active": user.PhoneActive,
			"avatar":       user.Avatar,
			"gender":       user.Gender,
		})
	}
}
