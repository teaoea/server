package vue

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"server/config/vars"
	"server/tools"
)

func Logger(c *gin.Context) {
	var logger struct {
		Method string `json:"method"`
		Path   string `json:"path"`
		Time   string `json:"time"`
		IPv4   string `json:"ipv4"`
	}
	err := c.ShouldBindJSON(&logger)
	if err != nil {
		tools.Err("logger", "解析前端传入的json数据失败")
		c.SecureJSON(404, nil)
		return
	}
	_, err = vars.MongoVueLogger.InsertOne(context.TODO(), bson.D{
		bson.E{Key: "_id", Value: tools.NewId()},
		bson.E{Key: "Method", Value: logger.Method},
		bson.E{Key: "Path", Value: logger.Path},
		bson.E{Key: "Time", Value: logger.Time},
		bson.E{Key: "IPv4", Value: logger.IPv4},
	})
	if err != nil {
		tools.Err("logger", "写入数据到mongo失败")
		c.SecureJSON(404, nil)
		return
	}
	c.SecureJSON(200, nil)
}
