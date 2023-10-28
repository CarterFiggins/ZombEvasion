package respond

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

func WithError(discord *discordgo.Session, interaction *discordgo.InteractionCreate, err error) {
	log.Printf("ERROR: %v\n", err)
	discord.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("ERROR: %v", err),
			Flags: discordgo.MessageFlagsEphemeral,
		},
	})
}

func EditWithError(discord *discordgo.Session, interaction *discordgo.InteractionCreate, err error) {
	message := fmt.Sprintf("ERROR: %v", err)
	log.Printf("ERROR: %v\n", err)
	discord.InteractionResponseEdit(interaction.Interaction, &discordgo.WebhookEdit{
		Content: &message,
	})
}

func EditWithMessage(discord *discordgo.Session, interaction *discordgo.InteractionCreate, message string) {
	discord.InteractionResponseEdit(interaction.Interaction, &discordgo.WebhookEdit{
		Content: &message,
	})
}

func WithMessage(discord *discordgo.Session, interaction *discordgo.InteractionCreate, message string) {
	discord.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: message,
			Flags: discordgo.MessageFlagsEphemeral,
		},
	})
}
