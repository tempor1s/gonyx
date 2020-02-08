package mux

import (
	"fmt"
	"net/http"

	"github.com/bwmarrin/discordgo"
	"github.com/gocolly/colly"
)

type weeklyURL struct {
	url []string
}

// Weekly gets Destiny 2 weekly information
func (m *Mux) Weekly(ds *discordgo.Session, dm *discordgo.Message, ctx *Context) {
	tempMessage, _ := ds.ChannelMessageSend(dm.ChannelID, "Fetching latest info...")

	// TODO: Clean this up
	// TODO: Set up caching to make this super fast
	// Setup web scraper collector and base struct
	c := colly.NewCollector()
	weeklyURL := &weeklyURL{}

	c.OnHTML("img", func(e *colly.HTMLElement) {
		weeklyURL.url = append(weeklyURL.url, e.Attr("data-src"))
	})

	// Start the web scraper
	c.Visit("https://www.niris.tv/blog/weekly-reset")

	// Send all the URLS :)
	for _, url := range weeklyURL.url {
		resp, err := http.Get(url)
		// Make sure that we get a response so we dont have nil errors.
		if err != nil {
			ds.ChannelMessageSend(dm.ChannelID, fmt.Sprintf("Error fetching weekly info from `%s` Please try again :)", url))
		} else {
			ds.ChannelFileSend(dm.ChannelID, "file.png", resp.Body)
		}
	}

	ds.ChannelMessageDelete(dm.ChannelID, tempMessage.ID)
}
