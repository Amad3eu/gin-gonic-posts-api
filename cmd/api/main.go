package main

import (
	"log"
	"os"

	"github.com/Amad3eu/gin-gonic-posts-api/internal/database"
	"github.com/Amad3eu/gin-gonic-posts-api/internal/http"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	connectionString := os.Getenv("DB_CONNECTION_STRING")
	if connectionString == "" {
		log.Fatal("DB_CONNECTION_STRING environment variable is required")
	}

	conn, err := database.NewConnection(connectionString)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	defer conn.Close()

	g := gin.Default()
	http.Configure()
	http.NewHandler(conn)
	http.SetRoutes(g)
	g.Run(":8080")
}
