package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) addGenre(c *gin.Context) {
	var request struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := h.service.GenreServiceInterface.AddGenre(c.Request.Context(), request.Name, request.Description); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add genre"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Genre added successfully"})
}

func (h *Handler) getGenreByName(c *gin.Context) {
	name := c.Param("name")
	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Genre name is required"})
		return
	}

	genre, err := h.service.GenreServiceInterface.GetGenreByName(c.Request.Context(), name)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Genre not found"})
		return
	}

	c.JSON(http.StatusOK, genre)
}

func (h *Handler) getAllGenres(c *gin.Context) {
	genres, err := h.service.GenreServiceInterface.GetAllGenres(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve genres"})
		return
	}

	c.JSON(http.StatusOK, genres)
}

func (h *Handler) deleteGenre(c *gin.Context) {
	genreId, err := strconv.Atoi(c.Param("genre_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid genre ID"})
		return
	}

	if err := h.service.GenreServiceInterface.DeleteGenre(c.Request.Context(), genreId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete genre"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Genre deleted successfully"})
}

func (h *Handler) getGenreByID(c *gin.Context) {
	genreId, err := strconv.Atoi(c.Param("genre_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid genre ID"})
		return
	}

	genre, err := h.service.GenreServiceInterface.GetGenreByID(c.Request.Context(), genreId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Genre not found"})
		return
	}

	c.JSON(http.StatusOK, genre)
}
