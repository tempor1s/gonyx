package mux

import (
	"github.com/bwmarrin/discordgo"
)

// Echo function just echos the passed in message
func (m *Mux) Echo(ds *discordgo.Session, dm *discordgo.Message, ctx *Context) {

	if len(ctx.Fields) < 2 {
		ds.ChannelMessageSend(dm.ChannelID, "Can't trick me with your empty messages :)")
	}

	ds.ChannelMessageSend(dm.ChannelID, ctx.Content[4:])
}
