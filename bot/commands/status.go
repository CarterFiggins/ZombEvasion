package commands

import (
	"fmt"

	"infection/bot/game"
	"infection/bot/respond"
	"infection/models"
	"github.com/bwmarrin/discordgo"
)

var StatusDetails = &discordgo.ApplicationCommand{
	Name: "status",
	Description: "gives a status of player",
}

func Status(discord *discordgo.Session, interaction *discordgo.InteractionCreate) {
	mongoUser, err := models.FindUser(interaction)
	if err != nil {
		respond.WithError(discord, interaction, err)
		return
	}
	
	sectorsToMove := game.GetMoveSectors(mongoUser)

	response := &discordgo.InteractionResponseData{
		Content: fmt.Sprintf("Role: %s\nCurrent Position: %s\nYour Turn: %t\nMax Moves: %d\nNext Possible Moves:%v", mongoUser.Role, mongoUser.Location.GetHexName(), mongoUser.TurnActive, mongoUser.MaxMoves, sectorsToMove),
		Flags: discordgo.MessageFlagsEphemeral,
	}

	discord.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: response,
	})
}
