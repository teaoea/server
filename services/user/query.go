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
func Query(c *gin.Context) {
	var (
		user  models.User
		query struct {
			Field string `json:"field"` // field name
			Value string `json:"value"` // the value to be queried
		}
	)

	_ = c.ShouldBindJSON(&query)

	switch {
	case query.Field == "" || query.Value == "":
		c.SecureJSON(452, nil)

	case !check(query.Field):
		c.SecureJSON(453, nil)

	default:
		affected := vars.DB0.Table("user").Where(fmt.Sprintf("%s = ?", query.Field), query.Value).Find(&user).RowsAffected
		if affected == 1 {
			c.SecureJSON(230, nil)
		} else {
			c.SecureJSON(231, nil)
		}
	}
}
