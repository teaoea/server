package tools

import (
	"context"

	"server/config/vars"

	"go.mongodb.org/mongo-driver/bson"
)

func Err(position, reason string) {
	_, _ = vars.MongoError.InsertOne(context.TODO(), bson.D{
		bson.E{Key: "_id", Value: NewId()},
		bson.E{Key: "position", Value: position},
		bson.E{Key: "reason", Value: reason},
	})
}
