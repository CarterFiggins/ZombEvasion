package commands

import (
	"strings"

	"infection/models"
	"infection/hexagonGrid/hexSectors"
	"infection/hexagonGrid"
	"infection/bot/game"
	"infection/bot/channel"
	"github.com/bwmarrin/discordgo"
)

var (
	minValue = 1.0
	MoveDetails = &discordgo.ApplicationCommand{
		Name: "move",
		Description: "move to another space on the board",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type: discordgo.ApplicationCommandOptionString,
				Name: "letter-column",
				Description: "The column of the sector",
				Required: true,
			},
			{
				Type: discordgo.ApplicationCommandOptionInteger,
				Name: "row-number",
				Description: "The row of the sector",
				MinValue: &minValue,
				Required: true,
			},
		},
	}
)

func Move(discord *discordgo.Session, interaction *discordgo.InteractionCreate) {
	mongoUser, err := models.FindUser(interaction, nil)
	if err != nil {
		RespondWithError(discord, interaction, err)
		return
	}

	if !mongoUser.TurnActive {
		RespondWithMessage(discord, interaction, "It is not your turn")
		return
	}

	if !mongoUser.CanMove {
		RespondWithMessage(discord, interaction, "You have already moved")
		return
	}

	response := &discordgo.InteractionResponseData{
		Content: "Can't move here try again",
		Flags: discordgo.MessageFlagsEphemeral,
	}

	options := interaction.ApplicationCommandData().Options
	letter := strings.ToUpper(options[0].Value.(string))
	moveY := int(options[1].Value.(float64)) - 1
	moveX := hexSectors.LetterToNumMap[letter]

	message, err := game.CanUserMoveHere(discord, interaction, moveX, moveY, mongoUser);
	if err != nil {
		RespondWithError(discord, interaction, err)
		return
	}
	if message != nil {
		RespondWithMessage(discord, interaction, *message)
		return
	} else {
		err = mongoUser.MoveUser(moveX, moveY)
		if err != nil {
			RespondWithError(discord, interaction, err)
			return
		}
		sector := hexagonGrid.Board.Grid[moveX][moveY]
		sectorName := sector.GetSectorName()
		alertsChannel := GetChannel(discord, interaction, channel.Alerts)
		if alertsChannel == nil {
			return
		}
		turnMessage, userMessage, err := game.MovedOnSectorMessages(discord, interaction, sectorName, hexSectors.GetHexName(moveX, moveY))
		if err != nil {
			RespondWithError(discord, interaction, err)
			return
		}

		response.Content = userMessage
		if (turnMessage != "") {
			_, err = discord.ChannelMessageSend(alertsChannel.ID, turnMessage)
			if err != nil {
				RespondWithError(discord, interaction, err)
				return
			}
		}
	}

	discord.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: response,
	})
}