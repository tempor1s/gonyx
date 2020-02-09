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
		embed := message.GetDefaultEmbed()
		embed.Title = "Log"
		embed.Description = fmt.Sprintf("The current log channel is `%s`", LoggerInstance.ChannelID)

		message.SendEmbed(ds, dm.ChannelID, embed)
		return
	}

	switch ctx.Fields[1] {
	case "update":
		updateLogChannel(ds, dm, ctx.Fields)
	}
	return
}

func updateLogChannel(ds *discordgo.Session, dm *discordgo.Message, fields []string) {
	embed := message.GetDefaultEmbed()
	embed.Title = "Update Log"

	if len(fields) < 3 {
		embed.Description = fmt.Sprintf("Please provide a channel ID to update the log channel.")
		message.SendEmbed(ds, dm.ChannelID, embed)
		return
	}

	oldLogChannel := LoggerInstance.ChannelID

	LoggerInstance.ChannelID = fields[2]

	embed.Description = fmt.Sprintf("Successfully updated log channel from `%s` --> `%s`", oldLogChannel, LoggerInstance.ChannelID)
	message.SendEmbed(ds, dm.ChannelID, embed)
}
