package postgres

import (
	"embed"

	"github.com/leapkit/core/db"
	"github.com/leapkit/template/config"

	_ "github.com/lib/pq"
)

var (
	//go:embed migrations/*.sql
	Migrations embed.FS

	//Connection is the database connection builder function
	//that will be used by the application based on the driver and
	//connection string.
	Connection = db.ConnectionFn(config.DatabaseURL)
)
