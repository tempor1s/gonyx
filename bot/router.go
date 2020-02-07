package bot

import (
	"github.com/tempor1s/gonyx/mux"
)

// registerRouter will register all the routes.
func (b *Bot) registerRouter() {
	mux := mux.New()
	// Register the mux OnMessageCreate handler that listens for and processes
	// all messages received.
	b.Session.AddHandler(mux.OnMessageCreate)

	// Register the built-in help command.
	mux.Route("help", "Display this message.", mux.Help)
	mux.Route("echo", "Echo the given message back at you.", mux.Echo)
	mux.Route("log", "Manage logging functionality.", mux.LogManager)
	mux.Route("xur", "Gives you a nice image of the weekly Xur information.", mux.Xur)
	mux.Route("weekly", "Gets Destiny 2 weekly information.", mux.Weekly)

	b.Mux = mux
}
