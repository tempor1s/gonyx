package tasks

import (
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
	return &WeeklyInfo{ChannelID: channelID, Enabled: true, Schedule: c, Session: session}
}

// RegisterXurInfo registers xur info command
func (w *WeeklyInfo) RegisterXurInfo() {
	w.Schedule.AddFunc("@every 5m", func() {
		// TODO: GoRoutine this
		requests.GetXurInfo(w.Session, w.ChannelID)
	})
}

// RegisterWeeklyInfo registers weekly info command
func (w *WeeklyInfo) RegisterWeeklyInfo() {
	w.Schedule.AddFunc("@every 5m", func() {
		// TODO: GoRoutine this
		requests.GetWeeklyInfo(w.Session, w.ChannelID)
	})
}
