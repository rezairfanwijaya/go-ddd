package helper

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

type responseAPI struct {
	Meta meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type meta struct {
	Status   int    `json:"status"`
	Messsage string `json:"message"`
}

func GenerateResponseAPI(status int, message string, data interface{}) *responseAPI {
	return &responseAPI{
		Meta: meta{
			Status:   status,
			Messsage: message,
		},
		Data: data,
	}
}

func HashPassowrd(rawPassword string) (string, error) {
	passByte, err := bcrypt.GenerateFromPassword([]byte(rawPassword), 10)
	if err != nil {
		return "", fmt.Errorf("hash password failed : %v", err.Error())
	}

	return string(passByte), nil
}

func VerifyPassword(rawPassword, hashedPassword string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(rawPassword)); err != nil {
		return err
	}

	return nil
}

func GenerateErrorBinding(err error) (errorBinding []string) {
	for _, e := range err.(validator.ValidationErrors) {
		errorBinding = append(errorBinding, e.Error())
	}

	return errorBinding
}
