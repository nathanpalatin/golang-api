package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)



func main() {

	server := gin.Default()

	server.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
      "message": "Hello, World!",
    })
	})

	server.Run(":3333")
}
