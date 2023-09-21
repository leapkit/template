package web

import (
	"context"
	"net/http"
)

// Allows to specify options to the server.
type Option func(*server)

// Allows to specify a key/value that should be set on each
// request context. This is useful for services that could be
// used by the handlers.
func WithRequestCtxVal(key string, value interface{}) Option {
	return func(s *server) {
		s.Use(func(next http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				r = r.WithContext(context.WithValue(r.Context(), key, value))
				next.ServeHTTP(w, r)
			})
		})
	}
}
