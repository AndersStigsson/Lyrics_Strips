package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type ServerConfig struct {
	httpEndpointPort      string
	httpCorsAllowedOrigin []string
	httpReadTimeout       time.Duration
	httpWriteTimeout      time.Duration

	spotifyClientId     string
	spotifyClientSecret string
	spotifyDc           string

	debug bool

	jwtSecret string
}

const httpDefaultTimeout time.Duration = 20 * time.Second

func parseBoolWithDefault(env string, d bool) bool {
	debug, err := strconv.ParseBool(os.Getenv(env))
	if err != nil {
		return d
	}

	return debug
}

func NewConfigFromEnv() (ServerConfig, error) {
	debug, err := strconv.ParseBool(os.Getenv("DEBUG"))
	if err != nil {
		return ServerConfig{}, fmt.Errorf("unable to parse environment variable DEBUG: %w", err)
	}

	corsAllowedOriginsString := os.Getenv("HTTP_CORS_ALLOWED_ORIGINS")
	corsAllowedOrigins := strings.Split(corsAllowedOriginsString, ",")

	config := ServerConfig{
		spotifyClientId:     os.Getenv("SPOTIFY_CLIENT_ID"),
		spotifyClientSecret: os.Getenv("SPOTIFY_CLIENT_SECRET"),
		spotifyDc:           os.Getenv("SPOTIFY_DC"),

		httpEndpointPort:      os.Getenv("HTTP_ENDPOINT_PORT"),
		httpCorsAllowedOrigin: corsAllowedOrigins,
		httpReadTimeout:       httpDefaultTimeout,
		httpWriteTimeout:      httpDefaultTimeout,

		debug: debug,

		jwtSecret: os.Getenv("JWT_SECRET"),
	}

	return config, nil
}
