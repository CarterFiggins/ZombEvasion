package commands

import (
	"fmt"
	"os"

	"infection/bot/game"
	"infection/hexagonGrid"
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
		response.Content = err.Error()
	}

	if hexagonGrid.Loaded {
		file, err := os.Open("./gameBoard.png")
		if err != nil {
			response.Flags = discordgo.MessageFlagsEphemeral
			response.Content = err.Error()
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

	fmt.Println(response.Files)

	discord.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: response,
	})
}