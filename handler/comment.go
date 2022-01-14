package handler

import (
	"alexandria/activity"
	"alexandria/comment"
	"alexandria/helper"
	"alexandria/note"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type commentHandler struct {
	commentService  comment.Service
	activityService activity.Service
}

func NewCommentHandler(commentService comment.Service, activityService activity.Service) *commentHandler {
	return &commentHandler{commentService, activityService}
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

	newComment, err := handler.commentService.CreateComment(input)
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

	if newComment.Mentions != nil {
		for _, mentionedID := range newComment.Mentions {
			activity := activity.Activity{
				ID:          primitive.NewObjectID(),
				ActivityID:  newComment.ID,
				AffiliateID: mentionedID,
				CreatedAt:   time.Now(),
				IsRead:      false,
				Message:     "You're mentioned at this comment.",
			}

			_, err := handler.activityService.CreateActivity(activity)
			if err != nil {
				response := helper.APIResponse(
					"Comment successfully created, but with failed activity creation!",
					http.StatusOK,
					"success",
					err.Error(),
				)

				context.JSON(http.StatusOK, response)
				return
			}
		}

		response := helper.APIResponse(
			"Comment successfully created with created activity!",
			http.StatusOK,
			"success",
			newComment,
		)
		context.JSON(http.StatusOK, response)
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

	comments, err := handler.commentService.GetCommentsByNoteID(noteID.ID)
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

	note, err := handler.commentService.GetCommentByID(commentID.ID)
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

	err = handler.commentService.DeleteComment(commentID.ID)
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
