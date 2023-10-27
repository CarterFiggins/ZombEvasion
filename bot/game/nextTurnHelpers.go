package game

import (
	"os"
	"fmt"

	"infection/models"
	"infection/bot/channel"
	"github.com/bwmarrin/discordgo"
)

func NextTurn(discord *discordgo.Session, interaction *discordgo.InteractionCreate, mongoUser *models.MongoUser) error {
	if err := mongoUser.EndTurn(); err != nil {
		return err
	}

	nextMongoUser, err := models.FindUserByIDs(interaction, &mongoUser.NextDiscordUserID, nil)
	if err != nil {
		return err
	}

	if err =  nextMongoUser.StartTurn(); err != nil {
		return err
	}

	if err = SetUpUsersTurn(discord, interaction, nextMongoUser); err != nil {
		return err
	}

	if err = channel.ShowMap(discord, interaction); err != nil {
		return err
	}
	return nil
}

func SetUpUsersTurn(discord *discordgo.Session, interaction *discordgo.InteractionCreate, nextMongoUser *models.MongoUser) error {
	userChannel, err := discord.UserChannelCreate(nextMongoUser.DiscordUserID)
	if err != nil {
		return err
	}

	guildID := interaction.Interaction.GuildID

	buttons := []discordgo.MessageComponent{
		discordgo.Button{
			Label: "Move",
			CustomID: fmt.Sprintf("move_%s", guildID),
		},
	}

	if nextMongoUser.Role == models.Zombie {
		buttons = append(
			buttons, 
			discordgo.Button{
			Label: "Attack",
			CustomID: fmt.Sprintf("attack_%s", guildID),
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

	nextUserMessage := fmt.Sprintf("It is your turn in the Infection game!\nCurrent Location: %s", nextMongoUser.Location.GetHexName())
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
