package bot

import (
	"log"
	"os"

	"github.com/tempor1s/gonyx/tasks"
)

func (b *Bot) registerTasks() {
	weeklyInfoChannel, exists := os.LookupEnv("WEEKLY_INFO_CHANNEL")

	if !exists {
		log.Println("No weekly info channel supplied.")
	}

	weeklyInfo := tasks.NewWeeklyInfo(weeklyInfoChannel, b.Session)

	// Register Xur weekly info task
	weeklyInfo.RegisterXurInfo()
	// Register weekly info
	weeklyInfo.RegisterWeeklyInfo()

	// Start all the registered tasks
	weeklyInfo.Schedule.Start()

	b.WeeklyInfo = weeklyInfo
}
