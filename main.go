package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
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

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "OK %s %s %s %s", Build, BuildDate, GoVersion, Version)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	srv := http.Server{
		Handler: r,
		Addr: "0.0.0.0:3004",
		WriteTimeout: WriteTimeout * time.Second,
		ReadTimeout: ReadTimeout * time.Second,
	}
	conn, err := Connect("mongodb://localhost:27017/tadi")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(conn)
	log.Printf("Chat Service listening on %s", srv.Addr)
	log.Fatal(srv.ListenAndServe())
}