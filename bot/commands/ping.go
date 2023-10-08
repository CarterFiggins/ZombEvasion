package commands

import (
	"github.com/bwmarrin/discordgo"
)

var PingDetails = &discordgo.ApplicationCommand{
	Name: "ping",
	Description: "ping the bot",
}

func Ping(discord *discordgo.Session, interaction *discordgo.InteractionCreate) {
	response := &discordgo.InteractionResponseData{
		Content: "PONG!",
	}

	discord.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: response,
	})
}