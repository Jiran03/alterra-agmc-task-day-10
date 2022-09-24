package middleware

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	userHandlerAPI "github.com/Jiran03/agmc/task/day5/user/handler/api"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type JWTCustomClaim struct {
	ID         int    `json:"id"`
	Email      string `json:"email"`
	Authorized bool   `json:"authorized"`
	jwt.StandardClaims
}

type ConfigJWT struct {
	SecretJWT       string
	ExpiresDuration int
}

func (cJWT ConfigJWT) Init() middleware.JWTConfig {
	return middleware.JWTConfig{
		Claims:     &JWTCustomClaim{},
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
	}
}

func (cJWT ConfigJWT) CreateToken(userID int, userEmail string) (token string, err error) {
	claims := JWTCustomClaim{
		userID,
		userEmail,
		true,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(int64(cJWT.ExpiresDuration))).Unix(),
		},
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err = t.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return token, err
	}

	return token, nil
}

func ExtractToken(ctx echo.Context) (userID, paramID int, userEmail string) {
	// user := ctx.Get("user").(*jwt.Token)
	// if user.Valid {
	// 	claims := user.Claims.(jwt.MapClaims)
	// 	userID = int(claims["userID"].(float64))
	// 	userEmail = claims["userEmail"].(string)

	// 	return userID, userEmail
	// }
	paramID, _ = strconv.Atoi(ctx.Param("id"))
	if temp := ctx.Get("user"); temp != nil {
		user := temp.(*jwt.Token)
		claims := user.Claims.(*JWTCustomClaim)
		userID = claims.ID
		userEmail = claims.Email
	}
	fmt.Println("userId, email", userID, userEmail)
	return userID, paramID, userEmail
}

func UserValidation(userController userHandlerAPI.UserHandler) echo.MiddlewareFunc {
	return func(hf echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			_, paramID, userEmail := ExtractToken(ctx)
			validUserID, err := userController.GetValidEmail(userEmail)
			if paramID != validUserID {
				return errors.New("anda tidak dapat mengakses endpoint ini")
			}

			if err != nil {
				return err
			}

			return hf(ctx)
		}
	}
}
