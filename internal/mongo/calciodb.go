package mongo

import (
	"calcio/config"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

func NewMongoDB(config *config.Config, client *mongo.Client) *mongo.Database {
	return client.Database(config.MongoConfig.Database)
}
