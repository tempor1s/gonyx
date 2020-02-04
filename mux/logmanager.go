package mux

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

// LogManager allows you to manage the logger
func (m *Mux) LogManager(ds *discordgo.Session, dm *discordgo.Message, ctx *Context) {
	fmt.Printf("Log instance: %v \n", LoggerInstance.ChannelID)

	if len(ctx.Fields) == 1 {
		ds.ChannelMessageSend(dm.ChannelID, fmt.Sprintf("The current log channel is %s", LoggerInstance.ChannelID))
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
		ds.ChannelMessageSend(dm.ChannelID, fmt.Sprintf("Please provide a channel ID to update the log channel."))
		return
	}

	oldLogChannel := LoggerInstance.ChannelID

	LoggerInstance.ChannelID = fields[2]
	ds.ChannelMessageSend(dm.ChannelID, fmt.Sprintf("Successfully updated log channel from %q --> %q", oldLogChannel, LoggerInstance.ChannelID))
}
