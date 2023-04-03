package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/korasdor/go-ether-test/internal/models"
	"github.com/korasdor/go-ether-test/pkg/auth"
)

func (h *Handler) initAuthRoutes(api *gin.RouterGroup) {
	auth := api.Group("auth")
	{
		auth.POST("user/sign-up", h.userSignUp)
		auth.POST("user/sign-in", h.userSignIn)
		auth.POST("user/refresh", h.userRefreshToken)
	}
}

func (h *Handler) userSignUp(ctx *gin.Context) {
	var signUpData models.SignUpData

	if err := ctx.BindJSON(&signUpData); err != nil {
		newResponse(ctx, http.StatusBadRequest, models.ErrBadRequestFormat.Error())
		return
	}

	err := h.services.AuthorizationService.SignUp(ctx.Request.Context(), signUpData)
	if err != nil {
		newResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "good",
	})
}

func (h *Handler) userSignIn(ctx *gin.Context) {
	var signInData models.SignInData

	if err := ctx.BindJSON(&signInData); err != nil {
		newResponse(ctx, http.StatusBadRequest, models.ErrBadRequestFormat.Error())
		return
	}

	tokenBinding := &auth.TokenBinding{}
	tokenBinding, err := tokenBinding.Parse(ctx.Request)
	if err != nil {
		newResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.services.AuthorizationService.SignIn(ctx.Request.Context(), signInData, tokenBinding)
	if err != nil {
		newResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	cookie := &http.Cookie{
		Name:     "refreshToken",
		Value:    token.RefreshToken,
		MaxAge:   int(h.config.Auth.JWT.RefreshTokenTTL.Seconds()),
		HttpOnly: true,
	}

	http.SetCookie(ctx.Writer, cookie)

	ctx.JSON(http.StatusOK, token)
}

func (h *Handler) userRefreshToken(ctx *gin.Context) {
	tokenBinding := &auth.TokenBinding{}
	tokenBinding, err := tokenBinding.Parse(ctx.Request)
	if err != nil {
		newResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	cookie, err := ctx.Request.Cookie("refreshToken")
	if err != nil {
		newResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	token, err := h.services.AuthorizationService.RefreshTokens(cookie.Value, tokenBinding)
	if err != nil {
		newResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	cookie = &http.Cookie{
		Name:     "refreshToken",
		Value:    token.RefreshToken,
		MaxAge:   int(h.config.Auth.JWT.RefreshTokenTTL.Seconds()),
		HttpOnly: true,
	}

	http.SetCookie(ctx.Writer, cookie)

	ctx.JSON(http.StatusOK, token)
}
