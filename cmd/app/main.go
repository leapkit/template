package main

import (
	"log/slog"
	"net/http"

	"go.leapkit.dev/template/internal"

	_ "github.com/mattn/go-sqlite3"       // sqlite3 driver
	_ "go.leapkit.dev/core/tools/envload" // Load environment variables
)

// main executable for the application server, it calls the internal package to
// create the HTTP server and starts listening for incoming requests
// on the specified address.
func main() {
	server, addr := internal.New()
	slog.Info("server started", "addr", addr)

	err := http.ListenAndServe(addr, server)
	slog.Error("server stopped", "error", err)
}
