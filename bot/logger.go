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

	logInstance := logger.New(logChannel, b.Session) // TODO: Disabled by default and then configure from db

	b.Session.AddHandler(logInstance.OnMessageDelete)
	b.Session.AddHandler(logInstance.OnMessageCreate)
	b.Session.AddHandler(logInstance.OnMessageEdit)

	b.Logger = logInstance
}
