package main

import (
	"flag"
	"github.com/brokenManager/maid/pkg/async"
	"github.com/brokenManager/maid/pkg/discord"
	"log"
)

var (
	Token string
)

func init() {
	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Fatal(err)
		}
	}()

	bot := discord.CreateBot(Token)
	bot.Start()
	async.Wait()
	bot.Stop()
}
