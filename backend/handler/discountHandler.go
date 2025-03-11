package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func (h *Handler) addDiscount(c *gin.Context) {
	gameId, err := strconv.Atoi(c.Param("game_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid game ID"})
		return
	}

	var request struct {
		DiscountValue int       `json:"discount_value"`
		StartDate     time.Time `json:"start_date"`
		CeaseDate     time.Time `json:"cease_date"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := h.service.DiscountServiceInterface.AddDiscount(c.Request.Context(), gameId, request.DiscountValue, request.StartDate, request.CeaseDate); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add discount"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Discount added successfully"})
}

func (h *Handler) getDiscountsByGameID(c *gin.Context) {
	gameId, err := strconv.Atoi(c.Param("game_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid game ID"})
		return
	}

	discounts, err := h.service.DiscountServiceInterface.GetDiscountsByGameID(c.Request.Context(), gameId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve discounts"})
		return
	}

	c.JSON(http.StatusOK, discounts)
}

func (h *Handler) deleteDiscount(c *gin.Context) {
	discountId, err := strconv.Atoi(c.Param("discount_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid discount ID"})
		return
	}

	if err := h.service.DiscountServiceInterface.DeleteDiscount(c.Request.Context(), discountId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete discount"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Discount deleted successfully"})
}
