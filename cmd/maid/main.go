package main

import (
	"github.com/brokenManager/maid/pkg/async"
	"github.com/brokenManager/maid/pkg/discord"
)

func main() {
	bot := discord.CreateBot("")
	bot.Start()
	async.Wait()
}
