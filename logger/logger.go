package logger

import (
	"fmt"
	"github.com/tempor1s/gonyx/message"

	"github.com/bwmarrin/discordgo"
)

// Logger represents a logger struct
type Logger struct {
	ChannelID  string
	LogDeletes bool
	LogEdits   bool
	LogImages  bool
}

// New creates a new logging instance
func New(channelID string) *Logger {
	l := &Logger{ChannelID: channelID, LogDeletes: false, LogEdits: false, LogImages: false}
	return l
}

// OnMessageDelete will log deleted messages
func (l *Logger) OnMessageDelete(ds *discordgo.Session, md *discordgo.MessageDelete) {
	if l.LogDeletes == false {
		return
	}

	// TODO: Create separate log for bots.
	// if md.Author.Bot {
	// 	return
	// }

	// TODO: Implement delete logging
	embed := message.GetDefaultEmbed()
	embed.Title = "Message Deleted"
	embed.Description = fmt.Sprintf("Message Deleted: %s", md.Content)

	message.SendEmbed(ds, l.ChannelID, embed)
}

// OnMessageEdit is called whenever a message is edited
func (l *Logger) OnMessageEdit(ds *discordgo.Session, mu *discordgo.MessageUpdate) {
	if l.LogEdits == false {
		return
	}

	// TODO: Create separate log for bots.
	// if mu.Author.Bot {
	// 	return
	// }

	embed := message.GetDefaultEmbed()
	embed.Title = "Message Edited"
	embed.Description = fmt.Sprintf("Message Edit: %v -> %v", mu.BeforeUpdate.Content, mu.Content)

	message.SendEmbed(ds, l.ChannelID, embed)
}

// OnMessageCreate is called whenever a new message is created
func (l *Logger) OnMessageCreate(ds *discordgo.Session, mc *discordgo.MessageCreate) {
	if l.LogImages == false {
		return
	}

	// TODO: Create separate log for bots.
	if mc.Author.Bot {
		return
	}

	if len(mc.Attachments) > 0 {
		message.SendMessage(ds, l.ChannelID, mc.Attachments[0].URL)
	}
}
