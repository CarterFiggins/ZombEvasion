package components

import (
	"fmt"

	"infection/bot/role"
	"infection/bot/respond"
	"github.com/bwmarrin/discordgo"
)

func JoinNextGame(discord *discordgo.Session, interaction *discordgo.InteractionCreate) {
	ok, err := role.UserHasOneRole(discord, interaction, []string{role.WaitingForNextGame, role.InGame})
	if err != nil {
		respond.WithError(discord, interaction, err)
		return 
	}
	if ok {
		respond.WithMessage(discord, interaction, "You are already in the queue")
		return
	}

	response := &discordgo.InteractionResponseData{
		Content: fmt.Sprintf("%v is in the queue for the next game", interaction.Interaction.Member.User.Mention()),
	}

	err = role.AddRole(discord, interaction, role.WaitingForNextGame)
	if err != nil {
		respond.WithError(discord, interaction, err)
		return
	}

	discord.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: response,
	})
}