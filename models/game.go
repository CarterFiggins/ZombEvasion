package models

import (
	"errors"
	"reflect"

	"infection/hexagonGrid"
	"infection/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"github.com/bwmarrin/discordgo"
)

type MongoGame struct {
	DiscordGuildId string
	Active bool
	BoardName string
}


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
			{Key: "board_name", Value: hexagonGrid.Board.Name},
			{Key: "current_discord_user_id", Value: hexagonGrid.Board.CurrentDiscordUserID},
		})
		if err != nil {
			return err
		}
	} else {
		var mongoGame bson.M
		if err = gameDb.FindOne(mongo.Ctx, bson.M{"discord_guild_id": guildID}).Decode(&mongoGame); err != nil {
			return err
		}
		active := reflect.ValueOf(mongoGame["active"]).Bool()
		if active {
			return errors.New("Game is already active")
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

func UpdateCurrentUserID(guildID, discordUserID string) error {
	gameDb := mongo.Db.Collection("games")
	_, err := gameDb.UpdateOne(
		mongo.Ctx,
		bson.M{"discord_guild_id": guildID},
		bson.D{
			{"$set", bson.D{{"current_discord_user_id", discordUserID}}},
		},
	)
	if err != nil {
		return err
	}
	return nil
}

func DeactivateGame(guildID string) error {
	gameDb := mongo.Db.Collection("games")

	_, err := gameDb.UpdateOne(
		mongo.Ctx,
		bson.M{"discord_guild_id": guildID},
		bson.D{
			{"$set", bson.D{{"active", false}}},
		},
	)
	if err != nil {
		return err
	}
	return nil
}

