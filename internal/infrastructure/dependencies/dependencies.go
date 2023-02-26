package dependencies

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"

	"github.com/marlonbarreto-git/rest-api-crud-go/internal/infrastructure/configurations"
)

type Container struct {
	Database *sql.DB
}

func StartDependencies() *Container {
	config := configurations.GetConfiguration()

	database, err := sql.Open("sqlite3", config.GetDBConfig().ConnectionPath)
	if err != nil || database.Ping() != nil {
		panic("error connecting to Database")
	}

	return &Container{
		Database: database,
	}
}
