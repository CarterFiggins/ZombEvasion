package commands

import (
	"fmt"
	
	"github.com/bwmarrin/discordgo"
)

var AnnounceNextGameDetails = &discordgo.ApplicationCommand{
	Name: "announce-next-game",
	Description: "adds a button where users can join the queue",
}

func AnnounceNextGame(discord *discordgo.Session, interaction *discordgo.InteractionCreate) {
	response := &discordgo.InteractionResponseData{
		Content: "Starting New Game! Click Here to join the queue for the next game",
		Components: []discordgo.MessageComponent{
			discordgo.ActionsRow{
				Components: []discordgo.MessageComponent{
					discordgo.Button{
						Label: "Join Queue",
						CustomID: fmt.Sprintf("joinQueue_%s", interaction.Interaction.GuildID),
					},
					discordgo.Button{
						Label: "Leave Queue",
						CustomID: fmt.Sprintf("leaveQueue_%s", interaction.Interaction.GuildID),
					},
				},
			},
		},
	}

	discord.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: response,
	})
}