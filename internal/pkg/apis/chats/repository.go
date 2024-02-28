package chats

import (
	"context"
	"tadi-chat-service/internal/database"
	"tadi-chat-service/internal/pkg/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	ChatCollection = "chats"
	ChatUsersCollection = "chatusers"
	UsersCollection = "users"
)

var (
	ctx = context.Background()
	DATABASE_NAME = "tadi"
	userModel = &database.Model[models.User]{}
)

type IChatRepository interface{
	Create(data *models.Chat) (*mongo.InsertOneResult, error)
	CreateChatUser(data *models.ChatUsers) (*mongo.InsertOneResult, error)
}

type ChatRepo struct {
	model *database.Model[models.Chat]
}

func NewChatRepo() IChatRepository {
	repo := &ChatRepo{
		model: &database.Model[models.Chat]{},
	}
	repository := IChatRepository(repo)
	return repository
}


func (cr *ChatRepo) Create(data *models.Chat) (*mongo.InsertOneResult, error) {
	insertedData, err := cr.model.Create(ctx, database.Client.Database(DATABASE_NAME), ChatCollection, data)
	 if err != nil {
		return nil, err
	 }
	 return insertedData, nil
}

func (cr *ChatRepo) CreateChatUser(data *models.ChatUsers) (*mongo.InsertOneResult, error) {
	insertedData, err := cr.model.Create(ctx, database.Client.Database(DATABASE_NAME), ChatUsersCollection, data)
	 if err != nil {
		return nil, err
	 }
	 return insertedData, nil
}

func UpdateUserLastSeen(userId primitive.ObjectID) ( error) {
	userUpdate := bson.M{"$set": bson.M{"lastSeen": time.Now()}}
	err := userModel.Update(ctx, database.Client.Database(DATABASE_NAME), UsersCollection, bson.M{"_id": userId}, userUpdate)
	 if err != nil {
		return err
	 }
	 return nil
}