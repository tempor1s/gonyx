package mux

import (
	"net/http"

	"github.com/bwmarrin/discordgo"
	"github.com/gocolly/colly"
)

type xurURL struct {
	url string
}

// Xur returns where xur is in the current week
func (m *Mux) Xur(ds *discordgo.Session, dm *discordgo.Message, ctx *Context) {
	tempMessage, _ := ds.ChannelMessageSend(dm.ChannelID, "Fetching latest info...")

	// TODO: Clean this up
	// TODO: Set up caching to make this super fast
	// Setup web scraper collector and base struct
	c := colly.NewCollector()
	xurURL := &xurURL{}

	// Find the URL to pull data from
	c.OnHTML("img", func(e *colly.HTMLElement) {
		xurURL.url = e.Attr("data-src")
	})

	// Start the web scraper
	c.Visit("https://www.niris.tv/blog/xurs-wares")

	// Get the response from the URL
	resp, err := http.Get(xurURL.url)
	// Make sure that we get a response so we dont have nil errors.
	if err != nil {
		ds.ChannelMessageSend(dm.ChannelID, "Error fetching Xur info. Please try again :)")
	} else {
		ds.ChannelFileSend(dm.ChannelID, "file.png", resp.Body)
	}

	ds.ChannelMessageDelete(dm.ChannelID, tempMessage.ID)
}
