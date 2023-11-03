package components

import (
	"fmt"
	"strings"

	"infection/bot/respond"
	"infection/bot/game"
	"infection/hexagonGrid/hexSectors"
	"infection/models"
	"github.com/bwmarrin/discordgo"
)

func SetOffAlarmText(discord *discordgo.Session, interaction *discordgo.InteractionCreate) {
	customID := interaction.MessageComponentData().CustomID
	splitID := strings.Split(customID, "_")
	guildID := splitID[1]

	mongoUser, err := models.FindUserByIDs(interaction, nil, &guildID)
	if err != nil {
		respond.WithError(discord, interaction, err)
		return
	}

	if !mongoUser.CanSetOffAlarm {
		respond.WithMessage(discord, interaction, "You can not set off the alarm right now")
		return
	}

	response := &discordgo.InteractionResponseData{
		Title: "Set Off Alarm",
		CustomID: fmt.Sprintf("set-off-alarm_%s", guildID),
		Components: []discordgo.MessageComponent{
			discordgo.ActionsRow{
				Components:  []discordgo.MessageComponent{
					discordgo.TextInput{
						CustomID: "sector",
						Label: "Set Off Alarm",
						Style: discordgo.TextInputShort,
						Placeholder: "A01",
						Required: true,
						MinLength: 3,
						MaxLength: 3,
					},
				},
			},
		},
	}

	discord.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseModal,
		Data: response,
	})
}

func SetOffAlarm(discord *discordgo.Session, interaction *discordgo.InteractionCreate) {
	customID := interaction.ModalSubmitData().CustomID
	splitID := strings.Split(customID, "_")
	guildID := splitID[1]
	value := interaction.ModalSubmitData().Components[0].(*discordgo.ActionsRow).Components[0].(*discordgo.TextInput).Value
	
	discord.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
	})

	mongoUser, err := models.FindUserByIDs(interaction, nil, &guildID)
	if err != nil {
		respond.WithError(discord, interaction, err)
		return
	}

	if !mongoUser.CanSetOffAlarm {
		respond.WithMessage(discord, interaction, "You can not set off the alarm right now")
		return
	}

	setOffX, setOffY, err := hexSectors.GetColAndRowFromName(value)
	if ok := game.CanSetOffAlarmHere(setOffX, setOffY); !ok {
		respond.EditWithMessage(discord, interaction, "The sector you selected is either off the grid or not a dangerous sector. Try again")
		return
	}


	if err = mongoUser.UpdateCanSetOffAlarm(false); err != nil {
		respond.WithError(discord, interaction, err)
		return
	}


	message := fmt.Sprintf("ALERT! Alarm set off at %s", hexSectors.GetHexName(setOffX, setOffY))
	game.SendAlarm(discord, interaction, mongoUser, guildID, message)


	if err = mongoUser.UpdateCanSetOffAlarm(false); err != nil {
		respond.WithError(discord, interaction, err)
		return
	}
	content := fmt.Sprintf("You set off the alarm at %s", hexSectors.GetHexName(setOffX, setOffY))
	response := &discordgo.WebhookEdit{
		Content: &content,
	}

	discord.InteractionResponseEdit(interaction.Interaction, response)
}