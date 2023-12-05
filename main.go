package main

import (
	"tts-go-api/routes"
)

func main() {
	router := routes.SetupRoutes()

	err := router.Run(":8080")
	if err != nil {
		panic("Error starting the server: " + err.Error())
	}
}
