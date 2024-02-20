package dataaccess

import (
	"tadi-chat-service/database"
	"tadi-chat-service/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)


const (
	UsersCollection = "users"
)

var (
	userModel = &database.Model[models.User]{}
)

func UpdateUser(filter primitive.M, data primitive.M) ( error) {
	err := userModel.Update(ctx, database.Client.Database(DATABASE_NAME), UsersCollection, filter, data)
	 if err != nil {
		return err
	 }
	 return nil
}