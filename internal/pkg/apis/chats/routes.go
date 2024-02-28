package chats

import "github.com/gin-gonic/gin"

const (
	WriteTimeout = 90
	ReadTimeout = 90
	ServicePrefix = "/api/v1/chats"
)

var (
	chatsRepo = NewChatRepo()
	service = NewChatService(chatsRepo)
	controller = NewChatController(service)
)

func Router(r *gin.Engine) {
	routeGroup := r.Group(ServicePrefix)

	routeGroup.POST("/groups", controller.CreateGroupChat)
	routeGroup.GET("/", controller.FetchChats)
}