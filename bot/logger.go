package bot

import (
	"log"
	"os"

	"github.com/tempor1s/gonyx/logger"
)

// registerLog registers all the logging handlers
func (b *Bot) registerLog() {
	logChannel, exists := os.LookupEnv("LOG_CHANNEL")

	if !exists {
		log.Println("No log channel supplied.")
	}

	logger := logger.New(logChannel) // TODO: Disabled by default and then configure from db

	b.Session.AddHandler(logger.OnMessageDelete)
	b.Session.AddHandler(logger.OnMessageCreate)
	b.Session.AddHandler(logger.OnMessageEdit)

	b.Logger = logger
}
