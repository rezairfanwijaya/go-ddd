package interfaces

import (
	"article/auth"
	repo "article/infrastructure/repository"
	service "article/infrastructure/service"
	handler "article/interfaces/handler"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewRouter(r *gin.Engine, db *gorm.DB) {
	jwtService := auth.NewServiceJWT()

	userRepo := repo.NewRepository(db)
	userService := service.NewUserSerivce(userRepo)
	userHandler := handler.NewHandlerUser(userService, jwtService)

	r.POST("/api/user/signup", userHandler.SignUp)
	r.POST("/api/user/login", userHandler.Login)
}
