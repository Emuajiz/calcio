package helper

import "go.mongodb.org/mongo-driver/v2/bson"

func GenerateID() bson.ObjectID {
	return bson.NewObjectID()
}

func ToStringID(id bson.ObjectID) string {
	return id.Hex()
}

func FromStringID(id string) bson.ObjectID {
	objID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return bson.NilObjectID
	}
	return objID
}
