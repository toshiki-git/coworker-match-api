package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
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

	endPoint := fmt.Sprintf(":%s", os.Getenv("API_PORT"))
	readTimeout := 10 * time.Second
	writeTimeout := 10 * time.Second

	server := &http.Server{
		Addr:         endPoint,
		Handler:      routers.InitRouter(db),
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
	}

	log.Printf("Server listening on %s", endPoint)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
