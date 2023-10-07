package commands

import (
	"github.com/bwmarrin/discordgo"
)

var (
	Commands = []*discordgo.ApplicationCommand{
		PingDetails,
		SetupServerDetails,
		JoinNextGameDetails,
		StartGameDetails,
		EndGameDetails,
	}

	CommandHandlers = map[string]func(discordS *discordgo.Session, interaction *discordgo.InteractionCreate){
		PingDetails.Name: Ping,
		SetupServerDetails.Name: SetupServer,
		JoinNextGameDetails.Name: JoinNextGame,
		StartGameDetails.Name: StartGame,
		EndGameDetails.Name: EndGame,
	}
)
