package handler

import (
	"log"
	"net/http"
	"strconv"

	entities "github.com/Communinst/GolangWebStore/backend/entity"
	"github.com/gin-gonic/gin"
)

func (h *Handler) getWalletByUserID(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	wallet, err := h.service.WalletServiceInterface.GetWalletByUserID(c.Request.Context(), userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve wallet"})
		return
	}

	c.JSON(http.StatusOK, wallet)
}

func (h *Handler) updateWalletBalance(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var wallet entities.Wallet

	if err := c.ShouldBindJSON(&wallet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	userWaller, err := h.service.WalletServiceInterface.GetWalletByUserID(c.Request.Context(), userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update wallet balance"})
		return
	}

	log.Printf("%d + %d", userWaller.Balance, wallet.Balance)

	if userWaller.Balance+wallet.Balance < 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Not enough GOLD"})
		return
	}

	if err := h.service.WalletServiceInterface.UpdateWalletBalance(c.Request.Context(), userId, int64(wallet.Balance)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update wallet balance"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Wallet balance updated successfully"})
}
