// Package hooks allows you to hook into different modules,
// such as the logger and the command handler
package hooks

import (
	"github.com/bwmarrin/discordgo"
	"github.com/tempor1s/gonyx/logger"
)

// RegisterLog registers all the logging handlers
func RegisterLog(session *discordgo.Session) *logger.Logger {
	logger := logger.New("536328234556588032")

	session.AddHandler(logger.OnMessageDelete)
	session.AddHandler(logger.OnMessageCreate)
	session.AddHandler(logger.OnMessageEdit)

	return logger
}
