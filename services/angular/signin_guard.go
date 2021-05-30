package angular

import (
	"context"
	"github.com/gin-gonic/gin"
	"server/config/vars"
	"server/models"
	"server/tools"
	"strconv"
)

func SigninGuard(c *gin.Context) {
	value := c.GetHeader("Authorization")
	if value == "" {
		c.SecureJSON(403, gin.H{
			"message": "not signin, please signin and visit again",
		})
		return
	}
	rows, _ := vars.DB0.Table("user").Model(&models.User{}).Where("id = ?", tools.Parse(value)).Rows()

	for rows.Next() {
		var user models.User
		_ = vars.DB0.ScanRows(rows, &user)

		result, _ := vars.RedisToken.Get(context.Background(), strconv.FormatInt(user.Id, 10)).Result()
		if result != value {
			c.JSON(403, gin.H{
				"message": "not signin, please signin and visit again",
			})
		} else {
			c.JSON(200, nil)
		}
	}
}
