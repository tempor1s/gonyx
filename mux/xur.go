package mux

import (
	"github.com/bwmarrin/discordgo"
	"github.com/gocolly/colly"
)

type dataURL struct {
	url string
}

// Xur returns where xur is in the current week
func (m *Mux) Xur(ds *discordgo.Session, dm *discordgo.Message, ctx *Context) {
	tempMessage, _ := ds.ChannelMessageSend(dm.ChannelID, "Fetching latest info...")

	c := colly.NewCollector()
	dataURL := &dataURL{}

	c.OnHTML("img", func(e *colly.HTMLElement) {
		dataURL.url = e.Attr("data-src")

		ds.ChannelMessageSend(dm.ChannelID, dataURL.url)
	})

	c.Visit("https://www.niris.tv/blog/xurs-wares")

	ds.ChannelMessageDelete(dm.ChannelID, tempMessage.ID)
}
