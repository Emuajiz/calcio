package mongo

import "go.mongodb.org/mongo-driver/v2/mongo"

func NewMongoClient() *mongo.Client {
	client, err := mongo.Connect()
	if err != nil {
		panic(err)
	}
	return client
}
