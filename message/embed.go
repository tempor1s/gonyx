package message

import (
	"github.com/bwmarrin/discordgo"
	"log"
	"time"
)

func GetDefaultEmbed() *discordgo.MessageEmbed {
	footer := &discordgo.MessageEmbedFooter{
		Text:         "Powered by GOnyx",
	}
	return &discordgo.MessageEmbed{
		Timestamp: time.Now().Format(time.RFC3339),
		Color:0x800080,
		Footer:footer,
	}
}

func SendEmbed(session *discordgo.Session, channelID string, embed *discordgo.MessageEmbed) *discordgo.Message {
	msg, err := session.ChannelMessageSendEmbed(channelID, embed)

	if err != nil {
		log.Printf("Error sending embed in %q. Error: %v", channelID, err)
	}

	return msg
}