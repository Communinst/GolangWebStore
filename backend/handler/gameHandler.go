package handler

import (
	"log"
	"net/http"
	"strconv"

	entities "github.com/Communinst/GolangWebStore/backend/entity"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createGame(c *gin.Context) {
	var game entities.Game
	if err := c.ShouldBindJSON(&game); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := h.service.GameServiceInterface.PostGame(c.Request.Context(), &game); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create game"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Game created successfully"})
}

func (h *Handler) getGame(c *gin.Context) {
	gameId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid game ID"})
		return
	}

	game, err := h.service.GameServiceInterface.GetGame(c.Request.Context(), gameId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Game not found"})
		return
	}

	c.JSON(http.StatusOK, game)
}

func (h *Handler) getAllGames(c *gin.Context) {
	games, err := h.service.GameServiceInterface.GetAllGames(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve games"})
		return
	}

	c.JSON(http.StatusOK, games)
}

func (h *Handler) deleteGame(c *gin.Context) {
	gameId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid game ID"})
		return
	}

	if err := h.service.GameServiceInterface.DeleteGame(c.Request.Context(), gameId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete game"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Game deleted successfully"})
}

func (h *Handler) updateGamePrice(c *gin.Context) {
	gameId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid game ID"})
		return
	}

	price, err := strconv.Atoi(c.Param("price"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid game price"})
		return
	}

	if err := h.service.GameServiceInterface.PutGamePrice(c.Request.Context(), gameId, price); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update game price"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Game price updated successfully"})
}

func (h *Handler) getGameByName(c *gin.Context) {
	name := c.Param("name")
	log.Printf("Extracted game name: %s", name)
	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Game name is required"})
		return
	}

	game, err := h.service.GameServiceInterface.GetGameByName(c.Request.Context(), name)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Game not found"})
		return
	}

	c.JSON(http.StatusOK, game)
}

func (h *Handler) deleteGameByName(c *gin.Context) {
	name := c.Param("name")
	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Game name is required"})
		return
	}

	if err := h.service.GameServiceInterface.DeleteGameByName(c.Request.Context(), name); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete game"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Game deleted successfully"})
}
