package router

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"server/config/vars"
	"server/models"
	"server/services/user/auth"
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

		t := c.GetHeader("Authorization")
		parse := auth.Parse(t)
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
