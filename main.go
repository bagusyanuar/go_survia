package main

import (
	"fmt"
	"go-survia/database"
	"go-survia/routes"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error : failed to load .env file")
	}

	err = database.Connect()
	if err != nil {
		panic("Error : failed to connect database ")
	}
	log.Println("success connect to database")

	// database.Migrate()
	// database.Seed()
	log.Println("success migrate database")

	appPort := os.Getenv("APP_PORT")
	server := routes.InitRoutes()
	port := fmt.Sprintf(":%s", appPort)
	server.Run(port)
}
