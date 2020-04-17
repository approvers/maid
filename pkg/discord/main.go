package discord

func CreateBot(token string) Bot {
	var bot Bot
	return *bot.Initialize(token)
}
