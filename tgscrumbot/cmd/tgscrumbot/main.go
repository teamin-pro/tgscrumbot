package main

import (
	"flag"
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/lmittmann/tint"

	"github.com/teamin-pro/tgscrumbot/tgscrumbot/internal"
)

var (
	token       = flag.String("token", "", "bot token")
	voteMessage = flag.String("vote-message", "Num votes: %d, average: %d", "Vote message")
)

func main() {
	flag.Parse()

	slog.SetDefault(slog.New(
		tint.NewHandler(os.Stdout, &tint.Options{
			Level:      slog.LevelDebug,
			TimeFormat: time.Kitchen,
		}),
	))

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
