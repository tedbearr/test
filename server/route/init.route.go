package route

import (
	"server/config"
	"server/controller"
	"server/repository"
	"server/service"

	"gorm.io/gorm"
)

var (
	db *gorm.DB = config.DatabaseInit()

	authRepository repository.AuthRepository = repository.NewAuthRepository(db)
	authService    service.AuthService       = service.NewAuthService(authRepository)
	authController controller.AuthController = controller.NewAuthController(authService)
)
