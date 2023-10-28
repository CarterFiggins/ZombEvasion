package game

import (
	"infection/bot/channel"
	"infection/models"
	"github.com/bwmarrin/discordgo"
)

func SendAlarm(discord *discordgo.Session, interaction *discordgo.InteractionCreate, mongoUser *models.MongoUser, guildID, alarmMessage string) error {
	gameChannel, err := channel.GetChannel(discord, guildID, channel.InfectionGameChannelName)
	if err != nil {
		return err
	}

	if (alarmMessage != "") {
		models.AddAlertMessage(alarmMessage, guildID)
		_, err = discord.ChannelMessageSend(gameChannel.ID, alarmMessage)
		if err != nil {
			return err
		}
		NextTurn(discord, interaction, mongoUser, guildID)
	}
	return nil
}