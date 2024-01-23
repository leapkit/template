// config package contains the configuration of the application,
// such as the database URL and the port the application will listen on
// other environment variables could in the future live here so that
// these can be used from other higher level packages.
package internal

import (
	"github.com/leapkit/core/assets"
	"github.com/leapkit/core/gloves"
	"github.com/paganotoni/tailo"
)

var (
	AssetsFolder = "./internal/assets"
	PublicFolder = "./public"

	// TailoOptions allow to define how to compile
	// the tailwind css files, which is the input and
	// what will be the output.
	TailoOptions = []tailo.Option{
		tailo.UseInputPath("internal/assets/application.css"),
		tailo.UseOutputPath("public/application.css"),
		tailo.UseConfigPath("tailwind.config.js"),
	}

	// GlovesOptions are the options that will be used by the gloves
	// tool to hot reload the application.
	GlovesOptions = []gloves.Option{
		// Run the tailo watcher so when changes are made to
		// the html code it rebuilds css.
		gloves.WithRunner(tailo.WatcherFn(TailoOptions...)),
		gloves.WithRunner(assets.Watcher(AssetsFolder, PublicFolder)),
		gloves.WatchExtension(".go", ".css", ".js"),
	}
)
