package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/teamin-pro/tgscrumbot/tgscrumbot/internal"
)

var (
	token       = flag.String("token", "", "bot token")
	voteMessage = flag.String("vote-message", "Num votes: %d, average: %d", "Vote message")
)

func main() {
	flag.Parse()

	bot, err := internal.NewBot(*token, *voteMessage)
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}

	if err := bot.Run(); err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
}
