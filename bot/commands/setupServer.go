package commands

import (
	"fmt"
	
	"infection/bot/role"
	"github.com/bwmarrin/discordgo"
)

var SetupServerDetails = &discordgo.ApplicationCommand{
	Name: "setup-server",
	Description: "adds the discord roles",
}

func SetupServer(discord *discordgo.Session, interaction *discordgo.InteractionCreate) {
	response := &discordgo.InteractionResponseData{
		Content: "Server is Ready!",
	}

	err := role.SetUpRoles(discord, interaction)
	if err != nil {
		response.Flags = discordgo.MessageFlagsEphemeral
		response.Content = fmt.Sprintf("ERROR: %v", err)
	}

	discord.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: response,
	})
}
