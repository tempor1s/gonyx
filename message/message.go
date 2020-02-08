// Package message is a light wrapper that allows you to print
package message

import (
	"github.com/bwmarrin/discordgo"
	"io"
	"log"
)

func SendMessage(session *discordgo.Session, channelID, content string) *discordgo.Message {
	msg, err := session.ChannelMessageSend(channelID, content)

	if err != nil {
		log.Printf("Error sending message in %q. Error: %v", channelID, err)
	}

	return msg
}

func SendFile(session *discordgo.Session, channelID, fileName string, reader io.Reader) *discordgo.Message {
	msg, err := session.ChannelFileSend(channelID, "file.png", reader)

	if err != nil {
		log.Printf("Error sending file in %q. Error: %v", channelID, err)
	}

	return msg
}

func DeleteMessage(session *discordgo.Session, channelID, messageID string) {
	err := session.ChannelMessageDelete(channelID, messageID)

	if err != nil {
		log.Printf("Error deleting message in %q. Error: %v", channelID, err)
	}
}