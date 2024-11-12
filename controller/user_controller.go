package controller

import (
	"apigo/model"
	usecase "apigo/use-cases"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userController struct {
	userUseCase usecase.User
}

func NewUserController(usecase usecase.User) userController {
	return userController{
		userUseCase: usecase,
	}
}

func (p *userController) GetUsers(ctx *gin.Context){
	 users := []model.User{
		{ID: 1, Name: "John Doe", Email: "john@example.com"},
    {ID: 2, Name: "Jane Doe", Email: "jane@example.com"},
    {ID: 3, Name: "Alice Johnson", Email: "alice@example.com"},
   }
	 ctx.JSON(http.StatusOK, users)
	 }
