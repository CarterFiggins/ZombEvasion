package channel

import (
	"fmt"
	"errors"

	"infection/bot/role"
	"github.com/bwmarrin/discordgo"
)

const (
	InfectionGameChannelName string = "infection-game"
)

func SetUpChannels(discord *discordgo.Session, interaction *discordgo.InteractionCreate) error {
	channelMap, err := CreateChannelMap(discord, interaction.Interaction.GuildID)
	if err != nil {
		return err
	}

	roleMap, err := role.CreateRoleMap(discord, interaction)

	everyonePermission := &discordgo.PermissionOverwrite {
		ID: roleMap["@everyone"].ID,
		Type: discordgo.PermissionOverwriteTypeRole,
		Deny: discordgo.PermissionAllText,
		Allow: discordgo.PermissionViewChannel | discordgo.PermissionReadMessageHistory,
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

	botPermissions := &discordgo.PermissionOverwrite {
		ID: roleMap[role.Bot].ID,
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
					botPermissions,
				},
			},
		)
		if err != nil {
			return err
		}
	}

	return nil
}

func CreateChannelMap(discord *discordgo.Session, guildID string) (map[string]*discordgo.Channel, error){
	channelMap := make(map[string]*discordgo.Channel)
	channels, err := discord.GuildChannels(guildID)
	if (err != nil) {
		return channelMap, err
	}

	for _, channel := range channels {
		channelMap[channel.Name] = channel
	}

	return channelMap, nil
}

func SendUserMessage(discord *discordgo.Session, discordUserID, message string) error {
	userChannel, err := discord.UserChannelCreate(discordUserID)
	if err != nil {
		return err
	}

	_, err = discord.ChannelMessageSend(userChannel.ID, message)
	if err != nil {
		return err
	}

	return nil
}

func GetChannel(discord *discordgo.Session, guildID string, channelName string) (*discordgo.Channel, error) {
	channelMap, err := CreateChannelMap(discord, guildID)
	if err != nil {
		return nil, err
	}
	channel, ok := channelMap[channelName]
	if !ok {
		return nil, errors.New(fmt.Sprintf("%s not found. Try running `/setup-server`", channelName))
	}

	return channel, nil
}
