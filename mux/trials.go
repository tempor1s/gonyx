package mux

import (
	"github.com/bwmarrin/discordgo"
	"github.com/tempor1s/gonyx/requests"
)

// Trials gets weekly information about trails stuff.
func (m *Mux) Trials(ds *discordgo.Session, dm *discordgo.Message, ctx *Context) {
	requests.GetTrialsInfo(ds, dm.ChannelID, false)
}
