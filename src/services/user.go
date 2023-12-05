package service

import (
	"database/sql"
	"log"

	"tts-go-api/database"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	// Add other fields as needed
}

func GetUser(userID string) (*User, error) {
	db := database.GetDB()

	var user User
	err := db.QueryRow("SELECT id, name FROM users WHERE id = ?", userID).Scan(&user.ID, &user.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("No user found")
			return nil, nil
		}
		log.Println("Error fetching user: ", err)
		return nil, err
	}

	return &user, nil
}
