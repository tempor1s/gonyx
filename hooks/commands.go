// Package hooks allows you to hook into different modules,
// such as the logger and the command handler
package hooks

import (
	"github.com/bwmarrin/discordgo"
	"github.com/tempor1s/gonyx/mux"
)

// RegisterRouter will register new routes.
func RegisterRouter(session *discordgo.Session) *mux.Mux {
	mux := mux.New()
	// Register the mux OnMessageCreate handler that listens for and processes
	// all messages received.
	session.AddHandler(mux.OnMessageCreate)

	// Register the built-in help command.
	mux.Route("help", "Display this message.", mux.Help)
	mux.Route("echo", "Echo the given message back at you.", mux.Echo)
	mux.Route("log", "Manage logging functionality.", mux.LogManager)

	return mux
}
