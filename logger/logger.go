package logger

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/tempor1s/gonyx/message"
	"log"
)

// Logger represents a logger struct
type Logger struct {
	ChannelID  string
	Channel *discordgo.Channel
	LogDeletes bool
	LogEdits   bool
	LogImages  bool
}

// New creates a new logging instance
func New(channelID string, session *discordgo.Session) *Logger {

	l := &Logger{
		ChannelID: channelID,
		LogDeletes: false,
		LogEdits: false,
		LogImages: false,
	}

	return l
}

// OnMessageDelete will log deleted messages
func (l *Logger) OnMessageDelete(ds *discordgo.Session, md *discordgo.MessageDelete) {
	if l.LogDeletes == false {
		return
	}

	// TODO: Create separate log for bots.
	if md.Author.ID == ds.State.User.ID {
		return
	}

	log.Printf("%+v", md)
	log.Printf("%+v", md.Message)

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

	// Get the message because fuck discord API
	msg, err := ds.ChannelMessage(mu.Message.ChannelID, mu.Message.ID)

	if err != nil {
		log.Printf("Error: %v", err)
	}

	if msg.Author.ID == ds.State.User.ID {
		return
	}

	log.Printf("message: %+v\n", mu)

	embed := message.GetDefaultEmbed()
	embed.Title = "Message Edited"
	embed.Fields = []*discordgo.MessageEmbedField{
		{
			Name:   "Before",
			Value:  mu.BeforeUpdate.Content,
			Inline: false,
		},
		{
			Name: "After",
			Value: mu.Content,
		},
		{
			 Name: "Message ID",
			 Value: mu.ID,
		},
		{
			Name: "Channel ID",
			Value: msg.ChannelID,
		},
	}
	//embed.Description = fmt.Sprintf("Message Edit: %v -> %v", mu.BeforeUpdate.Content, mu.Content)

	message.SendEmbed(ds, l.ChannelID, embed)
}

// OnMessageCreate is called whenever a new message is created
func (l *Logger) OnMessageCreate(ds *discordgo.Session, mc *discordgo.MessageCreate) {
	if l.LogImages == false {
		return
	}

	// TODO: Create separate log for bots.
	if mc.Author.ID == ds.State.User.ID {
		return
	}

	if len(mc.Attachments) > 0 {
		message.SendMessage(ds, l.ChannelID, mc.Attachments[0].URL)
	}
}
