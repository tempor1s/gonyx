package logger

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

// HandlerFunc is the function signature required for a message route handler.
type HandlerFunc func(*discordgo.Session, *discordgo.Message)

// Logger represents a logger struct
type Logger struct {
	ChannelID  string
	LogDeletes bool
	LogEdits   bool
	LogImages  bool
}

// New creates a new logging instance
func New(channelID string) *Logger {
	l := &Logger{ChannelID: channelID, LogDeletes: false, LogEdits: true, LogImages: false}
	return l
}

// OnMessageDelete will log deleted messages
func (l *Logger) OnMessageDelete(ds *discordgo.Session, md *discordgo.MessageDelete) {
	if l.LogDeletes == false {
		return
	}

	// TODO: Create seperate log for bots.
	if md.Author.Bot {
		return
	}

	// TODO: Implement delete logging
	ds.ChannelMessageSend(l.ChannelID, fmt.Sprintf("Content: %v", md.Content))
}

// OnMessageEdit is called whenever a message is edited
func (l *Logger) OnMessageEdit(ds *discordgo.Session, mu *discordgo.MessageUpdate) {
	if l.LogEdits == false {
		return
	}

	// TODO: Create seperate log for bots.
	if mu.Author.Bot {
		return
	}

	msg := fmt.Sprintf("Message Edit: %v -> %v", mu.BeforeUpdate.Content, mu.Content)
	ds.ChannelMessageSend(l.ChannelID, msg)
}

// OnMessageCreate is called whenever a new message is created
func (l *Logger) OnMessageCreate(ds *discordgo.Session, mc *discordgo.MessageCreate) {
	if l.LogImages == false {
		return
	}

	// TODO: Create seperate log for bots.
	if mc.Author.Bot {
		return
	}

	if len(mc.Attachments) > 0 {
		ds.ChannelMessageSend(l.ChannelID, mc.Attachments[0].URL)
	}
}
