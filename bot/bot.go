package bot

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/tempor1s/gonyx/logger"
	"github.com/tempor1s/gonyx/mux"
	"github.com/tempor1s/gonyx/tasks"
)

// Bot represents everything that has to do with the bot. Has different modules, the session, and soon the DB
type Bot struct {
	Logger     *logger.Logger
	Mux        *mux.Mux
	Session    *discordgo.Session
	WeeklyInfo *tasks.WeeklyInfo
}

// New creates a new bot instance and session and does some config stuff
func New() *Bot {
	// Create a new discordgo session.
	session, _ := discordgo.New()
	// Set the max amount of messages that the bot will keep in memory per channel.
	session.State.MaxMessageCount = 100
	// Create a new bot instance to return
	bot := &Bot{Session: session}
	// Set the bots token
	bot.setToken()
	// Register the handlers and tasks
	bot.registerHandlers()
	// Set the Mux's logger instance to be that of the session
	mux.LoggerInstance = bot.Logger // TODO: clean this up :(

	return &Bot{Session: session}
}

// registerHandlers registers the different bot handlers
func (b *Bot) registerHandlers() {
	b.registerRouter()
	b.registerLog()
	b.registerTasks()
}

// setToken will set the bots token from the environment
func (b *Bot) setToken() {
	botToken, exists := os.LookupEnv("BOT_TOKEN")

	if exists == false {
		log.Fatal("Discord bot token does not exist. Please set it in your .env!")
	}

	b.Session.Token = "Bot " + botToken
}

// Start will open the connection to discord and check for any errors and such :)
func (b *Bot) Start() {
	err := b.Session.Open()
	if err != nil {
		log.Fatal("Error opening websocket connection. Error: ", err)
	}

	printASCIIArt()

	log.Println("Bot is now running. Press CTRL-C to exit.")
}

// AwaitTermination awaits termination and then closes the bot and session cleanly.
func (b *Bot) AwaitTermination() {
	// This is used for closing the bot using various different termination signals.
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	b.Session.Close()
}

// for fun ascii art :)
func printASCIIArt() {
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
}
