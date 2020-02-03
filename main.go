package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

// Version is a constant that stores GOynx version information.
const Version = "v0.0.0-alpha"

// Session is declared in the global space so it can be easily used
// throughout this program.
// Im this use case, there is no error that would be returned.
var Session, _ = discordgo.New()

func init() {
	Session.State.MaxMessageCount = 100

	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found, therefor a discord token was not supplied.")
	}
}

func main() {
	botToken, exists := os.LookupEnv("BOT_TOKEN")

	if !exists {
		log.Fatal("Discord bot token does not exist. Please set it in your .env!")
	}

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

	err := Session.Open()
	if err != nil {
		log.Fatal("Error opening websocket connection. Error: ", err)
	}

	log.Println("Bot is now running. Press CTRL-C to exit.")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	Session.Close()
}
