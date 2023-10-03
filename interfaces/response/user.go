package interfaces

import "article/domain"

type userSignupResponse struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
}

type userLoginResponse struct {
	Token string `json:"token"`
}

func FormatUserSignUpResponse(user domain.User) *userSignupResponse {
	return &userSignupResponse{
		Id:    user.Id,
		Email: user.Email,
	}
}

func FormatUserLoginResponse(token string) *userLoginResponse {
	return &userLoginResponse{
		Token: token,
	}
}
