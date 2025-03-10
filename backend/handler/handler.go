package handler

import (
	"net/http"

	"github.com/Communinst/GolangWebStore/backend/service"
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

func (h *Handler) InitRoutes(middleware ...gin.HandlerFunc) error {
	router := gin.Default()
	router.Use(middleware...)
	apiRouter := router.Group("/api")
	apiRouter.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, "welcome page")
	})

	// auth := apiRouter.Group("/")
	// {
	// 	auth.POST("/sign-up", h.signUp)
	// 	auth.POST("/sign-in", h.signIn)
	// 	auth.GET("/roles", h.getAllRoles)
	// 	auth.POST("/sign-up-privileged", h.userIdentity, h.signUpPrivileged)
	// }

	return nil
}
