package chats

import (
	"context"
	"fmt"
	"log"
	"tadi-chat-service/internal/pkg/models"
	"tadi-chat-service/internal/utils"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IChatService interface{
	CreateChat(ctx context.Context, payload models.Chat, user any) (models.Chat, error)
}

type ChatService struct {
	repo	IChatRepository
}

func NewChatService(repo IChatRepository) IChatService {
	service := &ChatService{
		repo: repo,
	}
	chatService := IChatService(service)
	return chatService
}

func (s *ChatService) CreateChat(ctx context.Context, payload models.Chat, user any) (models.Chat, error) {
	chat := models.Chat{}

	userData, ok := user.(map[string]interface{})
	if !ok {
		return chat, fmt.Errorf("could not fetch user details")
	}
	
	if _, found := userData["_id"]; !found {
		return chat, fmt.Errorf("could not fetch user id")
	}
	userId := userData["_id"]

	userString, ok := userId.(string)
	if !ok {
		return chat, fmt.Errorf("could not parse user id to string")
	}

	userObjectId, err := utils.ConvertStringToObjId(userString)
	if err != nil {
		return chat, fmt.Errorf("could not parse user id")
	}
	payload.GroupAdmins = []primitive.ObjectID{*userObjectId}
	insertedData, err := s.repo.Create(&payload)
	if err != nil {
		return chat, err
	}

	// try to convertt to primitive.ObjectId
	id, ok := utils.AssertToObjectId(insertedData.InsertedID)
	if !ok {
		return chat, fmt.Errorf("unable to process chat id")
	}

	chatUser := &models.ChatUsers{
		ChatID: *id,
		UserID: *userObjectId,
		IsAdmin: true,
	}

	_, err = s.repo.CreateChatUser(chatUser)
	if err != nil {
		return chat, err
	}

	err = UpdateUserLastSeen(*userObjectId)
	if err != nil {
		log.Println(err.Error())
	}

	return chat, nil
}