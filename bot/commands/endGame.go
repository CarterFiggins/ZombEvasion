package commands

import (
	"infection/bot/game"
	"infection/bot/respond"
	"infection/bot/role"
	"github.com/bwmarrin/discordgo"
)

var EndGameDetails = &discordgo.ApplicationCommand{
	Name: "end-game",
	Description: "changes roles of players waiting in queue and starts the game",
}

func EndGame(discord *discordgo.Session, interaction *discordgo.InteractionCreate) {
	if ok := CheckPermissions(discord, interaction, []string{role.Admin}); !ok {
		return
	}

	discord.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Flags: discordgo.MessageFlagsEphemeral,
		},
	})

	message := "Game Ended"
	response := &discordgo.WebhookEdit{
		Content: &message,
	}

	err := game.End(discord, interaction)
	if err != nil {
		respond.EditWithError(discord, interaction, err)
		return
	}

	discord.InteractionResponseEdit(interaction.Interaction, response)

}