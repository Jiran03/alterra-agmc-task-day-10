package handlerAPI

import (
	"time"

	helperTime "github.com/Jiran03/agmc/task/day5/helpers/time"
	"github.com/Jiran03/agmc/task/day5/user/domain"
)

type RequestJSON struct {
	Name     string `json:"name" form:"name" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required,min=8"`
}

type RequestLoginJSON struct {
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required,min=8"`
}

type Token struct {
	Token string `json:"token"`
}

func toDomain(req RequestJSON) domain.User {
	return domain.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}
}

type ResponseJSON struct {
	ID        int       `json:"id"`
	Name      string    `json:"name" form:"name"`
	Email     string    `json:"email" form:"email"`
	CreatedAt time.Time `json:"created_at" form:"created_at"`
	UpdatedAt time.Time `json:"updated_at" form:"updated_at"`
}

func fromDomain(domain domain.User) ResponseJSON {
	//parse unix timestamp to time.Time
	tmCreatedAt := helperTime.NanoToTime(domain.CreatedAt)
	tmUpdatedAt := helperTime.NanoToTime(domain.UpdatedAt)

	return ResponseJSON{
		ID:        domain.ID,
		Name:      domain.Name,
		Email:     domain.Email,
		CreatedAt: tmCreatedAt,
		UpdatedAt: tmUpdatedAt,
	}
}
