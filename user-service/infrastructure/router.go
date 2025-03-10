package infrastructure

import (
	"user-service/internal/handler"
	"user-service/internal/repository"
	"user-service/internal/usecase"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	userRepo := repository.NewUserRepository(db)
	userUseCase := usecase.NewUserUseCase(userRepo)
	userHandler := handler.NewUserHandler(userUseCase)

	r.POST("/users/register", userHandler.Register)

	return r
}

