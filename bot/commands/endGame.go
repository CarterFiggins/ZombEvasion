package commands

import (
	"fmt"
	
	"infection/bot/game"
	"github.com/bwmarrin/discordgo"
)

var EndGameDetails = &discordgo.ApplicationCommand{
	Name: "end-game",
	Description: "changes roles of players waiting in queue and starts the game",
}

func EndGame(discord *discordgo.Session, interaction *discordgo.InteractionCreate) {
	response := &discordgo.InteractionResponseData{
		Content: "Game Ended",
	}
	err := game.End(discord, interaction)
	if err != nil {
		response.Content = fmt.Sprintf("ERROR: %v", err)
		response.Flags = discordgo.MessageFlagsEphemeral
	}

	discord.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: response,
	})
}