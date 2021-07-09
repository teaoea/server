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

// Query user table public query api
// true: this record exists
// false: this record doesn't exists
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
		c.SecureJSON(461, gin.H{
			"message": "Key and value can't be empty",
		})

	case !check(query.Key):
		c.SecureJSON(462, gin.H{
			"message": "Can't query this key",
		})

	default:
		affected := vars.DB0.Table("user").Where(fmt.Sprintf("%s = ?", query.Key), query.Value).Find(&user).RowsAffected
		if affected == 1 {
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
