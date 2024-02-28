package utils

import "go.mongodb.org/mongo-driver/bson/primitive"

func ConvertStringToObjId(str string) (*primitive.ObjectID, error) {
	objectId, err := primitive.ObjectIDFromHex(str)
	if err != nil {
		return nil, err
	}

	return &objectId, nil
}

func AssertToObjectId(data interface{}) (*primitive.ObjectID, bool) {
	id, ok := data.(primitive.ObjectID)
	if !ok {
		return nil, ok
	}

	return &id, ok
}