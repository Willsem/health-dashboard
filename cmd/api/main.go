package main

import (
	"context"

	"github.com/Willsem/health-dashboard/internal/app"
	"github.com/Willsem/health-dashboard/internal/app/build"
	"github.com/Willsem/health-dashboard/internal/app/config"
	"github.com/Willsem/health-dashboard/internal/app/startup"
	"github.com/Willsem/health-dashboard/internal/health"
	"github.com/Willsem/health-dashboard/internal/http/handlers"
	"github.com/Willsem/health-dashboard/internal/http/router"
	"github.com/Willsem/health-dashboard/internal/http/server"
	"github.com/Willsem/health-dashboard/internal/logger"
)

const appName = "health-dashboard"

// @title       Golang API Template
// @version     1.0
// @description golang api template

func main() {
	cfg, err := config.New()
	if err != nil {
		startup.NewFallbackLogger(appName).WithError(err).Fatal("failed to parse configuration")
	}

	logger := startup.NewLogger(appName, cfg.Log)

	if err := run(cfg, logger); err != nil {
		logger.WithError(err).Fatal("error during the running app")
	}
}

func run(cfg *config.Config, logger logger.Logger) error {
	logger.Infof(
		"%s has version %s built from %s on %s by %s",
		appName, build.Version, build.VersionCommit, build.BuildDate, build.GoVersion,
	)

	logger.With("config", cfg).Info("application is starting with config")

	probe := health.NewProbe()
	readyStatus := health.NewReadyStatus()

	router := router.NewHTTPRouter(
		logger,
		handlers.NewMetricsHandler(),
		handlers.NewSwaggerHandler(),
		handlers.NewLivenessHandler(probe),
		handlers.NewReadinessHandler(readyStatus),
	)

	return app.New(
		cfg.App, readyStatus, logger,
		server.NewHTTPServer(router, cfg.Server, logger),
	).Run(context.Background())
}
