package http

import (
	// "net/http" - not needed

	"github.com/Amad3eu/gin-gonic-posts-api/internal"
	"github.com/Amad3eu/gin-gonic-posts-api/internal/database"
	"github.com/Amad3eu/gin-gonic-posts-api/internal/post"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var service post.Service

func Configure() {
	service = post.Service{
		Repository: post.Repository{
			Conn: database.Conn,
		},
	}
}

func PostPosts(ctx *gin.Context) {
	var post internal.Post
	if err := ctx.BindJSON(&post); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := service.Create(post); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(201, nil)
}

func DeletePosts(ctx *gin.Context) {
	param := ctx.Param("id")
	id, err := uuid.Parse(param)
	if err != nil {
		ctx.JSON(404, nil)
	}

	if err := service.Delete(id); err != nil {
		statusCode := 500
		if err == post.ErrPostNotFound {
			statusCode = 404
		}

		ctx.JSON(statusCode, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(204, nil)
}

func GetPost(ctx *gin.Context) {
	param := ctx.Param("id")
	id, err := uuid.Parse(param)
	if err != nil {
		ctx.JSON(404, nil)
	}

	p, err := service.FindOneByID(id)
	if err != nil {
		statusCode := 500
		if err == post.ErrPostNotFound {
			statusCode = 404
		}

		ctx.JSON(statusCode, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(200, p)
}

func GetPosts(ctx *gin.Context) {
	posts, err := service.FindAll()
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, posts)
}
