package migrations

import (
	"context"
	"core/internal/db"
	"log"
)

func AddCharacterStats(c context.Context) {
	connection := db.GetDbConnection()

	addStatsQuery := `
		ALTER TABLE characters
		ADD COLUMN IF NOT EXISTS strength int NOT NULL,
		ADD COLUMN IF NOT EXISTS agility int NOT NULL,
		ADD COLUMN IF NOT EXISTS constitution int NOT NULL,
		ADD COLUMN IF NOT EXISTS intelligence int NOT NULL,
		ADD COLUMN IF NOT EXISTS wisdom int NOT NULL,
		ADD COLUMN IF NOT EXISTS charisma int NOT NULL;
	`

	if _, err := connection.Exec(c, addStatsQuery); err != nil {
		log.Panic("Error: AddCharacterStats", err)
	}

	log.Println("Success: AddCharacterStats")
}