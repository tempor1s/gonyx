package bot

import (
	"github.com/tempor1s/gonyx/logger"
)

// registerLog registers all the logging handlers
func (b *Bot) registerLog() {
	logger := logger.New("536328234556588032") // TODO: Disabled by default and then configure from db

	b.Session.AddHandler(logger.OnMessageDelete)
	b.Session.AddHandler(logger.OnMessageCreate)
	b.Session.AddHandler(logger.OnMessageEdit)

	b.Logger = logger
}
