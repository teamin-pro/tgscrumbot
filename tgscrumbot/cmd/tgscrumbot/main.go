package main

import (
	"flag"
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/lmittmann/tint"
	"github.com/mattn/go-isatty"

	"github.com/teamin-pro/tgscrumbot/tgscrumbot/internal"
)

var (
	token       = flag.String("token", "", "bot token")
	voteMessage = flag.String("vote-message", "Num votes: %d, average: %d", "Vote message")
)

func main() {
	flag.Parse()

	w := os.Stdout
	slog.SetDefault(slog.New(
		tint.NewHandler(w, &tint.Options{
			Level:      slog.LevelDebug,
			TimeFormat: time.Kitchen,
			NoColor:    !isatty.IsTerminal(w.Fd()),
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
