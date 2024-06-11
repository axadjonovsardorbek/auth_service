package main

import (
	"auth-service/api"
	"auth-service/api/handlers"
	"auth-service/service"
	"log"
)

func main() {
	userService := &service.UserService{}

	userHandler := handlers.NewUserHandler(userService)

	router := api.NewRouter(userHandler)

	log.Println("API Server is running...")
	if err := router.Run(":7070"); err != nil {
		log.Fatal(err)
	}
}
