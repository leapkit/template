package internal

import (
	"embed"
	"path"

	"github.com/leapkit/core/mdfs"
	"github.com/leapkit/template/internal/app/config"
)

var (
	//go:embed **/*.html
	tmpls embed.FS

	// Templates is a MDFS with the templates so we can use them in the application
	// and read them from disk in development.
	Templates = mdfs.New(
		tmpls,

		path.Join("internal"),
		config.Environment,
	)
)
