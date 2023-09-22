package web

import (
	"embed"
	"net/http"
	"path"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/gorilla/sessions"

	"github.com/leapkit/core/mdfs"
	"github.com/leapkit/core/render"
	"github.com/leapkit/core/server"
	"github.com/leapkit/core/session"
	"github.com/leapkit/template/internal/config"
	"github.com/leapkit/template/internal/web/helpers"
	"github.com/leapkit/template/internal/web/home"
	"github.com/leapkit/template/internal/web/public"
)

var (
	//go:embed **/*.html
	tmpls     embed.FS
	templates = mdfs.New(
		tmpls,
		path.Join("internal", "web"),
		config.Environment,
	)

	//sessions store
	store = sessions.NewCookieStore([]byte(config.SessionSecret))

	// the rendering engine for the application, this
	// is used to render each of the HTML responses
	// for the application.
	renderer = render.NewEngine(
		templates,
		render.WithHelpers(helpers.All),
	)
)

// routes sets up all the routes for the application.
func Routes(r *server.Instance) {
	// Base middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)

	// Custom General Middleware
	r.Use(session.InCtx(store, config.SessionName))
	r.Use(render.InCtx(renderer))
	r.Use(session.AddHelpers)

	r.Route("/", func(rx chi.Router) {
		rx.Get("/", home.Index)
	})

	// Public files that include anything thats on the
	// public folder. This is useful for files and assets.
	r.Handle("/public/*", http.StripPrefix("/public/", http.FileServer(http.FS(public.Folder))))
}
