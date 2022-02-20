package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Post
	router.GET("/post/read", postRead)
	router.GET("/post/read/:id", postReadByID)
	router.POST("/post/create", postCreate)
	router.GET("/post/delete/:id", postDelete)

	// Block
	router.GET("/block/delete/:id", blockDelete)
	router.POST("/block/create", blockCreate)
	router.GET("/block/read/:id", blockReadByUserID)

	// Like
	router.GET("/like/delete/:id", likeDelete)
	router.POST("/like/create", likeCreate)
	router.GET("/like/read/:id", likeReadByUserID)

	// Education
	router.GET("/education/delete/:id", educationDelete)
	router.POST("/education/create", educationCreate)
	router.POST("/education/update", educationUpdate)
	router.GET("/education/read/:id", educationReadByUserID)

	// Work
	router.GET("/work/delete/:id", workDelete)
	router.POST("/work/create", workCreate)
	router.POST("/work/update", workUpdate)
	router.GET("/work/read/:id", workReadByUserID)

	// Friend // friendReadByUserID not working
	router.POST("/work/delete", friendDelete)
	router.POST("/work/create", friendCreate)
	router.GET("/work/read/:id", friendReadByUserID)

	// Friend Request
	router.POST("/work/delete", friendRequestDelete)
	router.POST("/work/create", friendRequestCreate)
	router.POST("/work/update", friendRequestUpdate)

	// User
	router.GET("/work/delete", userDelete)
	router.POST("/work/create", userCreate)
	router.POST("/work/update", userUpdate)
	router.GET("/work/read", userRead)

	// Auth
	router.POST("/auth/signup", signup)
	router.POST("/auth/login", login)

	// Starting the Server
	router.Run("localhost:8080")
}
