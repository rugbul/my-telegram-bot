package main

import (
	"fmt"
	"log"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const telegramBotToken = "7504923807:AAEaVfzH7fV-6eh-SJMnHQJ6rJRM_h1IBf0"

func main() {
	bot, err := tgbotapi.NewBotAPI(telegramBotToken)
	if err != nil {
		log.Panicf("Ошибка при инициализации бота: %v", err)
	}

	bot.Debug = true
	log.Printf("Бот авторизован как %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		switch update.Message.Command() {
		case "start":
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Привет! Я бот, который может показать текущее время. Используй команду /time.")
			bot.Send(msg)
		case "time":
			currentTime := time.Now().Format("2006-01-02 15:04:05")
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("Текущее время: %s", currentTime))
			bot.Send(msg)
		default:
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Я не знаю такой команды.")
			bot.Send(msg)
		}
	}
}
