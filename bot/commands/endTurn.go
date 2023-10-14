package commands

import (
	"infection/models"
	"infection/bot/channel"
	"github.com/bwmarrin/discordgo"
)

var EndTurnDetails = &discordgo.ApplicationCommand{
	Name: "end-turn",
	Description: "Starts the next players turn",
}

func EndTurn(discord *discordgo.Session, interaction *discordgo.InteractionCreate) {
	mongoUser, err := models.FindUser(interaction, nil)
	if err != nil {
		RespondWithError(discord, interaction, err)
		return
	}

	if !mongoUser.TurnActive {
		RespondWithMessage(discord, interaction, "It is not your turn")
		return
	}

	if mongoUser.CanMove {
		RespondWithMessage(discord, interaction, "You have to `/move` before you end your turn")
		return
	}

	if mongoUser.CanSetOffAlarm {
		RespondWithMessage(discord, interaction, "You have to `/set-off-alarm` before you end your turn")
		return
	}

	if err = mongoUser.EndTurn(); err != nil {
		RespondWithError(discord, interaction, err)
		return
	}
	
	nextMongoUser, err := models.FindUser(interaction, &mongoUser.NextDiscordUserID)
	if err != nil {
		RespondWithError(discord, interaction, err)
		return
	}

	if err =  nextMongoUser.StartTurn(); err != nil {
		RespondWithError(discord, interaction, err)
		return
	}

	nextUserMessage := "It is your turn in the Infection game"
	if err = channel.SendUserMessage(discord, interaction, nextMongoUser.DiscordUserID, nextUserMessage); err != nil {
		RespondWithError(discord, interaction, err)
		return
	}

	response := &discordgo.InteractionResponseData{
		Content: "Your turn has ended",
		Flags: discordgo.MessageFlagsEphemeral,
	}

	discord.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: response,
	})
}