package handler

import (
	"net/http"

	"github.com/Communinst/GolangWebStore/backend/service"
	"github.com/gin-gonic/gin"
)

const (
	userRole = 0
)

type Handler struct {
	service *service.Service
}

func New(service *service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) InitRoutes(middleware ...gin.HandlerFunc) *gin.Engine {
	router := gin.Default()
	router.Use(middleware...)
	apiRouter := router.Group("/api")
	apiRouter.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, "welcome page")
	})

	auth := apiRouter.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}
	games := apiRouter.Group("/games")
	{
		games.POST("/", h.createGame)
		games.GET("/", h.getAllGames)
		ids := games.Group("/:id")
		{
			ids.GET("/", h.getGame)
			ids.DELETE("/", h.deleteGame)
			ids.PUT("/:price", h.updateGamePrice)
		}

		names := games.Group("/name")
		{
			names.GET("/:name", h.getGameByName)
			names.DELETE("/:name", h.deleteGameByName)
		}

	}
	companies := apiRouter.Group("/companies")
	{
		companies.POST("/", h.createCompany)
		companies.GET("/", h.getAllCompanies)
		ids := games.Group("/:id")
		{
			ids.GET("/:id", h.getCompany)
			ids.DELETE("/:id", h.deleteCompany)
		}
		names := games.Group("/name")
		{
			names.GET("/:name", h.getCompanyByName)
			names.DELETE("/:name", h.deleteCompanyByName)
		}

	}

	carts := apiRouter.Group("/carts")
	{
		carts.POST("/:user_id/games/:game_id", h.addGameToCart)
		carts.GET("/:user_id", h.getCartByUserID)
		carts.DELETE("/:user_id/games/:game_id", h.removeGameFromCart)
	}
	return router
}
