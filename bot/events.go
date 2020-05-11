package bot

import (
	"fmt"
	"log"
	"os"

	"github.com/tempor1s/gonyx/events"
)

func (b *Bot) registerEvents() {
	fmt.Println("Registering normal events..")

	joinRole, exists := os.LookupEnv("ROLE_ON_JOIN")

	if !exists {
		log.Fatal("Please specify role that will be added on join.")
	}

	e := events.New(joinRole, b.Session)

	b.Session.AddHandler(e.OnGuildMemberJoin)
}
