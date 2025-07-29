package dto

// RegisterRequest request body cho đăng ký
type RegisterRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Name     string `json:"name" binding:"required"`
}

// LoginRequest request body cho đăng nhập
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}
