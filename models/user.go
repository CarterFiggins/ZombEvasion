package models

import (
	"infection/mongo"
	"infection/hexagonGrid/hexTypes"

	"go.mongodb.org/mongo-driver/bson"
)

type MongoUser struct {
	Role string
	*hexTypes.Location
	MaxMoves int
	DiscordUserID string
	DiscordGuildID string
	DiscordUsername string
	NextDiscordUserID string
	PrevDiscordUserID string
}

func CreateMongoUsers(mongoUsers []*MongoUser) error {
	userDb := mongo.Db.Collection("users")
	for _, user := range mongoUsers {
		_, err := userDb.InsertOne(mongo.Ctx, bson.D{
			{Key: "role", Value: user.Role},
			{Key: "col", Value: user.Col},
			{Key: "row", Value: user.Row},
			{Key: "max_moves", Value: user.MaxMoves},
			{Key: "discord_user_id", Value: user.DiscordUserID},
			{Key: "discord_guild_id", Value: user.DiscordGuildID},
			{Key: "discord_username", Value: user.DiscordUsername},
			{Key: "next_discord_user_id", Value: user.NextDiscordUserID},
			{Key: "prev_discord_user_id", Value: user.PrevDiscordUserID},
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func FindUser(guildID, discordUserID string) (*MongoUser, error) {
	userDb := mongo.Db.Collection("users")
	var user bson.M
	if err := userDb.FindOne(mongo.Ctx, bson.M{"discord_guild_id": guildID, "discord_user_id": discordUserID}).Decode(&user); err != nil {
		return nil, err
	}
	return bsonUserToMongoUser(user), nil
}

func DeleteAllUsers(guildID string) error {
	userDb := mongo.Db.Collection("users")

	_, err := userDb.DeleteMany(mongo.Ctx, bson.M{"discord_guild_id": guildID})
	if err != nil {
		return err
	}
	return nil
}

func bsonUserToMongoUser(user bson.M) *MongoUser {
	return &MongoUser{
		Role: user["role"].(string),
		Location: &hexTypes.Location{
			Col: int(user["col"].(int32)),
			Row: int(user["row"].(int32)),
		},
		MaxMoves: int(user["max_moves"].(int32)),
		DiscordUserID: user["discord_user_id"].(string),
		DiscordGuildID: user["discord_guild_id"].(string),
		DiscordUsername: user["discord_username"].(string),
		NextDiscordUserID: user["next_discord_user_id"].(string),
		PrevDiscordUserID: user["prev_discord_user_id"].(string),
	}
}
