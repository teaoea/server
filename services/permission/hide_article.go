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
			// 查询用户是否已经创建权限表,未创建直接创建,已创建就更新
			userIdCheck := vars.DB0.Table("permission").Where(&models.Permission{UserId: hide.UserId}, "user_id").Find(&user).RowsAffected
			if userIdCheck == 0 {
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
			} else {
				var permission models.Permission
				vars.DB0.Table("permission").Model(&permission).Updates(models.Permission{
					HideArticle:     hide.HideArticle,
					HideArticleAuth: user.Name,
				})
				c.SecureJSON(200, gin.H{
					"message": fmt.Sprintf("用户%s隐藏文章的权限已修改为:%v", user.Name, hide.HideArticle),
				})
			}
		}
	}
}
