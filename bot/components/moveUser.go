package components

import (
	"strings"
	
	"infection/bot/game"
	"infection/bot/channel"
	"infection/bot/respond"
	"infection/models"
	"infection/hexagonGrid"
	"infection/hexagonGrid/hexSectors"
	"github.com/bwmarrin/discordgo"
)

func MoveUser(discord *discordgo.Session, interaction *discordgo.InteractionCreate) {
	customID := interaction.MessageComponentData().CustomID
	splitID := strings.Split(customID, "_")
	guildID := splitID[1]
	sectorLocationName := splitID[2]

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

	message, err := game.CanUserMoveHere(discord, interaction, sectorLocationName, mongoUser)
	if err != nil {
		respond.EditWithError(discord, interaction, err)
		return
	}
	if message != nil {
		respond.EditWithMessage(discord, interaction, *message)
		return
	}

	moveX, moveY, err := hexSectors.GetColAndRowFromName(sectorLocationName)
	if err != nil {
		respond.EditWithError(discord, interaction, err)
		return
	}
	err = mongoUser.MoveUser(moveX, moveY)
	if err != nil {
		respond.EditWithError(discord, interaction, err)
		return
	}

	sector := hexagonGrid.Board.Grid[moveX][moveY]
	sectorName := sector.GetSectorName()
	gameChannel, err := channel.GetChannel(discord, guildID, channel.InfectionGameChannelName)
	if err != nil {
		respond.EditWithError(discord, interaction, err)
		return
	}
	turnMessage, userMessage, err := game.MovedOnSectorMessages(discord, interaction, sectorName, sectorLocationName, guildID)
	if err != nil {
		respond.EditWithError(discord, interaction, err)
		return
	}

	response := &discordgo.WebhookEdit{
		Content: &userMessage,
	}

	if (turnMessage != "") {
		_, err = discord.ChannelMessageSend(gameChannel.ID, turnMessage)
		if err != nil {
			respond.EditWithError(discord, interaction, err)
			return
		}
	}



	discord.InteractionResponseEdit(interaction.Interaction, response)

}