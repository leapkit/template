package internal

import (
	"embed"

	"github.com/leapkit/core/mdfs"
	"github.com/leapkit/core/render"
)

var (
	//go:embed **/*.html **/*.html *.html
	tmpls embed.FS

	// templates is a MDFS with the templates so we can use them in the application
	// and read them from disk in development.
	templates = mdfs.New(
		tmpls,
		"internal",
		Environment,
	)

	// the rendering engine for the application, this
	// is used to render each of the HTML responses
	// for the application.
	renderer = render.NewEngine(templates)
)
