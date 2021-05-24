package user

import (
	"context"
	"github.com/gin-gonic/gin"
	"server/config/vars"
	"server/models"
	"server/tools"
	"strconv"
	"time"
)

func RefreshToken(c *gin.Context) {

	value := c.GetHeader("Authorization")
	rows, _ := vars.DB0.Table("user").Model(&models.User{}).Where("id = ?", tools.Parse(value)).Rows()

	for rows.Next() {
		var user models.User
		_ = vars.DB0.ScanRows(rows, &user)

		value := tools.Create(user.Id, user.Name)

		vars.RedisToken.Set(context.Background(), strconv.FormatInt(user.Id, 10), value, time.Hour*168)
		c.SecureJSON(200, gin.H{
			"message": value,
		})
	}
}
