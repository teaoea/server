package router

import (
	"context"
	"strconv"
	"time"

	"server/config/vars"
	"server/models"
	"server/tools"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

// Server
/// access log save to mongodb
func Server() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		if c.Request.Method == "OPTIONS" {
			c.Next()
		} else {
			c.Next()
			_, _ = vars.MongoHttp.InsertOne(context.TODO(), bson.D{
				bson.E{Key: "_id", Value: tools.NewId()},
				bson.E{Key: "method", Value: c.Request.Method},
				bson.E{Key: "path", Value: c.Request.URL.Path},
				bson.E{Key: "delay", Value: time.Since(start) / 1e6},
				bson.E{Key: "status", Value: c.Writer.Status()},
				bson.E{Key: "time", Value: time.Now().Format("2006-01-02 15:04:05")},
				bson.E{Key: "ipv4", Value: c.ClientIP()},
			})
		}
	}
}

// Authorization
/// not sign in don't access router group
func Authorization() gin.HandlerFunc {
	return func(c *gin.Context) {

		value := c.GetHeader("Authorization")
		rows, _ := vars.DB0.Table("user").Model(&models.User{}).Where("id = ?", tools.Parse(value)).Rows()

		for rows.Next() {
			var user models.User

			_ = vars.DB0.ScanRows(rows, &user)

			result, _ := vars.RedisToken.Get(context.Background(), strconv.FormatInt(user.Id, 10)).Result()

			if result != value {
				c.JSON(401, gin.H{
					"message": "Not signed in to access this page",
				})
				return
			} else {
				c.Next()
			}
		}
	}
}

func Cor() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		if method != "" {
			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "true")
		}

		if method == "OPTIONS" {
			c.AbortWithStatus(204)
		}

		c.Next()
	}
}

func ipCheck(ip string) bool {
	for _, value := range vars.Ip {
		if ip == value {
			return true
		}
	}
	return false
}

// ProxyAuth
/// Route group that can't be accessed without specifying ip
func ProxyAuth() gin.HandlerFunc {
	return func(c *gin.Context) {

		ip := c.ClientIP()
		if !ipCheck(ip) {
			c.JSON(305, gin.H{
				"message": "Please use designated proxy access",
			})
			return
		} else {
			c.Next()
		}
	}
}
