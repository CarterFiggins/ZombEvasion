package commands

import (
	"github.com/bwmarrin/discordgo"
)

var (
	Commands = []*discordgo.ApplicationCommand{
		PingDetails,
		SetupServerDetails,
		JoinNextGameDetails,
	}

	CommandHandlers = map[string]func(discordS *discordgo.Session, interaction *discordgo.InteractionCreate){
		"ping": Ping,
		"setup-server": SetupServer,
		"join-next-game": JoinNextGame,
	}
)
