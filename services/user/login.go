package user

import (
	"context"
	"database/sql"
	"fmt"
	"server/config/vars"
	"server/models"
	"server/services/user/auth"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {
	var login struct {
		models.User
	}
	var user models.User
	_ = c.ShouldBindJSON(&login)
	nameCheck := vars.DB0.Table("user").Where("name = @name OR email = @name", sql.Named("name", login.Name)).Find(&user).RowsAffected

	if nameCheck == 0 {
		c.SecureJSON(404, gin.H{
			"message": fmt.Sprintf("用户'%s'未注册", login.Name),
		})
		return
	}

	rows, _ := vars.DB0.Table("user").Model(&models.User{}).Where("name = @name OR email = @name", sql.Named("name", login.Name)).Rows()

	for rows.Next() {
		var user models.User
		_ = vars.DB0.ScanRows(rows, &user)
		/*
			encodePWD 获取数据库中已加密的密码
			decodePWD 校验密码是否正确
		*/
		decodePWD := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password))
		if decodePWD == nil {
			value := auth.Create(user.Id, user.Name)

			vars.RedisToken.Set(context.Background(), user.Name, value, time.Hour*168)
			// 返回token
			c.SecureJSON(200, gin.H{
				"message": value,
			})
			return
		}
		c.SecureJSON(403, gin.H{
			"message": "密码错误",
		})
	}
}
