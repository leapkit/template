package app

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/leapkit/core/render"
	"github.com/leapkit/core/server"
	"github.com/leapkit/core/session"

	"github.com/leapkit/template/internal"
	"github.com/leapkit/template/internal/app/config"
	"github.com/leapkit/template/internal/app/helpers"
	"github.com/leapkit/template/internal/app/public"
	"github.com/leapkit/template/internal/home"
)

var (
	// the rendering engine for the application, this
	// is used to render each of the HTML responses
	// for the application.
	renderer = render.NewEngine(
		internal.Templates,
		render.WithHelpers(helpers.All),
	)
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
	r.Use(session.Middleware(config.SessionSecret, config.SessionName))
	r.Use(render.Middleware(renderer))

	r.Route("/", func(rx chi.Router) {
		rx.Get("/", home.Index)
	})

	// Public files that include anything thats on the
	// public folder. This is useful for files and assets.
	r.Handle("/public/*", http.StripPrefix("/public/", http.FileServer(http.FS(public.Folder))))

	return nil
}
