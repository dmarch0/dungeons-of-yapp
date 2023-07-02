package controllers

import (
	"core/internal/services/character"

	"github.com/gofiber/fiber/v2"
)

func PostCharacterController(ctx *fiber.Ctx) error {

	return nil
}

func GetCharacterController(ctx *fiber.Ctx) error {

	return nil
}

type PostCharacterRandomResult struct {
	Success bool `json:"success"`
}

func PostCharacterRandomController(ctx *fiber.Ctx) error {
	user_id, _ := ctx.Locals("user_id").(int)
	err := character.CreateRandomCharacter(ctx.Context(), &character.CreateRandomCharacterProps{
		Name:    "Text",
		User_id: user_id,
	})

	if err != nil {
		return ctx.Status(500).Next()
	}

	return ctx.Status(200).JSON(PostCharacterRandomResult{
		Success: true,
	})
}
