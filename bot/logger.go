package bot

import (
	"github.com/bwmarrin/discordgo"
	"github.com/tempor1s/gonyx/logger"
)

// registerLog registers all the logging handlers
func registerLog(session *discordgo.Session) *logger.Logger {
	logger := logger.New("536328234556588032")

	session.AddHandler(logger.OnMessageDelete)
	session.AddHandler(logger.OnMessageCreate)
	session.AddHandler(logger.OnMessageEdit)

	return logger
}
