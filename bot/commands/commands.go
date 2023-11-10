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
		AnnounceNextGameDetails,
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
		AnnounceNextGameDetails.Name: AnnounceNextGame,
	}
)

func CheckPermissions(discord *discordgo.Session, interaction *discordgo.InteractionCreate, roles []string) bool {
	ok, err := role.UserHasOneRole(discord, interaction, roles)
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

func CheckChannel(discord *discordgo.Session, interaction *discordgo.InteractionCreate, channelID, channelName string) bool {
	channel, err := channel.GetChannel(discord, interaction.Interaction.GuildID, channelName)
	if err != nil {
		respond.EditWithError(discord, interaction, err)
		return false
	}

	if (channel.ID != channelID) {
		respond.WithMessage(discord, interaction, fmt.Sprintf("This command only works in %s", channelName))
		return false
	}
	return true
}
