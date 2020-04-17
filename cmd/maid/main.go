package main

import (
	"github.com/brokenManager/maid/pkg/async"
	"github.com/brokenManager/maid/pkg/discord"
	"log"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Fatal(err)
		}
	}()

	bot := discord.CreateBot("")
	bot.Start()
	async.Wait()
}
