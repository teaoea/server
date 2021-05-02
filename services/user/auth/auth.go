package auth

import (
	"Server/config/vars"
	"Server/models"
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func ipCheck(ip string) bool {
	var ipaddr struct{}

	filter := bson.D{
		bson.E{Key: "ip", Value: ip},
	}
	val := vars.MongoIpaddr.FindOne(context.TODO(), filter).Decode(&ipaddr)
	return val != mongo.ErrNoDocuments
}

// ProxyAuth 不是指定IP,无权访问的路由组
func ProxyAuth() gin.HandlerFunc {
	return func(c *gin.Context) {

		ip := c.ClientIP()
		switch {
		case !ipCheck(ip):
			c.JSON(403, gin.H{
				"message": fmt.Sprintf("ip地址%s,无权限访问", ip),
			})
		default:
			c.Next()
		}
	}
}

// LoginAuth 未登录无权访问的路由组
func LoginAuth() gin.HandlerFunc {
	return func(c *gin.Context) {

		t := c.GetHeader("Authorization")
		parse := Parse(t)
		id := parse.(jwt.MapClaims)["id"]

		rows, _ := vars.DB0.Table("user").Model(&models.User{}).Where("id = ?", id).Rows()

		for rows.Next() {
			var user models.User

			_ = vars.DB0.ScanRows(rows, &user)

			result, _ := vars.RedisToken.Get(context.Background(), user.Name).Result()

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
