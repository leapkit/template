package main

import (
	"fmt"
	"os"

	"github.com/leapkit/core/server"
	"github.com/leapkit/template/internal"
)

func main() {
	s := server.New("LeapKit")

	if err := internal.AddRoutes(s); err != nil {
		os.Exit(1)
	}

	if err := s.Start(); err != nil {
		fmt.Println(err)
	}
}
