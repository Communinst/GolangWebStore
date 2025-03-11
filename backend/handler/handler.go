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

	wallets := apiRouter.Group("/wallets")
	{
		wallets.GET("/:user_id", h.getWalletByUserID)
		wallets.PUT("/:user_id", h.updateWalletBalance)
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

	genres := apiRouter.Group("/genres")
	{
		genres.POST("/", h.addGenre)
		genres.GET("/name/:name", h.getGenreByName)
		genres.GET("/:genre_id", h.getGenreByID)
		genres.GET("/", h.getAllGenres)
		genres.DELETE("/:genre_id", h.deleteGenre)
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
		usersCart := carts.Group("/:user_id")
		{
			usersCart.POST("/games/:game_id", h.addGameToCart)
			usersCart.GET("/", h.getCartByUserID)
			usersCart.DELETE("/games/:game_id", h.removeGameFromCart)
		}
	}

	ownerships := apiRouter.Group("/ownerships")
	{
		ownerships.POST("/:user_id/games/:game_id", h.addOwnership)
		ownerships.GET("/user/:user_id", h.getOwnershipsByUserID)
		ownerships.GET("/game/:game_id", h.getOwnershipsByGameID)
		ownerships.DELETE("/:ownership_id", h.deleteOwnership)
	}

	discounts := apiRouter.Group("/discounts")
	{
		discounts.POST("/game/:game_id", h.addDiscount)
		discounts.GET("/game/:game_id", h.getDiscountsByGameID)
		discounts.DELETE("/:discount_id", h.deleteDiscount)
	}

	reviews := apiRouter.Group("/reviews")
	{
		reviews.POST("/:user_id/games/:game_id", h.addReview)
		reviews.GET("/game/:game_id", h.getReviewsByGameID)
		reviews.GET("/user/:user_id", h.getReviewsByUserID)
		reviews.DELETE("/:review_id", h.deleteReview)
	}

	return router
}
