package handler

import "github.com/gin-gonic/gin"

func (h *Handler) postUser(ctx *gin.Context) {
	h.signUp(ctx)
}
