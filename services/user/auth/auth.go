package auth

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"server/config/vars"
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
