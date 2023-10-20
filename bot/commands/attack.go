package commands

import (
	"fmt"
	"strings"

	"infection/models"
	"infection/hexagonGrid/hexSectors"
	"infection/hexagonGrid"
	"infection/bot/game"
	"infection/bot/channel"
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
	mongoUser, err := models.FindUser(interaction, nil)
	if err != nil {
		RespondWithError(discord, interaction, err)
		return
	}

	if mongoUser.Role != models.Zombie{
		RespondWithMessage(discord, interaction, "You are not a zombie! You can not attack")
		return
	}

	if !mongoUser.TurnActive {
		RespondWithMessage(discord, interaction, "It is not your turn")
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

	message, err := game.CanUserMoveHere(discord, interaction, attackX, attackY, mongoUser);
	if err != nil {
		RespondEditWithError(discord, interaction, err)
		return
	}
	if message != nil {
		RespondEditWithMessage(discord, interaction, *message)
		return
	} else {
		err = mongoUser.MoveUser(attackX, attackY)
		if err != nil {
			RespondEditWithError(discord, interaction, err)
			return
		}

		usersAttacked, err := models.FindUsersAtLocation(interaction, attackX, attackY)

		var usersAttackedRoles []string
		zombieUpgrade := false
		
		for _, user := range usersAttacked {
			if (user.DiscordUserID == mongoUser.DiscordUserID) {
				// Don't attack self
				continue
			}
			usersAttackedRoles = append(usersAttackedRoles, user.Role)
			zombieSectorCol := hexagonGrid.Board.ZombieSector.Col
			zombieSectorRow := hexagonGrid.Board.ZombieSector.Row

			if (user.Role == models.Human) {
				if (mongoUser.MaxMoves == 2) {
					zombieUpgrade = true
					mongoUser.UpgradeUsersMaxMoves(3)
				}
				attackedMessage := fmt.Sprintf("You have been bitten by a zombie! You have Respawned as a zombie at %s", hexSectors.GetHexName(zombieSectorCol, zombieSectorRow))
				if err = channel.SendUserMessage(discord, interaction, user.DiscordUserID, attackedMessage); err != nil {
					RespondEditWithError(discord, interaction, err)
					return
				}

				if err = user.TurnIntoZombie(); err != nil {
					RespondEditWithError(discord, interaction, err)
					return
				}

				if err = game.CheckGame(discord, interaction); err != nil {
					RespondEditWithError(discord, interaction, err)
					return
				} 
				
			} else if (user.Role == models.Zombie) {
				attackedMessage := fmt.Sprintf("A zombie mistaken you as a human and attacked you! You have Respawned at %s", hexSectors.GetHexName(zombieSectorCol, zombieSectorRow))
				if err = channel.SendUserMessage(discord, interaction, user.DiscordUserID, attackedMessage); err != nil {
					RespondEditWithError(discord, interaction, err)
					return
				}
			}
			if err = user.RespawnUser(zombieSectorCol, zombieSectorRow); err != nil {
				RespondEditWithError(discord, interaction, err)
				return
			}
		}

		content = "Missed!"
		if len(usersAttackedRoles) > 0 {
			content = fmt.Sprintf("You Attacked: %v", usersAttackedRoles)
			if zombieUpgrade {
				content += "\nYou have been upgraded! You can now move 3 sectors"
			}
		}
		response.Content = &content

		alertsChannel, err := channel.GetChannel(discord, interaction, channel.Alerts)
		if err != nil {
			RespondEditWithError(discord, interaction, err)
			return
		}

		_, err = discord.ChannelMessageSend(alertsChannel.ID, fmt.Sprintf("Sector %s was Attacked!", hexSectors.GetHexName(attackX, attackY)))
		if err != nil {
			RespondEditWithError(discord, interaction, err)
			return
		}

	}

	discord.InteractionResponseEdit(interaction.Interaction, response)

	game.NextTurn(discord, interaction, mongoUser)
}