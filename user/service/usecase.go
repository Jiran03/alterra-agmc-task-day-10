package service

import (
	"errors"

	encryptHelper "github.com/Jiran03/agmc/task/day5/helpers/encrypt"
	timeHelper "github.com/Jiran03/agmc/task/day5/helpers/time"
	authMiddleware "github.com/Jiran03/agmc/task/day5/middleware"
	"github.com/Jiran03/agmc/task/day5/user/domain"
)

type userService struct {
	repository domain.Repository
	jwtAuth    authMiddleware.ConfigJWT
}

// UpdateData implements domain.Service
func (us userService) UpdateData(id int, domain domain.User) (userObj domain.User, err error) {
	if userObj, err = us.GetByID(id); err != nil {
		return userObj, err
	}

	domain.ID = id
	if userObj, err = us.repository.Update(domain); err != nil {
		return userObj, err
	}

	return userObj, nil
}

// GetByEmail implements domain.Service
func (us userService) GetByEmail(email string) (userObj domain.User, err error) {
	userObj, err = us.repository.GetByEmail(email)

	if err != nil {
		return userObj, err
	}

	return userObj, nil
}

// GetByID implements domain.Service
func (us userService) GetByID(id int) (userObj domain.User, err error) {
	userObj, err = us.repository.GetByID(id)
	if err != nil {
		return userObj, err
	}

	return userObj, nil
}

func (us userService) CreateToken(email, password string) (token string, userObj domain.User, err error) {
	userObj, err = us.repository.GetByEmail(email)
	if err != nil {
		return token, userObj, err
	}

	if !encryptHelper.ValidateHash(password, userObj.Password) {
		return token, userObj, errors.New("email atau kata sandi salah")
	}

	id := userObj.ID
	token, err = us.jwtAuth.CreateToken(id, email)
	if err != nil {
		return token, userObj, err
	}

	userObj, err = us.GetByID(id)
	if err != nil {
		return token, userObj, err
	}

	return token, userObj, nil
}

func (us userService) InsertData(domain domain.User) (userObj domain.User, err error) {
	email := domain.Email
	_, errGetUser := us.repository.GetByEmail(email)
	if errGetUser == nil {
		return userObj, errors.New("email telah terdaftar")
	}

	domain.Password, err = encryptHelper.Hash(domain.Password)
	if err != nil {
		return userObj, err
	}

	domain.CreatedAt = timeHelper.Timestamp()
	domain.UpdatedAt = timeHelper.Timestamp()
	userObj, err = us.repository.Create(domain)

	if err != nil {
		return userObj, err
	}

	return userObj, nil
}

// GetAllData implements domain.Service
func (us userService) GetAllData() (userObj []domain.User, err error) {
	userObj, _ = us.repository.Get()

	if err != nil {
		return userObj, err
	}

	return userObj, nil
}

// DeleteData implements domain.Service
func (us userService) DeleteData(id int) (err error) {
	errResp := us.repository.Delete(id)

	if errResp != nil {
		return errResp
	}

	return nil
}

func NewUserService(repo domain.Repository, jwtAuth authMiddleware.ConfigJWT) domain.Service {
	return userService{
		repository: repo,
		jwtAuth:    jwtAuth,
	}
}
