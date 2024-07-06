package internal

import (
	"fmt"
	"log/slog"
	"os"
	"regexp"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func NewBot(token, voteMessage string) (*Bot, error) {
	if token == "" {
		return nil, fmt.Errorf("empty bot token")
	}

	api, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, fmt.Errorf("failed to create bot: %w", err)
	}

	return &Bot{
		api:         api,
		voteMessage: voteMessage,
		helpMessage: "https://github.com/teamin-pro/tgscrumbot",
		logger:      slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{})),
		chatVotes:   make(map[int64]votes),
		stopRegex:   regexp.MustCompile(`^[\s\-—–=]+$`),
	}, nil
}

type Bot struct {
	voteMessage string
	helpMessage string

	api       *tgbotapi.BotAPI
	chatVotes map[int64]votes
	stopRegex *regexp.Regexp
	logger    *slog.Logger
}

func (b Bot) Run() error {
	b.logger.Info("start!")
	b.logger.Info("help message: %s", b.helpMessage)
	b.logger.Info("vote message: %s", b.voteMessage)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 30

	for update := range b.api.GetUpdatesChan(u) {
		if update.Message == nil {
			continue
		}
		if err := b.handleUpdate(update.Message); err != nil {
			return fmt.Errorf("failed to handle update: %w", err)
		}
	}

	return nil
}

func (b Bot) handleUpdate(message *tgbotapi.Message) error {
	switch {
	case message.Chat.IsPrivate():
		return b.handlePrivateMessage(message)
	case b.stopRegex.MatchString(message.Text):
		return b.handleStopMessage(message)
	default:
		return b.handleVoteMessage(message)
	}
}

func (b Bot) handlePrivateMessage(message *tgbotapi.Message) error {
	b.logger.Info("send help to %d", message.From.ID)
	if _, err := b.api.Send(tgbotapi.NewMessage(message.Chat.ID, b.helpMessage)); err != nil {
		return fmt.Errorf("failed to send help: %w", err)
	}
	return nil
}

func (b Bot) handleStopMessage(message *tgbotapi.Message) error {
	state := b.chatVotes[message.Chat.ID]
	if state == nil {
		state = make(votes)
	}

	results := fmt.Sprintf(b.voteMessage, state.num(), state.avg())

	b.logger.Info("chat: %d results: %s", message.Chat.ID, results)
	if _, err := b.api.Send(tgbotapi.NewMessage(message.Chat.ID, results)); err != nil {
		return fmt.Errorf("failed to send results: %w", err)
	}

	delete(b.chatVotes, message.Chat.ID)
	return nil
}

func (b Bot) handleVoteMessage(message *tgbotapi.Message) error {
	userVote := forceInt(message.Text)
	if userVote <= 0 {
		return nil
	}

	if b.chatVotes[message.Chat.ID] == nil {
		b.chatVotes[message.Chat.ID] = make(votes)
	}

	b.logger.Info("chat: %d user: %d vote: %d", message.Chat.ID, message.From.ID, userVote)
	b.chatVotes[message.Chat.ID].add(message.From.ID, userVote)

	return nil
}
