package commands

import (
	"fmt"
	
	"infection/bot/game"
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
	response := &discordgo.InteractionResponseData{
		Content: "Game Ended",
		Flags: discordgo.MessageFlagsEphemeral,
	}

	err := game.End(discord, interaction)
	if err != nil {
		response.Content = fmt.Sprintf("ERROR: %v", err)
	}

	discord.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: response,
	})
}