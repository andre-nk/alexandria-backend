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

	oldUser, err := handler.service.GetUserByUID(input.UID)
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

	if oldUser.UID == context.MustGet("currentUID") {
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
		return
	}

	response := helper.APIResponse(
		"User can't be updated due to unauthorized request!",
		http.StatusUnauthorized,
		"failed",
		nil,
	)
	context.JSON(http.StatusUnauthorized, response)
}

func (handler *userHandler) DeleteUser(context *gin.Context) {
	var uid user.UserIDUri

	err := context.ShouldBindUri(&uid)
	if err != nil {
		response := helper.APIResponse(
			"Failed to delete user due to invalid UID",
			http.StatusUnprocessableEntity,
			"failed",
			err.Error(),
		)

		context.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	oldUser, err := handler.service.GetUserByUID(uid.UID)
	if err != nil {
		response := helper.APIResponse(
			"Failed to delete user due to bad inputs",
			http.StatusUnprocessableEntity,
			"failed",
			err.Error(),
		)

		context.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	if oldUser.UID == context.MustGet("currentUID") {
		err = handler.service.DeleteUser(uid.UID)
		if err != nil {
			response := helper.APIResponse(
				"Failed to delete user due to server error",
				http.StatusBadRequest,
				"failed",
				err.Error(),
			)

			context.JSON(http.StatusBadRequest, response)
			return
		}

		response := helper.APIResponse(
			"User deleted!",
			http.StatusOK,
			"success",
			nil,
		)

		context.JSON(http.StatusOK, response)
	}

	response := helper.APIResponse(
		"User can't be deleted due to unauthorized request!",
		http.StatusUnauthorized,
		"failed",
		nil,
	)
	context.JSON(http.StatusUnauthorized, response)
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

func (handler *userHandler) GetUserByEmail(context *gin.Context) {
	var email user.UserEmailUri

	err := context.ShouldBindUri(&email)
	if err != nil {
		response := helper.APIResponse(
			"Failed to fetch user due to invalid email",
			http.StatusUnprocessableEntity,
			"failed",
			err.Error(),
		)

		context.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	user, err := handler.service.GetUserByEmail(email.Email)

	if err != nil && err.Error() == "mongo: no documents in result" {
		response := helper.APIResponse(
			"No user found with this email",
			http.StatusNotFound,
			"failed",
			err.Error(),
		)

		context.JSON(http.StatusNoContent, response)
		return
	}

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
