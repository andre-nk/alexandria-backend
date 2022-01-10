package handler

import (
	"alexandria/comment"
	"alexandria/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

type commentHandler struct {
	service comment.Service
}

func NewCommentHandler(service comment.Service) *commentHandler {
	return &commentHandler{service}
}

func (handler *commentHandler) CreateComment(context *gin.Context) {
	var input comment.CreateCommentInput

	err := context.ShouldBindJSON(&input)
	if err != nil {
		response := helper.APIResponse(
			"Failed to create comment due to bad inputs",
			http.StatusUnprocessableEntity,
			"failed",
			err.Error(),
		)

		context.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newComment, err := handler.service.CreateComment(input)
	if err != nil {
		response := helper.APIResponse(
			"Failed to create comment due to server error",
			http.StatusBadRequest,
			"failed",
			err.Error(),
		)

		context.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse(
		"comment successfully created!",
		http.StatusOK,
		"success",
		newComment,
	)

	context.JSON(http.StatusOK, response)
}
