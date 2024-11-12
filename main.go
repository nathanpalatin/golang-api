package main

import (
	"apigo/controller"
	usecase "apigo/use-cases"

	"github.com/gin-gonic/gin"
)



func main() {

	server := gin.Default()
	UserUseCase := usecase.NewUserUseCase()
	UserController := controller.NewUserController(UserUseCase)

	server.GET("/users", UserController.GetUsers)

	server.Run(":3333")


}
