package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func (h *Handler) addOwnership(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	gameId, err := strconv.Atoi(c.Param("game_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid game ID"})
		return
	}

	var request struct {
		MinutesSpent int64     `json:"minutes_spent"`
		ReceiptDate  time.Time `json:"receipt_date"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := h.service.OwnershipServiceInterface.AddOwnership(c.Request.Context(), userId, gameId, request.MinutesSpent, request.ReceiptDate); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add ownership"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Ownership added successfully"})
}

func (h *Handler) getOwnershipsByUserID(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	ownerships, err := h.service.OwnershipServiceInterface.GetOwnershipsByUserID(c.Request.Context(), userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve ownerships"})
		return
	}

	c.JSON(http.StatusOK, ownerships)
}

func (h *Handler) getOwnershipsByGameID(c *gin.Context) {
	gameId, err := strconv.Atoi(c.Param("game_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid game ID"})
		return
	}

	ownerships, err := h.service.OwnershipServiceInterface.GetOwnershipsByGameID(c.Request.Context(), gameId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve ownerships"})
		return
	}

	c.JSON(http.StatusOK, ownerships)
}

func (h *Handler) deleteOwnership(c *gin.Context) {
	ownershipId, err := strconv.Atoi(c.Param("ownership_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ownership ID"})
		return
	}

	if err := h.service.OwnershipServiceInterface.DeleteOwnership(c.Request.Context(), ownershipId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete ownership"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Ownership deleted successfully"})
}
