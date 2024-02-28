package chats

import (
	"context"
	"net/http"
	"tadi-chat-service/internal/pkg/models"
	"tadi-chat-service/internal/utils"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ChatsController struct {
	service IChatService
}

func NewChatController(service IChatService) ChatsController {
	controller := &ChatsController{
		service: service,
	}
	return *controller
}

func (controller *ChatsController) CreateGroupChat(c *gin.Context) {
	var requestBody GroupChatRequest
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	user, found := c.Get("User")
	if !found {
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{"status": "failure", "message": "Could not fetch user details"})
		return
	}

	if err := c.BindJSON(&requestBody); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

	 requestBody.IsGroupChat = true

	 now := time.Now()

	 //  convert users strings to objectIds
	 usersIds := make([]primitive.ObjectID, 0, len(requestBody.Users))
	for _, id := range(requestBody.Users) {
		objId, err := utils.ConvertStringToObjId(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user id in the users array"})
        return
		}
		usersIds = append(usersIds, *objId)
	}

	payload := &models.Chat{
		ChatName: requestBody.ChatName,
		IsGroupChat: true,
		UsersIds: usersIds,
		CreatedAt: now,
		UpdatedAt: now,
	 }

	 _, err := controller.service.CreateChat(ctx, *payload, user)

	 if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
	 }

	 c.JSON(http.StatusOK, gin.H{"message": "Group chat created successfully"})
}

func (controller *ChatsController) FetchChats(c *gin.Context) {}