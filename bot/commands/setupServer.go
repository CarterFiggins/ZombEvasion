package commands

import (	
	"infection/bot/role"
	"infection/bot/channel"
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
		RespondWithError(discord, interaction, err)
		return
	}

	err = channel.SetUpChannels(discord, interaction)
	if err != nil {
		RespondWithError(discord, interaction, err)
		return
	}

	discord.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: response,
	})
}
