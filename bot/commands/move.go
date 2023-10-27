package commands

import (
	"strings"

	"infection/models"
	"infection/hexagonGrid/hexSectors"
	"infection/hexagonGrid"
	"infection/bot/game"
	"infection/bot/respond"
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
	mongoUser, err := models.FindUser(interaction)
	if err != nil {
		respond.WithError(discord, interaction, err)
		return
	}

	if !mongoUser.TurnActive {
		respond.WithMessage(discord, interaction, "It is not your turn")
		return
	}

	if !mongoUser.CanMove {
		respond.WithMessage(discord, interaction, "You have already moved")
		return
	}

	discord.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Flags: discordgo.MessageFlagsEphemeral,
		},
	})

	options := interaction.ApplicationCommandData().Options
	letter := strings.ToUpper(options[0].Value.(string))
	moveY := int(options[1].Value.(float64)) - 1
	moveX := hexSectors.LetterToNumMap[letter]

	content := "Can't move here try again"
	response := &discordgo.WebhookEdit{
		Content: &content,
	}

	message, err := game.CanUserMoveHere(discord, interaction, moveX, moveY, mongoUser);
	if err != nil {
		respond.EditWithError(discord, interaction, err)
		return
	}
	if message != nil {
		respond.EditWithMessage(discord, interaction, *message)
		return
	} else {
		err = mongoUser.MoveUser(moveX, moveY)
		if err != nil {
			respond.EditWithError(discord, interaction, err)
			return
		}
		sector := hexagonGrid.Board.Grid[moveX][moveY]
		sectorName := sector.GetSectorName()
		alertsChannel, err := channel.GetChannel(discord, interaction, channel.Alerts)
		if err != nil {
			respond.EditWithError(discord, interaction, err)
			return
		}
		turnMessage, userMessage, err := game.MovedOnSectorMessages(discord, interaction, sectorName, hexSectors.GetHexName(moveX, moveY))
		if err != nil {
			respond.EditWithError(discord, interaction, err)
			return
		}

		response.Content = &userMessage
		if (turnMessage != "") {
			_, err = discord.ChannelMessageSend(alertsChannel.ID, turnMessage)
			if err != nil {
				respond.EditWithError(discord, interaction, err)
				return
			}
		}
	}

	discord.InteractionResponseEdit(interaction.Interaction, response)
}