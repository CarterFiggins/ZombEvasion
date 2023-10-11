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
				Description: "The column of the hex gird",
				Required: true,
			},
			{
				Type: discordgo.ApplicationCommandOptionInteger,
				Name: "row-number",
				Description: "The row of the hex gird",
				MinValue: &minValue,
				Required: true,
			},
		},
	}
)

func Move(discord *discordgo.Session, interaction *discordgo.InteractionCreate) {
	response := &discordgo.InteractionResponseData{
		Content: "Can't move here try again",
		Flags: discordgo.MessageFlagsEphemeral,
	}

	options := interaction.ApplicationCommandData().Options
	letter := strings.ToUpper(options[0].Value.(string))
	moveY := int(options[1].Value.(float64)) - 1
	moveX := hexSectors.LetterToNumMap[letter]

	message, err := game.CanUserMoveHere(discord, interaction, moveX, moveY);
	if err != nil {
		RespondWithError(discord, interaction, err)
		return
	}
	if message != nil {
		response.Content = *message
	} else {
		err = models.MoveUser(interaction, moveX, moveY)
		if err != nil {
			RespondWithError(discord, interaction, err)
			return
		}
		sector := hexagonGrid.Board.Grid[moveX][moveY]
		sectorName := sector.GetSectorName()
		infectionGameChannel := GetChannel(discord, interaction, channel.InfectionGameChannelName)
		if infectionGameChannel == nil {
			return
		}
		turnMessage, userMessage := game.MovedOnSectorMessages(interaction, sectorName, hexSectors.GetHexName(moveX, moveY))
		response.Content = userMessage
		_, err = discord.ChannelMessageSend(infectionGameChannel.ID, turnMessage)
	}

	discord.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: response,
	})
}