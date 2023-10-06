package bot

import (
	"log"
	"os"
	"os/signal"
	"strings"
	
	"infection/bot/commands"
	"github.com/bwmarrin/discordgo"
)

var (
	BotToken string
)

func Run() {
	log.Println("Setting up discord bot...")
	
	discord, err := discordgo.New("Bot " + BotToken)
	if err != nil {
		log.Fatal(err)
	}

	// Add event handler for general messages
	discord.AddHandler(newMessage)

	// Add event handler for commands
	discord.AddHandler(func(discord *discordgo.Session, interaction *discordgo.InteractionCreate) {
		if handler, ok := commands.CommandHandlers[interaction.ApplicationCommandData().Name]; ok {
			handler(discord, interaction)
		}
	})
	
	// Open session
	discord.Open()
	defer discord.Close()

	// Run until code is terminated
	log.Println("Bot running...")
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop
	log.Println("Bot Stopped")
}

func newMessage(discord *discordgo.Session, message *discordgo.MessageCreate) {
	// Ignore bot message
	if message.Author.ID == discord.State.User.ID {
		return
	}

	// Respond to messages
	switch {
	case strings.Contains(message.Content, "bot"):
		discord.ChannelMessageSend(message.ChannelID, "Hello Human!")
	}

}