// config package contains the configuration of the application,
// such as the database URL and the port the application will listen on
// other environment variables could in the future live here so that
// these can be used from other higher level packages.
package config

import (
	"github.com/leapkit/core/envor"
	"github.com/leapkit/core/gloves"
	"github.com/paganotoni/tailo"
)

var (
	// Environment in which the application is running, this is useful
	// to determine the way we'll run the application, for example, if
	// we're running in production we might want to disable debug mode.
	Environment = envor.Get("GO_ENV", "development")

	// DatabaseURL to connect and interact with our database instance.
	DatabaseURL = envor.Get("DATABASE_URL", "postgres://postgres:postgres@localhost:5432/app?sslmode=disable")

	// Port in which the web application listens on.
	Port = envor.Get("PORT", "3000")

	// SessionSecret is the secret used to sign the session cookies.
	SessionSecret = envor.Get("SESSION_SECRET", "d720c059-9664-4980-8169-1158e167ae57")
	SessionName   = envor.Get("SESSION_NAME", "leapkit_session")

	// Options for tailo, this is used to setup the tailwindcss
	// watcher and builder.
	TailoOptions = []tailo.Option{
		tailo.UseInputPath("internal/assets/application.css"),
		tailo.UseOutputPath("internal/app/public/application.css"),
		tailo.UseConfigPath("tailwind.config.js"),
	}

	GlovesOptions = []gloves.Option{
		gloves.WithRunner(func() {
			// Run the tailo watcher so when changes are made to
			// the html code it rebuilds css.
			tailo.Watch(TailoOptions...)
		}),

		// Extensions to watch
		gloves.WatchExtension(".go", ".env", ".json", ".html", ".js"),
	}
)
