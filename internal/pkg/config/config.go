package config

import (
	"fmt"
	"log/slog"
	"os"
	"strconv"
	"strings"
)

var hasError bool

type Env struct {
	LogLevel  string
	LogFormat string
	Addr      string
}

func (e *Env) String() string {
	return fmt.Sprintf("%+v", *e)
}

type Config struct {
	// Logger is the slog logger to use
	Logger *slog.Logger

	Addr string `json:"addr"`
}

// ReadEnv reads the environment variables and returns an Env struct.
func ReadEnv() *Env {
	return &Env{
		LogLevel:  getEnv("LOG_LEVEL", "info"),
		LogFormat: getEnv("LOG_FORMAT", "logfmt"),
		Addr:      getEnv("ADDR", ":8080"),
	}
}

// New reads the configuration from the environment variables and returns a Config and Env struct.
func New(env *Env) *Config {
	cfg := &Config{
		Logger: parseLogger(env.LogLevel, env.LogFormat),
		Addr:   env.Addr,
	}

	if hasError {
		os.Exit(1)
	}

	return cfg
}

func parseLogger(level, format string) *slog.Logger {
	format = strings.ToLower(format)

	var handler slog.Handler

	switch format {
	case "json":
		handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: parseLogLevel(level)})
	case "logfmt", "text":
		handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: parseLogLevel(level)})
	default:
		configError("invalid log format", "log_format", format)
	}

	return slog.New(handler)
}

// parseLogLevel parses the log level from the environment variable and returns the corresponding slog.Level.
// If the level is invalid, it logs an error and returns slog.LevelInfo.
// Possible values are: debug, info, warn, error - default is info.
func parseLogLevel(level string) slog.Level {
	level = strings.ToLower(level)

	switch level {
	case "debug":
		return slog.LevelDebug
	case "info":
		return slog.LevelInfo
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		configError("invalid log level", "log_level", level)
		return slog.LevelInfo
	}
}

// parsePort parses the port from the environment variable.
func parsePort(port string) int {
	if port == "" {
		return 8080
	}

	if strings.HasPrefix(port, ":") {
		port = port[1:]
	}

	if p, err := strconv.Atoi(port); err == nil {
		return p
	} else {
		configError("invalid port", "port", port)
	}

	return 8080
}

// error logs an error message and sets the hasError flag to true, which will cause the program to exit with a non-zero exit code.
func configError(msg string, args ...any) {
	msg = "in config: " + msg
	slog.Error(msg, args...)

	hasError = true
}

// getEnv returns the value of a given environment variable, or the defined default value.
func getEnv(key, def string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}

	return def
}
