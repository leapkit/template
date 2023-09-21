package main

import (
	"fmt"
	"os"

	"github.com/leapkit/core/db"
)

func main() {
	usage := func() {
		fmt.Println("Usage: gen <generator_name>")
		fmt.Println("Available generators:")
		fmt.Println("  -  migration")
	}

	if len(os.Args) < 3 {
		usage()

		return
	}

	switch os.Args[1] {
	case "migration":
		err := db.GenerateMigration(os.Args[2])
		if err != nil {
			fmt.Println(err)

			return
		}
	default:
		fmt.Printf("Unknown generator `%v`.\n\n", os.Args[1])

		usage()
	}

}
