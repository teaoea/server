package user

import (
	"Server/models"
	"Server/tools/token"
	"Server/vars"
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
)

func RefreshToken(ctx *gin.Context) {
	t := ctx.GetHeader("Authorization")
	parse := token.Parse(t)
	id := parse.(jwt.MapClaims)["jti"]
	rows, _ := vars.PDB0.Table("user").Model(&models.User{}).Where("id = ?", id).Rows()

	for rows.Next() {
		var user models.User
		_ = vars.PDB0.ScanRows(rows, &user)

		newToken, _ := token.Create(user.Id, user.Name)

		vars.RDBTOKEN.Set(context.Background(), user.Name, newToken, time.Hour*168)
		ctx.SecureJSON(200, gin.H{
			"t": newToken,
		})
	}
}
