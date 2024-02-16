package endpoints

import (
	"context"
	"log"
	"net/http"
	"tadi-chat-service/database"
	"tadi-chat-service/models"

	"github.com/gin-gonic/gin"
)

type GroupChatRequest struct {
    ChatName   string   `json:"chatName"`
    Users      []string `json:"users"`
    GroupAdmin string   `json:"groupAdmin"`
	 IsGroupChat bool		`json:"isGroupChat"`
}

var ctx = context.Background()

func CreateGroupchat(c *gin.Context) {
	 var requestBody GroupChatRequest
    if err := c.BindJSON(&requestBody); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

	 requestBody.IsGroupChat = true
    // Process the group chat request
    log.Printf("Received group chat request: %+v\n", requestBody)

	 model := &database.Model{}

	 groupChat := &models.Chat{
		ChatName: requestBody.ChatName,
		IsGroupChat: requestBody.IsGroupChat,
		Users: requestBody.Users,
		GroupAdmin: requestBody.GroupAdmin,
	 }

	 err := model.Create(ctx, database.Client.Database("tadi"), "chats", groupChat)
	 if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{"status": "failure", "message": "Unable to create group chat"})
	 }

    // Respond with a success message
    c.JSON(http.StatusOK, gin.H{"message": "Group chat created successfully"})
}