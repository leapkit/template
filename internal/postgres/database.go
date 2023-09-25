package postgres

import (
	"github.com/leapkit/core/db"
	"github.com/leapkit/template/internal/config"

	_ "github.com/lib/pq"
)

// Connection is the database connection builder function
// that will be used by the application based on the driver and
// connection string.
var Connection = db.ConnectionFn(config.DatabaseURL)
