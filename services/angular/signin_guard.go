package angular

import (
	"context"
	"strconv"

	"server/config/vars"
	"server/models"
	"server/tools"

	"github.com/gin-gonic/gin"
)

func SigninGuard(c *gin.Context) {
	value := c.GetHeader("Authorization")
	if value == "" {
		c.SecureJSON(401, gin.H{
			"message": "Not signed in to access this page",
		})
		return
	}
	rows, _ := vars.DB0.Table("user").Model(&models.User{}).Where("id = ?", tools.Parse(value)).Rows()

	for rows.Next() {
		var user models.User
		_ = vars.DB0.ScanRows(rows, &user)

		result, _ := vars.RedisToken.Get(context.Background(), strconv.FormatInt(user.Id, 10)).Result()
		if result != value {
			c.SecureJSON(401, gin.H{
				"message": "Not signed in to access this page",
			})
		} else {
			c.JSON(200, gin.H{
				"message": user.Username,
			})
		}
	}
}
