package angular

import (
	"context"
	"time"

	"server/config/vars"
	"server/tools"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func Error(c *gin.Context) {
	var e struct {
		Position string `json:"position,omitempty"`
		Err      string `json:"err,omitempty"`
		Time     string `json:"time"`
	}
	_ = c.ShouldBindJSON(&e)

	_, _ = vars.MongoAngularError.InsertOne(context.TODO(), bson.D{
		bson.E{Key: "_id", Value: tools.NewId()},
		bson.E{Key: "position", Value: e.Position},
		bson.E{Key: "error", Value: e.Err},
		bson.E{Key: "time", Value: time.Now().Format("2006-01-02 15:04:05")}, // 请求时间
	})

	c.SecureJSON(200, nil)
}
