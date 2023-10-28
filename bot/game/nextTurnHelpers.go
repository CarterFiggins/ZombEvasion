package game

import (
	"os"
	"fmt"

	"infection/models"
	"github.com/bwmarrin/discordgo"
)

func NextTurn(discord *discordgo.Session, interaction *discordgo.InteractionCreate, mongoUser *models.MongoUser, guildID string) error {
	if err := mongoUser.EndTurn(); err != nil {
		return err
	}

	nextMongoUser, err := models.FindUserByIDs(interaction, &mongoUser.NextDiscordUserID, &guildID)
	if err != nil {
		return err
	}

	if err =  nextMongoUser.StartTurn(); err != nil {
		return err
	}

	if err = SetUpUsersTurn(discord, guildID, nextMongoUser); err != nil {
		return err
	}

	return nil
}

func SetUpUsersTurn(discord *discordgo.Session, guildID string, nextMongoUser *models.MongoUser) error {
	userChannel, err := discord.UserChannelCreate(nextMongoUser.DiscordUserID)
	if err != nil {
		return err
	}

	buttons := []discordgo.MessageComponent{
		discordgo.Button{
			Label: "Move",
			CustomID: fmt.Sprintf("move-buttons_%s", guildID),
		},
	}

	if nextMongoUser.Role == models.Zombie {
		buttons = append(
			buttons, 
			discordgo.Button{
			Label: "Attack",
			CustomID: fmt.Sprintf("attack-buttons_%s", guildID),
			},
		)
	}

	components := []discordgo.MessageComponent{
		discordgo.ActionsRow{
			Components: buttons,
		},
	}

	file, err := os.Open("./gameBoard.png")
	if err != nil {
		return err
	} 

	lastAlerts, err := models.LastRoundAlertMessages(guildID)
	if err != nil {
		return err
	} 
	
	lastAlertsMessage := ""
	for _, alert := range lastAlerts {
		lastAlertsMessage += fmt.Sprintf("%s\n", alert)
	}
	nextUserMessage := fmt.Sprintf("It is your turn in the Infection game!\nCurrent Location: %s\n%s", nextMongoUser.Location.GetHexName(), lastAlertsMessage)
	messageSend := &discordgo.MessageSend{
		Content: nextUserMessage,
		File: &discordgo.File{
			ContentType: "image/png",
			Name: "gameBoard.png",
			Reader: file,
		},
		Components: components,
	}

	_, err = discord.ChannelMessageSendComplex(userChannel.ID, messageSend)
	if err != nil {
		return err
	}

	return nil
}
