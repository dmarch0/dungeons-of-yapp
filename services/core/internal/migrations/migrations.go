package migrations

import (
	"context"
	"core/configs"
)

func RunMigrations(c context.Context) {
	config := configs.GetConfig()

	if config.MigrationVersion <= 0 {
		CreateTables(c)
	}

	if config.MigrationVersion <= 1 {
		AddCharacterStats(c)
	}
}