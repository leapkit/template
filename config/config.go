// config package contains the configuration of the application,
// such as the database URL and the port the application will listen on
// other environment variables could in the future live here so that
// these can be used from other higher level packages.
package config

import (
	"sync"
)

var (
	// oncer for the config loading
	loadOnce sync.Once

	// DatabaseURL to connect and interact with our database instance.
	DatabaseURL = envOr("DATABASE_URL", "postgres://postgres:postgres@localhost:5432/app?sslmode=disable")

	// Environment in which the application is running, this is useful
	// to determine the way we'll run the application, for example, if
	// we're running in production we might want to disable debug mode.
	Environment = envOr("GO_ENV", "development")

	// Port in which the web application listens on.
	Port = envOr("PORT", "3000")

	GlovesExtensionsToWatch = []string{".go", ".env", ".json", ".html", ".js"}
	GlovesExcludePaths      = []string{"internal/jumpkit"}

	// SessionSecret is the secret used to sign the session cookies.
	SessionSecret = envOr("SESSION_SECRET", "d720c059-9664-4980-8169-1158e167ae57")
	SessionName   = envOr("SESSION_NAME", "jumpkit_session")
)
