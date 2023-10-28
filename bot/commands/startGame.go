package commands

import (
	"os"

	"infection/bot/game"
	"infection/bot/respond"
	"infection/bot/role"
	"infection/bot/channel"
	"github.com/bwmarrin/discordgo"
)

var StartGameDetails = &discordgo.ApplicationCommand{
	Name: "start-game",
	Description: "changes roles of players waiting in queue and starts the game",
}

func StartGame(discord *discordgo.Session, interaction *discordgo.InteractionCreate) {
	if ok := CheckPermissions(discord, interaction, []string{role.Admin}); !ok {
		return
	}
	if ok := CheckChannel(discord, interaction, interaction.Interaction.ChannelID, channel.InfectionGameChannelName); !ok {
		return
	}

	discord.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
	})

	message := "Game Started"
	response := &discordgo.WebhookEdit{
		Content: &message,
	}

	err := game.Start(discord, interaction)
	if err != nil {
		respond.EditWithError(discord, interaction, err)
		return	
	} 
	file, err := os.Open("./gameBoard.png")
	if err != nil {
		respond.EditWithError(discord, interaction, err)
		return	
	} 
	response.Files = []*discordgo.File{
		{
			ContentType: "image/png",
			Name: "gameBoard.png",
			Reader: file,
		},
	}
	
	discord.InteractionResponseEdit(interaction.Interaction, response)
}