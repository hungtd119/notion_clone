package routes

import (
	"net/http"
	"notion/src/internal/delivery/http/user_controller"
	"notion/src/internal/repository/user_repository"
	"notion/src/internal/usecase/user_usecase"

	"github.com/gin-gonic/gin"
)

func Bootstrap(r *gin.Engine) {
	userRepository := user_repository.New()

	userUsecase := user_usecase.New()
	userUsecase.SetUserRepo(userRepository)

	userController := user_controller.New(userUsecase)

	routesWeb := r.Group("")
	{
		routesWeb.GET("/", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"message": "Service notion web",
			})
		})
	}

	routesApi := r.Group("api")
	{
		routesApi.GET("/", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"message": "Service notion api",
			})
		})
	}

	routesApiV1 := r.Group("api/v1")
	{
		routesApiV1.GET("/", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"message": "Service notion api v1",
			})
		})

		// User routes
		users := routesApiV1.Group("/users")
		{
			users.POST("/register", userController.Register)
			users.POST("/login", userController.Login)
			users.GET("/:id", userController.GetByID)
			users.GET("/", userController.List)
		}
	}
}
