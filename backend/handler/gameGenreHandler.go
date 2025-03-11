package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) addGenreToGame(c *gin.Context) {
	gameId, err := strconv.Atoi(c.Param("game_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid game ID"})
		return
	}

	genreId, err := strconv.Atoi(c.Param("genre_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid genre ID"})
		return
	}

	// Add genre to game
	if err := h.service.GameGenreServiceInterface.AddGenreToGame(c.Request.Context(), gameId, genreId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add genre to game"})
		return
	}

	// Increment the genre count
	if err := h.service.GameGenreServiceInterface.IncrementGenreCount(c.Request.Context(), gameId, genreId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to increment genre count"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Genre added to game successfully"})
}

func (h *Handler) getGenresByGameID(c *gin.Context) {
	gameId, err := strconv.Atoi(c.Param("game_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid game ID"})
		return
	}

	genres, err := h.service.GameGenreServiceInterface.GetGenresByGameID(c.Request.Context(), gameId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve genres"})
		return
	}

	c.JSON(http.StatusOK, genres)
}

func (h *Handler) getGamesByGenreID(c *gin.Context) {
	genreId, err := strconv.Atoi(c.Param("genre_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid genre ID"})
		return
	}

	games, err := h.service.GameGenreServiceInterface.GetGamesByGenreID(c.Request.Context(), genreId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve games"})
		return
	}

	c.JSON(http.StatusOK, games)
}

func (h *Handler) getGamesByGenreName(c *gin.Context) {
	genreName := c.Param("genre_name")
	if genreName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Genre name is required"})
		return
	}

	games, err := h.service.GameGenreServiceInterface.GetGamesByGenreName(c.Request.Context(), genreName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve games"})
		return
	}

	c.JSON(http.StatusOK, games)
}

func (h *Handler) deleteGameGenre(c *gin.Context) {
	gameId, err := strconv.Atoi(c.Param("game_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid game ID"})
		return
	}

	genreId, err := strconv.Atoi(c.Param("genre_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid genre ID"})
		return
	}

	if err := h.service.GameGenreServiceInterface.DeleteGameGenre(c.Request.Context(), gameId, genreId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete game genre"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Game genre deleted successfully"})
}
