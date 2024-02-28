package setup

import (
	"tadi-chat-service/internal/pkg/apis/chats"

	"github.com/gin-gonic/gin"
)

func HttpRouter(r *gin.Engine) {
	chats.Router(r)
}