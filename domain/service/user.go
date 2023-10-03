package service

import "article/domain"

type IUserService interface {
	SignUp(input domain.UserSignupRequest) (domain.User, int, error)
	Login(input domain.UserLoginRequest) (domain.User, int, error)
}
