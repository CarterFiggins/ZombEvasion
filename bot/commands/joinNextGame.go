package commands

import (
	"fmt"

	"infection/bot/role"
	"infection/bot/respond"
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
		respond.WithError(discord, interaction, err)
		return
	}

	discord.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: response,
	})
}