package user_controller

import (
	"net/http"
	"notion/src/internal/domain/models"
	"notion/src/internal/dto"

	"github.com/gin-gonic/gin"
)

// Register đăng ký user mới
func (c *controller) Register(ctx *gin.Context) {
	var req dto.RegisterRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := &models.User{
		Email:    req.Email,
		Password: req.Password,
		Name:     req.Name,
	}

	err := c.userUseCase.Register(ctx.Request.Context(), user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully",
		"user":    user,
	})
}
