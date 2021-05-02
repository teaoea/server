package permission

import (
	"Server/config/vars"
	"Server/models"
	"Server/services/user/auth"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func HideArticle(c *gin.Context) {
	var hide struct {
		models.Permission
	}

	token := c.GetHeader("Authorization")
	parse := auth.Parse(token)
	id := parse.(jwt.MapClaims)["id"]
	rows, _ := vars.DB0.Table("user").Model(&models.User{}).Where("id = ?", id).Rows()

	for rows.Next() {
		var user models.User
		_ = vars.DB0.ScanRows(rows, &user)
		_ = c.ShouldBindJSON(&hide)

		nameCheck := vars.DB0.Table("user").Where(&models.User{Name: hide.Name}, "name").Find(&user).RowsAffected
		switch {
		case user.IsAdmin == false:
			c.SecureJSON(403, gin.H{
				"message": "无权执行此操作",
			})

		case nameCheck == 0:
			c.SecureJSON(404, gin.H{
				"message": fmt.Sprintf("用户%s不存在", hide.Name),
			})

		case hide.Name != "":
			var permission models.Permission
			// 查询是否已经创建记录,已创建记录就执行更新操作,未创建,执行创建操作
			rowsAffected := vars.DB0.Table("permission").Where(&models.Permission{Name: hide.Name}).Find(&permission).RowsAffected

			if rowsAffected == 1 {
				vars.DB0.Table("permission").Model(&permission).Updates(models.Permission{
					HideArticle:     hide.HideArticle,
					HideArticleAuth: user.Name,
				})
				c.SecureJSON(200, gin.H{
					"message": fmt.Sprintf("用户%s隐藏文章的权限已修改为:%v", user.Name, hide.HideArticle),
				})
			} else {
				permission := models.Permission{
					UserId:          hide.UserId,
					Name:            hide.Name,
					HideArticle:     hide.HideArticle,
					HideArticleAuth: user.Name,
				}
				vars.DB0.Table("permission").Create(&permission)
				c.SecureJSON(200, gin.H{
					"message": fmt.Sprintf("用户%s隐藏文章的权限已修改为:%v", user.Name, hide.HideArticle),
				})
			}
		}
	}
}
