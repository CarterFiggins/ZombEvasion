package commands

import (
	"fmt"

	"infection/bot/role"
	"infection/bot/respond"
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
		SetOffAlarmDetails,
		StatusDetails,
		AttackDetails,
		LoadMapDetails,
	}

	CommandHandlers = map[string]func(discordS *discordgo.Session, interaction *discordgo.InteractionCreate){
		PingDetails.Name: Ping,
		SetupServerDetails.Name: SetupServer,
		JoinNextGameDetails.Name: JoinNextGame,
		StartGameDetails.Name: StartGame,
		EndGameDetails.Name: EndGame,
		MoveDetails.Name: Move,
		SetOffAlarmDetails.Name: SetOffAlarm,
		StatusDetails.Name: Status,
		AttackDetails.Name: Attack,
		LoadMapDetails.Name: LoadMap,
	}
)

func CheckPermissions(discord *discordgo.Session, interaction *discordgo.InteractionCreate, roles []string) bool {
	ok, err := role.UserHasRoles(discord, interaction, roles)
	if err != nil {
		respond.WithError(discord, interaction, err)
		return false
	}
	if !ok {
		respond.WithMessage(discord, interaction, "Unauthorized")
		return false
	}
	return true
}

func GetChannel(discord *discordgo.Session, interaction *discordgo.InteractionCreate, channelName string) *discordgo.Channel {
	channelMap, err := channel.CreateChannelMap(discord, interaction.Interaction.GuildID)
	if err != nil {
		respond.WithError(discord, interaction, err)
		return nil
	}
	channel, ok := channelMap[channelName]
	if !ok {
		respond.WithMessage(discord, interaction,  fmt.Sprintf("ERROR: %s not found. Try running `/setup-server`", channelName))
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
		respond.WithMessage(discord, interaction, fmt.Sprintf("This command only works in %s", channelName))
		return false
	}
	return true 
}
