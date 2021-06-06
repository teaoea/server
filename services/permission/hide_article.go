package permission

import (
	"fmt"

	"server/config/vars"
	"server/models"
	"server/tools"

	"github.com/gin-gonic/gin"
)

func HideArticle(c *gin.Context) {

	value := c.GetHeader("Authorization")
	rows, _ := vars.DB0.Table("user").Model(&models.User{}).Where("id = ?", tools.Parse(value)).Rows()

	for rows.Next() {
		var (
			user models.User
			hide struct {
				models.Permission
			}
		)

		_ = vars.DB0.ScanRows(rows, &user)
		_ = c.ShouldBindJSON(&hide)

		nameCheck := vars.DB0.Table("user").Where(&models.User{Username: hide.Name}, "username").Find(&user).RowsAffected
		switch {
		case user.IsAdmin == false:
			c.SecureJSON(403, gin.H{
				"message": "not authorized to perform this operation",
			})

		case nameCheck == 0:
			c.SecureJSON(404, gin.H{
				"message": fmt.Sprintf("user \"%s\" isn't exist", hide.Name),
			})

		case hide.Name != "":
			var permission models.Permission
			// query whether the record has been created,
			// perform the update operation when the record has been created,
			// and perform the creation operation if it isn't created
			rowsAffected := vars.DB0.Table("permission").Where(&models.Permission{Name: hide.Name}).Find(&permission).RowsAffected

			if rowsAffected == 1 {
				vars.DB0.Table("permission").Model(&permission).Updates(models.Permission{
					HideArticle:     hide.HideArticle,
					HideArticleAuth: user.Username,
				})
				c.SecureJSON(200, gin.H{
					"message": fmt.Sprintf("user \"%s\"'s permission to hide articles has been modified to \"%v\"", user.Username, hide.HideArticle),
				})
			} else {
				permission := models.Permission{
					UserId:          hide.UserId,
					Name:            hide.Name,
					HideArticle:     hide.HideArticle,
					HideArticleAuth: user.Username,
				}
				vars.DB0.Table("permission").Create(&permission)
				c.SecureJSON(200, gin.H{
					"message": fmt.Sprintf("The permission of user \"%s\" to hide articles has been modified to \"%v\"", user.Username, hide.HideArticle),
				})
			}
		}
	}
}
