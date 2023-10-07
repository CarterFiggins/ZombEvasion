package commands

import (
	"fmt"

	"infection/bot/role"
	"github.com/bwmarrin/discordgo"
)

var JoinNextGameDetails = &discordgo.ApplicationCommand{
	Name: "join-next-game",
	Description: "adds the WaitingForNextGame role to a user",
}

func JoinNextGame(discord *discordgo.Session, interaction *discordgo.InteractionCreate) {

	response := &discordgo.InteractionResponseData{
		Content: fmt.Sprintf("%v is in the queue for the next game", interaction.Interaction.Member.User.Mention()),
	}

	err := role.AddRole(discord, interaction, role.WaitingForNextGame)
	if err != nil {
		response.Content = fmt.Sprintf("ERROR: %v", err)
		response.Flags = discordgo.MessageFlagsEphemeral
	}

	discord.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: response,
	})
}