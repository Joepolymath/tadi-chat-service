package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

const (
	WriteTimeout = 90
	ReadTimeout = 90
)

// Global vars for versioning
var (
	Build     = "unset" // nolint
	BuildDate = "unset" // nolint
	GoVersion = "unset" // nolint
	Version   = "unset" // nolint
)

// Global vars for .env
var (
	mongoURI string
	baseUrl string
	userServiceApiKey string
	env string
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "OK %s %s %s %s", Build, BuildDate, GoVersion, Version)
}

func init() {
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
        log.Fatalf("Error reading config file: %s", err)
    }

	 mongoURI = viper.GetString("MONGO_URI")
	 baseUrl = viper.GetString("BASE_URL")

}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	srv := http.Server{
		Handler: r,
		Addr: baseUrl,
		WriteTimeout: WriteTimeout * time.Second,
		ReadTimeout: ReadTimeout * time.Second,
	}
	_, err := Connect(mongoURI)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Chat Service listening on %s", srv.Addr)
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}