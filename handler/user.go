package handler

import (
	"alexandria/helper"
	"alexandria/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	service user.Service
}

func NewUserHandler(service user.Service) *userHandler {
	return &userHandler{service}
}

func (handler *userHandler) RegisterUser(context *gin.Context) {
	var input user.UserInput

	err := context.ShouldBindJSON(&input)
	if err != nil {
		response := helper.APIResponse(
			"Failed to register user due to bad inputs",
			http.StatusUnprocessableEntity,
			"failed",
			err.Error(),
		)

		context.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newUser, err := handler.service.RegisterUser(input)
	if err != nil {
		response := helper.APIResponse(
			"Failed to register user due to server error",
			http.StatusBadRequest,
			"failed",
			err.Error(),
		)

		context.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse(
		"User successfully registered!",
		http.StatusOK,
		"success",
		newUser,
	)

	context.JSON(http.StatusOK, response)
}

func (handler *userHandler) UpdateUser(context *gin.Context) {
	var input user.UserInput

	err := context.ShouldBindJSON(&input)
	if err != nil {
		response := helper.APIResponse(
			"Failed to update user due to bad inputs",
			http.StatusUnprocessableEntity,
			"failed",
			err.Error(),
		)

		context.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	updatedUser, err := handler.service.UpdateUser(input)
	if err != nil {
		response := helper.APIResponse(
			"Failed to update user due to server error",
			http.StatusBadRequest,
			"failed",
			err.Error(),
		)

		context.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse(
		"User successfully updated!",
		http.StatusOK,
		"success",
		updatedUser,
	)

	context.JSON(http.StatusOK, response)
}

func (handler *userHandler) GetUserByUID(context *gin.Context) {
	var uid user.UserIDUri

	err := context.ShouldBindUri(&uid)
	if err != nil {
		response := helper.APIResponse(
			"Failed to fetch user due to invalid UID",
			http.StatusUnprocessableEntity,
			"failed",
			err.Error(),
		)

		context.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	user, err := handler.service.GetUserByUID(uid.UID)
	if err != nil {
		response := helper.APIResponse(
			"Failed to fetch user due to server error",
			http.StatusBadRequest,
			"failed",
			err.Error(),
		)

		context.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse(
		"User fetched!",
		http.StatusOK,
		"success",
		user,
	)

	context.JSON(http.StatusOK, response)
}
