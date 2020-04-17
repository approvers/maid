package discord

func CreateBot(token string) *Bot {
	bot := Bot{}
	return bot.Initialize(token)
}
