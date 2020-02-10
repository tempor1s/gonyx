package mux

import (
"fmt"
"github.com/bwmarrin/discordgo"
"github.com/tempor1s/gonyx/logger"
"github.com/tempor1s/gonyx/message"
)

// LoggerInstance is for commands to be able to access the Log for commands
var LoggerInstance *logger.Logger

// LogManager allows you to manage the logger
func (m *Mux) LogManager(ds *discordgo.Session, dm *discordgo.Message, ctx *Context) {
	// TODO: Make this shit work for the love of god...
	embed := message.GetDefaultEmbed()
	embed.Title = "Log"
	embed.Description = "Logging is currently disabled while we figure some stuff out. Try please try again later!"
	message.SendEmbed(ds, dm.ChannelID, embed)
	return

	//if len(ctx.Fields) == 1 {
	//	embed := message.GetDefaultEmbed()
	//	embed.Title = "Log"
	//	embed.Description = fmt.Sprintf("Current Log Channel: `%s`\nLog Edits: `%t`\nLog Deletes: `%t`\nLog Images: `%t`\n",
	//		LoggerInstance.ChannelID,
	//		LoggerInstance.LogEdits,
	//		LoggerInstance.LogDeletes,
	//		LoggerInstance.LogImages)
	//
	//	message.SendEmbed(ds, dm.ChannelID, embed)
	//	return
	//}
	//
	//switch ctx.Fields[1] {
	//case "update":
	//	updateLogChannel(ds, dm, ctx.Fields)
	//case "enable":
	//	enableLogging(ds, dm, ctx.Fields)
	//case "disable":
	//	disableLogging(ds, dm, ctx.Fields)
	//}
	//return
}

func updateLogChannel(ds *discordgo.Session, dm *discordgo.Message, fields []string) {
	embed := message.GetDefaultEmbed()
	embed.Title = "Update Log"

	if len(fields) < 3 {
		embed.Description = fmt.Sprintf("Please provide a channel ID to update the log channel.")
		message.SendEmbed(ds, dm.ChannelID, embed)
		return
	}

	// Get the old channel ID
	oldChannelId := LoggerInstance.ChannelID
	// Set it to be the new channel ID
	// TODO: Check to make sure it is an ID being passed in.
	LoggerInstance.ChannelID = fields[2]

	embed.Description = fmt.Sprintf("Successfully updated log channel from `%s` --> `%s`",
		oldChannelId, LoggerInstance.ChannelID)

	message.SendEmbed(ds, dm.ChannelID, embed)
}

func enableLogging(ds *discordgo.Session, dm *discordgo.Message, fields []string) {
	embed := message.GetDefaultEmbed()
	embed.Title = "Enabled Logging"

	if len(fields) < 3 {
		embed.Description = fmt.Sprintf("Please provide a type of logging to enable.")
		message.SendEmbed(ds, dm.ChannelID, embed)
		return
	}

	switch fields[2] {
	case "edits":
		LoggerInstance.LogEdits = true
		embed.Description = "Successfully enabled `edits` logging."
	case "deletes":
		LoggerInstance.LogDeletes = true
		embed.Description = "Successfully enabled `deletes` logging."
	case "images":
		LoggerInstance.LogImages = true
		embed.Description = "Successfully enabled `images` logging."
	default:
		embed.Description = "Please provide a logging type to enable. (`edits`, `deletes`, `images`)"
	}

	message.SendEmbed(ds, dm.ChannelID, embed)
}

func disableLogging(ds *discordgo.Session, dm *discordgo.Message, fields []string) {
	embed := message.GetDefaultEmbed()
	embed.Title = "Disable Logging"

	if len(fields) < 3 {
		embed.Description = fmt.Sprintf("Please provide a type of logging to disable.")
		message.SendEmbed(ds, dm.ChannelID, embed)
		return
	}
}