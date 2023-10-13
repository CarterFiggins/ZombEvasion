package commands

import (
	"infection/models"
	"github.com/bwmarrin/discordgo"
)

var PassTurnDetails = &discordgo.ApplicationCommand{
	Name: "pass-turn",
	Description: "Starts the next players turn",
}

func PassTurn(discord *discordgo.Session, interaction *discordgo.InteractionCreate) {
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
		RespondWithMessage(discord, interaction, "You have to move before you pass your turn")
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

	// send nextMongoUser a message saying it is there turn

	response := &discordgo.InteractionResponseData{
		Content: "Your turn has ended",
		Flags: discordgo.MessageFlagsEphemeral,
	}

	discord.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: response,
	})
}