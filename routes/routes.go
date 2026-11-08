package routes

import (
	"github.com/ahmed/shophub-backend/handlers"
	"github.com/ahmed/shophub-backend/infra"
	"github.com/ahmed/shophub-backend/repository"
	service "github.com/ahmed/shophub-backend/services"
	"github.com/gin-gonic/gin"
)

func RegisterRoute(r *gin.Engine) {

	userRepo := repository.RegisterUserRepo(infra.DB)
	userService := service.RegisterUserService(userRepo)
	userHandler := handlers.RegisterUserHandler(userService)

	authRoutes := r.Group("/api/v1/auth")
	{
		authRoutes.POST("/register", userHandler.RegisterUser)
	}

}
