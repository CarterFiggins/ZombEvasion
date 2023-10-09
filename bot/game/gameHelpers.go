package game

import (
	"math"
	"time"
	"math/rand"
	"errors"

	"infection/bot/role"
	"infection/models"
	"infection/hexagonGrid"
	"infection/hexagonGrid/hexTypes"
	"github.com/bwmarrin/discordgo"
)

const (
	Human string = "Human"
	Zombie = "Zombie"
)

func Start(discord *discordgo.Session, interaction *discordgo.InteractionCreate) error {
	err := hexagonGrid.Board.LoadBoard()
	if err != nil {
		return err
	}
	if (!hexagonGrid.Board.Loaded) {
		return errors.New("Board did not load")
	}

	users, err := role.GetDiscordUsersFromRole(discord, interaction, role.WaitingForNextGame)
	if err != nil {
		return err
	}

	err = role.MoveRoleToInGame(discord, interaction, users)
	if err != nil {
		return err
	}

	if len(users) < 2 {
		return errors.New("Not enough players (need more than one)")
	}

	// basic way to compute characters for now
	numOfUsers := float64(len(users))
	half := numOfUsers / 2.0
	numOfHumans := math.Ceil(half)

	// shuffle users
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(users), func(i, j int) { users[i], users[j] = users[j], users[i] })

	var mongoUsers []*models.MongoUser
	for i, user := range users {
		mongoUser := &models.MongoUser{
			DiscordUserID: user.ID,
			DiscordGuildID: interaction.Interaction.GuildID,
			DiscordUsername: user.Username,
		}

		if i + 1 <= int(numOfHumans) {
			// is human
			mongoUser.Role = Human
			mongoUser.Location = &hexTypes.Location{
				Col: hexagonGrid.Board.HumanSector.Col,
				Row: hexagonGrid.Board.HumanSector.Row,
			}
			mongoUser.MaxMoves = 1
		} else {
			// is zombie
			mongoUser.Role = Zombie
			mongoUser.Location = &hexTypes.Location{
				Col: hexagonGrid.Board.ZombieSector.Col,
				Row: hexagonGrid.Board.ZombieSector.Row,
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
			hexagonGrid.Board.CurrentDiscordUserID = mongoUser.DiscordUserID
			mongoUser.NextDiscordUserID = mongoUsers[i + 1].DiscordUserID
			mongoUser.PrevDiscordUserID = mongoUsers[len(mongoUsers) - 1].DiscordUserID
		} else if i == len(mongoUsers)-1 {
			mongoUser.NextDiscordUserID = mongoUsers[0].DiscordUserID
			mongoUser.PrevDiscordUserID = mongoUsers[i - 1].DiscordUserID
		} else {
			mongoUser.NextDiscordUserID = mongoUsers[i + 1].DiscordUserID
			mongoUser.PrevDiscordUserID = mongoUsers[i - 1].DiscordUserID
		}
	}

	err = models.CreateMongoUsers(mongoUsers)
	if err != nil {
		return err
	}

	err = models.CreateOrUpdateGameDocument(discord, interaction)
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

	err = models.DeactivateGame(guildID)
	if err != nil {
		return err
	}

	err = models.DeleteAllUsers(guildID)
	if err != nil {
		return err
	}

	hexagonGrid.Board.UnloadGame()

	return nil
}
