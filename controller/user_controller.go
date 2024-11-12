package controller

import (
	"apigo/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userController struct {

}

func NewUserController() userController {
	return userController{}
}

func (p *userController) GetUsers(ctx *gin.Context){
	 users := []model.User{
		{ID: 1, Name: "John Doe", Email: "john@example.com"},
    {ID: 2, Name: "Jane Doe", Email: "jane@example.com"},
    {ID: 3, Name: "Alice Johnson", Email: "alice@example.com"},
   }
	 ctx.JSON(http.StatusOK, users)
	 }
