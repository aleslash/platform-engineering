package routes

import (
	"net/http"
	"strconv"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func registerForEvent(context *gin.Context) {
	// Implementation for registering a user for an event
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}

	event, err := models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}

	err = event.RegisterUser(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register for event"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Successfully registered for the event"})
}

func unregisterForEvent(context *gin.Context) {
	// Implementation for unregistering a user from an event
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}

	event, err := models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}

	err = event.UnregisterUser(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unregister from event"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Successfully unregistered from the event"})
}
