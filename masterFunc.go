package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

var masterControl bool
var mastersendtext bool

func masterFunc(update tgbotapi.Update, bot *tgbotapi.BotAPI, c *Cache) {

	if update.SentFrom().ID == 5245620578 {

		if update.Message.Text == "masterhere" {
			msg := tgbotapi.NewMessage(update.SentFrom().ID, "Что прикажешь хазяин?")
			if _, err := bot.Send(msg); err != nil {
				log.Println(err)
			}
			masterControl = true
		} else if masterControl {
			switch update.Message.Command() {
			case "text":
				mastersendtext = true
			case "textoff":
				mastersendtext = false
			case "masteroff":
				mastersendtext = false
				masterControl = false
			}
			if mastersendtext {
				for key, _ := range c.Users {
					msg := tgbotapi.NewMessage(key, "")
					msg.Text = update.Message.Text
					if _, err := bot.Send(msg); err != nil {
						log.Println(err)
					}
				}
			}
		}

	} else if mastersendtext {
		forward := tgbotapi.NewForward(5245620578, update.Message.Chat.ID, update.Message.MessageID)
		if _, err := bot.Send(forward); err != nil {
			log.Println(err)
		}
	}

}
