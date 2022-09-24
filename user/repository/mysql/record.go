package repoMySQL

import (
	"github.com/Jiran03/agmc/task/day5/user/domain"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        int
	Name      string
	Email     string
	Password  string
	CreatedAt string
	UpdatedAt string
}

func toDomain(rec User) domain.User {
	return domain.User{
		ID:       rec.ID,
		Name:     rec.Name,
		Email:    rec.Email,
		Password: rec.Password,
		// Gender:    rec.Gender,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
	}
}

func fromDomain(rec domain.User) User {
	return User{
		ID:        rec.ID,
		Name:      rec.Name,
		Email:     rec.Email,
		Password:  rec.Password,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
	}
}
