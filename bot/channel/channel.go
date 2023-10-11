package channel

import (
	"github.com/bwmarrin/discordgo"
)

const (
	InfectionGameChannelName string = "infection-game"
)

func SetUpChannels(discord *discordgo.Session, interaction *discordgo.InteractionCreate) error {
	channelMap, err := CreateChannelMap(discord, interaction)
	if err != nil {
		return err
	}

	_, ok := channelMap[InfectionGameChannelName] 
	if !ok {
		_, err = discord.GuildChannelCreate(
			interaction.Interaction.GuildID,
			InfectionGameChannelName,
			discordgo.ChannelTypeGuildText,
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
