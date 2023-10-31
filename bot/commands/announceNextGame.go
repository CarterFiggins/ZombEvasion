package commands

import (
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
						CustomID: "joinQueue",
					},
					discordgo.Button{
						Label: "Leave Queue",
						CustomID: "leaveQueue",
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