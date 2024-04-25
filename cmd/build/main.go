package main

import (
	"os"
	"os/exec"

	"github.com/leapkit/template/internal"
	"github.com/paganotoni/tailo"

	// Load environment variables
	_ "github.com/leapkit/core/envload"
	// sqlite3 driver
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	tailo.Build(internal.TailoOptions...)

	cmd := exec.Command("go", "build")
	cmd.Args = append(
		cmd.Args,

		`--ldflags`, `-linkmode=external -extldflags="-static"`,
		`-tags`, `osusergo,netgo,musl`,
		`-buildvcs=false`,
		"-o", "bin/app",
		"cmd/app/main.go",
	)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}
