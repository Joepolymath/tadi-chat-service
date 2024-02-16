package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
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
)

// Global vars for .env
var (
	mongoURI string
	baseUrl string
	userServiceApiKey string
	env string
	port string
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
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
        log.Fatalf("Error reading config file: %s", err)
    }

	 mongoURI = viper.GetString("MONGO_URI")
	 baseUrl = viper.GetString("BASE_URL")
	 port = viper.GetString("PORT")

}

func main() {
	// r := mux.NewRouter()
	// r.HandleFunc(fmt.Sprintf("%s/health", ServicePrefix), HomeHandler).Methods("GET")
	r := gin.Default()
	r.GET(fmt.Sprintf("%s/health", ServicePrefix), HomeHandler)
	

	// connect to db
	_, err := Connect(mongoURI)
	if err != nil {
		log.Fatal(err)
	}

	r.Run(port)
}