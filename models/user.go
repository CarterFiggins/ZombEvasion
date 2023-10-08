package models

import (
	"infection/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

type MongoUser struct {
	Role string
	X int
	Y int
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
			{Key: "x", Value: user.X},
			{Key: "y", Value: user.Y},
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

func FindUser(discordUserID, guildID string) (*MongoUser, error) {
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
		X: user["x"].(int),
		Y: user["y"].(int),
		DiscordUserID: user["discord_user_id"].(string),
		DiscordGuildID: user["discord_guild_id"].(string),
		DiscordUsername: user["discord_username"].(string),
		NextDiscordUserID: user["next_discord_user_id"].(string),
		PrevDiscordUserID: user["prev_discord_user_id"].(string),
	}
}
