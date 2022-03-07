package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/rajanjaiswal/MyNewContactbook/internals/server"
)

// routes or endpoints
func main() {
	fmt.Print("Starting ContactBook Api")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	fmt.Println(os.Getenv("DB_NAME"))
	srv := server.GetServer()
	// //r := .Default()
	// srv.C = gin.Default()

	// //database connection
	// srv.DB = models.ConnectDatabase()

	srv.Run() // starting server,
}
