package log

import (
	"Server/tools"
	"Server/vars"
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

// Server
/// 日志中间件,访问日志保存到mongodb
func Server() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()
		ctx.Next()
		_, _ = vars.MDBHTTP.InsertOne(context.TODO(), bson.D{
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
