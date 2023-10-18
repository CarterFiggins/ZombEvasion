package bot

import (
	"log"
	"os"
	"os/signal"
	"strings"
	"math/rand"
	"time"
	
	"infection/bot/commands"
	"infection/hexagonGrid"
	"github.com/bwmarrin/discordgo"
)

func Run() {
	log.Println("Setting up discord bot...")
	botToken, ok := os.LookupEnv("BOT_TOKEN")
	if !ok {
		log.Fatal("Must set Discord token as env variable: BOT_TOKEN")
	}
	
	discord, err := discordgo.New("Bot " + botToken)
	if err != nil {
		log.Fatal(err)
	}

	// Add event handler for general messages
	discord.AddHandler(newMessage)

	// Add event handler for commands
	discord.AddHandler(func(discord *discordgo.Session, interaction *discordgo.InteractionCreate) {
		if interaction.Interaction.Member == nil {
			discord.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "DM commands are turned off",
					Flags: discordgo.MessageFlagsEphemeral,
				},
			})
			return
		}

		if handler, ok := commands.CommandHandlers[interaction.ApplicationCommandData().Name]; ok {
			handler(discord, interaction)
		}
	})
	
	// Open session
	discord.Open()
	defer discord.Close()


	hexagonGrid.Board.LoadBoard()

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
		discord.ChannelMessageSend(message.ChannelID, randomBotGif())
	}
}

func randomBotGif() string {
	gifs := []string{
		"https://tenor.com/Kaca.gif",
		"https://tenor.com/bSvv8.gif",
		"https://tenor.com/bVEQE.gif",
		"https://tenor.com/bFQTS.gif",
		"https://tenor.com/bfxl4.gif",
		"https://tenor.com/bkrQOwoAa7h.gif",
		"https://tenor.com/t60i.gif",
		"https://tenor.com/bE1Qx.gif",
		"https://tenor.com/xK1M.gif",
		"https://tenor.com/bkV6J.gif",
		"https://tenor.com/bdbfH.gif",
		"https://tenor.com/bHUoo.gif",
		"https://tenor.com/bUHlU.gif",
		"https://tenor.com/bM5nL.gif",
		"https://tenor.com/bRkE8.gif",
		"https://tenor.com/bwK3O.gif",
		"https://tenor.com/bcOeZ.gif",
		"https://tenor.com/bnYU9.gif",
		"https://tenor.com/bgvYq.gif",
		"https://tenor.com/bwKzq.gif",
		"https://tenor.com/wBGt.gif",
		"https://tenor.com/buiAL.gif",
		"https://tenor.com/uMFg.gif",
		"https://tenor.com/Mjar.gif",
		"https://tenor.com/bj0ti.gif",
		"https://tenor.com/vAus.gif",
		"https://tenor.com/bodTH.gif",
		"https://tenor.com/bvYx7.gif",
	}

	rand.Seed(time.Now().UnixNano())
	randNum := rand.Intn(len(gifs))
	return gifs[randNum]
}