package angular

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"server/config/vars"
	"server/tools"
	"time"
)

func Error(c *gin.Context) {
	var e struct {
		Position string `json:"position,omitempty"`
		Err      string `json:"err,omitempty"`
		Time     string `json:"time"`
	}
	err := c.ShouldBindJSON(&e)
	if err != nil {
		tools.Err("services/angular/error.go", fmt.Sprintf("%s", err))
		c.SecureJSON(503, nil)
		return
	}

	_, err = vars.MongoAngularError.InsertOne(context.TODO(), bson.D{
		bson.E{Key: "_id", Value: tools.NewId()},
		bson.E{Key: "position", Value: e.Position},
		bson.E{Key: "error", Value: e.Err},
		bson.E{Key: "time", Value: time.Now().Format("2006-01-02 15:04:05")}, // 请求时间
	})
	if err != nil {
		tools.Err("services/angular/error.go", fmt.Sprintf("%s", err))
		c.SecureJSON(503, nil)
		return
	}

	c.SecureJSON(200, nil)
}
