package requests

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/tempor1s/gonyx/message"

	"github.com/bwmarrin/discordgo"
	"github.com/gocolly/colly"
)

// GetXurInfo gets xur information and sends to to the given channel
func GetXurInfo(session *discordgo.Session, channelID string, task bool) {
	if task == false {
		tempMessage, _ := session.ChannelMessageSend(channelID, "Fetching latest info...")
		defer message.DeleteMessage(session, channelID, tempMessage.ID)
	}

	c := colly.NewCollector()
	var xurURL []string

	c.OnHTML("img", func(e *colly.HTMLElement) {
		xurURL = append(xurURL, e.Attr("data-src"))
	})

	if err := c.Visit("https://www.niris.tv/blog/xurs-wares"); err != nil {
		log.Println("Error occurred when scraping Xur's Wares.")
	}

	var sameUrl bool
	if task {
		sameUrl = CompareUrls("xur.txt", xurURL)
	} else {
		sameUrl = false
	}

	if sameUrl == false && task {
		message.DeleteMessagesWithAttachment(session, channelID, "xur.png")
	}

	// Get the response from the URL
	resp, err := http.Get(xurURL[0])

	// Make sure that we get a response so we dont have nil errors.
	if err != nil {
		message.SendMessage(session, channelID, "Error fetching Xur info. Please try again :)")
		return
	}

	if sameUrl == false {
		message.SendFile(session, channelID, "xur.png", resp.Body)
	}
}

// GetWeeklyInfo gets weekly information and sends to to the given channel
func GetWeeklyInfo(session *discordgo.Session, channelID string, task bool) {
	if task == false {
		tempMessage, _ := session.ChannelMessageSend(channelID, "Fetching latest info...")
		defer message.DeleteMessage(session, channelID, tempMessage.ID)
	}

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
		log.Println("Error occurred when scraping Weekly Reset.")
	}

	var sameUrls bool
	if task {
		sameUrls = CompareUrls("weekly.txt", weeklyURL)
	} else {
		sameUrls = false
	}

	if sameUrls == false && task {
		message.DeleteMessagesWithAttachment(session, channelID, "weekly.png")
	}

	// Send all the URLS :)
	for _, url := range weeklyURL {
		resp, err := http.Get(url)
		// Make sure that we get a response so we dont have nil errors.
		if err != nil {
			msg := fmt.Sprintf("Error fetching weekly info from `%s` Please try again :)", url)
			message.SendMessage(session, channelID, msg)
			return
		}

		if sameUrls == false {
			message.SendFile(session, channelID, "weekly.png", resp.Body)
		}
	}
}

func CompareUrls(fileName string, urls []string) bool {
	if fileExists(fileName) == false {
		writeLinks(fileName, urls)
		return false
	}

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	for i, line := range lines {
		if urls[i] != line {
			os.Remove(fileName)
			writeLinks(fileName, urls)
			return false
		}
	}

	fmt.Printf("URLS in file '%s' are the same. Skipping...\n", fileName)

	return true
}

func writeLinks(fileName string, urls []string) {

	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatalf("failed creating or reading file: %s", err)
	}

	defer file.Close()

	dataWriter := bufio.NewWriter(file)
	defer dataWriter.Flush()

	for _, url := range urls {
		_, _ = dataWriter.WriteString(url + "\n")
	}

	log.Printf("Successfully wrote URLS to file %q\n", fileName)
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}

	return !info.IsDir()
}
