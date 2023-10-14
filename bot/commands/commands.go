package commands

import (
	"fmt"

	"infection/bot/role"
	"infection/bot/channel"
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
		EndTurnDetails,
		SetOffAlarmDetails,
	}

	CommandHandlers = map[string]func(discordS *discordgo.Session, interaction *discordgo.InteractionCreate){
		PingDetails.Name: Ping,
		SetupServerDetails.Name: SetupServer,
		JoinNextGameDetails.Name: JoinNextGame,
		StartGameDetails.Name: StartGame,
		EndGameDetails.Name: EndGame,
		MoveDetails.Name: Move,
		EndTurnDetails.Name: EndTurn,
		SetOffAlarmDetails.Name: SetOffAlarm,
	}
)

func CheckPermissions(discord *discordgo.Session, interaction *discordgo.InteractionCreate, roles []string) bool {
	ok, err := role.UserHasRoles(discord, interaction, roles)
	if err != nil {
		RespondWithError(discord, interaction, err)
		return false
	}
	if !ok {
		RespondWithMessage(discord, interaction, "Unauthorized")
		return false
	}
	return true
}

func GetChannel(discord *discordgo.Session, interaction *discordgo.InteractionCreate, channelName string) *discordgo.Channel {
	channelMap, err := channel.CreateChannelMap(discord, interaction)
	if err != nil {
		RespondWithError(discord, interaction, err)
		return nil
	}
	channel, ok := channelMap[channelName]
	if !ok {
		RespondWithMessage(discord, interaction,  fmt.Sprintf("ERROR: %s not found. Try running `/setup-server`", channelName))
		return nil
	}

	return channel
}

func CheckChannel(discord *discordgo.Session, interaction *discordgo.InteractionCreate, channelID, channelName string) bool {
	channel := GetChannel(discord, interaction, channelName)
	if channel == nil {
		return false
	}

	if (channel.ID != channelID) {
		RespondWithMessage(discord, interaction, fmt.Sprintf("This command only works in %s", channelName))
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
	message := fmt.Sprintf("ERROR: %v", err)
	discord.InteractionResponseEdit(interaction.Interaction, &discordgo.WebhookEdit{
		Content: &message,
	})
}

func RespondWithMessage(discord *discordgo.Session, interaction *discordgo.InteractionCreate, message string) {
	discord.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: message,
			Flags: discordgo.MessageFlagsEphemeral,
		},
	})
}
