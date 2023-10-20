package commands

import (
	"infection/bot/role"
	"infection/hexagonGrid"
	"github.com/bwmarrin/discordgo"
)

var LoadMapDetails = &discordgo.ApplicationCommand{
	Name: "load-map",
	Description: "loads the map into memory",
}

func LoadMap(discord *discordgo.Session, interaction *discordgo.InteractionCreate) {
	if ok := CheckPermissions(discord, interaction, []string{role.Admin}); !ok {
		return
	}

	discord.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
	})

	if err := hexagonGrid.Board.LoadBoard(); err != nil {
		RespondWithError(discord, interaction, err)
		return
	}

	message := "Board Loaded"
	response := &discordgo.WebhookEdit{
		Content: &message,
	}
	
	discord.InteractionResponseEdit(interaction.Interaction, response)
}