package main

import (
	"github.com/Pgv03/CopyPastaMan/bot"
)

func main() {
	bot.Start()

	<-make(chan struct{})
}
