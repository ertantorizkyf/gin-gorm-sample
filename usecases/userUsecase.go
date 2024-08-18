package usecases

import (
	"errors"
	"fmt"

	"github.com/ertantorizkyf/gin-gorm-sample/dtos"
	"github.com/ertantorizkyf/gin-gorm-sample/helpers"
	"github.com/ertantorizkyf/gin-gorm-sample/repositories"
)

type UserUsecase struct {
	userRepository repositories.UserRepository
}

func NewUserUsecase(userRepository repositories.UserRepository) UserUsecase {
	return UserUsecase{
		userRepository: userRepository,
	}
}

func (uc *UserUsecase) RegisterUser(req dtos.RegisterReq) (int, error) {
	req.Password, _ = helpers.HashPassword(req.Password)
	existingUser := uc.userRepository.GetUserByEmail(req.Email)
	if existingUser.Email != "" {
		return 400, errors.New(fmt.Sprintf("[ERR] User with %s email already exists", req.Email))
	}

	err := uc.userRepository.RegisterUser(req)
	if err != nil {
		return 500, err
	}

	return 201, nil
}

func (uc *UserUsecase) Login(req dtos.LoginReq) (string, int, error) {
	user := uc.userRepository.GetUserByEmail(req.Email)
	if user.Email == "" {
		return "", 400, errors.New("[ERR] Incorrect email or password")
	}

	isMatch := helpers.CheckPasswordHash(req.Password, user.Password)
	if !isMatch {
		return "", 400, errors.New("[ERR] Incorrect email or password")
	}

	token, err := helpers.GenerateToken(user)
	if err != nil {
		return "", 500, err
	}

	return token, 200, nil
}
