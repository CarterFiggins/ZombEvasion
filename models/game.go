package models

import (
	"errors"
	"reflect"

	"infection/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"github.com/bwmarrin/discordgo"
)

func CreateOrUpdateGameDocument(discord *discordgo.Session, interaction *discordgo.InteractionCreate) error {
	gameDb := mongo.Db.Collection("games")
	guildID := interaction.Interaction.GuildID
	count, err := gameDb.CountDocuments(mongo.Ctx, bson.M{"discord_guild_id": guildID})
	if err != nil {
		return err
	}

	if count == 0 {
		_, err := gameDb.InsertOne(mongo.Ctx, bson.D{
			{Key: "discord_guild_id", Value: guildID},
			{Key: "active", Value: true},
			{Key: "board_name", Value: ""},
		})
		if err != nil {
			return err
		}
	} else {
		var theGame bson.M
		if err = gameDb.FindOne(mongo.Ctx, bson.M{"discord_guild_id": guildID}).Decode(&theGame); err != nil {
			return err
		}
		active := reflect.ValueOf(theGame["active"]).Bool()
		if active {
			return errors.New("Error: Game is already active")
		}
		_, err := gameDb.UpdateOne(
			mongo.Ctx,
			bson.M{"discord_guild_id": guildID},
			bson.D{
				{"$set", bson.D{{"active", true}}},
			},
		)
		if err != nil {
			return err
		}
	}
	return nil
}

