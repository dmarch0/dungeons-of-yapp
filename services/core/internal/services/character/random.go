package character

import (
	"context"
	"core/internal/db"
	"core/internal/utils"
	"log"
)

type CreateRandomCharacterProps struct {
	User_id int
	Name    string
}

func CreateRandomCharacter(ctx context.Context, props *CreateRandomCharacterProps) error {
	conn := db.GetDbConnection()

	createCharacterQuery := `
		INSERT INTO characters(name, user_id, strength, dexterity, constitution, intelligence, wisdom, charisma)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8);
	`

	log.Println("User id", props.User_id)
	if _, err := conn.Exec(ctx, createCharacterQuery, props.Name, props.User_id, utils.ND6(3), utils.ND6(3), utils.ND6(3), utils.ND6(3), utils.ND6(3), utils.ND6(3)); err != nil {
		log.Println("Error creating random char", err)
		return err
	}

	return nil
}
