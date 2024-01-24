package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/leapkit/core/db"
	"github.com/leapkit/core/db/migrations"
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
		err := db.GenerateMigration(
			os.Args[2], // name of the migration

			// This is the path to the migrations folder
			migrations.UseMigrationFolder(filepath.Join("internal", "migrations")),
		)

		if err != nil {
			fmt.Println(err)

			return
		}
	default:
		fmt.Printf("Unknown generator `%v`.\n\n", os.Args[1])

		usage()
	}

}
