package main

import (
    // "net/http"

    "github.com/gin-gonic/gin"
)


func main() {
    router := gin.Default()
    // router.GET("/", function)


    router.Run("localhost:8080")
}
