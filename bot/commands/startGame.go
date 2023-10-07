package commands

import (
	"infection/bot/game"
	"github.com/bwmarrin/discordgo"
)

var StartGameDetails = &discordgo.ApplicationCommand{
	Name: "start-game",
	Description: "changes roles of players waiting in queue and starts the game",
}

func StartGame(discord *discordgo.Session, interaction *discordgo.InteractionCreate) {
	message := "Game Started"
	err := game.Start(discord, interaction)
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