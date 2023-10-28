package components

import (
	"strings"
	"strconv"
	
	"infection/bot/respond"
	"infection/bot/game"
	"infection/models"
	"github.com/bwmarrin/discordgo"
)

func MaxMove(discord *discordgo.Session, interaction *discordgo.InteractionCreate) {
	customID := interaction.MessageComponentData().CustomID
	splitID := strings.Split(customID, "_")
	guildID := splitID[1]
	moveDistance, err := strconv.Atoi(splitID[2])
	if err != nil {
		respond.WithError(discord, interaction, err)
		return
	}

	mongoUser, err := models.FindUserByIDs(interaction, nil, &guildID)
	if err != nil {
		respond.WithError(discord, interaction, err)
		return
	}

	if !mongoUser.TurnActive {
		respond.WithMessage(discord, interaction, "It is not your turn")
		return
	}

	if !mongoUser.CanMove {
		respond.WithMessage(discord, interaction, "You have already moved this turn")
		return
	}

	discord.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Flags: discordgo.MessageFlagsEphemeral,
		},
	})

	game.RespondWithDistanceButtons(discord, interaction, mongoUser, guildID, moveDistance)
}