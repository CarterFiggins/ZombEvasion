package game

import (

	"infection/mongo"
	"infection/bot/role"
	"infection/models"
	"infection/hexagonGrid"
	"go.mongodb.org/mongo-driver/bson"
	"github.com/bwmarrin/discordgo"
)

func Start(discord *discordgo.Session, interaction *discordgo.InteractionCreate) error {
	err := models.CreateOrUpdateGameDocument(discord, interaction)
	if err != nil {
		return err
	}

	err = role.MoveRolesInGame(discord, interaction)
	if err != nil {
		return err
	}

	err = hexagonGrid.Board.LoadBoard()
	if err != nil {
		return err
	}

	return nil
}

func End(discord *discordgo.Session, interaction *discordgo.InteractionCreate) error {
	gameDb := mongo.Db.Collection("games")
	guildID := interaction.Interaction.GuildID

	err := role.RemoveAllInGameRoles(discord, interaction)
	if err != nil {
		return err
	}

	_, err = gameDb.UpdateOne(
		mongo.Ctx,
		bson.M{"discord_guild_id": guildID},
		bson.D{
			{"$set", bson.D{{"active", false}}},
		},
	)
	if err != nil {
		return err
	}

	hexagonGrid.Board.UnloadGame()

	return nil
}
