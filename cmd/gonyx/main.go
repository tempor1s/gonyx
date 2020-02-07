package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/tempor1s/gonyx/bot"
)

// Version is a constant that stores GOynx version information.
const Version = "v0.0.0-alpha"

func init() {
	// Initalize environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found, therefor a discord token was not supplied.")
	}
}

func main() {
	// Create a new bot instance and do all the stuff behind the scenes :)
	var Bot = bot.New()

	// Start the bot
	Bot.Start()

	// Close the websocket connection cleanly when a termination signal is given
	Bot.AwaitTermination()
}
