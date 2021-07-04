package user

import (
	"context"
	"fmt"
	"os"

	"server/config/vars"
	"server/models"
	"server/tools"

	"github.com/gin-gonic/gin"
)

func Logoff(c *gin.Context) {

	value := c.GetHeader("Authorization")
	rows, _ := vars.DB0.Table("user").Model(&models.User{}).Where("id = ?", tools.Parse(value)).Rows()

	for rows.Next() {
		var (
			user   models.User
			logoff struct {
				models.User
				Code string `json:"code"`
			}
		)
		_ = vars.DB0.ScanRows(rows, &user)
		_ = c.ShouldBindJSON(&logoff)
		value, _ := vars.RedisLogoff.Get(context.TODO(), user.Email).Result()
		if logoff.Code != value {
			c.SecureJSON(200, gin.H{
				"message": 1013,
			})
		} else {
			// delete account
			vars.DB0.Table("user").Delete(&models.User{}, user.Id)
			// remove github account binding
			vars.DB0.Table("github").Delete(&models.Github{}, user.Id)
			// delete draft article
			vars.DB0.Table("draft").Where("author = ?", user.Username).Delete(&models.Article{})
			// delete file
			_ = os.RemoveAll(fmt.Sprintf("./static/article/draft/%d", user.Id))

			vars.DB0.Table("article").Where("author = ?", user.Username).Updates(map[string]interface{}{"author": "account deleted"})
			vars.DB0.Table("comment").Where("user = ?", user.Username).Updates(map[string]interface{}{"user": "account deleted"})
			vars.DB0.Table("reply").Where("user = ?", user.Username).Updates(map[string]interface{}{"user": "account deleted"})
			c.SecureJSON(200, nil)
		}
	}
}
