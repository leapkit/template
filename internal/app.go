// Package internal integrates the app and loads general settings.
package internal

import (
	"cmp"
	"embed"
	"net/http"
	"os"

	"go.leapkit.dev/core/db"
	"go.leapkit.dev/core/render"
	"go.leapkit.dev/core/server"
	"go.leapkit.dev/template/internal/home"
	"go.leapkit.dev/template/internal/system/assets"
)

var (
	//go:embed **/*.html **/*.html
	tmpls embed.FS

	// dbURL is the database connection string
	dbURL = cmp.Or(os.Getenv("DATABASE_URL"), "database.db?_timeout=5000&_sync=1")

	// DBFn is the database connection builder function
	// that will be used by the application based on the driver and
	// connection string.
	DBFn = db.ConnectionFn(dbURL, db.WithDriver("sqlite3"))

	// Server configuration variables loaded from the environment
	// or default values.
	host          = cmp.Or(os.Getenv("HOST"), "0.0.0.0")
	port          = cmp.Or(os.Getenv("PORT"), "3000")
	sessionSecret = cmp.Or(os.Getenv("SESSION_SECRET"), "d720c059-9664-4980-8169-1158e167ae57")
	sessionName   = cmp.Or(os.Getenv("SESSION_NAME"), "leapkit_session")
)

// New creates the http handler using the Leapkit server package
// and returns it with the address it is listening on.
func New() (http.Handler, string) {
	// Creating a new server instance with the host and port
	// variables read from the environment or default values.
	r := server.New(
		server.WithHost(host),
		server.WithPort(port),
		server.WithSession(sessionSecret, sessionName),

		// Mounting the assets folder in the /assets URL path
		server.WithAssets(assets.Files, "/internal/system/assets"),
	)

	// Using the render middleware to load HTML templates
	// as well as setting a default layout for the application.
	r.Use(render.Middleware(
		render.TemplateFS(tmpls, "internal"),
		render.WithDefaultLayout("system/layout.html"),
	))

	// Defining the routes in the application.
	r.HandleFunc("GET /{$}", home.Index)
	// Add more routes here...

	return r.Handler(), r.Addr()
}
