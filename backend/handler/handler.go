package handler

import (
	"net/http"

	authToken "github.com/Communinst/GolangWebStore/backend/JSONWebTokens"
	"github.com/Communinst/GolangWebStore/backend/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
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

	// Configure CORS to allow all origins for testing
	router.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// open for eo
	welcome := router.Group("/welcome")
	welcome.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, "welcome page")
	})

	auth := welcome.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	apiRouter := router.Group("/api", authToken.JwtAuthMiddleware())

	wallets := apiRouter.Group("/wallets")
	{
		wallets.GET("/:user_id", h.getWalletByUserID)
		wallets.PUT("/:user_id", h.updateWalletBalance)
	}

	games := apiRouter.Group("/games")
	{
		//games.POST("/", h.createGame)
		games.GET("/", h.getAllGames)
		idsGames := games.Group("/:id")
		{
			idsGames.GET("/", h.getGame)
			//idsGames.DELETE("/", h.deleteGame)
			//idsGames.PUT("/:price", h.updateGamePrice)
		}

		namesGames := games.Group("/name")
		{
			namesGames.GET("/:name", h.getGameByName)
			//namesGames.DELETE("/:name", h.deleteGameByName)
		}
	}

	genres := apiRouter.Group("/genres")
	{
		//genres.POST("/", h.addGenre)
		genres.GET("/name/:name", h.getGenreByName)
		genres.GET("/:genre_id", h.getGenreByID)
		genres.GET("/", h.getAllGenres)
		//genres.DELETE("/:genre_id", h.deleteGenre)
	}

	gameGenres := apiRouter.Group("/games_genres")
	{

		gameGenres.GET("/genres/:genre_id", h.getGamesByGenreID)
		gameGenres.GET("/genres/name/:genre_name", h.getGamesByGenreName)

		idsGameGenres := gameGenres.Group("/:game_id")
		{
			genres := idsGameGenres.Group("/genres")
			{
				genres.POST("/:genre_id", h.addGenreToGame)
				genres.GET("/", h.getGenresByGameID)
				genres.DELETE("/:genre_id", h.deleteGameGenre)
			}
		}
	}

	companies := apiRouter.Group("/companies")
	{
		//companies.POST("/", h.createCompany)
		companies.GET("/", h.getAllCompanies)
		ids := companies.Group("/:id")
		{
			ids.GET("/:id", h.getCompany)
			//ids.DELETE("/:id", h.deleteCompany)
		}
		namesCompanies := companies.Group("/name")
		{
			namesCompanies.GET("/:name", h.getCompanyByName)
			//namesCompanies.DELETE("/:name", h.deleteCompanyByName)
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
		//ownerships.DELETE("/:ownership_id", h.deleteOwnership)
	}

	discounts := apiRouter.Group("/discounts")
	{
		discounts.GET("/game/:game_id", h.getDiscountsByGameID)
	}

	reviews := apiRouter.Group("/reviews")
	{
		reviews.POST("/:user_id/games/:game_id", h.addReview)
		reviews.GET("/game/:game_id", h.getReviewsByGameID)
		reviews.GET("/user/:user_id", h.getReviewsByUserID)
	}

	// ADMIN
	admin := router.Group("/admin")
	{
		wallets := admin.Group("/wallets")
		{
			wallets.GET("/:user_id", h.getWalletByUserID)
			wallets.PUT("/:user_id", h.updateWalletBalance)
		}

		games := admin.Group("/games")
		{
			games.POST("/", h.createGame)
			games.GET("/", h.getAllGames)
			idsGames := games.Group("/:id")
			{
				idsGames.GET("/", h.getGame)
				idsGames.DELETE("/", h.deleteGame)
				idsGames.PUT("/:price", h.updateGamePrice)
			}

			namesGames := games.Group("/name")
			{
				namesGames.GET("/:name", h.getGameByName)
				namesGames.DELETE("/:name", h.deleteGameByName)
			}
		}

		genres := admin.Group("/genres")
		{
			genres.POST("/", h.addGenre)
			genres.GET("/name/:name", h.getGenreByName)
			genres.GET("/:genre_id", h.getGenreByID)
			genres.GET("/", h.getAllGenres)
			genres.DELETE("/:genre_id", h.deleteGenre)
		}

		gameGenres := admin.Group("/games_genres")
		{

			gameGenres.GET("/genres/:genre_id", h.getGamesByGenreID)
			gameGenres.GET("/genres/name/:genre_name", h.getGamesByGenreName)

			idsGameGenres := gameGenres.Group("/:game_id")
			{
				genres := idsGameGenres.Group("/genres")
				{
					genres.POST("/:genre_id", h.addGenreToGame)
					genres.GET("/", h.getGenresByGameID)
					genres.DELETE("/:genre_id", h.deleteGameGenre)
				}
			}
		}

		companies := admin.Group("/companies")
		{
			companies.POST("/", h.createCompany)
			companies.GET("/", h.getAllCompanies)
			ids := companies.Group("/:id")
			{
				ids.GET("/:id", h.getCompany)
				ids.DELETE("/:id", h.deleteCompany)
			}
			namesCompanies := companies.Group("/name")
			{
				namesCompanies.GET("/:name", h.getCompanyByName)
				namesCompanies.DELETE("/:name", h.deleteCompanyByName)
			}

		}

		carts := admin.Group("/carts")
		{
			usersCart := carts.Group("/:user_id")
			{
				usersCart.POST("/games/:game_id", h.addGameToCart)
				usersCart.GET("/", h.getCartByUserID)
				usersCart.DELETE("/games/:game_id", h.removeGameFromCart)
			}
		}

		ownerships := admin.Group("/ownerships")
		{
			ownerships.POST("/:user_id/games/:game_id", h.addOwnership)
			ownerships.GET("/user/:user_id", h.getOwnershipsByUserID)
			ownerships.GET("/game/:game_id", h.getOwnershipsByGameID)
			ownerships.DELETE("/:ownership_id", h.deleteOwnership)
		}

		discounts := admin.Group("/discounts")
		{
			discounts.POST("/game/:game_id", h.addDiscount)
			discounts.GET("/game/:game_id", h.getDiscountsByGameID)
			discounts.DELETE("/:discount_id", h.deleteDiscount)
		}

		reviews := admin.Group("/reviews")
		{
			reviews.POST("/:user_id/games/:game_id", h.addReview)
			reviews.GET("/game/:game_id", h.getReviewsByGameID)
			reviews.GET("/user/:user_id", h.getReviewsByUserID)
			reviews.DELETE("/:review_id", h.deleteReview)
		}
		dumps := admin.Group("/dumps")
		{
			dumps.POST("/create", h.createDump)
			dumps.POST("/restore", h.restoreDump)
			dumps.GET("/", h.getAllDumps)
		}
		users := admin.Group("/users")
		{
			users.POST("/create", h.postUser)
			users.GET("/", h.getAllUsers)
			users.DELETE("/:id", h.deleteUser)
			users.PUT("/:id/role/:role_id", h.updateUserRole)
		}
	}

	return router
}
