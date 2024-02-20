package dataaccess

import (
	"context"
	"tadi-chat-service/database"
	"tadi-chat-service/models"

	"go.mongodb.org/mongo-driver/mongo"
)


const (
	ChatCollection = "chats"
)

var (
	DATABASE_NAME = "tadi"
	chatModel = &database.Model[models.Chat]{}
	ctx = context.Background()
)

func CreateChat(data *models.Chat) (*mongo.InsertOneResult, error) {
	insertedData, err := chatModel.Create(ctx, database.Client.Database(DATABASE_NAME), ChatCollection, data)
	 if err != nil {
		return nil, err
	 }
	 return insertedData, nil
}