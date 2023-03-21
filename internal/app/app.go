package app

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/korasdor/go-ether-test/internal/config"
	"github.com/korasdor/go-ether-test/internal/routes"
	"github.com/korasdor/go-ether-test/internal/server"
	"github.com/korasdor/go-ether-test/internal/services"
	"github.com/korasdor/go-ether-test/pkg/cache"
	"github.com/korasdor/go-ether-test/pkg/logger"
)

func Run() {
	logger := logger.NewLogrusLogger()

	cfg, err := config.NewConfig()
	if err != nil {
		logger.Errorf("Error occurred while reading env file, might fallback to OS env config %v", err)
	}

	services := services.NewServices(
		&services.Deps{
			Cache: cache.NewRedisCache(
				cfg.Reddis.Addr,
				cfg.Reddis.Password,
			),
			// Cache: cache.NewMemoryCache(),
		},
	)
	handlers := routes.NewHandlers(services)

	srv := server.NewServer(cfg, handlers.Init(cfg))

	go func() {
		if err := srv.Run(); !errors.Is(err, http.ErrServerClosed) {
			logger.Errorf("error occurred while running http server: %s\n", err.Error())
		}
	}()

	logger.Info("Server started")

	// Graceful Shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	if err := srv.Stop(ctx); err != nil {
		logger.Errorf("failed to stop server: %v", err)
	}
}
