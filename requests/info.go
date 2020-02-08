package requests

import (
	"fmt"
	"github.com/tempor1s/gonyx/message"
	"log"
	"net/http"

	"github.com/bwmarrin/discordgo"
	"github.com/gocolly/colly"
)

// GetXurInfo gets xur information and sends to to the given channel
func GetXurInfo(session *discordgo.Session, channelID string) {
	tempMessage := message.SendMessage(session, channelID, "Fetching latest info...")

	c := colly.NewCollector()
	var xurURL string

	c.OnHTML("img", func(e *colly.HTMLElement) {
		xurURL = e.Attr("data-src")
	})

	if err := c.Visit("https://www.niris.tv/blog/xurs-wares"); err != nil {
		log.Println("Error occured when scraping Xurs Wares.")
	}

	// Get the response from the URL
	resp, err := http.Get(xurURL)

	// Make sure that we get a response so we dont have nil errors.
	if err != nil {
		message.SendMessage(session, channelID, "Error fetching Xur info. Please try again :)")
	} else {
		message.SendFile(session, channelID, "xur.png", resp.Body)
	}

	message.DeleteMessage(session, channelID, tempMessage.ID)
}

// GetWeeklyInfo gets weekly information and sends to to the given channel
func GetWeeklyInfo(session *discordgo.Session, channelID string) {
	tempMessage, _ := session.ChannelMessageSend(channelID, "Fetching latest info...")

	// TODO: Clean this up
	// TODO: Set up caching or logging in database to speed this up most of the time.
	// Setup web scraper collector and base struct
	c := colly.NewCollector()
	var weeklyURL []string

	c.OnHTML("img", func(e *colly.HTMLElement) {
		weeklyURL = append(weeklyURL, e.Attr("data-src"))
	})

	// Start the web scraper
	if err := c.Visit("https://www.niris.tv/blog/weekly-reset"); err != nil {
		log.Println("Error occured when scraping Weekly Reset.")
	}

	// Send all the URLS :)
	for _, url := range weeklyURL {
		resp, err := http.Get(url)
		// Make sure that we get a response so we dont have nil errors.
		if err != nil {
			msg := fmt.Sprintf("Error fetching weekly info from `%s` Please try again :)", url)
			message.SendMessage(session, channelID, msg)
		} else {
			message.SendFile(session, channelID,"weekly.png", resp.Body)
		}
	}

	message.DeleteMessage(session, channelID, tempMessage.ID)
}
