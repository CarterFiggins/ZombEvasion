package models

import (
	"errors"
	"reflect"

	"infection/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"github.com/bwmarrin/discordgo"
)

type MongoGame struct {
	DiscordGuildID string
	Active bool
	BoardName string
	AlertMessages []string
}

func bsonGameToMongoGame(game bson.M) *MongoGame {
	return &MongoGame{
		DiscordGuildID: game["discord_guild_id"].(string),
		Active: game["active"].(bool),
		BoardName: game["board_name"].(string),
	}
}


func CreateOrUpdateGameDocument(discord *discordgo.Session, interaction *discordgo.InteractionCreate) error {
	gameDb := mongo.Db.Collection("games")
	boardName := interaction.ApplicationCommandData().Options[0].Value.(string)
	guildID := interaction.Interaction.GuildID
	count, err := gameDb.CountDocuments(mongo.Ctx, bson.M{"discord_guild_id": guildID})
	if err != nil {
		return err
	}

	if count == 0 {
		_, err := gameDb.InsertOne(mongo.Ctx, bson.D{
			{Key: "discord_guild_id", Value: guildID},
			{Key: "active", Value: true},
			{Key: "board_name", Value: boardName},
			{Key: "alert_messages", Value: []string{}},
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
				{"$set", bson.D{{"board_name", boardName}}},
				{"$set", bson.D{{"alert_messages", []string{}}}},
			},
		)
		if err != nil {
			return err
		}
	}
	return nil
}

func GetGame(guildID string) (*MongoGame, error) {
	gameDb := mongo.Db.Collection("games")
	var game bson.M
	if err := gameDb.FindOne(mongo.Ctx, bson.M{"discord_guild_id": guildID}).Decode(&game); err != nil {
		return nil, err
	}

	return bsonGameToMongoGame(game), nil
}

func DeactivateGame(guildID string) error {
	gameDb := mongo.Db.Collection("games")
	_, err := gameDb.UpdateOne(
		mongo.Ctx,
		bson.M{"discord_guild_id": guildID},
		bson.D{
			{"$set", bson.D{{"active", false}}},
			{"$set", bson.D{{"alert_messages", []string{}}}},
		},
	)
	if err != nil {
		return err
	}
	return nil
}

func AddAlertMessage(alertMessage, guildID string) error {
	gameDb := mongo.Db.Collection("games")
	_, err := gameDb.UpdateOne(
		mongo.Ctx,
		bson.M{"discord_guild_id": guildID},
		bson.M{"$push": bson.M{"alert_messages": alertMessage}},
	)
	if err != nil {
		return err
	}
	return nil
}

func LastRoundAlertMessages(guildID string) ([]string, error) {
	gameDb := mongo.Db.Collection("games")
	var game bson.M
	if err := gameDb.FindOne(mongo.Ctx, bson.M{"discord_guild_id": guildID}).Decode(&game); err != nil {
		return nil, err
	}

	numOfPlyersInGame, err := NumOfPlayersInGame(guildID)
	if err != nil {
		return nil, err
	}

	alerts := []string{}
	if messages, ok := game["alert_messages"].(bson.A); ok {
		for _, message := range messages {
			alerts = append(alerts, message.(string))
		}
	}

	if (len(alerts) < numOfPlyersInGame) {
		return alerts, nil
	}
	return alerts[len(alerts) - numOfPlyersInGame:], nil

}

