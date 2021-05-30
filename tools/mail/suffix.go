package mail

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"server/config/vars"
	"strings"
)

func SuffixCheck(email string) bool {
	var suffix struct {
		Suffix string
	}

	addr := strings.Split(email, "@") // string segmentation
	suf := "@" + addr[1]              // intercept email address suffix
	filter := bson.D{
		bson.E{Key: "suffix", Value: suf},
	}
	val := vars.MongoSuffix.FindOne(context.TODO(), filter).Decode(&suffix)
	return val != mongo.ErrNoDocuments
}
