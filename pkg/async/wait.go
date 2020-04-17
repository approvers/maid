package async

import (
	"os"
	"os/signal"
	"syscall"
)

func Wait() {
	channel := make(chan os.Signal, 1)
	signal.Notify(channel, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)

	<-channel
}
