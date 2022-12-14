package user

import (
	authMiddleware "github.com/Jiran03/agmc/task/day5/middleware"
	handlerAPI "github.com/Jiran03/agmc/task/day5/user/handler/api"
	repoMySQL "github.com/Jiran03/agmc/task/day5/user/repository/mysql"
	service "github.com/Jiran03/agmc/task/day5/user/service"
	"gorm.io/gorm"
)

func NewUserFactory(db *gorm.DB, configJWT authMiddleware.ConfigJWT) (userHandler handlerAPI.UserHandler) {
	userRepo := repoMySQL.NewUserRepository(db)
	userService := service.NewUserService(userRepo, configJWT)
	userHandler = handlerAPI.NewUserHandler(userService)
	return
}
