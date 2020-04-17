package handlers

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

func OnMessageCreate(session *discordgo.Session, event *discordgo.MessageCreate) {
	if event.Author.ID == session.State.User.ID {
		return
	}

	fmt.Printf(
		"[%s:%s] %s: %s\n",
		event.GuildID,
		event.ChannelID,
		event.Author.Username,
		event.Message.Content,
	)
}
