package discord

import (
	"errors"
	"github.com/brokenManager/maid/pkg/discord/handlers"
	"github.com/bwmarrin/discordgo"
)

type Bot struct {
	session *discordgo.Session
}

func (bot *Bot) Initialize(token string) *Bot {
	session, err := discordgo.New("Bot " + token)

	if err != nil {
		panic(err)
	}

	bot.session = session
	bot.RegisterHandlers()

	return bot
}

func (bot *Bot) Start() *Bot {
	err := bot.session.Open()

	if err != nil {
		panic(err)
	}

	return bot
}

func (bot *Bot) Stop() *Bot {
	err := bot.session.Close()

	if err != nil {
		panic(err)
	}

	return bot
}

func (bot *Bot) EnsureSession() *Bot {
	if bot.session == nil {
		panic(errors.New("no session activated"))
	}

	return bot
}

func (bot *Bot) RegisterHandlers() *Bot {
	bot.EnsureSession().session.AddHandler(handlers.OnMessageCreate)

	return bot
}
