package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"regexp"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var (
	helpText    = "https://github.com/shantilabs/shantiscrumbot"
	stopRegex   = regexp.MustCompile(`^[\s\-—–]+$`)
	chatVotes   = make(map[int64]votes)
)

func main() {
	if err := runBot(); err != nil {
		log.Panicln(err)
	}
}

func runBot() error {
	tokenPtr := flag.String("token", "", "bot token")
	voteMessagePtr := flag.String("vote-message", "Num votes: %d, average: %d", "Vote message")
	flag.Parse()

	if *tokenPtr == "" {
		return fmt.Errorf("empty bot token")
	}

	bot, err := tgbotapi.NewBotAPI(*tokenPtr)
	if err != nil {
		return err
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 30

	updates := bot.GetUpdatesChan(u)

	log.Printf("start!")

	for update := range updates {
		if update.Message == nil {
			continue
		}

		chatID := update.Message.Chat.ID

		if update.Message.Chat.IsPrivate() {
			log.Printf("send help to %d", update.Message.From.ID)
			if _, err := bot.Send(tgbotapi.NewMessage(chatID, helpText)); err != nil {
				return err
			}
			continue
		}

		if stopRegex.MatchString(update.Message.Text) {
			state := chatVotes[chatID]
			if state == nil {
				state = make(votes)
			}

			results := fmt.Sprintf(*voteMessagePtr, state.num(), state.avg())

			log.Printf("chat: %d results: %s", chatID, results)
			if _, err := bot.Send(tgbotapi.NewMessage(chatID, results)); err != nil {
				return err
			}

			delete(chatVotes, chatID)
			continue
		}

		// store or replace user's vote
		if userVote := forceInt(update.Message.Text); userVote > 0 {
			if chatVotes[chatID] == nil {
				chatVotes[chatID] = make(votes)
			}
			log.Printf("chat: %d vote %s from %d", chatID, update.Message.Text, update.Message.From.ID)
			chatVotes[chatID].add(update.Message.From.ID, userVote)
		}
	}

	return nil
}

// user - vote
type votes map[int64]int

func (v votes) add(user int64, vote int) {
	v[user] = vote
}

func (v votes) num() int {
	return len(v)
}

func (v votes) sum() int {
	sum := 0
	for _, vote := range v {
		sum += vote
	}
	return sum
}

func (v votes) avg() int {
	if len(v) == 0 {
		return 0
	}
	return int(math.Ceil(float64(v.sum()) / float64(v.num())))
}

func forceInt(s string) int {
	s = strings.TrimSpace(s)
	s = strings.TrimLeft(s, "0")
	if s == "" {
		return 0
	}

	val, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0
	}

	return int(val)
}
