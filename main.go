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
