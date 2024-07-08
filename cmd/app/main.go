package main

import (
	"fmt"
	"log"

	"github.com/coworker-match-api/internal/db"
	"github.com/coworker-match-api/internal/models"
)

func main() {
	fmt.Println("Hello, World!")

	db := db.InitDB()
	defer db.Close()

	getUserSQL := `SELECT * FROM users`
	rows, err := db.Query(getUserSQL)
	if err != nil {
		log.Fatalf("Error querying data: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.UserID, &user.UserName, &user.Email, &user.AvatarURL, &user.Age, &user.Gender, &user.Birthplace, &user.JobType, &user.LINEAccount, &user.DiscordAccount, &user.Biography, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			log.Fatalf("Error scanning data: %v", err)
		}
		fmt.Printf("User: %+v\n", user)
	}

	if err := rows.Err(); err != nil {
		log.Fatalf("Error with rows: %v", err)
	}
}
