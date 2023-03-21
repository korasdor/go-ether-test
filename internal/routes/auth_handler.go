package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handlers) userSignUp(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "good",
	})
}

func (h *Handlers) userSignIn(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "good",
	})
}

func (h *Handlers) userRefreshToken(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "good",
	})
}
