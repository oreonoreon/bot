package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"time"
)

// Вызов переданной функции раз в сутки в указанное время.
func callAt(hour, min, sec int, f func(), D time.Duration) error {
	loc, err := time.LoadLocation("Local")
	if err != nil {
		return err
	}

	// Вычисляем время первого запуска.
	now := time.Now().Local()
	firstCallTime := time.Date(
		now.Year(), now.Month(), now.Day(), hour, min, sec, 0, loc)
	if firstCallTime.Before(now) {
		// Если получилось время раньше текущего, прибавляем сутки.
		firstCallTime = firstCallTime.Add(time.Hour * 24)
	}

	// Вычисляем временной промежуток до запуска.
	durationUntilStart := firstCallTime.Sub(time.Now().Local())

	go func() {
		time.Sleep(durationUntilStart)
		for {
			f()
			// Следующий запуск через сутки или через D.
			time.Sleep(D) // периуд повтора time.Hour * 24 или любой другой периуд повтора

			//m = <-time.After(D)

		}
	}()

	return nil
}

// Ваша функция.
func (x xrr) sendMessageAtTime() {
	for k, value := range x.c.Users {
		msg := tgbotapi.NewMessage(k, textEnd) //отправка повторяющегося текстово сообщения
		for key, _ := range value.Tickets {    // in the end of lottery deleting all tickets value
			value.Tickets[key] = 0
		}
		if _, err := x.BOT.Send(msg); err != nil {
			log.Println("error occurred when trying to send msg to user at the end of timer ", err)
		}
		bank = 0
	}
}

type xrr struct {
	BOT *tgbotapi.BotAPI
	c   *Cache
}

func newXrr(b *tgbotapi.BotAPI, c *Cache) *xrr {
	return &xrr{BOT: b, c: c}
}
func callAtStartOfProgram(f func(), D time.Duration) {
	go func() {
		for {
			f()
			// Следующий запуск через через D.
			time.Sleep(D) // периуд повтора time.Hour * 24 или любой другой периуд повтора
		}
	}()
}
