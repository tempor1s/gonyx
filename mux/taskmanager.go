package mux

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/tempor1s/gonyx/message"
	"github.com/tempor1s/gonyx/tasks"
)

// WeeklyInstance just hooks into weekly info so we can manage tasks with commands
var WeeklyInstance *tasks.WeeklyInfo

func (m *Mux) TaskManager(ds *discordgo.Session, dm *discordgo.Message, ctx *Context) {
	if len(ctx.Fields) == 1 {
		embed := message.GetDefaultEmbed()
		embed.Title = "Weekly Info"
		embed.Description = fmt.Sprintf("Currently Info Channel: `%s`\nEnabled: %t\nJob Count (debugging): %d\n",
			WeeklyInstance.ChannelID,
			WeeklyInstance.Enabled,
			len(WeeklyInstance.Schedule.Entries()))

		message.SendEmbed(ds, dm.ChannelID, embed)
		return
	}

	switch ctx.Fields[1] {
	case "enable":
		enableTasks(ds, dm)
	case "disable":
		disableTasks(ds, dm)
	case "restart":
		restartTasks(ds, dm)
	}
}

func enableTasks(ds *discordgo.Session, dm *discordgo.Message) {
	embed := message.GetDefaultEmbed()
	embed.Title = "Enable Tasks"

	WeeklyInstance.Enabled = true
	embed.Description = "Successfully enabled tasks."
	message.SendEmbed(ds, dm.ChannelID, embed)
	return
}

func disableTasks(ds *discordgo.Session, dm *discordgo.Message) {
	embed := message.GetDefaultEmbed()
	embed.Title = "Disable Tasks"

	WeeklyInstance.Enabled = false
	embed.Description = "Successfully disabled tasks."
	message.SendEmbed(ds, dm.ChannelID, embed)
	return
}

func restartTasks(ds *discordgo.Session, dm *discordgo.Message) {
	embed := message.GetDefaultEmbed()
	embed.Title = "Restart Tasks"

	WeeklyInstance.Schedule.Stop()
	WeeklyInstance.Schedule.Start()

	embed.Description = fmt.Sprintf("Restarted tasks. Task Count: %d\n", len(WeeklyInstance.Schedule.Entries()))
	message.SendEmbed(ds, dm.ChannelID, embed)
	return
}
