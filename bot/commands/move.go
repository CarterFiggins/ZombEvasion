package commands

import (
	"fmt"
	"strings"

	"infection/hexagonGrid/hexTypes"
	"infection/bot/game"
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
	moveY := int(options[1].Value.(float64))
	moveX := hexTypes.LetterToNumMap[letter]

	message, err := game.CanUserMoveHere(discord, interaction, moveX, moveY);
	if err != nil {
		response.Content = fmt.Sprintf("ERROR: %v", err)
	} else {
		if message != nil {
			response.Content = *message
		} else {
			// move user here
			// Print out what happen at location (noise or all clear)
			response.Content = fmt.Sprintf("You moved to sector: %s%02d", letter, moveY)
		}
	}

	discord.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: response,
	})
}