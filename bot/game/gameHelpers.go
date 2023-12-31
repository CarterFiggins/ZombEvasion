package game

import (
	"math"
	"time"
	"math/rand"
	"errors"
	"fmt"

	"infection/bot/role"
	"infection/bot/channel"
	"infection/models"
	"infection/hexagonGrid"
	"infection/hexagonGrid/hexSectors"
	"github.com/bwmarrin/discordgo"
)

func Start(discord *discordgo.Session, interaction *discordgo.InteractionCreate) error {

	users, err := role.GetDiscordUsersFromRole(discord, interaction, role.WaitingForNextGame)
	if err != nil {
		return err
	}

	if len(users) < 2 {
		return errors.New("Not enough players (need more than one)")
	}

	err = models.CreateOrUpdateGameDocument(discord, interaction)
	if err != nil {
		return err
	}
	
	err = role.MoveRoleToInGame(discord, interaction, users)
	if err != nil {
		return err
	}

	// basic way to compute characters for now
	numOfUsers := float64(len(users))
	half := numOfUsers / 2.0
	numOfHumans := math.Ceil(half)

	// shuffle users
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(users), func(i, j int) { users[i], users[j] = users[j], users[i] })

	board, err := hexagonGrid.GetBoard(interaction.Interaction.GuildID)
	if err != nil {
		return err
	}

	zombieSector := hexagonGrid.FindZombieSector(board)
	humanSector := hexagonGrid.FindHumanSector(board)
	var mongoUsers []*models.MongoUser
	for i, user := range users {
		mongoUser := &models.MongoUser{
			DiscordUserID: user.ID,
			DiscordGuildID: interaction.Interaction.GuildID,
			DiscordUsername: user.Username,
			InGame: true,
		}
		if i + 1 <= int(numOfHumans) {
			// is human
			mongoUser.Role = models.Human
			mongoUser.Location = &hexSectors.Location{
				Col: humanSector.Col,
				Row: humanSector.Row,
			}
			mongoUser.MaxMoves = 1
		} else {
			// is zombie
			mongoUser.Role = models.Zombie
			mongoUser.Location = &hexSectors.Location{
				Col: zombieSector.Col,
				Row: zombieSector.Row,
			}
			mongoUser.MaxMoves = 2
		}
		mongoUsers = append(mongoUsers, mongoUser)
	}

	// shuffle mongoUsers for random turn
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(mongoUsers), func(i, j int) { mongoUsers[i], mongoUsers[j] = mongoUsers[j], mongoUsers[i] })
	for i, mongoUser := range mongoUsers {
		if i == 0 {
			mongoUser.NextDiscordUserID = mongoUsers[i + 1].DiscordUserID
			mongoUser.PrevDiscordUserID = mongoUsers[len(mongoUsers) - 1].DiscordUserID
			mongoUser.TurnActive = true
			mongoUser.CanMove = true
		} else if i == len(mongoUsers)-1 {
			mongoUser.NextDiscordUserID = mongoUsers[0].DiscordUserID
			mongoUser.PrevDiscordUserID = mongoUsers[i - 1].DiscordUserID
			mongoUser.TurnActive = false
		} else {
			mongoUser.NextDiscordUserID = mongoUsers[i + 1].DiscordUserID
			mongoUser.PrevDiscordUserID = mongoUsers[i - 1].DiscordUserID
			mongoUser.TurnActive = false
		}
	}

	err = models.CreateMongoUsers(mongoUsers)
	if err != nil {
		return err
	}

	if err = mongoUsers[0].StartTurn(); err != nil {
		return err
	}

	if err = SetUpUsersTurn(discord, interaction.Interaction.GuildID, mongoUsers[0]); err != nil {
		return err
	}

	return nil
}

func ClearDatabase(guildID string) error {
	err := models.DeactivateGame(guildID)
	if err != nil {
		return err
	}

	err = models.DeleteAllUsers(guildID)
	if err != nil {
		return err
	}

	return nil
}

func End(discord *discordgo.Session, interaction *discordgo.InteractionCreate) error {
	guildID := interaction.Interaction.GuildID

	users, err := role.GetDiscordUsersFromRole(discord, interaction, role.InGame)
	if err != nil {
		return err
	}

	err = role.RemoveInGameRole(discord, interaction, users)
	if err != nil {
		return err
	}

	err = role.AddRoleToUsers(discord, interaction, role.WaitingForNextGame, users)
	if err != nil {
		return err
	}

	ClearDatabase(guildID)

	return nil
}

func CheckGame(discord *discordgo.Session, guildID string) error {
	mongoUsers, err := models.GetAllUsersPlaying(guildID)
	if err != nil {
		return err
	}

	gameOver := true

	for _, mongoUser := range mongoUsers {
		if (mongoUser.Role == models.Human){
			gameOver = false
		}
	}

	if (gameOver) {
		// send game over message and who survived
		gameChannel, err := channel.GetChannel(discord, guildID, channel.InfectionGameChannelName)
		if err != nil {
			return errors.New("no channel")
		}

		survivers, err := models.GetSurvivers(guildID)

		var surviversNames []string
		for _, surviver := range survivers {
			surviversNames = append(surviversNames, surviver.DiscordUsername)
		}

		_, err = discord.ChannelMessageSend(gameChannel.ID, fmt.Sprintf("Game Over! Survivers: %v", surviversNames))
		if err != nil {
			return err
		}

		if err != nil {
			return err
		}

		ClearDatabase(guildID)
	}

	return nil
}
