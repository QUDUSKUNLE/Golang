package telegram

import (
	"fmt"
	"log"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type UserOrderState struct {
	Step      int
	Item      string
	Quantity  int
	Confirmed bool
}

var userStates = make(map[int64]*UserOrderState)

// StartOrderBot starts a simple order placement bot
func StartOrderBot(token string) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}
		chatID := update.Message.Chat.ID
		text := strings.TrimSpace(update.Message.Text)

		state, exists := userStates[chatID]
		if !exists {
			state = &UserOrderState{Step: 0}
			userStates[chatID] = state
		}

		switch state.Step {
		case 0:
			msg := tgbotapi.NewMessage(chatID, "Welcome! What item would you like to order?")
			bot.Send(msg)
			state.Step = 1
		case 1:
			state.Item = text
			msg := tgbotapi.NewMessage(chatID, fmt.Sprintf("How many '%s' would you like to order?", state.Item))
			bot.Send(msg)
			state.Step = 2
		case 2:
			var qty int
			_, err := fmt.Sscanf(text, "%d", &qty)
			if err != nil || qty <= 0 {
				msg := tgbotapi.NewMessage(chatID, "Please enter a valid quantity (number greater than 0).")
				bot.Send(msg)
				continue
			}
			state.Quantity = qty
			msg := tgbotapi.NewMessage(chatID, fmt.Sprintf("You want to order %d '%s'. Confirm? (yes/no)", state.Quantity, state.Item))
			bot.Send(msg)
			state.Step = 3
		case 3:
			if strings.ToLower(text) == "yes" {
				state.Confirmed = true
				msg := tgbotapi.NewMessage(chatID, fmt.Sprintf("✅ Order placed: %d '%s'. Thank you!", state.Quantity, state.Item))
				bot.Send(msg)
				delete(userStates, chatID) // Reset state
			} else if strings.ToLower(text) == "no" {
				msg := tgbotapi.NewMessage(chatID, "Order cancelled. To start again, type anything.")
				bot.Send(msg)
				delete(userStates, chatID)
			} else {
				msg := tgbotapi.NewMessage(chatID, "Please reply with 'yes' or 'no'.")
				bot.Send(msg)
			}
		}
	}
}
