package main

import (
	"fmt"

	"github.com/leapkit/core/gloves"
	"github.com/leapkit/template/internal"

	// Load environment variables
	_ "github.com/leapkit/core/envload"
)

func main() {
	err := gloves.Start(
		"cmd/app/main.go",

		internal.GlovesOptions...,
	)

	if err != nil {
		fmt.Println(err)
	}
}
