package commands

import (
	"os"
	"fmt"

	"infection/bot/game"
	"infection/bot/respond"
	"infection/bot/role"
	"infection/bot/channel"
	"infection/hexagonGrid"
	"github.com/bwmarrin/discordgo"
)

var StartGameDetails = &discordgo.ApplicationCommand{
	Name: "start-game",
	Description: "changes roles of players waiting in queue and starts the game",
	Options: []*discordgo.ApplicationCommandOption{
		{
			Type: discordgo.ApplicationCommandOptionString,
			Name: "board-name",
			Description: "Select the board you would like to play with",
			Required: true,
			Choices: []*discordgo.ApplicationCommandOptionChoice{
				{
					Name: hexagonGrid.ForestBoardName,
					Value: hexagonGrid.ForestBoardName,
				},
				{
					Name: hexagonGrid.GraveYardBoardName,
					Value: hexagonGrid.GraveYardBoardName,
				},
				{
					Name: hexagonGrid.HospitalBoardName,
					Value: hexagonGrid.HospitalBoardName,
				},
			},
		},
	},
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

	boardName := interaction.ApplicationCommandData().Options[0].Value.(string)
	message := "Game Started"
	response := &discordgo.WebhookEdit{
		Content: &message,
	}

	err := game.Start(discord, interaction)
	if err != nil {
		respond.EditWithError(discord, interaction, err)
		return	
	} 

	file, err := os.Open(fmt.Sprintf("./%s.png", boardName))
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