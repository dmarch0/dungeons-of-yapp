package migrations

import (
	"context"
	"core/internal/db"
	"log"
)

func RenameAgility(c context.Context) {
	connection := db.GetDbConnection()

	renameAgilityQuery := `
		ALTER TABLE characters
		RENAME COLUMN agility TO dexterity;
	`

	if _, err := connection.Exec(c, renameAgilityQuery); err != nil {
		log.Panic("Error: RenameAgility", err)
	}

	log.Println("Success: RenameAgility")
}
