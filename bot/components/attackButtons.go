package components

import (
	"strings"
	
	"infection/bot/respond"
	"infection/bot/game"
	"infection/models"
	"github.com/bwmarrin/discordgo"
)

func AttackButtons(discord *discordgo.Session, interaction *discordgo.InteractionCreate) {
	customID := interaction.MessageComponentData().CustomID
	guildID := strings.Split(customID, "_")[1]

	mongoUser, err := models.FindUserByIDs(interaction, nil, &guildID)
	if err != nil {
		respond.WithError(discord, interaction, err)
		return
	}

	if err := mongoUser.MarkAttacking(true); err != nil {
		respond.WithError(discord, interaction, err)
		return
	}

	if mongoUser.Role != models.Zombie{
		respond.WithMessage(discord, interaction, "You are not a zombie! You can not attack")
		return
	}

	if !mongoUser.TurnActive {
		respond.WithMessage(discord, interaction, "It is not your turn")
		return
	}

	if !mongoUser.CanMove {
		respond.WithMessage(discord, interaction, "You have already attacked or moved this turn")
		return
	}

	discord.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Flags: discordgo.MessageFlagsEphemeral,
		},
	})

	game.RespondWithLocationButtons(discord, interaction, mongoUser, guildID)
}