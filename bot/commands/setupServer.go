package commands

import (	
	"infection/bot/role"
	"infection/bot/respond"
	"infection/bot/channel"
	"infection/hexagonGrid"
	"github.com/bwmarrin/discordgo"
)

var SetupServerDetails = &discordgo.ApplicationCommand{
	Name: "setup-server",
	Description: "adds the discord roles",
}

func SetupServer(discord *discordgo.Session, interaction *discordgo.InteractionCreate) {
	discord.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Flags: discordgo.MessageFlagsEphemeral,
		},
	})

	err := role.SetUpRoles(discord, interaction)
	if err != nil {
		respond.EditWithError(discord, interaction, err)
		return
	}

	err = channel.SetUpChannels(discord, interaction)
	if err != nil {
		respond.EditWithError(discord, interaction, err)
		return
	}

	err = hexagonGrid.CreateAllGameImages()
	if err != nil {
		respond.EditWithError(discord, interaction, err)
		return
	}

	content := "Server is Ready!"
	response := &discordgo.WebhookEdit{
		Content: &content,
	}

	discord.InteractionResponseEdit(interaction.Interaction, response)
}
