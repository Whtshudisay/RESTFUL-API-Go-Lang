package routes

import (
	"awesomeProject/RESTFUL-API/Model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func registerForEvent(context *gin.Context) {
	userID := context.GetInt64("userID")
	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Message": "Could not parse event id !"})
		return
	}

	event, err := Model.GetEventByID(eventID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Message": "Could not Fetch Event !"})
		return
	}

	err = event.Register(userID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Message": "Could Not Register For Event !"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"Message": "Registered !"})
}

func cancelRegistrations(context *gin.Context) {
	userID := context.GetInt64("userID")
	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)

	var event Model.Event
	event.ID = eventID

	err = event.CancelRegistrations(userID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Message": "Could Not Cancel Registration For The Event !"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"Message": "Cancelled !"})
}
