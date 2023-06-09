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
	"github.com/korasdor/go-ether-test/internal/repository"
	"github.com/korasdor/go-ether-test/internal/routes"
	"github.com/korasdor/go-ether-test/internal/server"
	"github.com/korasdor/go-ether-test/internal/services"
	"github.com/korasdor/go-ether-test/pkg/auth"
	"github.com/korasdor/go-ether-test/pkg/cache"
	"github.com/korasdor/go-ether-test/pkg/database/mongodb"
	"github.com/korasdor/go-ether-test/pkg/hash"
	"github.com/korasdor/go-ether-test/pkg/logger"
)

func Run() {

	cfg, err := config.NewConfig()
	if err != nil {
		logger.Error(err)
		return
	}

	mongoClient, err := mongodb.NewClient(cfg.Mongo.URI, cfg.Mongo.User, cfg.Mongo.Password)
	if err != nil {
		logger.Error(err)
		return
	}

	db := mongoClient.Database(cfg.Mongo.Name)
	repos := repository.NewRepositories(db)
	cache := cache.NewRedisCache(cfg.Reddis.Addr, cfg.Reddis.Password)
	hasher := hash.NewSHA1Hasher(cfg.Auth.PasswordSalt)

	tokenManager, err := auth.NewManager(cfg.Auth.JWT.SigningKey)
	if err != nil {
		logger.Error(err)
		return
	}

	services := services.NewServices(
		&services.Deps{
			Repos:  repos,
			Cache:  cache,
			Hasher: hasher,
			// Cache: cache.NewMemoryCache(),

			TokenManager:    tokenManager,
			AccessTokenTTL:  cfg.Auth.JWT.AccessTokenTTL,
			RefreshTokenTTL: cfg.Auth.JWT.RefreshTokenTTL,
		},
	)
	handlers := routes.NewHandler(services, cfg, tokenManager)
	srv := server.NewServer(cfg, handlers.Init())
	go func() {
		if err := srv.Run(); !errors.Is(err, http.ErrServerClosed) {
			logger.Errorf("error occurred while running http server: %s\n", err.Error())
		}
	}()

	// run pprof
	// go func() {
	// 	if err := srv.RunPprof(); err != nil {
	// 		logger.Printf("error occurred while running pprof http server: %s\n", err.Error())
	// 	}
	// }()

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
