package user

import (
	"context"
	"fmt"

	"server/config/vars"
	"server/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func check(key string) bool {
	var query struct {
		Key string
	}
	filter := bson.D{
		bson.E{Key: "key", Value: key},
	}
	val := vars.MongoQuery.FindOne(context.TODO(), filter).Decode(&query)
	return val != mongo.ErrNoDocuments
}

// Query
/// user table public query api
/// return
/// false: can't pass
/// true: pass
func Query(c *gin.Context) {
	var (
		user  models.User
		query struct {
			Key   string `json:"key"`   // field name
			Value string `json:"value"` // the value to be queried
		}
	)

	_ = c.ShouldBindJSON(&query)

	switch {
	case query.Key == "" || query.Value == "":
		c.SecureJSON(200, gin.H{
			"message": false,
		})

	case !check(query.Key):
		c.SecureJSON(200, gin.H{
			"message": false,
		})

	default:
		affected := vars.DB0.Table("user").Where(fmt.Sprintf("%s = ?", query.Key), query.Value).Find(&user).RowsAffected
		if affected == 0 {
			c.SecureJSON(200, gin.H{
				"message": true,
			})
		} else {
			c.SecureJSON(200, gin.H{
				"message": false,
			})
		}
	}
}
