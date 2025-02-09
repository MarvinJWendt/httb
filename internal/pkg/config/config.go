package config

import (
	"fmt"
	"github.com/caarlos0/env/v11"
	"log/slog"
	"os"
	"strings"
	"time"
)

// Config holds application configuration.
type Config struct {
	LogLevel  string        `env:"LOG_LEVEL" envDefault:"info"`
	LogFormat string        `env:"LOG_FORMAT" envDefault:"logfmt"`
	Addr      string        `env:"ADDR" envDefault:"0.0.0.0:8080"`
	Timeout   time.Duration `env:"TIMEOUT" envDefault:"2m"`
}

// LoadConfig reads the environment variables and returns a Config struct.
func LoadConfig() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, fmt.Errorf("failed to parse environment variables: %w", err)
	}
	return cfg, nil
}

// SetupLogger initializes the slog.Logger based on the config.
func (cfg *Config) SetupLogger() (*slog.Logger, error) {
	format := strings.ToLower(cfg.LogFormat)
	level, err := parseLogLevel(cfg.LogLevel)
	if err != nil {
		return nil, err
	}

	var handler slog.Handler

	switch format {
	case "json":
		handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: level})
	case "logfmt", "text":
		handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: level})
	default:
		return nil, fmt.Errorf("invalid log format: %s", format)
	}

	return slog.New(handler), nil
}

// parseLogLevel converts a string to a slog.Level, returning an error if invalid.
func parseLogLevel(level string) (slog.Level, error) {
	switch strings.ToLower(level) {
	case "debug":
		return slog.LevelDebug, nil
	case "info":
		return slog.LevelInfo, nil
	case "warn":
		return slog.LevelWarn, nil
	case "error":
		return slog.LevelError, nil
	default:
		return slog.LevelInfo, fmt.Errorf("invalid log level: %s", level)
	}
}
