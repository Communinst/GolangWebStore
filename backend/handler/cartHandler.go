package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) addGameToCart(c *gin.Context) {
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

	if err := h.service.CartServiceInterface.AddGameToCart(c.Request.Context(), userId, gameId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add game to cart"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Game added to cart successfully"})
}

func (h *Handler) getCartByUserID(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	games, err := h.service.CartServiceInterface.GetCartByUserID(c.Request.Context(), userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve cart"})
		return
	}

	c.JSON(http.StatusOK, games)
}

func (h *Handler) removeGameFromCart(c *gin.Context) {
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

	if err := h.service.CartServiceInterface.RemoveGameFromCart(c.Request.Context(), userId, gameId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to remove game from cart"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Game removed from cart successfully"})
}
