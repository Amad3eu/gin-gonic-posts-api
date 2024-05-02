package main

import (
	"github.com/Amad3eu/gin-gonic-posts-api/internal/database"
	"github.com/Amad3eu/gin-gonic-posts-api/internal/http"
	"github.com/gin-gonic/gin"
)

func main() {
	connectionString := "postgresql://posts:p0stgr3s@localhost:5432/posts"
	conn, err := database.NewConnection(connectionString)
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	g := gin.Default()
	http.Configure()
	http.SetRoutes(g)
	g.Run(":8080")
}
