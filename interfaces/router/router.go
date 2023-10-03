package interfaces

import (
	repo "article/infrastructure/repository"
	service "article/infrastructure/service"
	handler "article/interfaces/handler"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewRouter(r *gin.Engine, db *gorm.DB) {
	userRepo := repo.NewRepository(db)
	userService := service.NewUserSerivce(userRepo)
	userHandler := handler.NewHandlerUser(userService)

	r.POST("/api/user/signup", userHandler.SignUp)
}
