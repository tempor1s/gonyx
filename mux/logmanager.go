package mux

import (
	"fmt"
	"github.com/tempor1s/gonyx/message"

	"github.com/bwmarrin/discordgo"
	"github.com/tempor1s/gonyx/logger"
)

// LoggerInstance is for commands to be able to access the Log for commands
var LoggerInstance *logger.Logger

// LogManager allows you to manage the logger
func (m *Mux) LogManager(ds *discordgo.Session, dm *discordgo.Message, ctx *Context) {
	if len(ctx.Fields) == 1 {
		msg := fmt.Sprintf("The current log channel is %s", LoggerInstance.ChannelID)
		message.SendMessage(ds, dm.ChannelID, msg)

		return
	}

	switch ctx.Fields[1] {
	case "update":
		updateLogChannel(ds, dm, ctx.Fields)
	}
	return
}

func updateLogChannel(ds *discordgo.Session, dm *discordgo.Message, fields []string) {
	if len(fields) < 3 {
		msg := fmt.Sprintf("Please provide a channel ID to update the log channel.")
		message.SendMessage(ds, dm.ChannelID, msg)
		return
	}

	oldLogChannel := LoggerInstance.ChannelID

	LoggerInstance.ChannelID = fields[2]
	ds.ChannelMessageSend(dm.ChannelID, fmt.Sprintf("Successfully updated log channel from %q --> %q", oldLogChannel, LoggerInstance.ChannelID))
}
