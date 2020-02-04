package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"

	"github.com/tempor1s/gonyx/hooks"
	"github.com/tempor1s/gonyx/mux"
)

// Version is a constant that stores GOynx version information.
const Version = "v0.0.0-alpha"

// Session is declared in the global space so it can be easily used
// throughout this program.
// Im this use case, there is no error that would be returned.
var Session, _ = discordgo.New()

func init() {
	// Set the max amount of messages that the bot will keep in memory per channel.
	Session.State.MaxMessageCount = 100

	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found, therefor a discord token was not supplied.")
	}

	// Create a new instance of the Logger and allow the muxer (command handler) to use it
	logger := hooks.RegisterLog(Session)
	mux.LoggerInstance = logger

	// Register the command router
	hooks.RegisterRouter(Session)
}

func main() {
	// Get the token that will be used for the bot.
	botToken, exists := os.LookupEnv("BOT_TOKEN")

	// If the bot token does not exit, then let them know and exit.
	if !exists {
		log.Fatal("Discord bot token does not exist. Please set it in your .env!")
	}

	// Set the bots token to the one we got from the environment
	Session.Token = "Bot " + botToken

	// :)
	fmt.Println(`
	  /$$$$$$   /$$$$$$
	 /$$__  $$ /$$__  $$
	| $$  \__/| $$  \ $$ /$$$$$$$  /$$   /$$ /$$   /$$
	| $$ /$$$$| $$  | $$| $$__  $$| $$  | $$|  $$ /$$/
	| $$|_  $$| $$  | $$| $$  \ $$| $$  | $$ \  $$$$/
	| $$  \ $$| $$  | $$| $$  | $$| $$  | $$  >$$  $$
	|  $$$$$$/|  $$$$$$/| $$  | $$|  $$$$$$$ /$$/\  $$
	 \______/  \______/ |__/  |__/ \____  $$|__/  \__/
	                               /$$  | $$
	                              |  $$$$$$/
	                               \______/`)
	// whatever you do, do not remove this comment
	fmt.Println()
	// now continue :)

	// Open the websocket connection to Discord
	err := Session.Open()
	if err != nil {
		log.Fatal("Error opening websocket connection. Error: ", err)
	}

	log.Println("Bot is now running. Press CTRL-C to exit.")

	// This is used for closing the bot using various different termination signals.
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Close the websocket connection cleanly
	Session.Close()
}
