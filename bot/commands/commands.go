package commands

import (
	"github.com/bwmarrin/discordgo"
)

var (
	Commands = []*discordgo.ApplicationCommand{
		{
			Name: "ping",
			Description: "ping the bot",
		},
		{
			Name: "setup-server",
			Description: "adds the discord roles",
		},
	}

	CommandHandlers = map[string]func(discordS *discordgo.Session, interaction *discordgo.InteractionCreate){
		"ping": Ping,
		"setup-server": SetupServer,
	}
)
