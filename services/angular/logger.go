package angular

import (
	"context"
	"time"

	"server/config/vars"
	"server/tools"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func LoggerAngular(c *gin.Context) {
	var logger struct {
		Method string `json:"method"`
		Path   string `json:"path"`
		Delay  string `json:"delay"`
	}

	_ = c.ShouldBindJSON(&logger)

	_, _ = vars.MongoAngularLogger.InsertOne(context.TODO(), bson.D{
		bson.E{Key: "_id", Value: tools.NewId()},
		bson.E{Key: "method", Value: logger.Method},
		bson.E{Key: "path", Value: logger.Path},
		bson.E{Key: "delay", Value: logger.Delay},
		bson.E{Key: "time", Value: time.Now().Format("2006-01-02 15:04:05")},
	})
}
