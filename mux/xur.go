package mux

import (
	"github.com/bwmarrin/discordgo"
	"github.com/tempor1s/gonyx/requests"
)

// Xur returns where xur is in the current week
func (m *Mux) Xur(ds *discordgo.Session, dm *discordgo.Message, ctx *Context) {
	// TODO: GoRoutine this
	requests.GetXurInfo(ds, dm.ChannelID)
}
