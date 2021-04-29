package user

import (
	vars2 "Server/config/vars"
	"Server/models"
	"Server/tools/token"
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
)

func RefreshToken(ctx *gin.Context) {
	t := ctx.GetHeader("Authorization")
	parse := token.Parse(t)
	id := parse.(jwt.MapClaims)["jti"]
	rows, _ := vars2.DB0.Table("user").Model(&models.User{}).Where("id = ?", id).Rows()

	for rows.Next() {
		var user models.User
		_ = vars2.DB0.ScanRows(rows, &user)

		newToken, _ := token.Create(user.Id, user.Name)

		vars2.RedisToken.Set(context.Background(), user.Name, newToken, time.Hour*168)
		ctx.SecureJSON(200, gin.H{
			"t": newToken,
		})
	}
}
