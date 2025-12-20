package web

import (
	"log/slog"
	"net/http"
	"time"
)

func loggingMiddleware(logger *slog.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		defer func(start time.Time) {
			if req.Method == http.MethodOptions {
				next.ServeHTTP(w, req)

				return
			}

			attributes := []slog.Attr{}

			attributes = []slog.Attr{
				slog.Attr{
					Key: "request",
					Value: slog.GroupValue(
						slog.String("method", req.Method),
						slog.String("path", req.URL.Path),
						slog.String("ip", req.RemoteAddr),
						slog.Duration("time", time.Since(start)),
					),
				},
			}
			logger.LogAttrs(req.Context(), slog.LevelInfo, "http", attributes...)
		}(time.Now())

		next.ServeHTTP(w, req)
	})
}
