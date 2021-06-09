package tools

import (
	"context"

	"server/config/vars"

	"go.mongodb.org/mongo-driver/bson"
)

func Err(path string, err error) {
	_, _ = vars.MongoError.InsertOne(context.TODO(), bson.D{
		bson.E{Key: "_id", Value: NewId()},
		bson.E{Key: "path", Value: path},
		bson.E{Key: "err", Value: err},
	})
}
