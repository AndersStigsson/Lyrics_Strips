package main

import (
	"context"
	"errors"
	"log"
	"log/slog"
	"net/http"

	spotifyauth "github.com/zmb3/spotify/v2/auth"

	"golang.org/x/oauth2/clientcredentials"

	"github.com/zmb3/spotify/v2"

	iSpotify "skafteresort.se/lyrics_strips/internal/spotify"
	"skafteresort.se/lyrics_strips/internal/web"
)

var ServiceVersion string

type Server struct {
	logger     *slog.Logger
	config     ServerConfig
	httpServer *http.Server

	spotifyService *iSpotify.Service
}

func NewServer(config ServerConfig, logger *slog.Logger) *Server {
	return &Server{
		config: config,
		logger: logger,
	}
}

func (s *Server) Start(ctx context.Context) {
	var err error

	s.logger.Info("Starting server", "version", ServiceVersion)

	config := &clientcredentials.Config{
		ClientID:     s.config.spotifyClientId,
		ClientSecret: s.config.spotifyClientSecret,
		TokenURL:     spotifyauth.TokenURL,
	}
	token, err := config.Token(ctx)
	if err != nil {
		log.Fatalf("couldn't get token: %v", err)
	}

	httpClient := spotifyauth.New().Client(ctx, token)
	client := spotify.New(httpClient)
	s.spotifyService = iSpotify.NewSpotifyService(s.logger, client)

	s.serveHTTP()

	<-ctx.Done()
	s.Shutdown(ctx)
}

func (s *Server) serveHTTP() {
	handler := web.NewServer(
		s.logger,
		s.config.httpCorsAllowedOrigin,
		s.config.jwtSecret,
		s.spotifyService,
	)
	s.httpServer = &http.Server{
		Addr:         s.config.httpEndpointPort,
		Handler:      handler,
		ReadTimeout:  s.config.httpReadTimeout,
		WriteTimeout: s.config.httpWriteTimeout,
	}

	s.logger.Info("Starting webserver", slog.String("Listening", s.httpServer.Addr))

	go func() {
		if err := s.httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			s.logger.Error("Webserver error", slog.String("error", err.Error()))
		}
	}()
}

func (s *Server) Shutdown(ctx context.Context) {

	// Shutdown context/withTimeout
	if err := s.httpServer.Shutdown(ctx); err != nil {
		s.logger.Error("HTTP Server shutdown", "error", err)
		return
	}

	s.logger.Info("Graceful shutdown completed")
}
