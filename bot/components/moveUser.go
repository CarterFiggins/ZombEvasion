package components

import (
	"fmt"
	"strings"
	
	"infection/bot/game"
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

	content := "Something went wrong!"
	var gameMessage string
	response := &discordgo.WebhookEdit{
		Content: &content,
	}

	if mongoUser.IsAttacking {
		usersAttackedRoles, zombieUpgrade, err := game.AttackSector(discord, guildID, mongoUser, moveX, moveY)
		if err != nil {
			respond.EditWithError(discord, interaction, err)
			return
		}

		content = fmt.Sprintf("You Attacked %s. Missed!", sectorLocationName)
		if len(usersAttackedRoles) > 0 {
			content = fmt.Sprintf("You Attacked %s. Roles Attacked: %v", sectorLocationName, usersAttackedRoles)
			if zombieUpgrade {
				content += "\nYou have been upgraded! You can now move 3 sectors"
			}
		}
		response.Content = &content
		gameMessage = fmt.Sprintf("Sector %s was Attacked!", hexSectors.GetHexName(moveX, moveY))
	} else {
		sector := hexagonGrid.Board.Grid[moveX][moveY]
		sectorName := sector.GetSectorName()

		turnMessage, userMessage, err := game.MovedOnSectorMessages(discord, interaction, sectorName, sectorLocationName, guildID)
		if err != nil {
			respond.EditWithError(discord, interaction, err)
			return
		}
		response.Content = &userMessage
		gameMessage = turnMessage
	}

	game.SendAlarm(discord, interaction, mongoUser, guildID, gameMessage)
	discord.InteractionResponseEdit(interaction.Interaction, response)
}