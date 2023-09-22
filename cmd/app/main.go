package main

import (
	"fmt"

	"github.com/leapkit/core/server"
	"github.com/leapkit/template/internal/web"
)

func main() {
	s := server.New(
		"LeapKit",

		// Routes are defined in internal/web/routes.go
		// we pass these to the newly created server
		// as an option.
		server.WithRoutesFn(web.Routes),

		// Add any other options here.
	)

	if err := s.Start(); err != nil {
		fmt.Println(err)
	}
}
