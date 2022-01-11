package handler

import (
	"alexandria/comment"
	"alexandria/helper"
	"alexandria/note"
	"fmt"
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
		"Comment successfully created!",
		http.StatusOK,
		"success",
		newComment,
	)

	context.JSON(http.StatusOK, response)
}

func (handler *commentHandler) GetCommentsByNoteID(context *gin.Context) {
	var noteID note.NoteIDUri

	err := context.ShouldBindUri(&noteID)
	if err != nil {
		response := helper.APIResponse(
			"Failed to fetch comments due to invalid NoteID",
			http.StatusUnprocessableEntity,
			"failed",
			err.Error(),
		)

		context.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	fmt.Println(noteID.ID)

	comments, err := handler.service.GetCommentsByNoteID(noteID.ID)
	if err != nil {
		response := helper.APIResponse(
			"Failed to fetch comments due to server error",
			http.StatusBadRequest,
			"failed",
			err.Error(),
		)

		context.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse(
		"Comments fetched!",
		http.StatusOK,
		"success",
		comments,
	)

	context.JSON(http.StatusOK, response)
}

func (handler *commentHandler) GetCommentByID(context *gin.Context) {
	var commentID comment.CommentIDUri

	err := context.ShouldBindUri(&commentID)
	if err != nil {
		response := helper.APIResponse(
			"Failed to fetch comment due to invalid ID",
			http.StatusUnprocessableEntity,
			"failed",
			err.Error(),
		)

		context.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	note, err := handler.service.GetCommentByID(commentID.ID)
	if err != nil {
		response := helper.APIResponse(
			"Failed to fetch comment due to server error",
			http.StatusBadRequest,
			"failed",
			err.Error(),
		)

		context.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse(
		"Comment by ID fetched!",
		http.StatusOK,
		"success",
		note,
	)

	context.JSON(http.StatusOK, response)
}

func (handler *commentHandler) DeleteCommentByID(context *gin.Context) {
	var commentID comment.CommentIDUri

	err := context.ShouldBindUri(&commentID)
	if err != nil {
		response := helper.APIResponse(
			"Failed to delete comment due to invalid ID",
			http.StatusUnprocessableEntity,
			"failed",
			err.Error(),
		)

		context.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	err = handler.service.DeleteComment(commentID.ID)
	if err != nil {
		response := helper.APIResponse(
			"Failed to delete comment due to server error",
			http.StatusBadRequest,
			"failed",
			err.Error(),
		)

		context.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse(
		"Comment by ID deleted!",
		http.StatusOK,
		"success",
		nil,
	)

	context.JSON(http.StatusOK, response)
}
