package boot

import (
	"context"
	"core/internal/db"
	"core/internal/http/server"
)

func Start() {
	c := context.Background()

	dbConnection := db.ConnectToDb(c)
	defer dbConnection.Close(c)
	db.InitTables(c)

	server.StartServer()
}