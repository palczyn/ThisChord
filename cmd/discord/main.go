package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/palczyn/phoebus/discord"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

var token string
var player *discord.Player

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("error while loading env file: %v", err)
		return
	}

	var ok bool
	token, ok = os.LookupEnv("API_TOKEN")

	if !ok || token == "" {
		fmt.Println("No token provided. Please ensure env is set.")
		return
	}

	// Load the sound file.
	player, err = discord.LoadSound()
	if err != nil {
		fmt.Println("Error loading sound: ", err)
		return
	}

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("Error creating Discord session: ", err)
		return
	}

	// Register ready as a callback for the ready events.
	dg.AddHandler(discord.Ready)

	// Register messageCreate as a callback for the messageCreate events.
	dg.AddHandler(discord.HandleMessageCreateWithPlayer(player))

	// Register guildCreate as a callback for the guildCreate events.
	dg.AddHandler(discord.GuildCreate)

	// We need information about guilds (which includes their channels),
	// messages and voice states.
	dg.Identify.Intents = discordgo.IntentsGuilds | discordgo.IntentsGuildMessages | discordgo.IntentsGuildVoiceStates

	// Open the websocket and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("Error opening Discord session: ", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Airhorn is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}
