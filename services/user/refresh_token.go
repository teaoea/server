package user

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"server/config/vars"
	"server/models"
	"server/tools"
	"time"
)

func RefreshToken(ctx *gin.Context) {
	token := ctx.GetHeader("Authorization")
	parse := tools.Parse(token)
	id := parse.(jwt.MapClaims)["jti"]
	rows, _ := vars.DB0.Table("user").Model(&models.User{}).Where("id = ?", id).Rows()

	for rows.Next() {
		var user models.User
		_ = vars.DB0.ScanRows(rows, &user)

		value := tools.Create(user.Id)

		vars.RedisToken.Set(context.Background(), user.Name, value, time.Hour*168)
		ctx.SecureJSON(200, gin.H{
			"message": value,
		})
	}
}
