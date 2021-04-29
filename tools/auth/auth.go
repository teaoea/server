package auth

import (
	"Server/models"
	"Server/tools/token"
	"Server/vars"
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"sort"
)

// ProxyAuth 不是指定IP,无权访问的路由组
func ProxyAuth() gin.HandlerFunc {
	return func(c *gin.Context) {

		ip := c.ClientIP()
		index := sort.SearchStrings(vars.PROXYADDR, ip)
		if ip != vars.PROXYADDR[index] {
			c.JSON(401, nil)
		} else {
			c.Next()
		}
	}
}

// LoginAuth 未登录无权访问的路由组
func LoginAuth() gin.HandlerFunc {
	return func(c *gin.Context) {

		t := c.GetHeader("Authorization")
		parse := token.Parse(t)
		id := parse.(jwt.MapClaims)["id"]

		rows, _ := vars.PDB0.Table("user").Model(&models.User{}).Where("id = ?", id).Rows()

		for rows.Next() {
			var user models.User

			_ = vars.PDB0.ScanRows(rows, &user)

			result, _ := vars.RDBTOKEN.Get(context.Background(), user.Name).Result()

			if result != t {
				c.JSON(403, gin.H{
					"message": "登录已过期,请重新登录",
				})
			} else {
				c.Next()
			}
		}
	}
}
