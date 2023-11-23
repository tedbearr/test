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

	userRepository repository.UserRepositoy  = repository.NewUserRepository(db)
	userService    service.UserService       = service.NewUserService(userRepository)
	userController controller.UserController = controller.NewUserController(userService)

	adminRepository repository.AdminRepositoy  = repository.NewAdminRepository(db)
	adminService    service.AdminService       = service.NewAdminService(adminRepository)
	adminController controller.AdminController = controller.NewAdminController(adminService)
)
