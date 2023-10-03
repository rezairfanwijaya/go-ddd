package infrastructure

import (
	"article/domain"
	"article/domain/repository"
	"article/domain/service"
	"article/helper"
	"errors"
	"fmt"
	"net/http"
)

type UserService struct {
	userRepository repository.IUserRepository
}

func NewUserSerivce(userRepository repository.IUserRepository) service.IUserService {
	return &UserService{userRepository}
}

func (u *UserService) SignUp(input domain.UserSignupRequest) (domain.User, int, error) {
	userByEmail, err := u.userRepository.FindByEmail(input.Email)
	if err != nil {
		return userByEmail, http.StatusInternalServerError, err
	}

	if userByEmail.Id != 0 {
		return userByEmail, http.StatusBadRequest, fmt.Errorf("email already taken")
	}

	var user domain.User
	user.Email = input.Email

	hashedPassword, err := helper.HashPassowrd(input.Password)
	if err != nil {
		return userByEmail, http.StatusBadRequest, err
	}
	user.Password = hashedPassword

	userSaved, err := u.userRepository.Save(user)
	if err != nil {
		return userSaved, http.StatusInternalServerError, err
	}

	return userSaved, http.StatusOK, nil
}

func (u *UserService) Login(input domain.UserLoginRequest) (domain.User, int, error) {
	userByEmail, err := u.userRepository.FindByEmail(input.Email)
	if err != nil {
		return userByEmail, http.StatusInternalServerError, err
	}

	if userByEmail.Id == 0 {
		return userByEmail, http.StatusBadRequest, errors.New("email not registered")
	}

	err = helper.VerifyPassword(input.Password, userByEmail.Password)
	if err != nil {
		return userByEmail, http.StatusInternalServerError, fmt.Errorf("wrong password")
	}

	return userByEmail, http.StatusOK, nil

}
