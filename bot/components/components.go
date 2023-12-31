package components

import (
	"github.com/bwmarrin/discordgo"
)

var (
	ComponentsHandlers = map[string]func(discordS *discordgo.Session, interaction *discordgo.InteractionCreate){
		"attack-buttons": AttackButtons,
		"move-buttons": MoveButtons,
		"max-move": MaxMove,
		"move-user": MoveUser,
		"joinQueue": JoinNextGame,
		"leaveQueue": LeaveQueue,
		"set-off-alarm-button": SetOffAlarmText,
		"set-off-alarm": SetOffAlarm,
	}
)
