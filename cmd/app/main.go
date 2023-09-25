package main

import (
	"fmt"

	"github.com/leapkit/core/server"
	"github.com/leapkit/template/internal/web"
)

func main() {
	s := server.New(
		"LeapKit",

		// Add any other options here.
		// Routes are defined in internal/web/routes.go
		// we pass these to the newly created server
		// as an option.
		server.WithRoutesFn(web.Routes),
	)

	if err := s.Start(); err != nil {
		fmt.Println(err)
	}
}
