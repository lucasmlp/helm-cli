package mongo

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type adapter struct {
	client   *mongo.Client
	database *mongo.Database
}

func NewAdapter(
	client *mongo.Client,
	database *mongo.Database,
) *adapter {
	return &adapter{
		client:   client,
		database: database,
	}
}
