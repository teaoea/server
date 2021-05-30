package user

import (
	"context"
	"database/sql"
	"fmt"
	"server/config/vars"
	"server/models"
	"server/tools"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func SignIn(c *gin.Context) {
	var (
		user  models.User
		login struct {
			models.User
		}
	)

	_ = c.ShouldBindJSON(&login)
	nameCheck := vars.DB0.Table("user").Where("name = @name OR email = @name", sql.Named("name", login.Name)).Find(&user).RowsAffected

	if nameCheck == 0 {
		c.SecureJSON(404, gin.H{
			"message": fmt.Sprintf("user \"%s\" isn't sign up", login.Name),
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
			value := tools.Create(user.Id, user.Name)

			vars.RedisToken.Set(context.Background(), strconv.FormatInt(user.Id, 10), value, time.Hour*168)
			// 返回token
			c.SecureJSON(200, gin.H{
				"message": value,
			})
			return
		}
		c.SecureJSON(403, gin.H{
			"message": "wrong password",
		})
	}
}
