package dataaccess

import (
	"tadi-chat-service/database"
	"tadi-chat-service/models"

	"go.mongodb.org/mongo-driver/mongo"
)


const (
	ChatUsersCollection = "chatusers"
)

var (
	chatUserModel = &database.Model[models.ChatUsers]{}
)

func CreateChatUser(data *models.ChatUsers) (*mongo.InsertOneResult, error) {
	insertedData, err := chatUserModel.Create(ctx, database.Client.Database(DATABASE_NAME), ChatUsersCollection, data)
	 if err != nil {
		return nil, err
	 }
	 return insertedData, nil
}