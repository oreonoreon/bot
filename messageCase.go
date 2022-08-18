package main

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func getCard(msg tgbotapi.MessageConfig) tgbotapi.MessageConfig {
	msg.Text = enterCardNumber
	msg.ReplyMarkup = tgbotapi.ForceReply{ForceReply: true, InputFieldPlaceholder: "Укажите карту"}
	return msg
}

func userLoby(update tgbotapi.Update, c *Cache, bot *tgbotapi.BotAPI) {
	if !update.Message.IsCommand() {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
		msg.ParseMode = "HTML"
		switch update.Message.Text {
		case myCard:
			if *c.Users[update.Message.Chat.ID].Card != "" {
				msg.Text = fmt.Sprintf("Номер вашей карты %v ", *c.Users[update.Message.Chat.ID].Card)
				msg.ReplyMarkup = numericKeyboard4
			} else {
				msg.Text = fmt.Sprintf(changeCardNumber)
				msg.ReplyMarkup = numericKeyboard4
			}

		case myTickets:
			msg.Text = c.getTickets(update.Message.Chat.ID)
		case buyTickets:
			msg.Text = text1
			msg.ReplyMarkup = numericKeyboard1
		case info:
			msg.Text = textInfo
		default:
			msg.ReplyMarkup = numericKeyboard5
		}
		if _, err := bot.Send(msg); err != nil {
			log.Println(err)
		}
	}

}

func incomingMessageIsCommand(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	if update.Message.IsCommand() {

		switch update.Message.Command() {
		case "start":
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
			msg.Text = text0
			msg.ParseMode = "HTML"
			msg.ReplyMarkup = numericKeyboard5

			if _, err := bot.Send(msg); err != nil {
				log.Println(err)
			}
			msg1 := tgbotapi.NewMessage(update.Message.Chat.ID, "")
			msg1.Text = text1
			msg1.ParseMode = "HTML"
			msg1.ReplyMarkup = numericKeyboard1

			if _, err := bot.Send(msg1); err != nil {
				log.Println(err)
			}

		default:
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
			msg.Text = text3
			msg.ReplyMarkup = numericKeyboard5
			if _, err := bot.Send(msg); err != nil {
				log.Println(err)
			}
		}

	}

}
func replyToMessageCase(update tgbotapi.Update, c *Cache, bot *tgbotapi.BotAPI) {
	if update.Message.ReplyToMessage != nil && update.Message.ReplyToMessage.MessageID == *c.Users[update.Message.Chat.ID].messageId {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
		*c.Users[update.Message.Chat.ID].Card = update.Message.Text
		*c.Users[update.Message.Chat.ID].messageId = 0
		msg.Text = cardAccepted
		msg.ReplyMarkup = numericKeyboard5
		if _, err := bot.Send(msg); err != nil {
			log.Println(err)
		}
	}
}
