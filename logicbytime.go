package main

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"time"
)

func logicByTime(update tgbotapi.Update, bot *tgbotapi.BotAPI, c *Cache) {

	if time.Now().Before(m.Add(time.Minute)) && time.Now().After(m) {
		fmt.Println(m)
		msg := tgbotapi.NewMessage(update.SentFrom().ID, "Введи номер карты")
		if _, err := bot.Send(msg); err != nil {
			log.Println(err)
		}
	} else {
		caseIncomingMessage(update, bot, c)
		callbackQueryCase(update, bot, c)
	}
}
