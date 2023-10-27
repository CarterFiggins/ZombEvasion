package components

import (
	"github.com/bwmarrin/discordgo"
)

var (
	ComponentsHandlers = map[string]func(discordS *discordgo.Session, interaction *discordgo.InteractionCreate){
		"attack": Attack,
	}
)
