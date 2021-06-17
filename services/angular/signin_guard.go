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
		c.SecureJSON(452, nil)
		return
	}
	rows, _ := vars.DB0.Table("user").Model(&models.User{}).Where("id = ?", tools.Parse(value)).Rows()

	for rows.Next() {
		var user models.User
		_ = vars.DB0.ScanRows(rows, &user)

		result, _ := vars.RedisToken.Get(context.Background(), strconv.FormatInt(user.Id, 10)).Result()
		if result != value {
			c.JSON(200, gin.H{
				"message": 1001,
			})
		} else {
			c.JSON(200, nil)
		}
	}
}
