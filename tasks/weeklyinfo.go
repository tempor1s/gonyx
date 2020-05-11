package tasks

import (
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/robfig/cron"
	"github.com/tempor1s/gonyx/requests"
)

// WeeklyInfo sets up configuration for weekly info
type WeeklyInfo struct {
	ChannelID string
	Enabled   bool
	Schedule  *cron.Cron
	Session   *discordgo.Session
}

// NewWeeklyInfo creates a new weekly info loop
func NewWeeklyInfo(channelID string, session *discordgo.Session) *WeeklyInfo {
	c := cron.New()
	return &WeeklyInfo{ChannelID: channelID, Enabled: false, Schedule: c, Session: session}
}

// RegisterXurInfo registers xur info command
func (w *WeeklyInfo) RegisterXurInfo() {

	err := w.Schedule.AddFunc("@every 1m", func() {
		// TODO: GoRoutine this
		if w.Enabled {
			requests.GetXurInfo(w.Session, w.ChannelID, true)
		}
	})

	if err != nil {
		log.Println("Did not register Xur Info task.")
	}
}

// RegisterWeeklyInfo registers weekly info command
func (w *WeeklyInfo) RegisterWeeklyInfo() {

	err := w.Schedule.AddFunc("@every 1m", func() {
		// TODO: GoRoutine this
		if w.Enabled {
			requests.GetWeeklyInfo(w.Session, w.ChannelID, true)
		}
	})

	if err != nil {
		log.Println("Did not register Weekly Info task.")
	}
}
