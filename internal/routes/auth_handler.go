package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/korasdor/go-ether-test/internal/models"
)

func (h *Handlers) userSignUp(ctx *gin.Context) {
	var signUpData models.SignUpData

	if err := ctx.BindJSON(&signUpData); err != nil {
		newResponse(ctx, http.StatusBadRequest, models.ErrBadRequestFormat.Error())
		return
	}

	err := h.services.AuthorizationService.SingUp(ctx.Request.Context(), signUpData)
	if err != nil {
		newResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "good",
	})
}

func (h *Handlers) userSignIn(ctx *gin.Context) {
	var signInData models.SignInData

	if err := ctx.BindJSON(&signInData); err != nil {
		newResponse(ctx, http.StatusBadRequest, models.ErrBadRequestFormat.Error())
		return
	}

	token, err := h.services.AuthorizationService.SingIn(ctx.Request.Context(), signInData)
	if err != nil {
		newResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, token)
}

func (h *Handlers) userRefreshToken(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "good",
	})
}
