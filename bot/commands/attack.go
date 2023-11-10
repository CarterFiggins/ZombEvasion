package commands

import (
	"fmt"
	"strings"

	"infection/models"
	"infection/bot/respond"
	"infection/hexagonGrid/hexSectors"
	"infection/bot/game"
	"github.com/bwmarrin/discordgo"
)

var (
	AttackDetails = &discordgo.ApplicationCommand{
		Name: "attack",
		Description: "move and attack sector",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type: discordgo.ApplicationCommandOptionString,
				Name: "letter-column",
				Description: "The column of the sector",
				Required: true,
			},
			{
				Type: discordgo.ApplicationCommandOptionInteger,
				Name: "row-number",
				Description: "The row of the sector",
				MinValue: &minValue,
				Required: true,
			},
		},
	}
)

func Attack(discord *discordgo.Session, interaction *discordgo.InteractionCreate) {
	mongoUser, err := models.FindUser(interaction)
	if err != nil {
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

	discord.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Flags: discordgo.MessageFlagsEphemeral,
		},
	})

	options := interaction.ApplicationCommandData().Options
	letter := strings.ToUpper(options[0].Value.(string))
	attackY := int(options[1].Value.(float64)) - 1
	attackX := hexSectors.LetterToNumMap[letter]

	content := "Can't attack here try again"
	response := &discordgo.WebhookEdit{
		Content: &content,
	}

	moveHexName := hexSectors.GetHexName(attackX, attackY)
	message, err := game.CanUserMoveHere(discord, interaction, moveHexName, interaction.Interaction.GuildID, mongoUser);
	if err != nil {
		respond.EditWithError(discord, interaction, err)
		return
	}
	if message != nil {
		respond.EditWithMessage(discord, interaction, *message)
		return
	} else {
		err = mongoUser.MoveUser(attackX, attackY)
		if err != nil {
			respond.EditWithError(discord, interaction, err)
			return
		}

		usersAttackedRoles, zombieUpgrade, err := game.AttackSector(discord, interaction.Interaction.GuildID, mongoUser, attackX, attackY)
		if err != nil {
			respond.EditWithError(discord, interaction, err)
			return
		}

		content = "Missed!"
		if len(usersAttackedRoles) > 0 {
			content = fmt.Sprintf("You Attacked: %v", usersAttackedRoles)
			if zombieUpgrade {
				content += "\nYou have been upgraded! You can now move 3 sectors"
			}
		}
		response.Content = &content
		game.SendAlarm(discord, interaction, mongoUser, interaction.Interaction.GuildID, fmt.Sprintf("Sector %s was Attacked!", hexSectors.GetHexName(attackX, attackY)))
	}

	discord.InteractionResponseEdit(interaction.Interaction, response)

}