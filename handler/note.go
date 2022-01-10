package handler

import (
	"alexandria/helper"
	"alexandria/note"
	"net/http"

	"github.com/gin-gonic/gin"
)

type noteHandler struct {
	service note.Service
}

func NewNoteHandler(service note.Service) *noteHandler {
	return &noteHandler{service}
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

	newNote, err := handler.service.CreateNote(input)
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

	updatedNote, err := handler.service.UpdateNote(input)
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

	response := helper.APIResponse(
		"Note successfully updated!",
		http.StatusOK,
		"success",
		updatedNote,
	)

	context.JSON(http.StatusOK, response)
}

func (handler *noteHandler) DeleteNote(context *gin.Context) {
	var noteID note.NoteIDUri
	var input note.UpdateNoteInput

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

	input.ID = noteID

	err = handler.service.DeleteNote(input)
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
}

func (handler *noteHandler) GetNoteByID(context *gin.Context) {
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

	note, err := handler.service.GetNoteByID(noteID.ID)
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
