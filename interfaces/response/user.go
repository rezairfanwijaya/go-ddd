package interfaces

import "article/domain"

type userResponseSignup struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
}

func FormatUserSignUpResponse(user domain.User) *userResponseSignup {
	return &userResponseSignup{
		Id:    user.Id,
		Email: user.Email,
	}
}
