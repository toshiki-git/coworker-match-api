package main

import (
	"log"
	"net/http"
	"time"

	"github.com/coworker-match-api/internal/db"
	"github.com/coworker-match-api/internal/routers"
)

func main() {
	db, err := db.InitDB()
	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}
	defer db.Close()

	endPoint := ":8080"
	readTimeout := 10 * time.Second
	writeTimeout := 10 * time.Second

	server := &http.Server{
		Addr:         endPoint,
		Handler:      routers.InitRouter(db),
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
	}

	log.Println("Starting server on :8080")
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
