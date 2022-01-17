package handler

import (
	"alexandria/activity"
	"alexandria/helper"
	"alexandria/note"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type noteHandler struct {
	noteService     note.Service
	activityService activity.Service
}

func NewNoteHandler(noteService note.Service, activityService activity.Service) *noteHandler {
	return &noteHandler{noteService, activityService}
}

func (handler *noteHandler) CreateNote(context *gin.Context) {
	var input note.CreateNoteInput

	err := context.ShouldBindJSON(&input)
	if err != nil {
		response := helper.APIResponse(
			"Failed to create note due to bad inputs",
			http.StatusUnprocessableEntity,
			"failed",
			err.Error(),
		)

		context.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newNote, err := handler.noteService.CreateNote(input)
	if err != nil {
		response := helper.APIResponse(
			"Failed to create note due to server error",
			http.StatusBadRequest,
			"failed",
			err.Error(),
		)

		context.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse(
		"Note successfully created!",
		http.StatusOK,
		"success",
		newNote,
	)

	context.JSON(http.StatusOK, response)
}

func (handler *noteHandler) UpdateNote(context *gin.Context) {
	var noteID note.NoteIDUri
	var input note.UpdateNoteInput

	err := context.ShouldBindUri(&noteID)
	if err != nil {
		response := helper.APIResponse(
			"Failed to update note due to invalid ID",
			http.StatusUnprocessableEntity,
			"failed",
			err.Error(),
		)

		context.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	input.ID = noteID

	err = context.ShouldBindJSON(&input)
	if err != nil {
		response := helper.APIResponse(
			"Failed to update note due to bad inputs",
			http.StatusUnprocessableEntity,
			"failed",
			err.Error(),
		)

		context.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	oldNote, err := handler.noteService.GetNoteByID(noteID.ID)
	if err != nil {
		response := helper.APIResponse(
			"Failed to update note due to server error",
			http.StatusBadRequest,
			"failed",
			err.Error(),
		)

		context.JSON(http.StatusBadRequest, response)
		return
	}

	if oldNote.CreatorUID == context.MustGet("currentUID") {
		updatedNote, err := handler.noteService.UpdateNote(input)
		if err != nil {
			response := helper.APIResponse(
				"Failed to update note due to server error",
				http.StatusBadRequest,
				"failed",
				err.Error(),
			)
			context.JSON(http.StatusBadRequest, response)
			return
		}

		fmt.Println(oldNote.Collaborators, updatedNote.Collaborators)

		if len(oldNote.Collaborators) < len(updatedNote.Collaborators) {
			var globalError error

			fmt.Println("Added collab")

			for _, collaboratorID := range updatedNote.Collaborators {
				activity := activity.Activity{
					ID:          primitive.NewObjectID(),
					ActivityID:  oldNote.ID,
					AffiliateID: collaboratorID,
					CreatedAt:   time.Now(),
					IsRead:      false,
					Message:     "You have been invited as collaborator in this note.",
				}

				_, err := handler.activityService.CreateActivity(activity)
				if err != nil {
					globalError = err
				}
			}

			if globalError != nil {
				response := helper.APIResponse(
					"Note successfully updated, but with failed activity creation!",
					http.StatusOK,
					"success",
					err.Error(),
				)

				context.JSON(http.StatusOK, response)
				return
			}

			response := helper.APIResponse(
				"Note successfully updated with created activity!",
				http.StatusOK,
				"success",
				updatedNote,
			)
			context.JSON(http.StatusOK, response)
			return
		} else if len(oldNote.Collaborators) > len(updatedNote.Collaborators) {
			var globalError error

			for _, collaboratorID := range oldNote.Collaborators {
				activity := activity.Activity{
					ID:          primitive.NewObjectID(),
					ActivityID:  oldNote.ID,
					AffiliateID: collaboratorID,
					CreatedAt:   time.Now(),
					IsRead:      false,
					Message:     "You have been removed as collaborator in this note.",
				}

				_, err := handler.activityService.CreateActivity(activity)
				if err != nil {
					globalError = err
				}
			}

			if globalError != nil {
				response := helper.APIResponse(
					"Note successfully updated, but with failed activity creation!",
					http.StatusOK,
					"success",
					err.Error(),
				)

				context.JSON(http.StatusOK, response)
				return
			}

			response := helper.APIResponse(
				"Note successfully updated with created activity!",
				http.StatusOK,
				"success",
				updatedNote,
			)

			context.JSON(http.StatusOK, response)
			return
		} else {
			fmt.Println("Collaborators is not updated")
		}

		response := helper.APIResponse(
			"Note successfully updated!",
			http.StatusOK,
			"success",
			updatedNote,
		)
		context.JSON(http.StatusOK, response)
		return
	}

	response := helper.APIResponse(
		"Note can't be updated due to unauthorized request!",
		http.StatusUnauthorized,
		"failed",
		nil,
	)
	context.JSON(http.StatusUnauthorized, response)
}

func (handler *noteHandler) DeleteNote(context *gin.Context) {
	var noteID note.NoteIDUri

	err := context.ShouldBindUri(&noteID)
	if err != nil {
		response := helper.APIResponse(
			"Failed to delete note due to invalid ID",
			http.StatusUnprocessableEntity,
			"failed",
			err.Error(),
		)

		context.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	oldNote, err := handler.noteService.GetNoteByID(noteID.ID)
	if err != nil {
		response := helper.APIResponse(
			"Failed to delete note due to invalid ID",
			http.StatusUnprocessableEntity,
			"failed",
			err.Error(),
		)

		context.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	if oldNote.CreatorUID == context.MustGet("currentUID") {
		err = handler.noteService.DeleteNote(noteID.ID)
		if err != nil {
			response := helper.APIResponse(
				"Failed to delete note due to server error",
				http.StatusBadRequest,
				"failed",
				err.Error(),
			)

			context.JSON(http.StatusBadRequest, response)
			return
		}

		response := helper.APIResponse(
			"This note has been deleted",
			http.StatusOK,
			"success",
			nil,
		)

		context.JSON(http.StatusOK, response)
		return
	}

	response := helper.APIResponse(
		"Failed to delete note due to unauthorized request!",
		http.StatusUnauthorized,
		"failed",
		err.Error(),
	)

	context.JSON(http.StatusUnauthorized, response)
}

func (handler *noteHandler) GetNotes(context *gin.Context) {
	uid := context.Query("uid")
	isFeatured := context.Query("featured")
	isRecent := context.Query("recent")
	isStarred := context.Query("starred")
	isArchived := context.Query("archived")

	if isFeatured != "" && uid != "" {
		notes, err := handler.noteService.GetFeaturedNotes(uid)
		if err != nil {
			response := helper.APIResponse(
				"Featured notes fetching failed due to server error",
				http.StatusBadGateway,
				"failed",
				err.Error(),
			)
			context.JSON(http.StatusBadGateway, response)
			return
		}

		response := helper.APIResponse(
			"Featured notes fetched!",
			http.StatusOK,
			"success",
			notes,
		)
		context.JSON(http.StatusOK, response)
		return
	}

	if isRecent != "" && uid != "" {
		notes, err := handler.noteService.GetRecentNotes(uid)
		if err != nil {
			response := helper.APIResponse(
				"Recent notes fetching failed due to server error",
				http.StatusBadGateway,
				"failed",
				err.Error(),
			)
			context.JSON(http.StatusBadGateway, response)
			return
		}

		response := helper.APIResponse(
			"Recent notes fetched!",
			http.StatusOK,
			"success",
			notes,
		)
		context.JSON(http.StatusOK, response)
		return
	}

	if isStarred != "" && uid != "" {
		notes, err := handler.noteService.GetStarredNotes(uid)
		if err != nil {
			response := helper.APIResponse(
				"Starred notes fetching failed due to server error",
				http.StatusBadGateway,
				"failed",
				err.Error(),
			)
			context.JSON(http.StatusBadGateway, response)
			return
		}

		response := helper.APIResponse(
			"Starred notes fetched!",
			http.StatusOK,
			"success",
			notes,
		)
		context.JSON(http.StatusOK, response)
		return
	}

	if isArchived != "" && uid != "" {
		notes, err := handler.noteService.GetArchivedNotes(uid)
		if err != nil {
			response := helper.APIResponse(
				"Archived notes fetching failed due to server error",
				http.StatusBadGateway,
				"failed",
				err.Error(),
			)
			context.JSON(http.StatusBadGateway, response)
			return
		}

		response := helper.APIResponse(
			"Archived notes fetched!",
			http.StatusOK,
			"success",
			notes,
		)
		context.JSON(http.StatusOK, response)
		return
	}

	if uid != "" {
		notes, err := handler.noteService.GetNotesByUserID(uid)
		if err != nil {
			response := helper.APIResponse(
				"Notes fetching by UID failed due to server error",
				http.StatusBadGateway,
				"failed",
				err.Error(),
			)
			context.JSON(http.StatusBadGateway, response)
			return
		}

		response := helper.APIResponse(
			"Notes fetching by UID success!",
			http.StatusOK,
			"success",
			notes,
		)
		context.JSON(http.StatusOK, response)
		return
	}

	//FALLBACK
	response := helper.APIResponse(
		"False endpoint",
		http.StatusBadRequest,
		"success",
		nil,
	)
	context.JSON(http.StatusBadRequest, response)
}

func (handler *noteHandler) GetNoteByID(context *gin.Context) {
	var noteID note.NoteIDUri

	err := context.ShouldBindUri(&noteID)
	if err != nil {
		response := helper.APIResponse(
			"Failed to fetch note due to invalid ID",
			http.StatusUnprocessableEntity,
			"failed",
			err.Error(),
		)

		context.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	note, err := handler.noteService.GetNoteByID(noteID.ID)
	if err != nil {
		response := helper.APIResponse(
			"Failed to fetch note due to server error",
			http.StatusBadRequest,
			"failed",
			err.Error(),
		)

		context.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse(
		"Note by ID fetched!",
		http.StatusOK,
		"success",
		note,
	)

	context.JSON(http.StatusOK, response)
}
