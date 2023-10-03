package interfaces

import (
	"article/auth"
	userdomain "article/domain"
	userservice "article/domain/service"
	"article/helper"
	interfaces "article/interfaces/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HandlerUser struct {
	userService userservice.IUserService
	jwtService  auth.IJWTService
}

func NewHandlerUser(userService userservice.IUserService, jwtService auth.IJWTService) *HandlerUser {
	return &HandlerUser{userService, jwtService}
}

func (h *HandlerUser) SignUp(c *gin.Context) {
	var request userdomain.UserSignupRequest

	if err := c.BindJSON(&request); err != nil {
		errBinding := helper.GenerateErrorBinding(err)
		response := helper.GenerateResponseAPI(
			http.StatusBadRequest,
			"error binding",
			errBinding,
		)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	userSignedUp, httpCode, err := h.userService.SignUp(request)
	if err != nil {
		response := helper.GenerateResponseAPI(
			httpCode,
			"failed signup",
			err.Error(),
		)

		c.JSON(httpCode, response)
		return
	}

	userFormatted := interfaces.FormatUserSignUpResponse(userSignedUp)

	response := helper.GenerateResponseAPI(
		httpCode,
		"success",
		userFormatted,
	)

	c.JSON(httpCode, response)
}

func (h *HandlerUser) Login(c *gin.Context) {
	var request userdomain.UserLoginRequest

	if err := c.BindJSON(&request); err != nil {
		errBinding := helper.GenerateErrorBinding(err)
		response := helper.GenerateResponseAPI(
			http.StatusBadRequest,
			"error binding",
			errBinding,
		)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	userLoggedin, httpCode, err := h.userService.Login(request)
	if err != nil {
		response := helper.GenerateResponseAPI(
			httpCode,
			"failed login",
			err.Error(),
		)

		c.JSON(httpCode, response)
		return
	}

	token, err := h.jwtService.GenerateToken(userLoggedin.Id)
	if err != nil {
		response := helper.GenerateResponseAPI(
			httpCode,
			"failed generate token",
			err.Error(),
		)

		c.JSON(httpCode, response)
		return
	}

	userFormatted := interfaces.FormatUserLoginResponse(token)

	response := helper.GenerateResponseAPI(
		httpCode,
		"success",
		userFormatted,
	)

	c.JSON(httpCode, response)
}
