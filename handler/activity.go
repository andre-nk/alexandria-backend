package handler

import (
	"alexandria/activity"
	"alexandria/helper"
	"alexandria/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type activityHandler struct {
	service activity.Service
}

func NewActivityHandler(service activity.Service) *activityHandler {
	return &activityHandler{service}
}

func (handler *activityHandler) GetActivityByAffiliateID(context *gin.Context) {
	var userID user.UserIDUri

	err := context.ShouldBindUri(&userID)
	if err != nil {
		response := helper.APIResponse(
			"Failed to fetch activities due to invalid Affiliate ID",
			http.StatusUnprocessableEntity,
			"failed",
			err.Error(),
		)

		context.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	activities, err := handler.service.GetActivityByAffiliateID(userID.UID)
	if err != nil {
		response := helper.APIResponse(
			"Failed to fetch activities due to server error",
			http.StatusBadRequest,
			"failed",
			err.Error(),
		)

		context.JSON(http.StatusBadRequest, response)
		return
	}

	if activities[0].AffiliateID == context.MustGet("currentUID") {
		response := helper.APIResponse(
			"Activities fetched!",
			http.StatusOK,
			"success",
			activities,
		)

		context.JSON(http.StatusOK, response)
		return
	}

	response := helper.APIResponse(
		"Activities fetching failed due to unauthorized request!",
		http.StatusUnauthorized,
		"failed",
		nil,
	)

	context.JSON(http.StatusUnauthorized, response)
}

func (handler *activityHandler) MarkActivityAsRead(context *gin.Context) {
	var activityID activity.ActivityIDUri

	err := context.ShouldBindUri(&activityID)
	if err != nil {
		response := helper.APIResponse(
			"Failed to mark activity due to invalid Activity ID",
			http.StatusUnprocessableEntity,
			"failed",
			err.Error(),
		)

		context.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	oldActivity, err := handler.service.GetActivityByID(activityID.ID)
	if err != nil {
		response := helper.APIResponse(
			"Failed to mark activity due to invalid Activity ID",
			http.StatusUnprocessableEntity,
			"failed",
			err.Error(),
		)

		context.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	if oldActivity.AffiliateID == context.MustGet("currentUID") {
		err = handler.service.MarkActivityAsRead(activityID.ID)
		if err != nil {
			response := helper.APIResponse(
				"Failed to mark activity due to server error",
				http.StatusBadRequest,
				"failed",
				err.Error(),
			)

			context.JSON(http.StatusBadRequest, response)
			return
		}

		response := helper.APIResponse(
			"Activities fetched!",
			http.StatusOK,
			"success",
			nil,
		)

		context.JSON(http.StatusOK, response)
		return
	}

	response := helper.APIResponse(
		"Failed to mark activity due to unauthorized request",
		http.StatusUnauthorized,
		"failed",
		nil,
	)

	context.JSON(http.StatusUnauthorized, response)
}
