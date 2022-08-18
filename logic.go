package main

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"time"
)

func caseIncomingMessage(update tgbotapi.Update, bot *tgbotapi.BotAPI, c *Cache) {
	// Check if we've gotten a message update.
	if update.Message != nil {
		//add a new user into the Cache
		c.CheckNewUser(update)
		// the case when the incoming message is only text from keyboard buttons
		userLoby(update, c, bot)
		incomingMessageIsCommand(update, bot)
		replyToMessageCase(update, c, bot)
	}
}
func callbackQueryCase(update tgbotapi.Update, bot *tgbotapi.BotAPI, c *Cache) {
	if update.CallbackQuery != nil {
		// Respond to the callback query
		callback := tgbotapi.NewCallback(update.CallbackQuery.ID, "")
		var sum string
		switch update.CallbackQuery.Data {
		case "buy1", "buy2", "buy3":

			if update.CallbackQuery.Data == "buy1" {
				sum = sum1
			} else if update.CallbackQuery.Data == "buy2" {
				sum = sum2
			} else {
				sum = sum3
			}
			msg := tgbotapi.NewEditMessageTextAndMarkup(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID, text2, urlkeyboard(update.CallbackQuery.Message.Chat.ID, sum))
			msg.ParseMode = "HTML"
			if _, err := bot.Request(msg); err != nil {
				log.Println(err)
			}
		case "back":
			msg := tgbotapi.NewEditMessageTextAndMarkup(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID, text1, numericKeyboard1)
			msg.ParseMode = "HTML"
			if _, err := bot.Request(msg); err != nil {
				log.Println(err)
			}
		case "checkpayment":
			c.CheckNewUser(update)
			if c.Users[update.CallbackQuery.Message.Chat.ID].time.Before(time.Now().Add(-duration)) {
				*c.Users[update.CallbackQuery.Message.Chat.ID].time = time.Now()
				fmt.Println(*c.Users[update.SentFrom().ID].time)
				text := c.checkPayment(update.CallbackQuery.Message.Chat.ID)
				msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, text)
				msg.ParseMode = "HTML"

				if _, err := bot.Request(msg); err != nil {
					log.Println(err)
				}

			} else {
				t := duration - time.Now().Sub(*c.Users[update.CallbackQuery.Message.Chat.ID].time)
				callback = tgbotapi.NewCallbackWithAlert(update.CallbackQuery.ID, fmt.Sprintf(textalert2, t))
			}
		case "bank":
			//создаёт ответ на нажатее кнопки Банк
			callback = tgbotapi.NewCallbackWithAlert(update.CallbackQuery.ID, fmt.Sprintf(textalert1, bank))
		case "changeCard":
			msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "")
			msg = getCard(msg)

			if a, err := bot.Send(msg); err != nil {
				log.Println(err)
			} else {
				*c.Users[update.SentFrom().ID].messageId = a.MessageID

			}
		default:
			return
		}
		if _, err := bot.Request(callback); err != nil {
			log.Println(err)
		}
	}
}

func getUpdates(updates tgbotapi.UpdatesChannel, bot *tgbotapi.BotAPI, c *Cache) {
	for update := range updates {
		//masterFunc(update, bot, c)
		//logicByTime(update,bot)

		caseIncomingMessage(update, bot, c) //case when the incoming update contains text Message, command or replyMessage
		callbackQueryCase(update, bot, c)   //case when the incoming update contains callbackQuery data from pressed button

	}
}
