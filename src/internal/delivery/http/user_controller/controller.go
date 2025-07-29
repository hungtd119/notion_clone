package user_controller

import (
	"notion/src/internal/domain/usecases"

	"github.com/gin-gonic/gin"
)

type IController interface {
	Register(ctx *gin.Context)
	Login(ctx *gin.Context)
	List(ctx *gin.Context)
	GetByID(ctx *gin.Context)
}

type controller struct {
	userUseCase usecases.UserUseCase
}

func New(userUseCase usecases.UserUseCase) IController {
	return &controller{userUseCase: userUseCase}
}
