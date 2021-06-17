package user

import (
	"context"
	"database/sql"
	"strconv"
	"time"

	"server/config/vars"
	"server/models"
	"server/tools"

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
	nameCheck := vars.DB0.Table("user").Where("username = @signin OR email = @signin", sql.Named("signin", login.Username)).Find(&user).RowsAffected

	if nameCheck == 0 {
		c.SecureJSON(200, gin.H{
			"message": 1009,
		})
		return
	}

	rows, _ := vars.DB0.Table("user").Model(&models.User{}).Where("username = @signin OR email = @signin", sql.Named("signin", login.Username)).Rows()

	for rows.Next() {
		var user models.User
		_ = vars.DB0.ScanRows(rows, &user)
		/*
			encodePWD get the encrypted password in the database
			decodePWD verify that the password is correct
		*/
		decodePWD := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password))
		if decodePWD == nil {
			value := tools.Create(user.Id, user.Username)

			vars.RedisToken.Set(context.Background(), strconv.FormatInt(user.Id, 10), value, time.Hour*168)
			// return token
			c.SecureJSON(200, gin.H{
				"message": value,
			})
			return
		}
		c.SecureJSON(200, gin.H{
			"message": 1010,
		})
	}
}
