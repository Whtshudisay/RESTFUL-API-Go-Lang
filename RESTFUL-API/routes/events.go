package routes

import (
	"awesomeProject/RESTFUL-API/Model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func getEvents(context *gin.Context) {
	events, err := Model.GetEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Message": "Could not fetch Events. Try again Later !"})
		return
	}
	context.JSON(http.StatusOK, events)
}

func getEvent(context *gin.Context) {
	eventid, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Message": "Could not parse event id !"})
		return
	}
	event, err := Model.GetEventByID(eventid)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Message": "Could not fetch event !"})
		return
	}

	context.JSON(http.StatusOK, event)
}

func createEvent(context *gin.Context) {

	var event Model.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request data."})
		return
	}
	userid := context.GetInt64("userID")
	event.UserId = userid

	err = event.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Message": "Could not create Event. Try again Later !"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Event Created! ", "Event": event})
}

func updateEvent(context *gin.Context) {
	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Message": "Could not parse event id !"})
		return
	}

	userid := context.GetInt64("userID")
	event, err := Model.GetEventByID(eventID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Message": "Could not fetch the event !"})
		return
	}

	if event.UserId != userid {
		context.JSON(http.StatusUnauthorized, gin.H{"Message": "Not Authorized to Update Events !"})
		return
	}

	var updatedEvent Model.Event
	err = context.ShouldBindJSON(&updatedEvent)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request data."})
		return
	}

	updatedEvent.ID = eventID

	err = updatedEvent.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Message": "Could not update event !"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"Message": "Event Updated Successfully !"})
}

func deleteEvent(context *gin.Context) {
	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Message": "Could not parse event id !"})
		return
	}
	userid := context.GetInt64("userID")
	event, err := Model.GetEventByID(eventID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Message": "Could not fetch the event !"})
		return
	}

	if event.UserId != userid {
		context.JSON(http.StatusUnauthorized, gin.H{"Message": "Not Authorized to Delete Events !"})
		return
	}

	err = event.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Message": "Could not delete the event !"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"Message": "Event Deleted Successfully !"})
}
