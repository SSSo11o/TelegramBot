package main

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strings"
	"telegrambot/cmd/config"
	dowloader2 "telegrambot/cmd/internal/dowloader"
)

func handlerDownload(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	url := update.Message.Text

	if url == "" {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Please enter a URL")
		bot.Send(msg)
		return
	}

	var err error
	var videoPath string
	if strings.Contains(url, "youtube.com") || strings.Contains(url, "youtu.be") {
		videoPath, err = dowloader2.DownloadVideo(url)
	} else if strings.Contains(url, "instagram.com") {
		err = dowloader2.DownloadInstagram(url)
	} else {
		err = fmt.Errorf("не поддерживаемый формат ссылки")
	}
	if err != nil {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, err.Error())
		bot.Send(msg)
		return
	}

	if videoPath != "" {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, videoPath)
		bot.Send(msg)
	} else {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Не удалось получить ссылку для скачивания.")
		bot.Send(msg)
	}
}

func main() {
	config, err := config.LoadConfig("config/config.yaml")
	if err != nil {
		log.Fatalf("Error loading config.yaml: %v", err)
	}
	bot, err := tgbotapi.NewBotAPI(config.Telegram.Token)
	if err != nil {
		log.Fatalf("Error creating telegram bot: %v", err)
	}

	bot.Debug = true
	log.Printf("Авторизован под именем %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
			if strings.HasPrefix(update.Message.Text, "https://") {
				handlerDownload(update, bot)
			} else {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Пожалуйста, отправьте корректную ссылку.")
				bot.Send(msg)
			}
		}
	}
}
