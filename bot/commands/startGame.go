package commands

import (
	"os"
	"fmt"

	"infection/bot/game"
	"github.com/bwmarrin/discordgo"
)

var StartGameDetails = &discordgo.ApplicationCommand{
	Name: "start-game",
	Description: "changes roles of players waiting in queue and starts the game",
}

func StartGame(discord *discordgo.Session, interaction *discordgo.InteractionCreate) {
	response := &discordgo.InteractionResponseData{
		Content: "Game Started",
	}

	err := game.Start(discord, interaction)
	if err != nil {
		response.Flags = discordgo.MessageFlagsEphemeral
		response.Content = fmt.Sprintf("ERROR: %v", err)
	} else {
		file, err := os.Open("./gameBoard.png")
		if err != nil {
			response.Flags = discordgo.MessageFlagsEphemeral
			response.Content = fmt.Sprintf("ERROR: %v", err)
		} else {
			response.Files = []*discordgo.File{
				{
					ContentType: "image/png",
					Name: "gameBoard.png",
					Reader: file,
				},
			}
		}
	}

	discord.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: response,
	})
}