package main

import (
	"apigo/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)



func main() {

	server := gin.Default()

	UserController := controller.NewUserController()

	server.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
      "message": "Hello, World!",
    })
	})

	server.GET("/users", UserController.GetUsers)

	server.Run(":3333")
}
