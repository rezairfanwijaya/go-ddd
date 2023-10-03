package interfaces

import (
	userdomain "article/domain"
	userservice "article/domain/service"
	"article/helper"
	interfaces "article/interfaces/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HandlerUser struct {
	userService userservice.IUserService
}

func NewHandlerUser(userService userservice.IUserService) *HandlerUser {
	return &HandlerUser{userService}
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
