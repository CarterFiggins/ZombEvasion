package commands

import (
	"infection/bot/game"
	"github.com/bwmarrin/discordgo"
)

var EndGameDetails = &discordgo.ApplicationCommand{
	Name: "end-game",
	Description: "changes roles of players waiting in queue and starts the game",
}

func EndGame(discord *discordgo.Session, interaction *discordgo.InteractionCreate) {
	message := "Game Ended"
	err := game.End(discord, interaction)
	if err != nil {
		message = err.Error()
	}

	discord.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: message,
			Flags: discordgo.MessageFlagsEphemeral,
		},
	})
}