package requests

import (
	"fmt"
	"net/http"

	"github.com/bwmarrin/discordgo"
	"github.com/gocolly/colly"
)

// GetXurInfo gets xur information and sends to to the given channel
func GetXurInfo(session *discordgo.Session, channelID string) {
	tempMessage, _ := session.ChannelMessageSend(channelID, "Fetching latest info...")

	c := colly.NewCollector()
	var xurURL string

	c.OnHTML("img", func(e *colly.HTMLElement) {
		xurURL = e.Attr("data-src")
	})

	c.Visit("https://www.niris.tv/blog/xurs-wares")

	// Get the response from the URL
	resp, err := http.Get(xurURL)

	// Make sure that we get a response so we dont have nil errors.
	if err != nil {
		session.ChannelMessageSend(channelID, "Error fetching Xur info. Please try again :)")
	} else {
		session.ChannelFileSend(channelID, "file.png", resp.Body)
	}

	session.ChannelMessageDelete(channelID, tempMessage.ID)
}

// GetWeeklyInfo gets weekly information and sends to to the given channel
func GetWeeklyInfo(session *discordgo.Session, channelID string) {
	tempMessage, _ := session.ChannelMessageSend(channelID, "Fetching latest info...")

	// TODO: Clean this up
	// TODO: Set up caching to make this super fast
	// Setup web scraper collector and base struct
	c := colly.NewCollector()
	var weeklyURL []string

	c.OnHTML("img", func(e *colly.HTMLElement) {
		weeklyURL = append(weeklyURL, e.Attr("data-src"))
	})

	// Start the web scraper
	c.Visit("https://www.niris.tv/blog/weekly-reset")

	// Send all the URLS :)
	for _, url := range weeklyURL {
		resp, err := http.Get(url)
		// Make sure that we get a response so we dont have nil errors.
		if err != nil {
			session.ChannelMessageSend(channelID, fmt.Sprintf("Error fetching weekly info from `%s` Please try again :)", url))
		} else {
			session.ChannelFileSend(channelID, "file.png", resp.Body)
		}
	}

	session.ChannelMessageDelete(channelID, tempMessage.ID)
}
