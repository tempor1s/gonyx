package mux

import (
	"github.com/bwmarrin/discordgo"
	"github.com/tempor1s/gonyx/requests"
)

// Weekly gets Destiny 2 weekly information
func (m *Mux) Weekly(ds *discordgo.Session, dm *discordgo.Message, ctx *Context) {
	// TODO: GoRoutine this
	requests.GetWeeklyInfo(ds, dm.ChannelID)
}
