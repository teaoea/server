package router

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"server/config/vars"
	"server/models"
	"server/tools"
	"time"
)

// Server
/// 日志中间件,访问日志保存到mongodb
func Server() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()
		ctx.Next()
		_, _ = vars.MongoHttp.InsertOne(context.TODO(), bson.D{
			bson.E{Key: "_id", Value: tools.NewId()},
			bson.E{Key: "Method", Value: ctx.Request.Method},                     // 请求方式
			bson.E{Key: "Path", Value: ctx.Request.URL.Path},                     // 请求路径
			bson.E{Key: "Delay", Value: time.Since(start) / 1e6},                 // 延迟
			bson.E{Key: "Status", Value: ctx.Writer.Status()},                    // 请求状态
			bson.E{Key: "Time", Value: time.Now().Format("2006-01-02 15:04:05")}, // 请求时间
			bson.E{Key: "IPv4", Value: ctx.ClientIP()},                           // 客户端ip
		})
	}
}

// LoginAuth 未登录无权访问的路由组
func LoginAuth() gin.HandlerFunc {
	return func(c *gin.Context) {

		value := c.GetHeader("Authorization")
		rows, _ := vars.DB0.Table("user").Model(&models.User{}).Where("id = ?", tools.Parse(value)).Rows()

		for rows.Next() {
			var user models.User

			_ = vars.DB0.ScanRows(rows, &user)

			result, _ := vars.RedisToken.Get(context.Background(), user.Name).Result()

			if result != value {
				c.JSON(403, gin.H{
					"message": "登录已过期,请重新登录",
				})
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
