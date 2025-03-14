package handler

import (
	"net/http"
	"strconv"
	"time"

	entities "github.com/Communinst/GolangWebStore/backend/entity"
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

	wallet, err := h.service.GetWalletByUserID(c.Request.Context(), userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to add ownership"})
		return
	}
	game, err := h.service.GetGame(c.Request.Context(), gameId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add ownership"})
		return
	}

	if game.Price > wallet.Balance {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add ownership. Not enough balance"})
		return
	}

	request := &entities.Ownership{
		ReceiptDate:  time.Now(),
		MinutesSpent: 0,
	}

	if err := h.service.OwnershipServiceInterface.AddOwnership(c.Request.Context(), userId, gameId, request.MinutesSpent, request.ReceiptDate); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add ownership"})
		return
	}
	h.service.UpdateWalletBalance(c.Request.Context(), userId, int64(wallet.Balance)-int64(game.Price))
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
