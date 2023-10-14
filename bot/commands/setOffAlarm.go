package commands

import (
	"fmt"
	"strings"

	"infection/models"
	"infection/hexagonGrid/hexSectors"
	"infection/bot/game"
	"infection/bot/channel"
	"github.com/bwmarrin/discordgo"
)

var (
	SetOffAlarmDetails = &discordgo.ApplicationCommand{
		Name: "set-off-alarm",
		Description: "set off alarm in a sector",
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

func SetOffAlarm(discord *discordgo.Session, interaction *discordgo.InteractionCreate) {
	mongoUser, err := models.FindUser(interaction, nil)
	if err != nil {
		RespondWithError(discord, interaction, err)
		return
	}

	if !mongoUser.TurnActive {
		RespondWithMessage(discord, interaction, "It is not your turn")
		return
	}

	if !mongoUser.CanSetOffAlarm {
		RespondWithMessage(discord, interaction, "You are not able to set off alarm")
		return
	}

	options := interaction.ApplicationCommandData().Options
	letter := strings.ToUpper(options[0].Value.(string))
	setOffY := int(options[1].Value.(float64)) - 1
	setOffX := hexSectors.LetterToNumMap[letter]

	if ok := game.CanSetOffAlarmHere(setOffX, setOffY); !ok {
		RespondWithMessage(discord, interaction, "The sector you selected is either off the grid or not a dangerous sector")
		return
	}

	alertsChannel := GetChannel(discord, interaction, channel.Alerts)
	if alertsChannel == nil {
		return
	}

	message := fmt.Sprintf("ALERT! Alarm set off at %s", hexSectors.GetHexName(setOffX, setOffY))
	_, err = discord.ChannelMessageSend(alertsChannel.ID, message)
	if err != nil {
		RespondWithError(discord, interaction, err)
		return
	}

	if err = mongoUser.UpdateCanSetOffAlarm(false); err != nil {
		RespondWithError(discord, interaction, err)
		return
	}

	response := &discordgo.InteractionResponseData{
		Content: fmt.Sprintf("You set off the alarm at %s", hexSectors.GetHexName(setOffX, setOffY)),
		Flags: discordgo.MessageFlagsEphemeral,
	}


	discord.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: response,
	})
}