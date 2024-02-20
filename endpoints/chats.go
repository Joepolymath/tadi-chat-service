package endpoints

import (
	"context"
	"fmt"
	"net/http"
	"tadi-chat-service/database"
	"tadi-chat-service/models"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	 user, found := c.Get("User")
	if !found {
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{"status": "failure", "message": "Could not fetch user details"})
		return
	}

	userData, ok := user.(map[string]interface{})
	if !ok {
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{"status": "failure", "message": "Could not fetch user details"})
		return
	}
	userId := userData["_id"]
	userString, ok := userId.(string)
	if !ok {
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{"status": "failure", "message": "Could not parse user id to string"})
		return
	}

	// Convert the string to an ObjectId
    userObjectId, err := primitive.ObjectIDFromHex(userString)
    if err != nil {
        fmt.Println("Error converting string to ObjectId:", err)
        return
    }

    if err := c.BindJSON(&requestBody); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

	 requestBody.IsGroupChat = true
    // Process the group chat request

	 model := &database.Model[models.Chat]{}

	 now := time.Now()
	 groupChat := &models.Chat{
		ChatName: requestBody.ChatName,
		IsGroupChat: requestBody.IsGroupChat,
		GroupAdmins: []primitive.ObjectID{userObjectId},
		CreatedAt: now,
		UpdatedAt: now,
	 }

	 insertedData, err := model.Create(ctx, database.Client.Database("tadi"), "chats", groupChat)
	 if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{"status": "failure", "message": "Unable to create group chat"})
		return
	 }


	//  try to convert to primitive.ObjectId
	id, ok := insertedData.InsertedID.(primitive.ObjectID)
	if !ok {
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{"status": "failure", "message": "Unable to parse chat id to primitive.ObjectId"})
		return
	}
	 
	 chatUserModel := &database.Model[models.ChatUsers]{}
	 chatUser := &models.ChatUsers{
		ChatID: id,
		UserID: userObjectId,
		IsAdmin: true,
	 }

	 _, err = chatUserModel.Create(ctx, database.Client.Database("tadi"), "chatusers", chatUser)
	 if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{"status": "failure", "message": "Unable to add chat_user"})
		return
	 }

	 userModel := &database.Model[models.User]{}
	 userUpdate := bson.M{"$set": bson.M{"lastSeen": time.Now()}}

	 err = userModel.Update(ctx, database.Client.Database("tadi"), "users", bson.M{"_id": userObjectId}, userUpdate)
	 if err != nil {
		fmt.Println(err.Error())
	 }

    // Respond with a success message
    c.JSON(http.StatusOK, gin.H{"message": "Group chat created successfully"})
}

// fetching chats for authenticated user
func FetchChats(c *gin.Context) {
	user, found := c.Get("User")
	if !found {
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{"status": "failure", "message": "Could not fetch user details"})
		return
	}

	userData, ok := user.(map[string]interface{})
	if !ok {
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{"status": "failure", "message": "Could not fetch user details"})
		return
	}

	filter := bson.M{"users": bson.M{"$elemMatch": bson.M{"$eq": userData["_id"]}}}

	model := &database.Model[models.Chat]{}
	var fetchedChats []models.Chat

	err := model.ReadMany(ctx, database.Client.Database("tadi"), "chats", filter, &fetchedChats)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{"status": "failure", "message": "Could not fetch chats"})
		return
	}


	c.JSON(http.StatusOK, fetchedChats)
}