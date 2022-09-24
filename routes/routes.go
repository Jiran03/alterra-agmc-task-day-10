package routes

import (
	"os"
	"strconv"

	"github.com/Jiran03/agmc/task/day5/book"
	bookRepo "github.com/Jiran03/agmc/task/day5/book/repository/mysql"
	"github.com/Jiran03/agmc/task/day5/config"
	authMiddleware "github.com/Jiran03/agmc/task/day5/middleware"
	"github.com/Jiran03/agmc/task/day5/user"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	config.Init()
	db := config.DBInit()
	bookDB := bookRepo.Book{}
	config.DBMigrate(db)
	expDuration, _ := strconv.Atoi(os.Getenv("JWT_EXPIRED"))
	configJWT := authMiddleware.ConfigJWT{
		SecretJWT:       os.Getenv("JWT_SECRET"),
		ExpiresDuration: expDuration,
	}

	user := user.NewUserFactory(db, configJWT)
	book := book.NewBookFactory(bookDB)
	e := echo.New()
	authMiddleware.LogMiddleware(e)
	v1 := e.Group("/v1")
	authGroup := v1.Group("/auth")
	bookGroup := v1.Group("/books")
	//not authenticated
	bookGroup.GET("", book.GetAll)
	bookGroup.GET("/:id", book.GetByID)
	//authenticated
	authBookGroup := authGroup.Group("/books")
	authBookGroup.Use(middleware.JWT([]byte(os.Getenv("JWT_SECRET"))))
	authBookGroup.POST("", book.Create)
	authBookGroup.PUT("/:id", book.Update)
	authBookGroup.DELETE("/:id", book.Delete)

	userGroup := v1.Group("/users")
	//not authenticated
	userGroup.POST("", user.Create)
	userGroup.POST("/login", user.Login)
	//authenticated
	authUserGroup := authGroup.Group("/users")
	authUserGroup.Use(middleware.JWT([]byte(os.Getenv("JWT_SECRET"))))
	authUserGroup.GET("", user.GetAll)
	authUserGroup.GET("/:id", user.GetByID)
	//UserValidation function's make the user can only change his/her data
	authUserGroup.PUT("/:id", user.Update, authMiddleware.UserValidation(user))
	authUserGroup.DELETE("/:id", user.Delete, authMiddleware.UserValidation(user))
	return e
}
