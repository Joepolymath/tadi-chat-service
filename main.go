package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"runtime"
	"tadi-chat-service/configs"
	"tadi-chat-service/database"
	"tadi-chat-service/endpoints"
	"tadi-chat-service/middlewares"

	"github.com/gin-gonic/gin"
)

const (
	WriteTimeout = 90
	ReadTimeout = 90
	ServicePrefix = "/api/v1/chats"
)

// Global vars for versioning
var (
	Build     = "1.0" // nolint
	BuildDate = "unset" // nolint
	GoVersion = runtime.Version() // nolint
	Version   = "unset" // nolint
	err error
)

// Global vars for .env
var (
	mongoURI string
	baseUrl string
	userServiceApiKey string
	env string
	port string
	envs *configs.EnvsStruct
)


func HomeHandler(c *gin.Context) {
	data := make(map[string]any)
	data["Build"] = Build
	data["BuildDate"] = BuildDate
	data["GoVersion"] = GoVersion
	data["Version"] = Version
	c.JSON(http.StatusOK, gin.H{"status": "success", "data": data})
}

func init() {
	 envs, err = configs.LoadEnv()
	 if err != nil {
		log.Fatal(err)
	 }

}

func main() {
	ctx := context.Background()
	r := gin.Default()
	r.Use(middlewares.TokenMiddleware)

	// Endpoints
	r.GET(fmt.Sprintf("%s/health", ServicePrefix), HomeHandler)
	r.POST(ServicePrefix + "/groups", endpoints.CreateGroupchat)

	// connect to db
	client, err := database.Connect(envs.MongoURI)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	r.Run(envs.Port)
}