package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/korasdor/go-ether-test/internal/config"
	"github.com/korasdor/go-ether-test/internal/services"
	"github.com/korasdor/go-ether-test/pkg/limiter"
)

type Handlers struct {
	services *services.Services
	config   *config.Config
}

func NewHandlers(services *services.Services, config *config.Config) *Handlers {
	return &Handlers{
		services: services,
		config:   config,
	}
}

func (h *Handlers) Init(cfg *config.Config) *gin.Engine {
	gin.SetMode(cfg.Gin.GinMode)

	router := gin.Default()

	router.Use(
		gin.Recovery(),
		gin.Logger(),
		limiter.Limit(cfg.Limiter.RPS, cfg.Limiter.Burst, cfg.Limiter.TTL),
		corsMiddleware,
	)

	router.GET("/ping", func(ctx *gin.Context) {
		newResponse(ctx, http.StatusForbidden, "error message")
	})

	api := router.Group("v1/api")
	{
		user := api.Group("users")
		{
			auth := user.Group("auth")
			{
				auth.POST("/sing-up", h.userSignUp)
				auth.POST("/sign-in", h.userSignIn)
				auth.POST("/refresh", h.userRefreshToken)
			}
		}
	}

	return router
}
