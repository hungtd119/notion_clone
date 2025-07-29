package user_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// GetByID lấy user theo ID
func (c *controller) GetByID(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	user, err := c.userUseCase.GetByID(ctx.Request.Context(), uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"user": user})
}
