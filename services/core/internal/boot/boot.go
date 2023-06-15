package boot

import (
	"context"
	"core/internal/db"
	"core/internal/http/server"
	"core/internal/migrations"
)

func Start() {
	c := context.Background()

	dbConnection := db.ConnectToDb(c)
	defer dbConnection.Close(c)

	migrations.RunMigrations(c)

	server.StartServer()
}