package main

import(
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
)


func main() {
	botToken, ok := os.LookupEnv("BOT_TOKEN")
	if !ok {
		log.Fatal("Must set Discord token as env variable: BOT_TOKEN")
	}

	discord, err := discordgo.New("Bot " + botToken)
	if err != nil {
		log.Fatal(err)
	}

	err = discord.Open()
	if err != nil {
		log.Fatalf("cannot open the session: %v", err)
	}
	
	defer discord.Close()
	
	discord.AddHandler(func(discord *discordgo.Session, r *discordgo.Ready) {
		log.Printf("Logged in as: %v#%v", discord.State.User.Username, discord.State.User.Discriminator)
	})

	log.Println("Adding commands...")
	command := &discordgo.ApplicationCommand{
		Name: "ping",
		Description: "ping the bot",
	}

	

	_, err = discord.ApplicationCommandCreate(discord.State.User.ID, "", command)
	if err != nil {
		log.Panicf("Cannot create '%v' command: %v", command.Name, err)
	}


	// stop := make(chan os.Signal, 1)
	// signal.Notify(stop, os.Interrupt)
	// log.Println("Press Ctrl+C to exit")
	// <-stop



}