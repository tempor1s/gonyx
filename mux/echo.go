package mux

import (
	"github.com/bwmarrin/discordgo"
	"github.com/tempor1s/gonyx/message"
)

// Echo function just echos the passed in message
func (m *Mux) Echo(ds *discordgo.Session, dm *discordgo.Message, ctx *Context) {

	if len(ctx.Fields) < 2 {
		message.SendMessage(ds, dm.ChannelID, "Can't trick me with your empty messages :)")
		return
	}

	embed := message.GetDefaultEmbed()
	embed.Title = "Echo"
	embed.Description = ctx.Content[5:]

	message.SendEmbed(ds, dm.ChannelID, embed)
}
