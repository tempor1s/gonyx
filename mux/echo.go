package mux

import (
	"github.com/bwmarrin/discordgo"
	"github.com/tempor1s/gonyx/message"
)

// Echo function just echos the passed in message
func (m *Mux) Echo(ds *discordgo.Session, dm *discordgo.Message, ctx *Context) {

	if len(ctx.Fields) < 2 {
		message.SendMessage(ds, dm.ChannelID, "Can't trick me with your empty messages :)")
	}

	message.SendMessage(ds, dm.ChannelID, ctx.Content[5:])
}
