package config

import (
	"github.com/kelseyhightower/envconfig"

	"github.com/Willsem/health-dashboard/internal/app"
	"github.com/Willsem/health-dashboard/internal/app/startup"
	"github.com/Willsem/health-dashboard/internal/http/server"
)

type Config struct {
	Log    startup.LogConfig `envconfig:"LOG"`
	App    app.Config        `envconfig:"APP"`
	Server server.Config     `envconfig:"SERVER"`
}

func New() (*Config, error) {
	var config Config

	err := envconfig.Process("", &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
