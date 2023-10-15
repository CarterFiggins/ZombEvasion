package channel

import (
	"fmt"
	"errors"
	"os"

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
	_, ok = channelMap[Alerts] 
	if !ok {
		_, err = discord.GuildChannelCreateComplex(
			interaction.Interaction.GuildID,
			discordgo.GuildChannelCreateData{
				Name: Alerts,
				Type: discordgo.ChannelTypeGuildText,
				PermissionOverwrites: []*discordgo.PermissionOverwrite{
					everyonePermission,
					botPermissions,
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

func SendUserMessage(discord *discordgo.Session, interaction *discordgo.InteractionCreate, discordUserID, message string) error {
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

func ShowMap(discord *discordgo.Session, interaction *discordgo.InteractionCreate) error {
	channelMap, err := CreateChannelMap(discord, interaction)
	if err != nil {
		return err
	}
	channel, ok := channelMap[InfectionGameChannelName]
	if !ok {
		return errors.New(fmt.Sprintf("%s not found. Try running `/setup-server`", InfectionGameChannelName))
	}

	file, err := os.Open("./gameBoard.png")
	if err != nil {
		return	err
	} 

	_, err = discord.ChannelMessageSendComplex(channel.ID, &discordgo.MessageSend{
		Files: []*discordgo.File{
			{
				ContentType: "image/png",
				Name: "gameBoard.png",
				Reader: file,
			},
		},
	})
	if err != nil {
		return	err
	} 

	return nil
}

func GetChannel(discord *discordgo.Session, interaction *discordgo.InteractionCreate, channelName string) (*discordgo.Channel, error) {
	channelMap, err := CreateChannelMap(discord, interaction)
	if err != nil {
		return nil, err
	}
	channel, ok := channelMap[channelName]
	if !ok {
		return nil, errors.New(fmt.Sprintf("ERROR: %s not found. Try running `/setup-server`", channelName))
	}

	return channel, nil
}
