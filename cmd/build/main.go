package main

import (
	"os"
	"os/exec"

	"github.com/paganotoni/tailo"
)

func main() {
	tailo.Build()

	cmd := exec.Command("go", "build", `--ldflags`, `-linkmode=external -extldflags="-static"`, `-tags`, `osusergo,netgo,musl`, `-buildvcs=false`, "-o", "bin/app", "cmd/app/main.go")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}
