package components

import (	
	"infection/bot/role"
	"infection/bot/respond"
	"github.com/bwmarrin/discordgo"
)

func LeaveQueue(discord *discordgo.Session, interaction *discordgo.InteractionCreate) {
	ok, err := role.UserHasOneRole(discord, interaction, []string{role.WaitingForNextGame})
	if err != nil {
		respond.WithError(discord, interaction, err)
		return 
	}
	if !ok {
		respond.WithMessage(discord, interaction, "You are not in the queue")
		return
	}

	response := &discordgo.InteractionResponseData{
		Content: "You have left the queue",
		Flags: discordgo.MessageFlagsEphemeral,
	}

	err = role.RemoveRole(discord, interaction, role.WaitingForNextGame)
	if err != nil {
		respond.WithError(discord, interaction, err)
		return
	}

	discord.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: response,
	})
}