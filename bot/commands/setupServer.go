package commands

import (
	"infection/bot/role"
	"github.com/bwmarrin/discordgo"
)

var SetupServerDetails = &discordgo.ApplicationCommand{
	Name: "setup-server",
	Description: "adds the discord roles",
}

func SetupServer(discord *discordgo.Session, interaction *discordgo.InteractionCreate) {
	message := "Server is Ready!"

	err := role.SetUpRoles(discord, interaction)
	if err != nil {
		message = err.Error()
	}

	discord.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: message,
		},
	})
}
