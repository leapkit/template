package internal

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/leapkit/core/render"
	"github.com/leapkit/core/server"
	"github.com/leapkit/core/session"

	"github.com/leapkit/template/internal/home"
	"github.com/leapkit/template/public"
)

// AddRoutes mounts the routes for the application,
// it assumes that the base services have been injected
// in the creation of the server instance.
func AddRoutes(r *server.Instance) error {
	// Base middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)

	// LeapKit Middleware
	r.Use(session.Middleware(SessionSecret, SessionName))
	r.Use(render.Middleware(
		templates,

		// options
		render.WithDefaultLayout("layout.html"),
	))

	r.Route("/", func(rx chi.Router) {
		rx.Get("/", home.Index)
	})

	// Public files that include anything thats on the
	// public folder. This is useful for files and assets.
	r.Folder("/public/", public.Files)

	return nil
}
