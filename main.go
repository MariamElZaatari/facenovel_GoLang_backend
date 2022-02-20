package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()



	// Auth
	router.POST("/auth/signup", signup)
	router.POST("/auth/login", login)

	// Starting the Server
	router.Run("localhost:8080")
}