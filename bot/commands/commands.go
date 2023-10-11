package commands

import (
	"fmt"

	"infection/bot/role"
	"github.com/bwmarrin/discordgo"
)

var (
	Commands = []*discordgo.ApplicationCommand{
		PingDetails,
		SetupServerDetails,
		JoinNextGameDetails,
		StartGameDetails,
		EndGameDetails,
		MoveDetails,
	}

	CommandHandlers = map[string]func(discordS *discordgo.Session, interaction *discordgo.InteractionCreate){
		PingDetails.Name: Ping,
		SetupServerDetails.Name: SetupServer,
		JoinNextGameDetails.Name: JoinNextGame,
		StartGameDetails.Name: StartGame,
		EndGameDetails.Name: EndGame,
		MoveDetails.Name: Move,
	}
)

func CheckPermissions(discord *discordgo.Session, interaction *discordgo.InteractionCreate, roles []string) bool {
	ok, err := role.UserHasRoles(discord, interaction, roles)
	if err != nil {
		response := &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("ERROR: %v", err),
			Flags: discordgo.MessageFlagsEphemeral,
		}
		discord.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: response,
		})
		return false
	}
	if !ok {
		response := &discordgo.InteractionResponseData{
			Content: "Unauthorized",
			Flags: discordgo.MessageFlagsEphemeral,
		}
		discord.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: response,
		})
		return false
	}
	return true
}

func RespondWithError(discord *discordgo.Session, interaction *discordgo.InteractionCreate, err error) {
	discord.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("ERROR: %v", err),
			Flags: discordgo.MessageFlagsEphemeral,
		},
	})
}

func RespondEditWithError(discord *discordgo.Session, interaction *discordgo.InteractionCreate, err error) {
	message = fmt.Sprintf("ERROR: %v", err)
	discord.InteractionResponseEdit(interaction.Interaction, &discordgo.WebhookEdit{
		Content: &message,
	})
}
