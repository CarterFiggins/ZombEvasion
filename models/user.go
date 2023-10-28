package models

import (
	"infection/mongo"
	"infection/hexagonGrid/hexSectors"

	"go.mongodb.org/mongo-driver/bson"
	"github.com/bwmarrin/discordgo"
)

type MongoUser struct {
	Role string
	*hexSectors.Location
	InGame bool
	MaxMoves int
	DiscordUserID string
	DiscordGuildID string
	DiscordUsername string
	NextDiscordUserID string
	PrevDiscordUserID string
	TurnActive bool
	CanMove bool
	IsAttacking bool
	CanSetOffAlarm bool
	IsSafe bool
}

const (
	Human string = "Human"
	Zombie = "Zombie"
)

func CreateMongoUsers(mongoUsers []*MongoUser) error {
	userDb := mongo.Db.Collection("users")
	for _, user := range mongoUsers {
		_, err := userDb.InsertOne(mongo.Ctx, bson.D{
			{Key: "role", Value: user.Role},
			{Key: "col", Value: user.Col},
			{Key: "row", Value: user.Row},
			{Key: "in_game", Value: user.InGame},
			{Key: "is_safe", Value: user.IsSafe},
			{Key: "can_set_off_alarm", Value: user.CanSetOffAlarm},
			{Key: "turn_active", Value: user.TurnActive},
			{Key: "can_move", Value: user.CanMove},
			{Key: "is_attacking", Value: user.IsAttacking},
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

func bsonUserToMongoUser(user bson.M) *MongoUser {
	return &MongoUser{
		Role: user["role"].(string),
		Location: &hexSectors.Location{
			Col: int(user["col"].(int32)),
			Row: int(user["row"].(int32)),
		},
		InGame: user["in_game"].(bool),
		IsSafe: user["is_safe"].(bool),
		CanSetOffAlarm: user["can_set_off_alarm"].(bool),
		TurnActive: user["turn_active"].(bool),
		CanMove: user["can_move"].(bool),
		IsAttacking: user["is_attacking"].(bool),
		MaxMoves: int(user["max_moves"].(int32)),
		DiscordUserID: user["discord_user_id"].(string),
		DiscordGuildID: user["discord_guild_id"].(string),
		DiscordUsername: user["discord_username"].(string),
		NextDiscordUserID: user["next_discord_user_id"].(string),
		PrevDiscordUserID: user["prev_discord_user_id"].(string),
	}
}

func FindUserByIDs(interaction *discordgo.InteractionCreate, discordUserID *string, guildID *string) (*MongoUser, error) {
	if discordUserID == nil {
		if interaction.Interaction.Member != nil {
			discordUserID = &interaction.Interaction.Member.User.ID
		} else {
			discordUserID = &interaction.Interaction.User.ID
		}
	}

	if guildID == nil {
		guildID = &interaction.Interaction.GuildID
	}

	userDb := mongo.Db.Collection("users")
	var user bson.M
	if err := userDb.FindOne(mongo.Ctx, bson.M{"discord_guild_id": *guildID, "discord_user_id": *discordUserID}).Decode(&user); err != nil {
		return nil, err
	}
	return bsonUserToMongoUser(user), nil
}

func FindUser(interaction *discordgo.InteractionCreate) (*MongoUser, error) {
	discordUserID := interaction.Interaction.Member.User.ID
	guildID := interaction.Interaction.GuildID
	userDb := mongo.Db.Collection("users")
	var user bson.M
	if err := userDb.FindOne(mongo.Ctx, bson.M{"discord_guild_id": guildID, "discord_user_id": discordUserID}).Decode(&user); err != nil {
		return nil, err
	}
	return bsonUserToMongoUser(user), nil
}

func FindUsersAtLocation(interaction *discordgo.InteractionCreate, col, row int) ([]*MongoUser, error) {
	userDb := mongo.Db.Collection("users")
	guildID := interaction.Interaction.GuildID

	filterCursor, err := userDb.Find(mongo.Ctx, bson.M{
		"discord_guild_id": guildID,
		"col": col,
		"row": row,
	})
	if err != nil {
		return nil, err
	}
	var mongoUsersFiltered []bson.M
	if err = filterCursor.All(mongo.Ctx, &mongoUsersFiltered); err != nil {
		return nil, err
	}
	var mongoUsers []*MongoUser
	for _, bsonUser := range mongoUsersFiltered {
		mongoUsers = append(mongoUsers, bsonUserToMongoUser(bsonUser))
	}

	return mongoUsers, nil
}

func GetAllUsersPlaying(guildID string) ([]*MongoUser, error){
	userDb := mongo.Db.Collection("users")

	filterCursor, err := userDb.Find(mongo.Ctx, bson.M{"discord_guild_id": guildID, "in_game": true})
	if err != nil {
		return nil, err
	}
	var mongoUsersFiltered []bson.M
	if err = filterCursor.All(mongo.Ctx, &mongoUsersFiltered); err != nil {
		return nil, err
	}
	var mongoUsers []*MongoUser
	for _, bsonUser := range mongoUsersFiltered {
		mongoUsers = append(mongoUsers, bsonUserToMongoUser(bsonUser))
	}

	return mongoUsers, nil
}

func GetSurvivers(guildID string) ([]*MongoUser, error){
	userDb := mongo.Db.Collection("users")

	filterCursor, err := userDb.Find(mongo.Ctx, bson.M{"discord_guild_id": guildID, "is_safe": true})
	if err != nil {
		return nil, err
	}
	var mongoUsersFiltered []bson.M
	if err = filterCursor.All(mongo.Ctx, &mongoUsersFiltered); err != nil {
		return nil, err
	}
	var mongoUsers []*MongoUser
	for _, bsonUser := range mongoUsersFiltered {
		mongoUsers = append(mongoUsers, bsonUserToMongoUser(bsonUser))
	}

	return mongoUsers, nil
}

func DeleteAllUsers(guildID string) error {
	userDb := mongo.Db.Collection("users")

	_, err := userDb.DeleteMany(mongo.Ctx, bson.M{"discord_guild_id": guildID})
	if err != nil {
		return err
	}
	return nil
}

func (u *MongoUser) UpgradeUsersMaxMoves(maxMoves int) error {
	userDb := mongo.Db.Collection("users")

	_, err := userDb.UpdateOne(
		mongo.Ctx,
		bson.M{
			"discord_guild_id": u.DiscordGuildID,
			"discord_user_id": u.DiscordUserID,
		},
		bson.D{
			{"$set", bson.D{{"max_moves", maxMoves}}},
		},
	)
	if err != nil {
		return err
	}
	return nil
}

func (u *MongoUser) RespawnUser(col, row int) error {
	userDb := mongo.Db.Collection("users")

	_, err := userDb.UpdateOne(
		mongo.Ctx,
		bson.M{
			"discord_guild_id": u.DiscordGuildID,
			"discord_user_id": u.DiscordUserID,
		},
		bson.D{
			{"$set", bson.D{{"col", col}}},
			{"$set", bson.D{{"row", row}}},
		},
	)
	if err != nil {
		return err
	}
	return nil
}

func (u *MongoUser) MoveUser(col, row int) error {
	userDb := mongo.Db.Collection("users")

	_, err := userDb.UpdateOne(
		mongo.Ctx,
		bson.M{
			"discord_guild_id": u.DiscordGuildID,
			"discord_user_id": u.DiscordUserID,
		},
		bson.D{
			{"$set", bson.D{{"col", col}}},
			{"$set", bson.D{{"row", row}}},
			{"$set", bson.D{{"can_move", false}}},
		},
	)
	if err != nil {
		return err
	}
	return nil
}

func (u *MongoUser) UpdateCanSetOffAlarm(canSetOffAlarm bool) error {
	userDb := mongo.Db.Collection("users")
	
	_, err := userDb.UpdateOne(
		mongo.Ctx,
		bson.M{
			"discord_guild_id": u.DiscordGuildID,
			"discord_user_id": u.DiscordUserID,
		},
		bson.D{
			{"$set", bson.D{{"can_set_off_alarm", canSetOffAlarm}}},
		},
	)
	if err != nil {
		return err
	}
	return nil
}

func (u *MongoUser) EndTurn() error {
	userDb := mongo.Db.Collection("users")

	_, err := userDb.UpdateOne(
		mongo.Ctx,
		bson.M{
			"discord_guild_id": u.DiscordGuildID,
			"discord_user_id": u.DiscordUserID,
		},
		bson.D{
			{"$set", bson.D{{"turn_active", false}}},
		},
	)
	if err != nil {
		return err
	}
	return nil
}

func (u *MongoUser) ExitGame() error {
	userDb := mongo.Db.Collection("users")

	_, err := userDb.UpdateOne(
		mongo.Ctx,
		bson.M{
			"discord_guild_id": u.DiscordGuildID,
			"discord_user_id": u.PrevDiscordUserID,
		},
		bson.D{
			{"$set", bson.D{
				{"next_discord_user_id", u.NextDiscordUserID},
			}},
		},
	)
	if err != nil {
		return err
	}

	_, err = userDb.UpdateOne(
		mongo.Ctx,
		bson.M{
			"discord_guild_id": u.DiscordGuildID,
			"discord_user_id": u.NextDiscordUserID,
		},
		bson.D{
			{"$set", bson.D{
				{"prev_discord_user_id", u.PrevDiscordUserID},
			}},
		},
	)
	if err != nil {
		return err
	}

	return nil
}

func (u *MongoUser) TurnIntoZombie() error {
	userDb := mongo.Db.Collection("users")

	_, err := userDb.UpdateOne(
		mongo.Ctx,
		bson.M{
			"discord_guild_id": u.DiscordGuildID,
			"discord_user_id": u.DiscordUserID,
		},
		bson.D{
			{"$set", bson.D{
				{"role", Zombie},
				{"max_moves", 2},
			}},
		},
	)
	if err != nil {
		return err
	}

	return nil
}

func (u *MongoUser) MarkAttacking(isAttacking bool) error {
	userDb := mongo.Db.Collection("users")

	_, err := userDb.UpdateOne(
		mongo.Ctx,
		bson.M{
			"discord_guild_id": u.DiscordGuildID,
			"discord_user_id": u.DiscordUserID,
		},
		bson.D{
			{"$set", bson.D{
				{"is_attacking", isAttacking},
			}},
		},
	)
	if err != nil {
		return err
	}
	u.IsAttacking = isAttacking
	return nil
}

func (u *MongoUser) EnterSafeHouse() error {
	userDb := mongo.Db.Collection("users")

	_, err := userDb.UpdateOne(
		mongo.Ctx,
		bson.M{
			"discord_guild_id": u.DiscordGuildID,
			"discord_user_id": u.DiscordUserID,
		},
		bson.D{
			{"$set", bson.D{
				{"in_game", false},
				{"is_safe", true},
			}},
		},
	)
	if err != nil {
		return err
	}

	u.ExitGame()

	return nil
}

func (u *MongoUser) StartTurn() error {
	userDb := mongo.Db.Collection("users")

	set := bson.D{
		{"turn_active", true},
		{"can_move", true},
	}

	_, err := userDb.UpdateOne(
		mongo.Ctx,
		bson.M{
			"discord_guild_id": u.DiscordGuildID,
			"discord_user_id": u.DiscordUserID,
		},
		bson.D{
			{"$set", set},
		},
	)
	if err != nil {
		return err
	}
	return nil
}
