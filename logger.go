package main

import "github.com/tempor1s/gonyx/logger"

// Logger is registered as a global variable to allow easy access to the
// logger throughout the bot.
var Logger = logger.New("536328234556588032")

func init() {
	Session.AddHandler(Logger.OnMessageDelete)
	Session.AddHandler(Logger.OnMessageCreate)
	Session.AddHandler(Logger.OnMessageEdit)
}
