package commands

import (
	"github.com/bwmarrin/discordgo"
)

func Ping(discordS *discordgo.Session, interaction *discordgo.InteractionCreate) {
	discordS.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "PONG!",
		},
	})
}