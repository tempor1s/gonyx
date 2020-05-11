package events

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

type Events struct {
	onJoinRole string
	enabled    bool
}

func New(onJoinRole string, session *discordgo.Session) *Events {
	e := &Events{
		onJoinRole: onJoinRole,
		enabled:    true,
	}

	return e
}

func (e *Events) OnGuildMemberJoin(ds *discordgo.Session, gma *discordgo.GuildMemberAdd) {
	err := ds.GuildMemberRoleAdd(gma.GuildID, gma.User.ID, e.onJoinRole)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Added role with id (%s) to user %s\n", gma.User.ID, gma.User.Username)
}
