package channel

import (
	"infection/bot/role"
	"github.com/bwmarrin/discordgo"
)

const (
	InfectionGameChannelName string = "infection-game"
	Alerts = "alerts"
)

func SetUpChannels(discord *discordgo.Session, interaction *discordgo.InteractionCreate) error {
	channelMap, err := CreateChannelMap(discord, interaction)
	if err != nil {
		return err
	}

	roleMap, err := role.CreateRoleMap(discord, interaction)

	everyonePermission := &discordgo.PermissionOverwrite {
		ID: roleMap["@everyone"].ID,
		Type: discordgo.PermissionOverwriteTypeRole,
		Deny: discordgo.PermissionAllText,
		Allow: discordgo.PermissionViewChannel,
	}

	adminPermission := &discordgo.PermissionOverwrite {
		ID: roleMap[role.Admin].ID,
		Type: discordgo.PermissionOverwriteTypeRole,
		Deny: 0,
		Allow: discordgo.PermissionAllText,
	}

	inGamePermission := &discordgo.PermissionOverwrite {
		ID: roleMap[role.InGame].ID,
		Type: discordgo.PermissionOverwriteTypeRole,
		Deny: 0,
		Allow: discordgo.PermissionAllText,
	}

	_, ok := channelMap[InfectionGameChannelName] 
	if !ok {
		_, err = discord.GuildChannelCreateComplex(
			interaction.Interaction.GuildID,
			discordgo.GuildChannelCreateData{
				Name:InfectionGameChannelName,
				Type: discordgo.ChannelTypeGuildText,
				PermissionOverwrites: []*discordgo.PermissionOverwrite{
					inGamePermission,
					adminPermission,
					everyonePermission,
				},
			},
		)
		if err != nil {
			return err
		}
	}
	_, ok = channelMap[Alerts] 
	if !ok {
		_, err = discord.GuildChannelCreateComplex(
			interaction.Interaction.GuildID,
			discordgo.GuildChannelCreateData{
				Name: Alerts,
				Type: discordgo.ChannelTypeGuildText,
				PermissionOverwrites: []*discordgo.PermissionOverwrite{
					everyonePermission,
					adminPermission,
				},
			},
		)
		if err != nil {
			return err
		}
	}

	return nil
}

func CreateChannelMap(discord *discordgo.Session, interaction *discordgo.InteractionCreate) (map[string]*discordgo.Channel, error){
	channelMap := make(map[string]*discordgo.Channel)
	channels, err := discord.GuildChannels(interaction.Interaction.GuildID)
	if (err != nil) {
		return channelMap, err
	}

	for _, channel := range channels {
		channelMap[channel.Name] = channel
	}

	return channelMap, nil
}
