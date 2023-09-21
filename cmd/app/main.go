package main

import (
	"fmt"

	"github.com/leapkit/template/internal/web"
)

func main() {
	// conn, err := postgres.Connection()
	// if err != nil {
	// 	panic(err)
	// }

	server := web.NewServer(
		"LeapKit",

		// Services we want to be set in the context
	)

	err := server.Start()
	if err != nil {
		fmt.Println(err)
	}
}
